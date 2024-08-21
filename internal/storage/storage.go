package storage

import "tevian/internal/models"

type Storage interface {
	CreateTask(models.Task) error
	DeleteTask(models.Task) error
	TaskStatus(string) (models.TaskStatus, error)
	SetTaskStatus(string, models.TaskStatus) error

	AddImage(string, string) (uint64, error)
	AddFaces(models.Image) error
	FacesByImage(int64) ([]models.Face, error)
}
type DiskStorage interface {
	SaveImage(string, string, uint64, []byte) error
	DeleteImages(string) error
	Images(string) ([]models.Image, error)
}
