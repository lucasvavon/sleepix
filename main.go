package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/db"
	"github.com/lucasvavon/slipx-api/services"
	"log"
)

type user struct {
	ID    string
	Name  string
	Email string
}

func main() {
	db, err := db.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	router.GET("/users", services.GetUsers)
	router.GET("/user/:id", services.GetUser)
	router.POST("/user", services.PostUser)
	router.PUT("/user/:id", services.UpdateUser)
	router.DELETE("/user/:id", services.DeleteUser)

	log.Fatal(router.Run(":10000"))
}
