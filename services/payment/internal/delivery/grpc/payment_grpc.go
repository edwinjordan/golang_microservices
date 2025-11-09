package grpc

import (
	"context"

	"github.com/edwinjordan/golang_microservices/services/payment/internal/domain"
	pb "github.com/edwinjordan/golang_microservices/services/payment/pkg/pb"
)

type PaymentGRPCHandler struct {
	pb.UnimplementedPaymentServiceServer
	paymentUsecase domain.PaymentUsecase
}

func NewPaymentGRPCHandler(paymentUsecase domain.PaymentUsecase) *PaymentGRPCHandler {
	return &PaymentGRPCHandler{
		paymentUsecase: paymentUsecase,
	}
}

func (h *PaymentGRPCHandler) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	payment, err := h.paymentUsecase.ProcessPayment(req.OrderId, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.ProcessPaymentResponse{
		Id:      payment.ID,
		OrderId: payment.OrderID,
		Amount:  payment.Amount,
		Status:  payment.Status,
	}, nil
}

func (h *PaymentGRPCHandler) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	payment, err := h.paymentUsecase.GetPayment(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPaymentResponse{
		Id:      payment.ID,
		OrderId: payment.OrderID,
		Amount:  payment.Amount,
		Status:  payment.Status,
	}, nil
}
