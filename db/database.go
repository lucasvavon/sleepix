package db

import (
	"github.com/lucasvavon/slipx-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Database() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=admin dbname=slipx port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Println(err)
	}

	return db, err

}
