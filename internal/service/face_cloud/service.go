package facecloud

import (
	"sync"
	"tevian/internal/config"
	def "tevian/internal/service"
	"tevian/internal/storage"
)

var _ def.Service = (*service)(nil)

type service struct {
	storage         storage.Storage
	diskStorage     storage.DiskStorage
	url             string
	email, password string
	faceCloudToken  string
	workersForTask  int
	cancelTasks     sync.Map
}

func NewService(storage storage.Storage, cfg *config.Config, diskStorage storage.DiskStorage) *service {
	s := &service{
		storage:        storage,
		diskStorage:    diskStorage,
		url:            cfg.FaceCloud.Url,
		email:          cfg.FaceCloud.Email,
		password:       cfg.FaceCloud.Password,
		workersForTask: 4,
		cancelTasks:    sync.Map{},
	}
	return s
}
