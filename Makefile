.PHONY: build run test clean

BINARY_NAME=golang-playground

build:
	@go build -o $(BINARY_NAME)

run: build
	@./$(BINARY_NAME)

test:
	@go test ./...

clean:
	@go clean
	@rm $(BINARY_NAME)
