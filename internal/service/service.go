package service

import "tevian/internal/models"

type Service interface {
	CreateTask() (string, error)
	Task(string) (models.Task, error)
	StartTask(string) error
	DeleteTask(string) error
	AddImageToTask(string, string, []byte) error
}
