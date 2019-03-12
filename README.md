# Go Skeleton

A go web application skeleton.

## Prerequisites

To run the application create `.env.*` file for setting up the environment. Example of `.env.dev` file:
```
GIN_MODE=debug
PORT=8080
```

## Usage

Most common operations can be done with make. To specify app name or environment use environment variables:
```
# Makes docker image called superapp
sudo APP=superapp make docker-build
# Produces the binary /tmp/superapp
APP=/tmp/superapp make build
# Runs app with configuration from .env.prod
ENV=prod make run
```

### Download
```
go get github.com/ivanterekh/go-skeleton
```

### Build app
```
make build
```
or simply
```
make
```

### Run app
```
make run
```

### Build docker image
```
sudo make docker-build
```

### Run in docker container
Note: port must be specified in this command unless you are using `8080` in `.env`  file.
```
sudo PORT=8082 make docker-run
```

### Run all tests
```
make test
```

### Run linters and formatting tools
```
make lint
```

