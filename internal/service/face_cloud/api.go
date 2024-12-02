package facecloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"tevian/internal/converter"
	"tevian/internal/models"

	"github.com/valyala/fasthttp"
)

func (s *service) login() error {
	loginJSON := fmt.Sprintf(`{
		"email": "%s",
		"password": "%s"
  	}`, s.email, s.password)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(s.url + "/api/v1/login")
	req.Header.SetMethod("POST")
	req.SetBody([]byte(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}
	data := models.FaceServiceLogin{}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return err
	}
	if data.StatusCode != 200 {
		return errors.New(data.Message)
	}
	s.faceCloudToken = "Bearer " + data.Data.Token
	return nil
}

func (s *service) processImage(data []byte) (models.Image, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(s.url + "/api/v1/detect?demographics=true")
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Authorization", s.faceCloudToken)
	req.SetBody(data)
	req.PostArgs().Add("demographics", "true")
	req.PostArgs().Add("attributes", "true")

	resp := fasthttp.AcquireResponse()

	err := fasthttp.Do(req, resp)
	if err != nil {
		return models.Image{}, err
	}
	if resp.StatusCode() != 200 {
		return models.Image{}, errors.New(resp.String())
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
