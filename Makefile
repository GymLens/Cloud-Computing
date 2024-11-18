APP_NAME=myapp
BINARY=bin/$(APP_NAME)

.PHONY: all build run clean test

all: build

build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o $(BINARY) ./app/cmd/app

run: build
	@echo "Running..."
	@./$(BINARY)

clean:
	@echo "Cleaning..."
	@rm -rf bin

test:
	@echo "Running tests..."
	@go test ./...
