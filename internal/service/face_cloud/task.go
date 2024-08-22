package facecloud

import (
	"tevian/internal/models"

	"github.com/google/uuid"
)

func (s *service) Task(uuid string) (models.Task, error) {
	task := models.Task{}
	status, err := s.storage.TaskStatus(uuid)
	if err != nil {
		return task, err
	}
	images, err := s.diskStorage.Images(uuid)
	if err != nil {
		return task, err
	}
	for i := range images {
		images[i].Faces, err = s.storage.FacesByImage(images[i].Id)
		if err != nil {
			return task, err
		}
	}
	task.UUID = uuid
	task.Status = status
	task.Images = images
	task.CalcStats()
	return task, nil
}

func (s *service) CreateTask() (string, error) {
	task := models.Task{
		UUID:   uuid.NewString(),
		Status: models.Pending,
	}
	err := s.storage.CreateTask(task)
	if err != nil {
		return "", err
	}
	return task.UUID, nil
}

func (s *service) DeleteTask(uuid string) error {
	status, err := s.storage.TaskStatus(uuid)
	if err != nil {
		return err
	}
	if status == models.Processed {
		s.cancelTask(uuid)
	}
	task := models.Task{
		UUID:   uuid,
		Status: status,
	}
	err = s.storage.DeleteTask(task)
	if err != nil {
		return err
	}

	err = s.diskStorage.DeleteImages(uuid)
	if err != nil {
		return err
	}
	return nil
}
