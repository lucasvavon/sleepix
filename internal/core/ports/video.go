package ports

import (
	"github.com/lucasvavon/slipx-api/internal/core/domain"
)

// VideoRepository is an interface for interacting with user-related data
type VideoRepository interface {
	GetVideos() ([]domain.Video, error)
	GetVideo(id *int) (domain.Video, error)
	CreateVideo(user *domain.Video) error
	UpdateVideo(id *int, user *domain.Video) error
	DeleteVideo(id *int) error
}
