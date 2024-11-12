package facecloud

import (
	"context"
	"tevian/internal/models"

	"github.com/google/uuid"
)

func (s *service) Task(ctx context.Context, uuid string) (models.Task, error) {
	task := models.Task{}
	status, err := s.storage.TaskStatus(ctx, uuid)
	if err != nil {
		return task, err
	}
	images, err := s.diskStorage.Images(ctx, uuid)
	if err != nil {
		return task, err
	}
	for i := range images {
		images[i].Faces, err = s.storage.FacesByImage(ctx, images[i].Id)
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

func (s *service) CreateTask(ctx context.Context) (string, error) {
	task := models.Task{
		UUID:   uuid.NewString(),
		Status: models.Pending,
	}
	err := s.storage.CreateTask(ctx, task)
	if err != nil {
		return "", err
	}
	return task.UUID, nil
}

func (s *service) DeleteTask(ctx context.Context, uuid string) error {
	status, err := s.storage.TaskStatus(ctx, uuid)
	if err != nil {
		return err
	}
	if status == models.Processed {
		s.cancelTask(ctx, uuid)
	}
	task := models.Task{
		UUID:   uuid,
		Status: status,
	}
	err = s.storage.DeleteTask(ctx, task)
	if err != nil {
		return err
	}

	err = s.diskStorage.DeleteImages(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
