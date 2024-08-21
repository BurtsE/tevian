package facecloud

import (
	"context"
	"log"
	"tevian/internal/models"
)

func (s *service) worker(ctx context.Context, jobs <-chan models.Image, results chan<- models.Image, errChan chan error) {
	for j := range jobs {
		log.Printf("worker started processing image with id: %d", j.Id)
		processedImage, err := s.processImage(j.Data)
		if err != nil {
			errChan <- err
		}
		processedImage.Id = j.Id

		err = s.storage.AddFaces(processedImage)
		if err != nil {
			errChan <- err
		}
		log.Printf("image with id %d processed", j.Id)
		select {
		case <-ctx.Done():
			return
		case results <- processedImage:
		}
	}
}
func (s *service) producer(ctx context.Context, cancel context.CancelFunc,
	images []models.Image, jobs chan<- models.Image, errChan chan error) {
	var (
		err  error
		ok   bool
		uuid string
	)
	defer func() {
		close(jobs)
		if uuid, ok = ctx.Value("uuid").(string); !ok {
			panic("task uuid was not provided to producer")
		}
		if err != nil {
			errStatusUpd := s.storage.SetTaskStatus(uuid, models.Failed)
			if errStatusUpd != nil {
				log.Printf("could not update task status! uuid: %s, err:%v", uuid, errStatusUpd)
				return
			}
			log.Printf("task with id %s failed: %v", uuid, err)
			return
		}

		errStatusUpd := s.storage.SetTaskStatus(uuid, models.Completed)
		if errStatusUpd != nil {
			log.Printf("could not update task status! uuid: %s, err:%v", uuid, errStatusUpd)
			return
		} else {
			log.Printf("task with id %s finished successfully", uuid)
		}
	}()

	for _, image := range images {
		select {
		case err = <-errChan:
			log.Println(err)
			cancel()
			return
		case jobs <- image:
		}
	}
	select {
	case <-ctx.Done():
	case err = <-errChan:
		log.Println(err)
		cancel()
	}
}
