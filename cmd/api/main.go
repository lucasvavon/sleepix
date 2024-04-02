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

	userHandler := handlers.NewUserHandler(*app.userService)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	r.PUT("/users/:id", userHandler.UpdateUser)

	videoHandler := handlers.NewVideoHandler(*app.videoService)
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
