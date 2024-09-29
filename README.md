# bccli

bccli is a simple cli tool for setting up server project structure based on thee-layered archictecture that is controller, usecase, and repository.

## Installation

Make sure you have [Go](https://golang.org/dl/) installed on your device.

```bash
go install github.com/bccfilkom-be/bccli
```

## Usage
```
Usage:
  bccli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  domain      Generate domain components like handler, usecase, and repository.
  help        Help about any command
  infra       Infra layer command
  init        Initialize a new Go REST server project structure.

Flags:
  -h, --help   help for bccli
```

## Command Overview

### Initialize a new project
The ``init`` command sets up the initial project structure.

```bash
mkdir go-server
cd go-server
bccli init github.com/bccfilkom-be/go-server
```

This will bootstrap the current directory with the following structure:
```text
go-server
├── cmd
│   └── api
│       └── main.go
├── Dockerfile
├── go.mod
├── go.sum
└── Makefile
```

### Create a domain entity
The ``domain generate`` command used to create a new domain entity.

```bash
bccli domain generate todo
```
This command generates a domain for todo and creates related files in the following structure:
```bash
internal/
├── domain/
│   └── todo.go
├── todo/
│   ├── interface/
│   │   └── rest/
│   │       └── todo.go
│   ├── repository/
│   └── service/
│       └── todo.go
└── infra/
```

### Set up Infrastucture
The ``infra generate`` command is used to generate 3rd party service libraries config or other external component of the app

```bash
bccli infra generate mysql
```
This command generates a basic MySQL configuration file:
```bash
internal/
├── domain/
└── infra/
    └── mysql.go
```
You can choose another database like postgres or mariadb to scaffold the corresponding setup.
## License
This repository is licensed under the [MIT License](LICENSE).
