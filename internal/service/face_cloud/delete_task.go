package facecloud

import "tevian/internal/models"

func (s *service) DeleteTask(uuid string) error {
	task := models.Task{
		UUID: uuid,
	}
	err := s.storage.DeleteTask(task)
	if err != nil {
		return err
	}
	
	err = s.diskStorage.DeleteImages(uuid)
	if err != nil {
		return err
	}
	return nil
}
