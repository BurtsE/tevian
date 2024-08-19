package facecloud

import (
	"tevian/internal/models"

	"github.com/google/uuid"
)

func (s *service) CreateTask() (string, error) {
	task := models.Task{
		UUID:   uuid.NewString(),
		Status: models.Processed,
	}
	err := s.storage.CreateTask(task)
	if err != nil {
		return "", err
	}
	return task.UUID, nil
}
