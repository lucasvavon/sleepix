package services

import (
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"github.com/lucasvavon/slipx-api/internal/core/ports"
)

type VideoService struct {
	repo ports.VideoRepository
}

func NewVideoService(repo ports.VideoRepository) *VideoService {
	return &VideoService{repo: repo}
}

func (s *VideoService) GetVideos() ([]domain.Video, error) {
	return s.repo.GetVideos()
}

func (s *VideoService) GetVideo(id *int) (domain.Video, error) {
	return s.repo.GetVideo(id)
}

func (s *VideoService) CreateVideo(video *domain.Video) error {
	return s.repo.CreateVideo(video)
}

func (s *VideoService) UpdateVideo(id *int, video *domain.Video) error {
	return s.repo.UpdateVideo(id, video)
}

func (s *VideoService) DeleteVideo(id *int) error {
	return s.repo.DeleteVideo(id)
}
