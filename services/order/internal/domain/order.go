package domain

import "time"

type Order struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Product   string    `json:"product" db:"product"`
	Amount    float64   `json:"amount" db:"amount"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id string) (*Order, error)
	Update(order *Order) error
}

type OrderUsecase interface {
	CreateOrder(userID, product string, amount float64) (*Order, error)
	GetOrder(id string) (*Order, error)
}
