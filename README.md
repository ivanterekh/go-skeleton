# Go Skeleton

A go web application skeleton.

## Prerequisites

To run the application create `.env` files for setting up the environment. Currently `dev`, `prod` and `staging` modes are supported.
```
# .env.dev file example
ENV=dev
GIN_MODE=debug
PORT=8080
```

## Usage

Most common operations can be done with make. To specify binary or docker image name use environment variables:
```
APP=your-name make something
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

