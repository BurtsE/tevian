package facecloud

import (
	"context"
	"sync"
	"tevian/internal/config"
	def "tevian/internal/service"
	"tevian/internal/storage"

	"github.com/sirupsen/logrus"
)

var _ def.Service = (*service)(nil)

type service struct {
	storage         storage.Storage
	diskStorage     storage.DiskStorage
	url             string
	email, password string
	faceCloudToken  string
	workersForTask  int
	cancelTasks     map[string]context.CancelFunc
	mu              sync.Mutex
	logger          *logrus.Logger
}

func NewService(storage storage.Storage, cfg *config.Config, diskStorage storage.DiskStorage, logger *logrus.Logger) *service {
	s := &service{
		storage:        storage,
		diskStorage:    diskStorage,
		url:            cfg.FaceCloud.Url,
		email:          cfg.FaceCloud.Email,
		password:       cfg.FaceCloud.Password,
		logger:         logger,
		workersForTask: 4,
		cancelTasks:    make(map[string]context.CancelFunc),
	}
	return s
}
