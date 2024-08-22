package facecloud

import (
	"context"
	"errors"
	"log"
	"sync"
	"tevian/internal/models"
)

func (s *service) initWorkers(ctx context.Context, uuid string, images []models.Image) {
	jobsChan := make(chan models.Image, len(images))
	wg := new(sync.WaitGroup)

	for range s.workersForTask {
		wg.Add(1)
		go s.worker(ctx, uuid, jobsChan, wg)
	}
	go func() {
		defer close(jobsChan)
		for _, image := range images {
			jobsChan <- image
		}
	}()
	wg.Wait()
	select {
	case <-ctx.Done():
		s.stopExecution(uuid, errors.New("task was cancelled"))
	default:
		s.completeExecution(uuid)
	}
}
func (s *service) worker(ctx context.Context, uuid string, jobs <-chan models.Image, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			s.logger.Printf("worker started processing image with id: %d", j.Id)

			processedImage, err := s.processImage(j.Data)
			if err != nil {
				s.cancelTask(uuid)
				return
			}
			processedImage.Id = j.Id

			err = s.storage.AddFaces(processedImage)
			if err != nil {
				s.cancelTask(uuid)
				return
			}
			s.logger.Printf("image with id %d processed", j.Id)
		}
	}
}

func (s *service) stopExecution(uuid string, err error) {
	errStatusUpd := s.storage.SetTaskStatus(uuid, models.Failed)
	if errStatusUpd != nil {
		s.logger.Printf("could not update task status! uuid: %s, err:%v", uuid, errStatusUpd)
		return
	}
	s.logger.Printf("task with id %s failed: %v", uuid, err)
}

func (s *service) completeExecution(uuid string) {
	errStatusUpd := s.storage.SetTaskStatus(uuid, models.Completed)
	if errStatusUpd != nil {
		s.logger.Printf("could not update task status! uuid: %s, err:%v", uuid, errStatusUpd)
		return
	}
	s.logger.Printf("task with id %s finished successfully", uuid)
}
