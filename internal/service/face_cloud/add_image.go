package facecloud

import (
	"context"
	"errors"
	"tevian/internal/models"
)

func (s *service) AddImageToTask(ctx context.Context, uuid, title string, img []byte) error {
	status, err := s.storage.TaskStatus(ctx, uuid)
	if err != nil {
		return err
	}
	if status != models.Pending {
		return errors.New("task changes unavailable")
	}

	imageId, err := s.storage.AddImage(ctx, uuid, title)
	if err != nil {
		return err
	}

	err = s.diskStorage.SaveImage(ctx, uuid, title, imageId, img)
	if err != nil {
		return err
	}
	return nil
}
