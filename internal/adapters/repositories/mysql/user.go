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

func (r *UserGORMRepository) GetUser(id *int) (domain.User, error) {
	var user domain.User
	req := r.db.First(&user, id)

	if req.Error != nil {
		// Use fmt.Errorf for error formatting and return the zero value of domain.User.
		return domain.User{}, fmt.Errorf("user not found: %v", req.Error)
	}
	return user, nil
}

func (r *UserGORMRepository) CreateUser(user *domain.User) error {
	req := r.db.Create(&user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *UserGORMRepository) UpdateUser(id *int, user *domain.User) error {

	user.ID = id

	req := r.db.Save(user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *UserGORMRepository) DeleteUser(id *int) error {
	var user domain.User

	req := r.db.Delete(&user, &id)

	if req.Error != nil {
		return req.Error
	}

	return nil
}
