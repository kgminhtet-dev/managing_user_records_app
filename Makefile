GO_VERSION := 1.23.1
APP_NAME := api
BUILD_DIR := bin
SRC := ./cmd/api/main.go

GO := go
GOTEST := go test
GOLINT := golangci-lint

# Targets
.PHONY: all fmt lint test build run clean

all: fmt build run clean

build:
	@echo "Building the application..."
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(SRC)

run:
	@echo "Running the application..."
	$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

lint:
	@echo "Linting the code..."
	$(GOLINT) run ./...

fmt:
	@echo "Formatting the code..."
	$(GO) fmt ./...

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
