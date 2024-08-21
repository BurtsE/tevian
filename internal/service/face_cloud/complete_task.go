package facecloud

import (
	"context"
	"encoding/json"
	"errors"
	"tevian/internal/converter"
	"tevian/internal/models"
	"time"

	"github.com/valyala/fasthttp"
)

func (s *service) StartTask(uuid string) error {
	status, err := s.storage.TaskStatus(uuid)
	if err != nil {
		return err
	}
	if status != models.Pending {
		return errors.New("task already started")
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

	resultChan := make(chan models.Image, len(images))
	jobsChan := make(chan models.Image, len(images))
	errChan := make(chan error)
	ctx, cancel := context.WithTimeout(context.WithValue(context.Background(), "uuid", uuid), time.Second*15)
	go s.producer(ctx, cancel, images, jobsChan, errChan)
	for range s.workersForTask {
		go s.worker(ctx, jobsChan, resultChan, errChan)
	}

	return nil
}

func (s *service) processImage(data []byte) (models.Image, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(s.url + "/api/v1/detect?demographics=true")
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Authorization", s.token)
	req.SetBody(data)
	req.PostArgs().Add("demographics", "true")
	req.PostArgs().Add("attributes", "true")

	resp := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, resp)
	if err != nil {
		return models.Image{}, err
	}
	result := models.FaceServiceTask{}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return models.Image{}, err
	}
	if result.StatusCode != 200 {
		return models.Image{}, errors.New(result.Message)
	}
	image := converter.ImageFromFaceApi(result)
	return image, nil
}
