package storage

import "tevian/internal/models"

type Storage interface {
	CreateTask(models.Task) error
	AddImage(string, models.Image) error
	FinishTask(models.Task) error
	DeleteTask(models.Task) error
}
