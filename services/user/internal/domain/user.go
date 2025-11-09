package domain

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

type UserUsecase interface {
	CreateUser(name, email string) (*User, error)
	GetUser(id string) (*User, error)
	ValidateUser(id string) (bool, string, error)
}
