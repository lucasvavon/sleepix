package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/adapters/handlers"
	"github.com/lucasvavon/slipx-api/internal/adapters/repositories/mysql"
	"github.com/lucasvavon/slipx-api/internal/core/services"
	"log"
)

var (
	userHandler *handlers.UserHandler
	us          *services.UserService
)

func main() {

	db := mysql.InitDB()

	store := mysql.NewUserGORMRepository(db) // Assuming NewUserGORMRepository expects *gorm.DB
	if store == nil {
		log.Fatalf("Failed to create UserGORMRepository")
	}

	us = services.NewUserService(store)

	InitRoutes()

}

func InitRoutes() {
	r := gin.Default()
	handler := handlers.NewUserHandler(*us)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.CreateUser)
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
