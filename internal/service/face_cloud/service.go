package face_cloud

import (
	"tevian/internal/config"
	"tevian/internal/models"
	def "tevian/internal/service"
)

var _ def.Service = (*Service)(nil)

type Service struct {
}

// AddImageToTask implements service.Service.
func (s *Service) AddImageToTask(uuid string, img []byte) error {
	panic("unimplemented")
}



// DeleteTask implements service.Service.
func (s *Service) DeleteTask(string) error {
	panic("unimplemented")
}

// GetTask implements service.Service.
func (s *Service) GetTask(string) (models.Task, error) {
	panic("unimplemented")
}

// StartTask implements service.Service.
func (s *Service) StartTask(string) error {
	panic("unimplemented")
}

func NewService(cfg config.Config) *Service {
	return &Service{}
}
