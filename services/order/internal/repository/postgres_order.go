package repository

import (
	"database/sql"
	"time"

	"github.com/edwinjordan/golang_microservices/services/order/internal/domain"
	"github.com/google/uuid"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) domain.OrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Create(order *domain.Order) error {
	order.ID = uuid.New().String()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.Status = "pending"

	query := `INSERT INTO orders (id, user_id, product, amount, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, order.ID, order.UserID, order.Product, order.Amount, order.Status, order.CreatedAt, order.UpdatedAt)
	return err
}

func (r *PostgresOrderRepository) GetByID(id string) (*domain.Order, error) {
	order := &domain.Order{}
	query := `SELECT id, user_id, product, amount, status, created_at, updated_at FROM orders WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.Product, &order.Amount, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *PostgresOrderRepository) Update(order *domain.Order) error {
	order.UpdatedAt = time.Now()
	query := `UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(query, order.Status, order.UpdatedAt, order.ID)
	return err
}
