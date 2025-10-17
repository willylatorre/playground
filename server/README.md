# Playground Golang Server

A modern Golang REST API server built with Gin framework for learning purposes.

## Features

- RESTful API with CRUD operations
- CORS support for frontend communication
- In-memory data storage (perfect for learning)
- Docker containerization
- Sample data for immediate testing

## Quick Start

### Prerequisites

- Go 1.21 or later
- Docker (optional)

### Installation

1. Navigate to the server directory:
   ```bash
   cd server
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /api/health` - Server health status

### Items CRUD
- `GET /api/items` - Get all items
- `GET /api/items/:id` - Get item by ID
- `POST /api/items` - Create new item
- `PUT /api/items/:id` - Update item
- `DELETE /api/items/:id` - Delete item

### Sample Item Data Structure
```json
{
  "id": "uuid-string",
  "name": "Item Name",
  "data": "Item content",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Development

### Available Make Commands

- `make build` - Build the application
- `make run` - Run the application
- `make clean` - Clean build artifacts
- `make docker-build` - Build Docker image
- `make docker-run` - Run Docker container
- `make test` - Run tests
- `make tidy` - Tidy dependencies
- `make fmt` - Format code
- `make dev-setup` - Setup development environment

### Project Structure

```
server/
├── main.go           # Application entry point
├── go.mod            # Go module file
├── go.sum            # Dependency checksums
├── Makefile          # Build automation
├── Dockerfile        # Docker configuration
├── handlers/         # HTTP request handlers
│   └── handlers.go
└── models/           # Data models and storage
    └── models.go
```

## Docker Usage

Build and run with Docker:

```bash
make docker-build
make docker-run
```

Or manually:

```bash
docker build -t playground-server .
docker run -p 8080:8080 playground-server
```

## Testing the API

The server includes sample data. You can test endpoints using curl:

```bash
# Health check
curl http://localhost:8080/api/health

# Get all items
curl http://localhost:8080/api/items

# Create new item
curl -X POST http://localhost:8080/api/items \
  -H "Content-Type: application/json" \
  -d '{"name": "New Item", "data": "Item data"}'
```

## Learning Path

This skeleton provides a foundation for learning Golang web development:

1. **HTTP Routing**: Learn how Gin handles routes and middleware
2. **JSON Handling**: Understand request/response marshaling
3. **Data Modeling**: See basic struct definitions and operations
4. **Concurrency**: Notice the use of sync.RWMutex for thread safety
5. **Error Handling**: Observe proper HTTP status codes and error responses
6. **Docker**: Learn containerization basics

## Next Steps

- Add authentication/authorization
- Integrate with a real database (PostgreSQL, MongoDB)
- Add logging and monitoring
- Implement input validation
- Add unit and integration tests
- Set up CI/CD pipeline
