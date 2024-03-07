package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/services"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.NewUserService().GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
