package ports

import (
	"github.com/lucasvavon/slipx-api/internal/core/domain"
)

// UserRepository is an interface for interacting with user-related data
type UserRepository interface {
	GetUsers() ([]domain.User, error)
	GetUser(id *int) (domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUser(id *int, user *domain.User) error
	DeleteUser(id *int) error
}
