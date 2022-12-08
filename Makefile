BINARY:=sentinel

default: build

.PHONY: build
build:
	@mkdir -p bin/
	go build -o bin/$(BINARY) main.go

.PHONY: test
test:
	go test main.go

.PHONY: run
run:
	@echo "Running sentinel server..."
	@go run main.go

redis-server:
	@echo "Running redis server..."
	resredis-server

clean:
	rm -rf bin/
