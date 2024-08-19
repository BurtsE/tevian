package facecloud

import (
	"tevian/internal/models"
)

func (s *service) AddImageToTask(uuid, title string, img []byte) error {
	image := models.Image{
		Name: title,
	}
	imageId, err := s.storage.AddImage(uuid, image)
	if err != nil {
		return err
	}

	err = s.diskStorage.SaveImage(uuid, title, imageId, img)
	if err != nil {
		return err
	}
	return nil
}
