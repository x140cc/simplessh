package simplessh

import "testing"
import "github.com/stretchr/testify/assert"
import "code.google.com/p/go.crypto/ssh"

func TestMakePrivateKeys(t *testing.T) {
	assert := assert.New(t)
	var keyPaths []string
	var keys []ssh.Signer
	var err error

	keyPaths = []string{}
	keys, err = makePrivateKeys(keyPaths)
	assert.Equal(len(keys), 0)
	assert.Nil(err)

	keyPaths = []string{"./test/id_rsa"}
	keys, err = makePrivateKeys(keyPaths)
	assert.Equal(len(keys), 1)
	assert.Nil(err)
}

func TestMakePrivateKey(t *testing.T) {
	assert := assert.New(t)
	var keyPath string
	var key ssh.Signer
	var err error

	keyPath = "./test/id_rsa"
	key, err = makePrivateKey(keyPath)
	assert.NotNil(key)
	assert.Nil(err)

	keyPath = "./test/id_rsa_404"
	key, err = makePrivateKey(keyPath)
	assert.Nil(key)
	assert.NotNil(err)
}
