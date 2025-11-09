package http

import (
	"net/http"

	"github.com/edwinjordan/golang_microservices/services/payment/internal/domain"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUsecase domain.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase domain.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUsecase: paymentUsecase,
	}
}

type ProcessPaymentRequest struct {
	OrderID string  `json:"order_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}

type PaymentResponse struct {
	ID      string  `json:"id"`
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req ProcessPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.paymentUsecase.ProcessPayment(req.OrderID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, PaymentResponse{
		ID:      payment.ID,
		OrderID: payment.OrderID,
		Amount:  payment.Amount,
		Status:  payment.Status,
	})
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	payment, err := h.paymentUsecase.GetPayment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, PaymentResponse{
		ID:      payment.ID,
		OrderID: payment.OrderID,
		Amount:  payment.Amount,
		Status:  payment.Status,
	})
}

func (h *PaymentHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "payment-service",
	})
}
