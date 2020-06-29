BINARY_PATH=bin/
SERVER_MAIN_PATH=cmd/
SERVER_MAIN_FILE=$(SERVER_MAIN_PATH)main.go
SERVER_BINARY_NAME=$(BINARY_PATH)main

run: build run-server

build:
	go build -o ./$(SERVER_BINARY_NAME) ./$(SERVER_MAIN_PATH)

test:
	go test ./... -cover -tags=unit

run-server:
	./$(SERVER_BINARY_NAME)