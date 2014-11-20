package simplessh

import "bytes"
import "code.google.com/p/go.crypto/ssh"
import "errors"
import "io/ioutil"
import "os/user"
import "strconv"

type Config struct {
	User     string
	Host     string
	Port     int
	Password string
	KeyPaths []string
}

type Client struct {
	SshClient *ssh.Client
}

func NewClient(config *Config) (client *Client, err error) {
	clientConfig := &ssh.ClientConfig{
		User: config.User,
	}

	if config.Password != "" {
		password := ssh.Password(config.Password)
		clientConfig.Auth = append(clientConfig.Auth, password)
	}

	if len(config.KeyPaths) == 0 {

		usr, err := user.Current()
		if err != nil {
			return client, err
		}

		keyPath := usr.HomeDir + "/.ssh/id_rsa"
		key, err := makePrivateKey(keyPath)

		if err != nil {
			return client, err
		}

		clientConfig.Auth = append(clientConfig.Auth, ssh.PublicKeys(key))
	} else {

		keys, err := makePrivateKeys(config.KeyPaths)

		if err != nil {
			return client, err
		}

		clientConfig.Auth = append(clientConfig.Auth, ssh.PublicKeys(keys...))
	}

	if config.Port == 0 {
		config.Port = 22
	}

	hostAndPort := config.Host + ":" + strconv.Itoa(config.Port)
	sshClient, err := ssh.Dial("tcp", hostAndPort, clientConfig)

	if err != nil {
		return client, errors.New("Failed to dial: " + err.Error())
	}

	return &Client{sshClient}, nil
}

func (client *Client) Run(command string) (stdoutStr string, stderrStr string, err error) {
	session, err := client.SshClient.NewSession()
	if err != nil {
		return
	}
	defer session.Close()

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	session.Stdout = &stdout
	session.Stderr = &stderr

	err = session.Run(command)
	stdoutStr = stdout.String()
	stderrStr = stderr.String()

	return
}

func (client *Client) Close() {
	client.SshClient.Close()
}

func makePrivateKeys(keyPaths []string) (keys []ssh.Signer, err error) {
	for _, keyPath := range keyPaths {

		if key, err := makePrivateKey(keyPath); err == nil {
			keys = append(keys, key)
		} else {
			return keys, err
		}
	}

	return
}

func makePrivateKey(keyPath string) (key ssh.Signer, err error) {
	buffer, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	key, err = ssh.ParsePrivateKey(buffer)
	if err != nil {
		return
	}

	return
}
