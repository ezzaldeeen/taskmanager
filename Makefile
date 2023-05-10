.PHONY: all clean build run test

# Set the name of your Go binary
BINARY_NAME := todo-app

all: clean build run

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)

build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) .

run:
	@echo "Running..."
	@./$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test ./...

coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
