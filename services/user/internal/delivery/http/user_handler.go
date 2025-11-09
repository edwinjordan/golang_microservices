package http

import (
	"net/http"

	"github.com/edwinjordan/golang_microservices/services/user/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.CreateUser(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userUsecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (h *UserHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "user-service",
	})
}
