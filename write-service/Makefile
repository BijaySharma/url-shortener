BINARY_NAME=write-service
CMD_PATH=./cmd/main.go

.PHONY: build run test clean tidy fmt vet lint lint-install docker-build

build:
	go build -o bin/$(BINARY_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)

test:
	go test ./... -v

clean:
	rm -rf bin

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

lint-install:
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: lint-install
	golangci-lint run ./...

docker-build:
	docker build -t $(BINARY_NAME) .
