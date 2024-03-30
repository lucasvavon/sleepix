package services

import (
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"github.com/lucasvavon/slipx-api/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() ([]domain.User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) GetUser(id *int) (domain.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) CreateUser(user *domain.User) (domain.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id *int) error {
	return s.repo.DeleteUser(id)
}
