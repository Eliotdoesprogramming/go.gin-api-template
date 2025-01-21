.PHONY: build run test clean

# Build the application
build:
	go build -o bin/api main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies and tools
deps:
	go mod tidy
	go install golang.org/x/lint/golint@latest
	go install github.com/swaggo/swag/cmd/swag@latest

# Run linter
lint:
	go vet ./...
	golint ./...

# Generate API documentation (requires swag)
docs:
	swag init

# Build and run in development mode
dev: build
	./bin/api

.DEFAULT_GOAL := build