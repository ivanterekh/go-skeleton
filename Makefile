APP ?= go-skeleton

build:lint
	@GO111MODULE=on go build -o ${APP}

run: build
	@./${APP}

test:lint
	@GO111MODULE=on go test -v ./...

lint:
	@go fmt 
	@goimports -w .
	@GO111MODULE=on go vet
	
