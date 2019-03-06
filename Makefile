APP ?= bin/go-skeleton

.PHONY: build
build:
	@GO111MODULE=on CGO_ENABLED=0 go build -o ${APP}

.PHONY: run
run: build
	@./${APP}

.PHONY: docker-build
docker-build:
	@docker build -t ${APP} .

.PHONY: docker-run
docker-run: docker-build
	@docker run -p 8080:8080 ${APP} #TODO: make port configurable

.PHONY: test
test:lint
	@GO111MODULE=on go test -v ./...

.PHONY: lint
lint:
	@go fmt 
	@goimports -w .
	@GO111MODULE=on go vet
	
