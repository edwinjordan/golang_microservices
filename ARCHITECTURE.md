# Microservices Architecture Documentation

## Overview

This project is a microservices-based application built with **Go 1.22+** following **Clean Architecture** principles. It consists of three independent services: User, Order, and Payment, each with its own PostgreSQL database.

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                          Client Layer                           │
└─────────────────────────────────────────────────────────────────┘
                                 │
                    ┌────────────┼────────────┐
                    │            │            │
         ┌──────────▼──────┐  ┌──▼──────────┐  ┌──▼──────────────┐
         │  User Service   │  │Order Service│  │ Payment Service │
         │  HTTP: 8081     │  │ HTTP: 8082  │  │  HTTP: 8083     │
         │  gRPC: 9091     │  │ gRPC: 9092  │  │  gRPC: 9093     │
         └────────┬────────┘  └──────┬──────┘  └──────┬──────────┘
                  │                  │                 │
         ┌────────▼────────┐  ┌──────▼──────┐  ┌──────▼──────────┐
         │  PostgreSQL     │  │ PostgreSQL  │  │  PostgreSQL     │
         │  users_db       │  │ orders_db   │  │  payments_db    │
         │  Port: 5433     │  │ Port: 5434  │  │  Port: 5435     │
         └─────────────────┘  └─────────────┘  └─────────────────┘
```

## Clean Architecture Layers

Each service follows Clean Architecture with the following layers:

### 1. Domain Layer (`internal/domain`)
- **Entities**: Core business objects (User, Order, Payment)
- **Interfaces**: Repository and use case interfaces
- **Business Rules**: Core business logic independent of frameworks

### 2. Use Case Layer (`internal/usecase`)
- **Application Logic**: Orchestrates data flow between layers
- **Business Rules**: Application-specific business rules
- **Interfaces**: Defines contracts for external dependencies

### 3. Repository Layer (`internal/repository`)
- **Data Access**: Implements database operations
- **Persistence Logic**: PostgreSQL-specific implementations
- **Interface Implementation**: Implements domain repository interfaces

### 4. Delivery Layer (`internal/delivery`)
- **HTTP Handlers**: REST API endpoints using Gin framework
- **gRPC Handlers**: gRPC service implementations
- **Input/Output**: Request/Response transformation

### 5. Configuration Layer (`internal/config`)
- **Environment Variables**: Configuration management
- **Database Connection**: Connection string generation

## Services

### User Service

**Responsibilities:**
- User management (create, retrieve, update, delete)
- User validation for other services

**Endpoints:**
- `POST /users` - Create a new user
- `GET /users/:id` - Get user by ID
- `GET /health` - Health check

**gRPC Methods:**
- `GetUser` - Retrieve user information
- `CreateUser` - Create a new user
- `ValidateUser` - Validate user existence

**Database:** `users_db` (PostgreSQL)

### Order Service

**Responsibilities:**
- Order management
- User validation via gRPC call to User Service

**Endpoints:**
- `POST /orders` - Create a new order
- `GET /orders/:id` - Get order by ID
- `GET /health` - Health check

**gRPC Methods:**
- `GetOrder` - Retrieve order information
- `CreateOrder` - Create a new order

**Database:** `orders_db` (PostgreSQL)

**Dependencies:**
- User Service (gRPC) - for user validation

### Payment Service

**Responsibilities:**
- Payment processing
- Order validation via gRPC call to Order Service

**Endpoints:**
- `POST /payments` - Process a payment
- `GET /payments/:id` - Get payment by ID
- `GET /health` - Health check

**gRPC Methods:**
- `ProcessPayment` - Process a new payment
- `GetPayment` - Retrieve payment information

**Database:** `payments_db` (PostgreSQL)

**Dependencies:**
- Order Service (gRPC) - for order validation

## Technology Stack

### Core Technologies
- **Language**: Go 1.22+
- **Web Framework**: Gin (HTTP/REST API)
- **RPC Framework**: gRPC (inter-service communication)
- **Database**: PostgreSQL 15
- **Containerization**: Docker & Docker Compose

### Go Packages
- `github.com/gin-gonic/gin` - HTTP web framework
- `google.golang.org/grpc` - gRPC framework
- `google.golang.org/protobuf` - Protocol Buffers
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/google/uuid` - UUID generation

## Database Schema

### users_db
```sql
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

### orders_db
```sql
CREATE TABLE orders (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    product VARCHAR(255) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

### payments_db
```sql
CREATE TABLE payments (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

## Communication Patterns

### REST API (Client ↔ Services)
- **Protocol**: HTTP/JSON
- **Framework**: Gin
- **Usage**: External client interactions

### gRPC (Service ↔ Service)
- **Protocol**: HTTP/2 + Protocol Buffers
- **Framework**: gRPC
- **Usage**: Inter-service communication
- **Benefits**: Type-safe, high-performance, bi-directional streaming

## Project Structure

```
golang_microservices/
├── services/
│   ├── user/
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── config/
│   │   │   ├── domain/
│   │   │   ├── repository/
│   │   │   ├── usecase/
│   │   │   └── delivery/
│   │   │       ├── http/
│   │   │       └── grpc/
│   │   └── pkg/
│   │       └── pb/
│   ├── order/
│   │   └── (same structure as user)
│   └── payment/
│       └── (same structure as user)
├── proto/
│   ├── user.proto
│   ├── order.proto
│   └── payment.proto
├── docker-compose.yml
├── Makefile
├── go.work
├── .env
└── ARCHITECTURE.md
```

## Getting Started

### Prerequisites
- Go 1.22+
- Docker & Docker Compose
- Protocol Buffers compiler (protoc)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/edwinjordan/golang_microservices.git
cd golang_microservices
```

2. Generate Protocol Buffer files:
```bash
make proto
```

3. Build all services:
```bash
make build
```

4. Start all services with Docker:
```bash
make up
```

### Testing

Run all tests:
```bash
make test
```

### Available Make Commands

- `make help` - Show all available commands
- `make proto` - Generate protobuf files
- `make build` - Build all services
- `make test` - Run tests for all services
- `make up` - Start all services with docker-compose
- `make down` - Stop all services
- `make clean` - Clean up build artifacts and docker volumes
- `make logs` - Show logs from all services
- `make restart` - Restart all services
- `make rebuild` - Rebuild and restart all services

## Health Checks

Each service provides a health check endpoint:

- User Service: `http://localhost:8081/health`
- Order Service: `http://localhost:8082/health`
- Payment Service: `http://localhost:8083/health`

## Environment Variables

All configuration is managed through the `.env` file:

### Database Configuration
- `USER_DB_HOST`, `USER_DB_PORT`, `USER_DB_USER`, `USER_DB_PASSWORD`, `USER_DB_NAME`
- `ORDER_DB_HOST`, `ORDER_DB_PORT`, `ORDER_DB_USER`, `ORDER_DB_PASSWORD`, `ORDER_DB_NAME`
- `PAYMENT_DB_HOST`, `PAYMENT_DB_PORT`, `PAYMENT_DB_USER`, `PAYMENT_DB_PASSWORD`, `PAYMENT_DB_NAME`

### Service Ports
- `USER_SERVICE_HTTP_PORT`, `USER_SERVICE_GRPC_PORT`
- `ORDER_SERVICE_HTTP_PORT`, `ORDER_SERVICE_GRPC_PORT`
- `PAYMENT_SERVICE_HTTP_PORT`, `PAYMENT_SERVICE_GRPC_PORT`

### gRPC Service Addresses
- `USER_GRPC_ADDR`, `ORDER_GRPC_ADDR`, `PAYMENT_GRPC_ADDR`

## API Examples

### Create User
```bash
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Create Order
```bash
curl -X POST http://localhost:8082/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user-uuid", "product": "Laptop", "amount": 1500.00}'
```

### Process Payment
```bash
curl -X POST http://localhost:8083/payments \
  -H "Content-Type: application/json" \
  -d '{"order_id": "order-uuid", "amount": 1500.00}'
```

## Design Principles

### 1. Clean Architecture
- **Separation of Concerns**: Each layer has a specific responsibility
- **Dependency Rule**: Dependencies point inward toward business logic
- **Testability**: Business logic is independent of frameworks and databases

### 2. Microservices Principles
- **Single Responsibility**: Each service manages one business domain
- **Database Per Service**: Each service has its own database
- **Independent Deployment**: Services can be deployed independently
- **Technology Heterogeneity**: Services can use different technologies

### 3. API Design
- **RESTful**: HTTP endpoints follow REST principles
- **gRPC**: Inter-service communication uses gRPC for performance
- **Versioning**: APIs can be versioned independently

## Scalability Considerations

- **Horizontal Scaling**: Each service can be scaled independently
- **Database Isolation**: Prevents database bottlenecks
- **Stateless Services**: Services maintain no session state
- **Load Balancing**: Ready for load balancer integration

## Security Considerations

- **Environment Variables**: Sensitive data stored in .env
- **Database Credentials**: Separate credentials per service
- **Network Isolation**: Services communicate via Docker network
- **Input Validation**: Request validation at HTTP/gRPC layer

## Future Enhancements

- API Gateway for unified entry point
- Service discovery (e.g., Consul, etcd)
- Distributed tracing (e.g., Jaeger)
- Message queue for async communication (e.g., RabbitMQ, Kafka)
- Authentication & Authorization (e.g., JWT)
- Circuit breaker pattern
- Rate limiting
- Monitoring & Logging (e.g., Prometheus, Grafana, ELK)

## Troubleshooting

### Services can't connect to database
- Ensure PostgreSQL containers are healthy: `docker-compose ps`
- Check logs: `make logs`

### gRPC communication failures
- Verify service dependencies in docker-compose.yml
- Check network connectivity between services

### Port conflicts
- Ensure ports 8081-8083, 9091-9093, 5433-5435 are available
- Modify .env file if needed

## License

This project is open source and available under the MIT License.
