# Tourism-FullStack-admin Makefile
# Created by: Khaled Mohamed (mo7amed4522@gmail.com)

.PHONY: all build run clean test docker-build docker-run swagger help

# Default target
all: build

# Build the server application
build:
	cd server && go build -o server main.go

# Run the server application
run:
	cd server && go run main.go

# Clean build artifacts
clean:
	rm -f server/server

# Run tests
test:
	cd server && go test ./... -v

# Build Docker image
docker-build:
	docker build -t tourism-fullstack-admin -f deploy/docker/Dockerfile .

# Run Docker container
docker-run:
	docker-compose -f deploy/docker-compose/docker-compose.yml up

# Generate Swagger documentation
swagger:
	cd server && swag init

# Install dependencies
deps:
	cd server && go mod tidy
	@echo "Installing web dependencies..."
	cd web && npm install

# Run frontend development server
web-dev:
	cd web && npm run dev

# Build frontend for production
web-build:
	cd web && npm run build

# Run database migrations
db-migrate:
	cd server && go run cmd/main.go init_db

# Check database status
db-check:
	cd server && go run cmd/main.go check_data

# Help command
help:
	@echo "Tourism-FullStack-admin Makefile"
	@echo "Available commands:"
	@echo "  make build         - Build the server application"
	@echo "  make run           - Run the server application"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make test          - Run tests"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make swagger       - Generate Swagger documentation"
	@echo "  make deps          - Install dependencies"
	@echo "  make web-dev       - Run frontend development server"
	@echo "  make web-build     - Build frontend for production"
	@echo "  make db-migrate    - Run database migrations"
	@echo "  make db-check      - Check database status"
	@echo "  make help          - Show this help message"