# Variables
APP_NAME = app
SRC_DIR = cmd/app
BUILD_DIR = bin
GOFILES = $(wildcard $(SRC_DIR)/*.go)

# Build the application
.PHONY: build
build: 
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(GOFILES)

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(APP_NAME)

# Clean build artifacts
.PHONY: clean
clean: 
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Install dependencies
.PHONY: install
install: 
	@echo "Installing dependencies..."
	@go mod tidy

# Display help
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Usage:"
	@echo "  make build       Build the application"
	@echo "  make run         Run the application"
	@echo "  make clean       Clean build artifacts"
	@echo "  make test        Run tests"
	@echo "  make install     Install dependencies"
	@echo "  make help        Display this help message"
