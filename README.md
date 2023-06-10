<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# User Login And Register API

## About the project

The project Created REST API Example using golang, gin, and tools JWT. This project is created two end point. Login and Register user 
## Getting started

- [Setup Golang 19](https://go.dev/doc/install)
- [Gin-gonic](https://gin-gonic.com/docs/quickstart/)
- [Mysql or Mysql Docker](https://hub.docker.com/_/mysql)
- [JWT](https://jwt.io/)

### Layout

```tree
├── .husky
│   └── hooks
│       └── pre-commit
├── .vscode
|   └── settings.json
├── config
|   └── config.go
|   └── jwt.private.pem
|   └── jwt.public.pem
├── database
|   └── .migrations
|   └── .query
|   └── mysql.go
├── docs
|   └── docs.go
|   └── swagger.json
|   └── swagger.yaml
|   └── Documentations about the repository.pdf
├── internal
|   └── constants
|   └── entity
|   └── handler
|       └── mocks
|   └── helper
|   └── repository
|       └── mocks
|   └── request
|   └── response
|   └── router
|   └── usecase
|       └── mocks
|   └── utils
├── package
|   └── server
|       └── server.go
├── .env
├── .env.example
├── .gitignore
├── main.go
├── makefile
└── README.md
```

A brief description of the layout:

- `.husky` 
- `.vscode` is the vscode settings.json add a new vocabulary.
- `config` contains configure for this project.
- `database` contains database of the project.
- `docs` contains documentations of the project.
- `internal` contains main directory, each subdirectory for this project.
- `package` contains directory server for create server to this project.
- `.env/.env.example` scripts or environment for develop.
- `.gitignore` varies per project, but all projects need to ignore `bin` directory.
- `Makefile` is used to build the project. **You need to tweak the variables based on your project**.
- `README.md` is a detailed description of the project.
<!-- END doctoc generated TOC please keep comment here to allow auto update -->

