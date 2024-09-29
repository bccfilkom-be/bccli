# bccli

bccli is a simple cli tool for setting up go server project based on Three-layered Archictecture combined with Domain Driven Design.

## Installation

Make sure you have [Go](https://go.dev/doc/install) installed on your device.

```bash
go install github.com/bccfilkom-be/bccli
```

## Usage
```text
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
cmd/
└── api/
    └── main.go
Dockerfile
go.mod
go.sum
Makefile
```

### Create a domain entity
The ``domain generate`` command used to create a new domain entity with its layer.
```bash
bccli domain generate todo
```
This command generates todo domain representation and with its three-layered app in the following structure:
```text
internal/
├── domain/
│   └── todo.go
└── todo/
    ├── interface/
    │   └── rest/
    │       └── todo.go
    ├── repository/
    └── usecase/
        └── todo.go
```

### Set up Infrastucture
The ``infra generate`` command is used to generate 3rd party service libraries config or other external component of the app.
```bash
bccli infra generate postgresql
```
This command generates a basic MySQL configuration file:
```text
internal/
└── infra/
    └── postgresql.go
```
You can choose another database like mysql or mariadb.

## License
This repository is licensed under the [MIT License](LICENSE).
