# SimpleSSH

[![Build Status](https://travis-ci.org/meskyanichi/simplessh.svg)](https://travis-ci.org/meskyanichi/simplessh)

Lightweight and easy-to-use client wrapper around [go.crypto/ssh] making it
easy to ssh into- and perform commands on a remote machine.


### Installation

Simply import when you want to use it and `go get` it:

```go
import "github.com/meskyanichi/simplessh"
```


### Usage

Create a client and execute commands on the remote machine:

```go
package main

import "fmt"
import "os"
import "github.com/meskyanichi/simplessh"

func main() {
	config := simplessh.Config{
		User: "git",
		Host: "repositories.example.com",
	}

	client, err := simplessh.NewClient(&config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer client.Close()

	stdout, stderr, err := client.Run("whoami")

	if err != nil {
		fmt.Fprintln(os.Stderr, stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(stdout)
}
```

By default the above `config` will:

* Use the default ssh port (port 22)
* Use your `$HOME/.ssh/id_rsa` to authenticate

To additionally use a password for authentication:

```go
	config := simplessh.Config{
		Password: "example",
	}
```

To specify one or more different RSA keys:

```go
	usr, _ := user.Current()
	config := simplessh.Config{
		KeyPaths: []string{
			usr.HomeDir + "/.ssh/id_rsa",
			usr.HomeDir + "/.ssh/id_rsa_two",
			usr.HomeDir + "/.ssh/id_rsa_three",
		},
	}
```

### Contributing

Fork/Clone the repository:

```sh
git clone https://github.com/meskyanichi/simplessh.git
cd simplessh
```

Install dependencies:

```sh
go get -t
```

Run test suite:

```sh
go test ./...
```

Create branch, add/improve/test feature, submit pull request.


### Author / License

Copyright (c) 2014 Michael van Rooijen ( [@meskyanichi] )<br />
Released under the MIT [License].

[@meskyanichi]: https://twitter.com/meskyanichi
[License]: https://github.com/meskyanichi/simplessh/blob/master/LICENSE
[go.crypto/ssh]: https://code.google.com/p/go.crypto/ssh
