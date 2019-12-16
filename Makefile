.PHONY: all golangci-lint golangci-lint-fix test install-golangci-lint

all: golangci-lint-fix test

golangci-lint:
	golangci-lint run

golangci-lint-fix:
	golangci-lint run --fix

test:
	go test -race -count 1 -cover

install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.21.0
