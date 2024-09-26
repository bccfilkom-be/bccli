# bccli - A CLI for Back-End Engineers

`bccli` is a command-line interface (CLI) tool designed for back-end engineers to quickly set up projects based on clean architecture principles. It provides commands to generate the essential structure of an application, including handlers, services, repositories, and infrastructure components.

## Features

- **Project Initialization**: Quickly set up a new project directory with all the necessary folders.
- **Application Components**: Generate handler, service, and repository files for your application modules.
- **Domain Entities**: Create domain-specific structs with ease.
- **Infrastructure Setup**: Scaffold infrastructure like databases inside your project.

## Installation

Make sure you have [Go](https://golang.org/dl/) installed on your device.

```bash
go install github.com/bccfilkom-be/bccli
```

## Quick Start

### Initialize a new project


```bash
bccli init todoapp
```

This will generate the structure like this
```bash
todoapp/
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

### Create a new app module

Generates the handler, service, and repository files for a specific app module.

- Create the app handler, service, and repository files inside the `internal/<app_name>/` folder.
- If you also want to generate data storage implementation (MySQL, PostgreSQL, etc.) for the repository, you can pass a database flag and choose your database.

```bash
bccli app todo
```

This will generate
```bash
internal/
├── domain/
├── infra/
└── todo/
    ├── interface/
    │   └── rest/
    │       └── todo.go
    ├── repository/
    └── service/
        └── todo.go
```
### Create a domain entity

Creates a domain entity file containing the struct for a given domain name.
- Creates a new file in the internal/domain/ folder for the given domain.

```bash
bccli domain todo
```
This will generate:
```bash
internal/
├── domain/
│   └── todo.go
└── infra/
```

### Set up Infrastucture
Sets up the infrastructure for your project.
- --db: Initializes a database setup file inside the internal/infra/ folder.

```bash
bccli infra --db=mysql
```
This will generate:
```bash
internal/
├── domain/
└── infra/
    └── mysql.go
```

## License
This repository is licensed under the [MIT License](LICENSE).
