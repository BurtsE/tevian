package storage

import "tevian/internal/models"

type Storage interface {
	CreateTask(models.Task) error
	AddImage(string, models.Image) error
	FinishTask(models.Task) error
	DeleteTask(models.Task) error
}
type DiskStorage interface {
	SaveImage(string, string, []byte) error
	// DeleteImages(string) error
}
