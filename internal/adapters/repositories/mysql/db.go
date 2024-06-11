package mysql

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"regexp"
)

type DB struct {
	Db *gorm.DB
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*sleepix)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func InitDB() *gorm.DB {

	loadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database! %s\n", dsn)
	}

	fmt.Println("Database connection successfully established")

	var models = []interface{}{&domain.User{}, &domain.Video{}}

	db.AutoMigrate(models...)

	return db
}
