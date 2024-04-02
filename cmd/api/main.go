package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/adapters/handlers"
	"github.com/lucasvavon/slipx-api/internal/adapters/repositories/mysql"
	"github.com/lucasvavon/slipx-api/internal/core/services"
	"log"
)

var (
	us *services.UserService
	vs *services.VideoService
)

func main() {

	db := mysql.InitDB()

	userStore := mysql.NewUserGORMRepository(db)   // Assuming NewUserGORMRepository expects *gorm.DB
	videoStore := mysql.NewVideoGORMRepository(db) // Assuming NewUserGORMRepository expects *gorm.DB
	if userStore == nil || videoStore == nil {
		log.Fatalf("Failed to create store")
	}

	us = services.NewUserService(userStore)
	vs = services.NewVideoService(videoStore)

	InitRoutes()

}

func InitRoutes() {
	r := gin.Default()
	userHandler := handlers.NewUserHandler(*us)
	videoHandler := handlers.NewVideoHandler(*vs)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	r.PUT("/users/:id", userHandler.UpdateUser)

	r.GET("/videos", videoHandler.GetVideos)
	r.GET("/videos/:id", videoHandler.GetVideo)
	r.POST("/videos", videoHandler.CreateVideo)
	r.DELETE("/videos/:id", videoHandler.DeleteVideo)
	r.PUT("/videos/:id", videoHandler.UpdateVideo)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
