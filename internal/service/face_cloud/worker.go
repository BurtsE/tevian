package facecloud

import (
	"context"
	"log"
	"tevian/internal/models"
)

func (s *service) worker(ctx context.Context, jobs <-chan models.Image, results chan<- models.Image, errChan chan error) {
	for j := range jobs {
		log.Println(j)
		processedImage, err := s.processImage(j.Data)
		if err != nil {
			errChan <- err
		}
		processedImage.Id = j.Id
		err = s.storage.AddFaces(processedImage)
		if err != nil {
			errChan <- err
		}

		select {
		case results <- processedImage:
		case <-ctx.Done():
			return
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
		if err == nil {
			err = s.storage.SetTaskStatus(uuid, models.Completed)
		} else {
			err = s.storage.SetTaskStatus(uuid, models.Failed)
		}
		if err != nil {
			log.Printf("could not update task status! uuid: %s, err:%v", uuid, err)
		}
		log.Printf("task with id %s finished successfully", uuid)

	}()
	for _, image := range images {
		select {
		case err = <-errChan:
			log.Println(err)
			cancel()
			return
		default:
			jobs <- image
		}
	}
	<-ctx.Done()
}
