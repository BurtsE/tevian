package facecloud

import (
	"tevian/internal/config"
	"tevian/internal/models"
	def "tevian/internal/service"
	"tevian/internal/storage"
)

var _ def.Service = (*service)(nil)

type service struct {
	storage     storage.Storage
	diskStorage storage.DiskStorage
}


// GetTask implements service.Service.
func (s *service) GetTask(string) (models.Task, error) {
	panic("unimplemented")
}

// StartTask implements service.Service.
func (s *service) StartTask(string) error {
	panic("unimplemented")
}

func NewService(storage storage.Storage, cfg *config.Config, diskStorage storage.DiskStorage) *service {
	return &service{
		storage:     storage,
		diskStorage: diskStorage,
	}
}
