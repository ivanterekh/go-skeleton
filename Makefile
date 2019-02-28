APP ?= go-skeleton

build:
	@GO111MODULE=on go build -o ${APP}

run: build
	@./${APP}

test:
	@GO111MODULE=on go test -v ./...
