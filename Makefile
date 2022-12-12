BINARY:=sentinel

default: build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	@mkdir -p bin/
	go build -o bin/$(BINARY) cmd/sentinel/main.go
	go build -o bin/example-starter examples/simple-go/main.go
	go build -o bin/example-worker examples/simple-go/worker/main.go

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	@echo "Running sentinel server..."
	@go run cmd/sentinel/main.go

redis-server:
	@echo "Running redis server..."
	resredis-server

temporal:
	@echo "Running temporal server..."
	@temporalite start --ephemeral --namespace default

clean:
	rm -rf bin/
