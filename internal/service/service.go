package service

import (
	"context"
	"tevian/internal/models"
)

type Service interface {
	CreateTask(context.Context) (string, error)
	Task(context.Context, string) (models.Task, error)
	StartTask(context.Context, string) error
	DeleteTask(context.Context, string) error
	AddImageToTask(context.Context, string, string, []byte) error
}
