# golang_microservices

A Golang microservices project built with **Clean Architecture**, featuring three independent services (User, Order, Payment) that communicate via REST and gRPC.

## 🚀 Quick Start

### Prerequisites
- Go 1.22+
- Docker & Docker Compose
- Make

### Start All Services

```bash
# Start all services with Docker Compose
make up

# View logs
make logs
```

### Access Services

- **User Service**: http://localhost:8081/health
- **Order Service**: http://localhost:8082/health  
- **Payment Service**: http://localhost:8083/health

## 📚 Documentation

See [ARCHITECTURE.md](ARCHITECTURE.md) for detailed architecture documentation.

## 🛠️ Available Commands

```bash
make help       # Show all available commands
make build      # Build all services
make test       # Run tests
make up         # Start services with Docker
make down       # Stop all services
make clean      # Clean up everything
make rebuild    # Rebuild and restart
```

## 🏗️ Architecture

This project follows **Clean Architecture** principles with:

- **Domain Layer**: Business entities and interfaces
- **Use Case Layer**: Application business rules  
- **Repository Layer**: Data access implementations
- **Delivery Layer**: HTTP (Gin) and gRPC handlers
- **Config Layer**: Configuration management

Each service has its own PostgreSQL database and can be deployed independently.

## 📦 Services

- **User Service** (8081/9091): User management
- **Order Service** (8082/9092): Order processing with user validation
- **Payment Service** (8083/9093): Payment processing with order validation

## 🔗 Inter-Service Communication

Services communicate via gRPC:
- Order Service → User Service (user validation)
- Payment Service → Order Service (order validation)

## 📝 API Examples

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
  -d '{"user_id": "<user-id>", "product": "Laptop", "amount": 1500.00}'
```

### Process Payment
```bash
curl -X POST http://localhost:8083/payments \
  -H "Content-Type: application/json" \
  -d '{"order_id": "<order-id>", "amount": 1500.00}'
```

## 🧪 Testing

```bash
make test
```

## 📄 License

MIT License