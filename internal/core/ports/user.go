package ports

import (
	"context"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
)

// UserRepository is an interface for interacting with user-related data
type UserRepository interface {
	// CreateUser inserts a new user into the database
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUserByID selects a user by id
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	// GetUserByEmail selects a user by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	// GetUsers selects a list of users with pagination
	GetUsers(ctx context.Context) ([]domain.User, error)
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, id uint) error
}

// UserService is an interface for interacting with user-related business logic
type UserService interface {
	// Register registers a new user
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUser returns a user by id
	GetUser(ctx context.Context, id uint) (*domain.User, error)
	// GetUsers returns a list of users with pagination
	GetUsers(ctx context.Context) ([]domain.User, error)
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, id uint) error
}
