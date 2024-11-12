package facecloud

import (
	"context"
	"errors"
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
			select {
			case <-ctx.Done():
				return
			case jobsChan <- image:
			}
		}
	}()
	wg.Wait()
	select {
	case <-ctx.Done():
		s.stopExecution(ctx, uuid, errors.New("task was cancelled"))
	default:
		s.completeExecution(ctx, uuid)
	}
}
func (s *service) worker(ctx context.Context, uuid string, jobs <-chan models.Image, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		s.logger.Printf("worker started processing image with id: %d", j.Id)

		processedImage, err := s.processImage(ctx, j.Data)
		if err != nil {
			s.cancelTask(uuid)
			return
		}
		processedImage.Id = j.Id

		err = s.storage.AddFaces(ctx, processedImage)
		if err != nil {
			s.cancelTask(uuid)
			return
		}
		s.logger.Printf("image with id %d processed", j.Id)
	}
}

func (s *service) stopExecution(ctx context.Context, uuid string, err error) {
	errStatusUpd := s.storage.SetTaskStatus(ctx, uuid, models.Failed)
	if errStatusUpd != nil {
		s.logger.Printf("could not update task status! uuid: %s, err:%v", uuid, errStatusUpd)
		return
	}
	s.logger.Printf("task with id %s failed: %v", uuid, err)
}

func (s *service) completeExecution(ctx context.Context, uuid string) {
	errStatusUpd := s.storage.SetTaskStatus(ctx, uuid, models.Completed)
	if errStatusUpd != nil {
		s.logger.Printf("could not update task status! uuid: %s, err:%v", uuid, errStatusUpd)
		return
	}
	s.logger.Printf("task with id %s finished successfully", uuid)
}
