APP := go-skeleton

all: build

clean:
	@rm -f ${APP}

build: clean
	@go build -o ${APP}

run: build
	@./${APP}

test:
	@go test -v ./...
