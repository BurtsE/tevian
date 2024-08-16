package facecloud

import (
	"tevian/internal/config"
	"tevian/internal/models"
	def "tevian/internal/service"
	"tevian/internal/storage"
)

var _ def.Service = (*service)(nil)

type service struct {
	storage storage.Storage
}

// AddImageToTask implements service.Service.
func (s *service) AddImageToTask(string, []byte) error {
	panic("unimplemented")
}

// DeleteTask implements service.Service.
func (s *service) DeleteTask(string) error {
	panic("unimplemented")
}

// GetTask implements service.Service.
func (s *service) GetTask(string) (models.Task, error) {
	panic("unimplemented")
}

// StartTask implements service.Service.
func (s *service) StartTask(string) error {
	panic("unimplemented")
}

func NewService(storage storage.Storage, cfg *config.Config) *service {
	return &service{
		storage: storage,
	}
}
