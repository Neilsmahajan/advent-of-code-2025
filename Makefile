.PHONY: build run test clean all

# Default target
build:
	go build -o bin/aoc ./cmd/aoc/main.go

# Run specific day and part (optional INPUT=<filename or path>)
run:
	@if [ -z "$(DAY)" ]; then \
		echo "Usage: make run DAY=<day> [PART=<part>] [INPUT=<input file>]"; \
		echo "Example: make run DAY=4 PART=1 INPUT=example_input.txt"; \
		exit 1; \
	fi
	@PART=$${PART:-1}; \
	if [ -n "$(INPUT)" ]; then \
		go run ./cmd/aoc/main.go -day=$(DAY) -part=$$PART -input="$(INPUT)"; \
	else \
		go run ./cmd/aoc/main.go -day=$(DAY) -part=$$PART; \
	fi

# Run all implemented solutions
all:
	go run ./cmd/aoc/main.go -all

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Check for updates to dependencies
check-updates:
	go list -u -m all

# Help
help:
	@echo "Available targets:"
	@echo "  build      - Build the project binary"
	@echo "  run        - Run specific day: make run DAY=1 PART=1 [INPUT=example_input.txt]"
	@echo "  all        - Run all implemented solutions"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  deps       - Install/update dependencies"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter"
	@echo "  help       - Show this help"
