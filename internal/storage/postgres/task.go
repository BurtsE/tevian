package postgres

import (
	"context"
	"tevian/internal/converter"
	"tevian/internal/models"
)

func (s *Storage) CreateTask(ctx context.Context, task models.Task) error {
	query := `
		INSERT INTO tasks(uuid, progress)
		VALUES($1,$2)
	`
	_, err := s.db.Exec(query, task.UUID, task.Status.String())
	return err
}

func (s *Storage) DeleteTask(ctx context.Context, task models.Task) error {
	query := `
		DELETE FROM tasks
		WHERE uuid = $1 
	`
	_, err := s.db.Exec(query, &task.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) TaskStatus(ctx context.Context, uuid string) (models.TaskStatus, error) {
	var status string
	query := `
		SELECT progress
		FROM tasks
		where uuid = $1
	`
	row := s.db.QueryRow(query, &uuid)
	err := row.Scan(&status)
	if err != nil {
		return nil, err
	}

	taskStatus, err := converter.TaskStatusFromString(status)
	if err != nil {
		return nil, err
	}
	return taskStatus, nil
}

func (s *Storage) SetTaskStatus(ctx context.Context, uuid string, status models.TaskStatus) error {
	query := `
		UPDATE tasks
		SET progress = $2
		where uuid = $1
	`
	statusStr := status.String()
	_, err := s.db.Exec(query, &uuid, &statusStr)
	if err != nil {
		return err
	}
	return nil
}
