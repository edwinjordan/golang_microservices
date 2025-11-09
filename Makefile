.PHONY: help build test up down clean proto

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

proto: ## Generate protobuf files
	@echo "Generating protobuf files..."
	@protoc --go_out=services/user/pkg/pb --go_opt=paths=source_relative \
		--go-grpc_out=services/user/pkg/pb --go-grpc_opt=paths=source_relative \
		proto/user.proto
	@protoc --go_out=services/order/pkg/pb --go_opt=paths=source_relative \
		--go-grpc_out=services/order/pkg/pb --go-grpc_opt=paths=source_relative \
		proto/order.proto
	@protoc --go_out=services/payment/pkg/pb --go_opt=paths=source_relative \
		--go-grpc_out=services/payment/pkg/pb --go-grpc_opt=paths=source_relative \
		proto/payment.proto
	@echo "Protobuf files generated successfully"

build: ## Build all services
	@echo "Building all services..."
	@cd services/user && go build -o ../../bin/user-service ./cmd/main.go
	@cd services/order && go build -o ../../bin/order-service ./cmd/main.go
	@cd services/payment && go build -o ../../bin/payment-service ./cmd/main.go
	@echo "All services built successfully"

test: ## Run tests for all services
	@echo "Running tests..."
	@cd services/user && go test -v ./...
	@cd services/order && go test -v ./...
	@cd services/payment && go test -v ./...
	@echo "All tests completed"

up: ## Start all services with docker-compose
	@echo "Starting all services..."
	@docker-compose up -d
	@echo "All services started. Access:"
	@echo "  User Service:    http://localhost:8081/health"
	@echo "  Order Service:   http://localhost:8082/health"
	@echo "  Payment Service: http://localhost:8083/health"

down: ## Stop all services
	@echo "Stopping all services..."
	@docker-compose down
	@echo "All services stopped"

clean: ## Clean up build artifacts and docker volumes
	@echo "Cleaning up..."
	@rm -rf bin/
	@docker-compose down -v
	@echo "Cleanup completed"

logs: ## Show logs from all services
	@docker-compose logs -f

restart: down up ## Restart all services

rebuild: ## Rebuild and restart all services
	@echo "Rebuilding all services..."
	@docker-compose up -d --build
	@echo "All services rebuilt and restarted"
