package http

import (
	"net/http"

	"github.com/edwinjordan/golang_microservices/services/order/internal/domain"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase domain.OrderUsecase
}

func NewOrderHandler(orderUsecase domain.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}

type CreateOrderRequest struct {
	UserID  string  `json:"user_id" binding:"required"`
	Product string  `json:"product" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}

type OrderResponse struct {
	ID      string  `json:"id"`
	UserID  string  `json:"user_id"`
	Product string  `json:"product"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderUsecase.CreateOrder(req.UserID, req.Product, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, OrderResponse{
		ID:      order.ID,
		UserID:  order.UserID,
		Product: order.Product,
		Amount:  order.Amount,
		Status:  order.Status,
	})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderUsecase.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, OrderResponse{
		ID:      order.ID,
		UserID:  order.UserID,
		Product: order.Product,
		Amount:  order.Amount,
		Status:  order.Status,
	})
}

func (h *OrderHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "order-service",
	})
}
