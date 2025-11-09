package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/edwinjordan/golang_microservices/services/order/internal/config"
	grpcHandler "github.com/edwinjordan/golang_microservices/services/order/internal/delivery/grpc"
	httpHandler "github.com/edwinjordan/golang_microservices/services/order/internal/delivery/http"
	"github.com/edwinjordan/golang_microservices/services/order/internal/repository"
	"github.com/edwinjordan/golang_microservices/services/order/internal/usecase"
	pb "github.com/edwinjordan/golang_microservices/services/order/pkg/pb"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database with retry
	var db *sql.DB
	var err error
	for i := 0; i < 30; i++ {
		db, err = sql.Open("postgres", cfg.GetDSN())
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Failed to connect to database, retrying in 2 seconds... (%d/30)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to database successfully")

	// Initialize database schema
	initSchema(db)

	// Initialize layers
	orderRepo := repository.NewPostgresOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, cfg.UserGRPCAddr)

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port: %v", err)
		}

		grpcServer := grpc.NewServer()
		orderGRPCHandler := grpcHandler.NewOrderGRPCHandler(orderUsecase)
		pb.RegisterOrderServiceServer(grpcServer, orderGRPCHandler)

		log.Printf("gRPC server listening on port %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server
	router := gin.Default()
	orderHandler := httpHandler.NewOrderHandler(orderUsecase)

	router.GET("/health", orderHandler.Health)
	router.POST("/orders", orderHandler.CreateOrder)
	router.GET("/orders/:id", orderHandler.GetOrder)

	log.Printf("HTTP server listening on port %s", cfg.HTTPPort)
	if err := router.Run(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func initSchema(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS orders (
		id VARCHAR(36) PRIMARY KEY,
		user_id VARCHAR(36) NOT NULL,
		product VARCHAR(255) NOT NULL,
		amount DECIMAL(10, 2) NOT NULL,
		status VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	);
	`
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}
	log.Println("Database schema initialized")
}
