package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"github.com/lucasvavon/slipx-api/internal/core/services"
	"net/http"
	"strconv"
)

type VideoHandler struct {
	us services.VideoService
}

func NewVideoHandler(VideoService services.VideoService) *VideoHandler {
	return &VideoHandler{
		us: VideoService,
	}
}

func (h *VideoHandler) GetVideos(ctx *gin.Context) {
	videos, err := h.us.GetVideos()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, videos)
}

func (h *VideoHandler) GetVideo(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	video, err := h.us.GetVideo(&id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, video)
}

func (h *VideoHandler) CreateVideo(ctx *gin.Context) {
	var video domain.Video

	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.us.CreateVideo(&video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, video)
}

func (h *VideoHandler) DeleteVideo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := h.us.DeleteVideo(&id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Video %d has been deleted", id))
}

func (h *VideoHandler) UpdateVideo(ctx *gin.Context) {
	var updateVideo domain.Video
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Bind the JSON body to the updateVideo variable.
	if err := ctx.ShouldBindJSON(&updateVideo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the VideoService to update the video.
	err := h.us.UpdateVideo(&id, &updateVideo)
	if err != nil {
		// Handle errors, e.g., video not found or validation errors.
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success.
	ctx.JSON(http.StatusOK, gin.H{"message": "Video updated successfully"})
}
