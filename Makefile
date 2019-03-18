APP ?= bin/go-skeleton
ENV ?= dev
PORT ?= 8080

MODULE = github.com/ivanterekh/go-skeleton
VERSION = $(shell git describe --abbrev=0 --tags $(git rev-list --tags --max-count=1))

FILES := $(shell find . -type f -name "*.go")

.PHONY: build
build:
	@GO111MODULE=on CGO_ENABLED=0 go build -ldflags " \
		-X ${MODULE}/version.Version=${VERSION} \
		-X ${MODULE}/version.Commit=$(shell git rev-parse HEAD) \
		-X ${MODULE}/version.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S')" \
		-o ${APP} \
		./cmd/go-skeleton

.PHONY: run
run: build
	@bash -ac 'source .env.${ENV} && ENV=${ENV} ./${APP}'

.PHONY: docker-build
docker-build:
	@GO111MODULE=on go mod vendor
	@docker build -t ${APP}:${VERSION} .

.PHONY: docker-run
docker-run: docker-build
	@docker run -p ${PORT}:${PORT} -e PORT -e ENV=${ENV} --env-file .env.${ENV} ${APP}

.PHONY: test
test:lint
	@GO111MODULE=on go test -v ./...

.PHONY: lint
lint:
	@go fmt ./...
	@gogroup -order std,other,prefix=github.com/ivanterekh/ -rewrite $(shell find . -type f -name "*.go") 
	@GO111MODULE=on golangci-lint run
	@golint ./...
