package usecase

import (
	"errors"

	"github.com/edwinjordan/golang_microservices/services/user/internal/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUser(name, email string) (*domain.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &domain.User{
		Name:  name,
		Email: email,
	}

	err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetUser(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) ValidateUser(id string) (bool, string, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return false, "", err
	}

	return true, user.Name, nil
}
