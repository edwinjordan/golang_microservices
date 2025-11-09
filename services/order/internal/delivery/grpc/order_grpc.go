package grpc

import (
	"context"

	"github.com/edwinjordan/golang_microservices/services/order/internal/domain"
	pb "github.com/edwinjordan/golang_microservices/services/order/pkg/pb"
)

type OrderGRPCHandler struct {
	pb.UnimplementedOrderServiceServer
	orderUsecase domain.OrderUsecase
}

func NewOrderGRPCHandler(orderUsecase domain.OrderUsecase) *OrderGRPCHandler {
	return &OrderGRPCHandler{
		orderUsecase: orderUsecase,
	}
}

func (h *OrderGRPCHandler) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := h.orderUsecase.GetOrder(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponse{
		Id:      order.ID,
		UserId:  order.UserID,
		Product: order.Product,
		Amount:  order.Amount,
		Status:  order.Status,
	}, nil
}

func (h *OrderGRPCHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order, err := h.orderUsecase.CreateOrder(req.UserId, req.Product, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:      order.ID,
		UserId:  order.UserID,
		Product: order.Product,
		Amount:  order.Amount,
		Status:  order.Status,
	}, nil
}
