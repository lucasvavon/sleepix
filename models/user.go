package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Name     string
	Username string
	Email    string
	Password string
}
