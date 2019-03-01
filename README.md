# Go Skeleton

A go web application skeleton.

## Usage

### Download
```
go get github.com/ivanterekh/go-skeleton
```

### Build app
```
make
```

By default binary is put to the source code directory. But you can specify output app location:
```
APP=/path/to/binary make
```

### Run app
```
make run
```

Or with specified binary destination:
```
APP=/path/to/binary make run
```

### Run all tests
```
make test
```

### Run linters and formatting tools
```
make lint
```

