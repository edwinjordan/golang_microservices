package repository

import (
	"database/sql"
	"time"

	"github.com/edwinjordan/golang_microservices/services/user/internal/domain"
	"github.com/google/uuid"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `INSERT INTO users (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *PostgresUserRepository) GetByID(id string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) Update(user *domain.User) error {
	user.UpdatedAt = time.Now()
	query := `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`
	_, err := r.db.Exec(query, user.Name, user.Email, user.UpdatedAt, user.ID)
	return err
}

func (r *PostgresUserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
