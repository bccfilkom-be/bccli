# ``bccli`` - A CLI for Back-End Engineers

``bccli`` is a simple cli tool for setting up server project structure based on thee-layered archictecture that is controller, usecase, and repository.

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
bccli init todo
```

This will create the project structure inside ``todo`` folder with the following structure:
```bash
todo/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   └── infra/
├── Dockerfile
├── go.mod
├── go.sum
└── Makefile
```

### Create a domain entity
The ``domain`` command used to create a new domain entity. This command scaffolds domain files in a structured manner, ensuring separation of concerns and clear layers.

```bash
bccli domain generate todo
```
This command generates a domain for ``todo`` and creates related files in the following structure:
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
Setting up databases and other infrastructure components with the ``infra`` command. Use it to scaffold infrastructure setup, such as databases, caching systems, or external services:

```bash
bccli infra generate --database=mysql
```
This command generates a basic MySQL configuration file:
```bash
internal/
├── domain/
└── infra/
    └── mysql.go
```
You can use the --database flag with other options like postgres or mariadb to scaffold the corresponding setup.
## License
This repository is licensed under the [MIT License](LICENSE).
