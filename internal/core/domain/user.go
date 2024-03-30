package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	Lastname  string
	Firstname string
	Email     string
	Password  string
}
