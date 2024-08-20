package facecloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"tevian/internal/config"
	"tevian/internal/models"
	def "tevian/internal/service"
	"tevian/internal/storage"

	"github.com/valyala/fasthttp"
)

var _ def.Service = (*service)(nil)

type service struct {
	storage         storage.Storage
	diskStorage     storage.DiskStorage
	url             string
	email, password string
	token           string
}

func NewService(storage storage.Storage, cfg *config.Config, diskStorage storage.DiskStorage) *service {
	s := &service{
		storage:     storage,
		diskStorage: diskStorage,
		url:         cfg.FaceCloud.Url,
		email:       cfg.FaceCloud.Email,
		password:    cfg.FaceCloud.Password,
	}
	return s
}

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
	s.token = "Bearer " + data.Data.Token
	return nil
}


// GetTask implements service.Service.
func (s *service) GetTask(string) (models.Task, error) {
	panic("unimplemented")
}
