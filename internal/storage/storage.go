package storage

import (
	"context"
	"tevian/internal/models"
)

type Storage interface {
	CreateTask(context.Context, models.Task) error
	DeleteTask(context.Context, models.Task) error
	TaskStatus(context.Context, string) (models.TaskStatus, error)
	SetTaskStatus(context.Context, string, models.TaskStatus) error

	AddImage(context.Context, string, string) (uint64, error)
	AddFaces(context.Context, models.Image) error
	FacesByImage(context.Context, int64) ([]models.Face, error)
}
type DiskStorage interface {
	SaveImage(context.Context, string, string, uint64, []byte) error
	DeleteImages(context.Context, string) error
	Images(context.Context, string) ([]models.Image, error)
}
