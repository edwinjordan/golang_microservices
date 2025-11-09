package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/edwinjordan/golang_microservices/services/order/internal/domain"
	userpb "github.com/edwinjordan/golang_microservices/services/user/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type orderUsecase struct {
	orderRepo      domain.OrderRepository
	userGRPCClient userpb.UserServiceClient
}

func NewOrderUsecase(orderRepo domain.OrderRepository, userGRPCAddr string) domain.OrderUsecase {
	// Connect to user service
	conn, err := grpc.NewClient(userGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to user service: %v", err)
	}

	userClient := userpb.NewUserServiceClient(conn)

	return &orderUsecase{
		orderRepo:      orderRepo,
		userGRPCClient: userClient,
	}
}

func (u *orderUsecase) CreateOrder(userID, product string, amount float64) (*domain.Order, error) {
	if userID == "" || product == "" || amount <= 0 {
		return nil, errors.New("userID, product, and amount are required")
	}

	// Validate user exists via gRPC
	if u.userGRPCClient != nil {
		resp, err := u.userGRPCClient.ValidateUser(context.Background(), &userpb.ValidateUserRequest{UserId: userID})
		if err != nil || !resp.Valid {
			return nil, errors.New("invalid user")
		}
	}

	order := &domain.Order{
		UserID:  userID,
		Product: product,
		Amount:  amount,
	}

	err := u.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *orderUsecase) GetOrder(id string) (*domain.Order, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	order, err := u.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
