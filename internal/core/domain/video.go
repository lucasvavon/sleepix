package domain

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID     uint `gorm:"primary_key"`
	UserId uint
	Title  string
	Link   string
	Moment string
}
