package facecloud

import (
	"encoding/json"
	"errors"
	"log"
	"tevian/internal/converter"
	"tevian/internal/models"

	"github.com/valyala/fasthttp"
)

func (s *service) StartTask(uuid string) error {
	err := s.login()
	if err != nil {
		return err
	}
	images, err := s.diskStorage.GetImages(uuid)
	if err != nil {
		return err
	}
	for _, image := range images {
		processedImage, err := s.processImage(image.Data)
		if err != nil {
			return err
		}
		processedImage.Id = image.Id
		err = s.storage.AddFaces(processedImage)
		if err != nil {
			return err
		}
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

	log.Println(req.PostArgs())

	resp := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, resp)
	if err != nil {
		return models.Image{}, err
	}
	result := models.FaceServiceTask{}
	log.Println(string(resp.Body()))

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
