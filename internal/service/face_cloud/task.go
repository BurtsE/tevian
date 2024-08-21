package facecloud

import (
	"errors"
	"tevian/internal/models"
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

func (s *service) DeleteTask(uuid string) error {
	status, err := s.storage.TaskStatus(uuid)
	if err != nil {
		return err
	}
	if status == models.Processed {
		return errors.New("task is being processed")
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
