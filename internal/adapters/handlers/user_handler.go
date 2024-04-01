package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"github.com/lucasvavon/slipx-api/internal/core/services"
	"net/http"
	"strconv"
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

func (h *UserHandler) GetUser(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	user, err := h.us.GetUser(&id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.us.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := h.us.DeleteUser(&id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("User %d has been deleted", id))
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var updateUser domain.User
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Bind the JSON body to the updateUser variable.
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the UserService to update the user.
	err := h.us.UpdateUser(&id, &updateUser)
	if err != nil {
		// Handle errors, e.g., user not found or validation errors.
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success.
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
