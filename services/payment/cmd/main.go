package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/edwinjordan/golang_microservices/services/payment/internal/config"
	grpcHandler "github.com/edwinjordan/golang_microservices/services/payment/internal/delivery/grpc"
	httpHandler "github.com/edwinjordan/golang_microservices/services/payment/internal/delivery/http"
	"github.com/edwinjordan/golang_microservices/services/payment/internal/repository"
	"github.com/edwinjordan/golang_microservices/services/payment/internal/usecase"
	pb "github.com/edwinjordan/golang_microservices/services/payment/pkg/pb"
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
	paymentRepo := repository.NewPostgresPaymentRepository(db)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, cfg.OrderGRPCAddr)

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port: %v", err)
		}

		grpcServer := grpc.NewServer()
		paymentGRPCHandler := grpcHandler.NewPaymentGRPCHandler(paymentUsecase)
		pb.RegisterPaymentServiceServer(grpcServer, paymentGRPCHandler)

		log.Printf("gRPC server listening on port %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server
	router := gin.Default()
	paymentHandler := httpHandler.NewPaymentHandler(paymentUsecase)

	router.GET("/health", paymentHandler.Health)
	router.POST("/payments", paymentHandler.ProcessPayment)
	router.GET("/payments/:id", paymentHandler.GetPayment)

	log.Printf("HTTP server listening on port %s", cfg.HTTPPort)
	if err := router.Run(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func initSchema(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS payments (
		id VARCHAR(36) PRIMARY KEY,
		order_id VARCHAR(36) NOT NULL,
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
