package ports

import (
	"github.com/lucasvavon/slipx-api/internal/core/domain"
)

// UserRepository is an interface for interacting with user-related data
type UserRepository interface {
	// GetUsers selects a list of users with pagination
	GetUsers() ([]domain.User, error)
}
