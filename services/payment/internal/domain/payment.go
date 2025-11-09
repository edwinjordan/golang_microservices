package domain

import "time"

type Payment struct {
	ID        string    `json:"id" db:"id"`
	OrderID   string    `json:"order_id" db:"order_id"`
	Amount    float64   `json:"amount" db:"amount"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type PaymentRepository interface {
	Create(payment *Payment) error
	GetByID(id string) (*Payment, error)
	Update(payment *Payment) error
}

type PaymentUsecase interface {
	ProcessPayment(orderID string, amount float64) (*Payment, error)
	GetPayment(id string) (*Payment, error)
}
