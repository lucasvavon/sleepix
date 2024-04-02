package mysql

import (
	"errors"
	"fmt"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"gorm.io/gorm"
)

type VideoGORMRepository struct {
	db *gorm.DB
}

func NewVideoGORMRepository(db *gorm.DB) *VideoGORMRepository {
	return &VideoGORMRepository{db: db}
}

func (r *VideoGORMRepository) GetVideos() ([]domain.Video, error) {
	var videos []domain.Video
	req := r.db.Find(&videos)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}
	return videos, nil
}

func (r *VideoGORMRepository) GetVideo(id *int) (domain.Video, error) {
	var video domain.Video
	req := r.db.First(&video, id)

	if req.Error != nil {
		// Use fmt.Errorf for error formatting and return the zero value of domain.Video.
		return domain.Video{}, fmt.Errorf("video not found: %v", req.Error)
	}
	return video, nil
}

func (r *VideoGORMRepository) CreateVideo(video *domain.Video) error {
	req := r.db.Create(&video)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *VideoGORMRepository) UpdateVideo(id *int, video *domain.Video) error {

	video.ID = id

	req := r.db.Save(video)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *VideoGORMRepository) DeleteVideo(id *int) error {
	var video domain.Video

	req := r.db.Delete(&video, &id)

	if req.Error != nil {
		return req.Error
	}

	return nil
}
