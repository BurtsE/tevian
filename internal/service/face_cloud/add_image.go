package facecloud

import (
	"errors"
	"tevian/internal/models"
)

func (s *service) AddImageToTask(uuid, title string, img []byte) error {
	status, err := s.storage.TaskStatus(uuid)
	if err != nil {
		return err
	}
	if status != models.Pending {
		return errors.New("task changes unavailable")
	}

	imageId, err := s.storage.AddImage(uuid, title)
	if err != nil {
		return err
	}

	err = s.diskStorage.SaveImage(uuid, title, imageId, img)
	if err != nil {
		return err
	}
	return nil
}
