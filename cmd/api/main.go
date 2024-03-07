package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/adapters/postgres"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"github.com/lucasvavon/slipx-api/internal/repositories"
	"github.com/lucasvavon/slipx-api/internal/services"
	"log"
	"net/http"
)

func main() {
	database, err := postgres.Database()
	if err != nil {
		log.Println(err)
	}
	database.DB()

	router := gin.Default()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)

	router.GET("/users", userService.GetUsers)
	router.GET("/user/:id", GetUser)
	router.POST("/user", PostUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)

	log.Fatal(router.Run(":10000"))
}

type NewUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetUser(c *gin.Context) {

	var user domain.User

	database, err := postgres.Database()
	if err != nil {
		log.Println(err)
	}

	if err := database.Where("id= ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func PostUser(c *gin.Context) {
	var input NewUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := domain.User{Name: input.Name, Username: input.Username, Email: input.Email, Password: input.Password}

	db, err := postgres.Database()
	if err != nil {
		log.Println(err)
	}

	if err := newUser.HashPassword(input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c *gin.Context) {

	var user domain.User

	db, err := postgres.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	var updateUser UserUpdate

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&user).Updates(domain.User{Name: user.Name, Username: user.Username, Email: user.Email, Password: user.Password}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {

	var user domain.User

	db, err := postgres.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}
