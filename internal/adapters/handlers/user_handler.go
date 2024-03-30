package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/core/services"
	"net/http"
)

type UserHandler struct {
	us services.UserService
}

func NewUserHandler(UserService services.UserService) *UserHandler {
	return &UserHandler{
		us: UserService,
	}
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.us.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
