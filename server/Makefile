.PHONY: build run clean docker-build docker-run test tidy

# Build the application
build:
	go build -o bin/server ./main.go

# Run the application
run:
	go run ./main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -t playground-server .

# Run Docker container
docker-run:
	docker run -p 8080:8080 playground-server

# Run tests
test:
	go test ./...

# Tidy dependencies
tidy:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Development setup
dev-setup: tidy fmt
	@echo "Development environment setup complete"

# Full build and run
all: clean build run
