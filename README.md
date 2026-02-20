# GIN-framework

A small example web service built with Gin (Go) demonstrating a simple layered architecture: handlers, services, repositories, and middleware.

**Status:** Example / learning project

## Features
- Simple user creation flow
- Middleware for logging and JWT (example)
- In-memory repository implementation

## Prerequisites
- Go 1.20+ installed

## Quickstart

1. Fetch dependencies:

```
go mod tidy
```

2. Run the server:

```
go run ./cmd/server
```

The server starts and listens according to the code in `cmd/server/main.go`.

## Build

Build a binary:

```
go build -o bin/server ./cmd/server
```

## Project Structure

- `cmd/server` - application entrypoint
- `internal/handler` - HTTP handlers
- `internal/service` - business logic
- `internal/repository` - data storage (in-memory)
- `internal/middleware` - middleware (logger, jwt)
- `internal/model` - data models

## Contributing

Open issues or pull requests — low friction changes and improvements are welcome.

## License

This repository doesn't include a license file. Add one if you intend to publish or share this project.

---
Created and committed by your assistant.
