package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/adapters/handlers"
	"github.com/lucasvavon/slipx-api/internal/adapters/repositories/mysql"
	"github.com/lucasvavon/slipx-api/internal/core/services"
)

type App struct {
	userService  *services.UserService
	videoService *services.VideoService
}

func main() {
	// Initialisation de la DB et des services
	db := mysql.InitDB()
	userStore := mysql.NewUserGORMRepository(db)
	videoStore := mysql.NewVideoGORMRepository(db)

	app := &App{
		userService:  services.NewUserService(userStore),
		videoService: services.NewVideoService(videoStore),
	}

	app.InitRoutes()
}

func (app *App) InitRoutes() {
	r := gin.Default()

	api := r.Group("/api")

	userHandler := handlers.NewUserHandler(*app.userService)
	api.GET("/users", userHandler.GetUsers)
	api.GET("/users/:id", userHandler.GetUser)
	api.POST("/users", userHandler.CreateUser)
	api.DELETE("/users/:id", userHandler.DeleteUser)
	api.PUT("/users/:id", userHandler.UpdateUser)

	videoHandler := handlers.NewVideoHandler(*app.videoService)
	api.GET("/videos", videoHandler.GetVideos)
	api.GET("/videos/:id", videoHandler.GetVideo)
	api.POST("/videos", videoHandler.CreateVideo)
	api.DELETE("/videos/:id", videoHandler.DeleteVideo)
	api.PUT("/videos/:id", videoHandler.UpdateVideo)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
