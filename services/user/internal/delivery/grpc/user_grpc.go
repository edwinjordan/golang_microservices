package grpc

import (
	"context"

	"github.com/edwinjordan/golang_microservices/services/user/internal/domain"
	pb "github.com/edwinjordan/golang_microservices/services/user/pkg/pb"
)

type UserGRPCHandler struct {
	pb.UnimplementedUserServiceServer
	userUsecase domain.UserUsecase
}

func NewUserGRPCHandler(userUsecase domain.UserUsecase) *UserGRPCHandler {
	return &UserGRPCHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserGRPCHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userUsecase.GetUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserGRPCHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.userUsecase.CreateUser(req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserGRPCHandler) ValidateUser(ctx context.Context, req *pb.ValidateUserRequest) (*pb.ValidateUserResponse, error) {
	valid, name, err := h.userUsecase.ValidateUser(req.UserId)
	if err != nil {
		return &pb.ValidateUserResponse{Valid: false}, nil
	}

	return &pb.ValidateUserResponse{
		Valid: valid,
		Name:  name,
	}, nil
}
