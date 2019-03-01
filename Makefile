APP ?= go-skeleton

.PHONY: build
build:lint
	@GO111MODULE=on go build -o ${APP}

.PHONY: run
run: build
	@./${APP}

.PHONY: test
test:lint
	@GO111MODULE=on go test -v ./...

.PHONY: lint
lint:
	@go fmt 
	@goimports -w .
	@GO111MODULE=on go vet
	
