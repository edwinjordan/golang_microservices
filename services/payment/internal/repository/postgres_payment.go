package repository

import (
	"database/sql"
	"time"

	"github.com/edwinjordan/golang_microservices/services/payment/internal/domain"
	"github.com/google/uuid"
)

type PostgresPaymentRepository struct {
	db *sql.DB
}

func NewPostgresPaymentRepository(db *sql.DB) domain.PaymentRepository {
	return &PostgresPaymentRepository{db: db}
}

func (r *PostgresPaymentRepository) Create(payment *domain.Payment) error {
	payment.ID = uuid.New().String()
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	query := `INSERT INTO payments (id, order_id, amount, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, payment.ID, payment.OrderID, payment.Amount, payment.Status, payment.CreatedAt, payment.UpdatedAt)
	return err
}

func (r *PostgresPaymentRepository) GetByID(id string) (*domain.Payment, error) {
	payment := &domain.Payment{}
	query := `SELECT id, order_id, amount, status, created_at, updated_at FROM payments WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Status, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PostgresPaymentRepository) Update(payment *domain.Payment) error {
	payment.UpdatedAt = time.Now()
	query := `UPDATE payments SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(query, payment.Status, payment.UpdatedAt, payment.ID)
	return err
}
