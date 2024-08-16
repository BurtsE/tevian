package app

import (
	"log"
	"tevian/internal/config"
	"tevian/internal/service"
	facecloud "tevian/internal/service/face_cloud"
	"tevian/internal/storage"
	"tevian/internal/storage/postgres"

	"tevian/internal/api"

	"github.com/sirupsen/logrus"
)

type serviceProvider struct {
	cfg      *config.Config
	postgres storage.Storage
	service  service.Service
	router   *api.Router
}

func NewSericeProvider() *serviceProvider {
	s := &serviceProvider{}
	s.Router()
	return s
}
func (s *serviceProvider) Config() *config.Config {
	if s.cfg == nil {
		cfg, err := config.InitConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.cfg = &cfg
	}
	return s.cfg
}
func (s *serviceProvider) Postgres() storage.Storage {
	if s.postgres == nil {
		storage, err := postgres.NewStorage(s.Config())
		if err != nil {
			log.Fatal(err)
		}
		s.postgres = storage
	}
	return s.postgres
}
func (s *serviceProvider) Service() service.Service {
	if s.service == nil {
		s.service = facecloud.NewService(s.Postgres(), s.Config())
	}
	return s.service
}

func (s *serviceProvider) Router() *api.Router {
	if s.router == nil {
		s.router = api.NewRouter(s.Config(), s.Service(), logrus.New())
	}
	return s.router
}
