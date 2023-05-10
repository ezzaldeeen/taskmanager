.PHONY: all clean build run test

# Set the name of your Go binary
BINARY_NAME := todo-app

# Set the Go compiler command
GO := go

all: clean build run

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)

build:
	@echo "Building..."
	@$(GO) build -o $(BINARY_NAME) ./cmd/todo

run:
	@echo "Running..."
	@./$(BINARY_NAME)

test:
	@echo "Running tests..."
	@$(GO) test ./...
