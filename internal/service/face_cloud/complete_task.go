package facecloud

import (
	"context"
	"errors"
	"tevian/internal/models"
	"time"
)

func (s *service) StartTask(uuid string) error {
	var filter bool
	status, err := s.storage.TaskStatus(uuid)
	if err != nil {
		return err
	}
	switch status {
	case models.Processed:
		return errors.New("task already started")
	case models.Completed:
		return nil
	case models.Failed:
		filter = true
	}

	err = s.storage.SetTaskStatus(uuid, models.Processed)
	if err != nil {
		return err
	}

	err = s.login()
	if err != nil {
		return err
	}

	images, err := s.diskStorage.Images(uuid)
	if err != nil {
		return err
	}
	if filter {
		images, err = s.filterProcessedImages(images)
		if err != nil {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	s.storeFunc(uuid, cancel)
	go s.initWorkers(ctx, uuid, images)
	return nil
}

func (s *service) cancelTask(uuid string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	cancel, ok := s.cancelTasks[uuid]
	if ok {
		cancel()
		delete(s.cancelTasks, uuid)
	}
}

func (s *service) filterProcessedImages(images []models.Image) ([]models.Image, error) {
	filteredImages := make([]models.Image, 0)
	for _, image := range images {
		faces, err := s.storage.FacesByImage(image.Id)
		if err != nil {
			return nil, err
		}
		if len(faces) != 0 {
			continue
		}
		filteredImages = append(filteredImages, image)
	}
	return filteredImages, nil
}

func (s *service) storeFunc(uuid string, cancel context.CancelFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cancelTasks[uuid] = cancel
}
