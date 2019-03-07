APP ?= bin/go-skeleton
ENV ?= dev
PORT ?= 8080

FILES := $(shell find . -type f -name "*.go")

.PHONY: build
build:
	@GO111MODULE=on CGO_ENABLED=0 go build -o ${APP}

.PHONY: run
run: build
	@bash -ac 'source .env.${ENV} && ENV=${ENV} ./${APP}'

.PHONY: docker-build
docker-build:
	@docker build -t ${APP} .

.PHONY: docker-run
docker-run: docker-build
	@docker run -p ${PORT}:${PORT} -e PORT -e ENV=${ENV} --env-file .env.${ENV} ${APP}

.PHONY: test
test:lint
	@GO111MODULE=on go test -v ./...

.PHONY: lint
lint:
	@go fmt ./...
	@gogroup -order std,other,prefix=github.com/ivanterekh/ -rewrite ${FILES} 
	@GO111MODULE=on golangci-lint run
	@golint ./...
