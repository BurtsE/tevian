package storage

import "tevian/internal/models"

type Storage interface {
	CreateTask(models.Task) error
	AddImage(string, string) (uint64, error)
	AddFaces(models.Image) error
	FinishTask(models.Task) error
	DeleteTask(models.Task) error
}
type DiskStorage interface {
	SaveImage(string, string, uint64, []byte) error
	DeleteImages(string) error
	GetImages(string) ([]models.Image, error)
}
