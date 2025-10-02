# Makefile for Go Task Manager CLI

BINARY_NAME=Taskmanager

# Build the CLI
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) ./cmd/taskmanager
	@echo "Build complete!"

# Run the CLI
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME) list

# Clean the binary
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	@echo "Clean complete!"

# Test 
test:
	@echo "Running tests..."
	go test ./...
