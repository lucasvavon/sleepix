package mysql

import (
	"errors"
	"fmt"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"gorm.io/gorm"
)

type UserGORMRepository struct {
	db *gorm.DB
}

func NewUserGORMRepository(db *gorm.DB) *UserGORMRepository {
	return &UserGORMRepository{db: db}
}

func (r *UserGORMRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	req := r.db.Find(&users)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}
	return users, nil
}
