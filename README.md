# ``bccli`` - A CLI for Back-End Engineers

`bccli` is a command-line interface (CLI) tool designed to assist back-end engineers in setting up project structures based on clean architecture. It helps developers quickly scaffold the necessary components such as handlers, services, repositories, and infrastructure, reducing the overhead of manual setup and keeping your codebase organized and maintainable.

## Why use bccli?
bccli is designed to simplify the workflow of back-end engineers by automating the setup of key components in a project, all based on clean architecture principles.

### Key Features

- **Project Initialization**: Set up a new project directory with the required structure and files, ready for development.
- **Domain Components**: Generate handlers, services, repositories, and domain-specific entities in one go, maintaining a clean separation of concerns.
- **Infrastructure Setup**: Quickly scaffold infrastructure components, such as databases, and integrate them into your project.

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
The ``init`` command sets up the initial project structure, making it easy to start developing right away. By simply running the following command:

```bash
bccli init todo
```

This will create a new folder ``todo`` with the following structure:
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
bccli domain todo
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
bccli infra --db=mysql
```
This command generates a basic MySQL configuration file:
```bash
internal/
├── domain/
└── infra/
    └── mysql.go
```
You can use the --db flag with other options like postgres or mariadb to scaffold the corresponding setup.
## License
This repository is licensed under the [MIT License](LICENSE).
