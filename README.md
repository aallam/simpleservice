# Simple Web Service

A simple and configurable service for testing container orchestration setups.

## Endpoints

- `/` - returns a simple message
- `/health` - returns a simple health check message
- `/env` - returns the environment variables
- `/info` - returns the service information
- `/status` - returns the service status

## Run

Start by cloning the repository.

### Using Go

1. Install [Go 1.22](https://golang.org/dl/)
2. Run `make run` to start the server

### Using Docker

1. Install [Docker](https://www.docker.com/products/docker-desktop)
2. Run `make docker-boot` to build and start the server

## Prior Art

This project is inspired by the [mhausenblas/simpleservice](https://github.com/mhausenblas/simpleservice).
