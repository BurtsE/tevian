package app

import (
	"log"
	"tevian/internal/config"
	"tevian/internal/service"
	facecloud "tevian/internal/service/face_cloud"
	"tevian/internal/storage"
	"tevian/internal/storage/disk"
	"tevian/internal/storage/postgres"
	"tevian/internal/telegram"

	"tevian/internal/api"

	"github.com/sirupsen/logrus"
)

type serviceProvider struct {
	cfg         *config.Config
	postgres    storage.Storage
	service     service.Service
	router      *api.Router
	diskStorage storage.DiskStorage
	logger      *logrus.Logger
	bot         *telegram.TelegramBot
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
func (s *serviceProvider) Logger() *logrus.Logger {
	if s.logger == nil {
		s.logger = logrus.New()
	}
	return s.logger
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
		s.service = facecloud.NewService(s.Postgres(), s.Config(), s.DiskStorage(), s.Logger())
	}
	return s.service
}
func (s *serviceProvider) DiskStorage() storage.DiskStorage {
	if s.diskStorage == nil {
		s.diskStorage = disk.NewStorage()
	}
	return s.diskStorage
}

func (s *serviceProvider) Router() *api.Router {
	if s.router == nil {
		s.router = api.NewRouter(s.Config(), s.Service(), logrus.New())
	}
	return s.router
}

func (s *serviceProvider) TelegramBot() *telegram.TelegramBot {
	if s.bot == nil {
		bot, err := telegram.NewTGBot(s.Config())
		if err != nil {
			log.Fatal(err)
		}
		s.bot = bot
	}
	return s.bot
}
