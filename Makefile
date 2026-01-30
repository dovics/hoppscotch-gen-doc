# Binary name
BINARY_NAME=hoppscotch-gen-doc
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOINSTALL=$(GOCMD) install

# Build variables
VERSION?=dev
COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: all
all: clean build

## build: Build the application
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v
	@echo "Build complete: $(BINARY_NAME)"

## install: Install the application to $GOPATH/bin
.PHONY: install
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GOINSTALL) $(LDFLAGS) -v
	@echo "Install complete"

## clean: Clean build files
.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	@echo "Clean complete"

## test: Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

## test-coverage: Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool coverage -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## lint: Run linter
.PHONY: lint
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

## fmt: Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...
	@echo "Format complete"

## tidy: Tidy go modules
.PHONY: tidy
tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy
	@echo "Tidy complete"

## deps: Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOGET) -v ./...
	@echo "Dependencies downloaded"

## run: Build and run the application
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

## generate: Run example generation
.PHONY: generate
generate: build
	@echo "Generating example documentation..."
	./$(BINARY_NAME) generate -i example.json -o API.md
	@echo "Documentation generated: API.md"

## release: Build release binaries for multiple platforms
.PHONY: release
release:
	@echo "Building release binaries..."
	@mkdir -p dist
	@echo "Building for linux-amd64..."
	@GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64
	@echo "Building for linux-arm64..."
	@GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64
	@echo "Building for darwin-amd64..."
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64
	@echo "Building for darwin-arm64..."
	@GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64
	@echo "Building for windows-amd64..."
	@GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe
	@echo "Release binaries built in dist/"

## help: Show this help message
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.DEFAULT_GOAL := help
