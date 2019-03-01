# Go Skeleton

A go web application skeleton.

## Usage

Most common operations can be done with make. To spicify binary or docker image name use environment variables:
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
```
sudo make docker-run
```

### Run all tests
```
make test
```

### Run linters and formatting tools
```
make lint
```

