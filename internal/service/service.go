package service

import "tevian/internal/models"

type Service interface {
	CreateTask() (string, error)
	GetTask(string) (models.Task, error)
	StartTask(string) error
	DeleteTask(string) error
	AddImageToTask(string, string, []byte) error
}

type FaceCloud interface {
	Login()
}
