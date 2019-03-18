FILES=$(find . -type f -name "*.go" | grep -v "vendor/")

go fmt ./...

if [ -f "$(go env GOPATH)/bin/gogroup"  ]; then
    $(go env GOPATH)/bin/gogroup -order std,other,prefix=github.com/ivanterekh/ -rewrite  $FILES;
else
    echo "warning: gogroup not found, imports won't be sorted
    you can get gogroup by running(in GOPATH):
    go get -u -v github.com/Bubblyworld/gogroup\n";
fi

if [ -f "$(go env GOPATH)/bin/golangci-lint"  ]; then
    GO111MODULE=on $(go env GOPATH)/bin/golangci-lint run -D errcheck;#errcheck disabled because of incorrect work with gin's c.Error
else
    echo "warning: golangci-lint not found, some linters won't be run
    you can get golangci-lint by running:
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.15.0\n";
fi

if [ -f "$(go env GOPATH)/bin/golint"  ]; then
    $(go env GOPATH)/bin/golint ./internal/...;
    $(go env GOPATH)/bin/golint ./cmd/...;
else
    echo "warning: golint not found
    you can get golint by running:
    go get -u -v golang.org/x/lint/golint";
fi

