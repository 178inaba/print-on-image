.PHONY: all golangci-lint test install-golangci-lint

all: golangci-lint test

golangci-lint:
	golangci-lint run

test:
	go test -race -count 1 -cover

install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.21.0
