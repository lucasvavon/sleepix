package domain

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID     *int `gorm:"primary_key"`
	UserId *int
	Title  string
	Link   string
	Moment *int
}
