package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/edwinjordan/golang_microservices/services/payment/internal/domain"
	orderpb "github.com/edwinjordan/golang_microservices/services/order/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type paymentUsecase struct {
	paymentRepo     domain.PaymentRepository
	orderGRPCClient orderpb.OrderServiceClient
}

func NewPaymentUsecase(paymentRepo domain.PaymentRepository, orderGRPCAddr string) domain.PaymentUsecase {
	// Connect to order service
	conn, err := grpc.NewClient(orderGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to order service: %v", err)
	}

	orderClient := orderpb.NewOrderServiceClient(conn)

	return &paymentUsecase{
		paymentRepo:     paymentRepo,
		orderGRPCClient: orderClient,
	}
}

func (u *paymentUsecase) ProcessPayment(orderID string, amount float64) (*domain.Payment, error) {
	if orderID == "" || amount <= 0 {
		return nil, errors.New("orderID and amount are required")
	}

	// Validate order exists via gRPC
	if u.orderGRPCClient != nil {
		_, err := u.orderGRPCClient.GetOrder(context.Background(), &orderpb.GetOrderRequest{Id: orderID})
		if err != nil {
			return nil, errors.New("invalid order")
		}
	}

	payment := &domain.Payment{
		OrderID: orderID,
		Amount:  amount,
		Status:  "completed",
	}

	err := u.paymentRepo.Create(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (u *paymentUsecase) GetPayment(id string) (*domain.Payment, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	payment, err := u.paymentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
