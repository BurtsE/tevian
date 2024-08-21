package postgres

import (
	"log"
	"tevian/internal/converter"
	"tevian/internal/models"
)

func (s *Storage) CreateTask(task models.Task) error {
	query := `
		INSERT INTO tasks(uuid, progress)
		VALUES($1,$2)
	`
	_, err := s.db.Exec(query, task.UUID, task.Status.String())
	return err
}

func (s *Storage) DeleteTask(task models.Task) error {
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

func (s *Storage) TaskStatus(uuid string) (models.TaskStatus, error) {
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

func (s *Storage) SetTaskStatus(uuid string, status models.TaskStatus) error {
	query := `
		UPDATE tasks
		SET progress = $2
		where uuid = $1
	`
	statusStr := status.String()
	log.Println(uuid, statusStr)
	_, err := s.db.Exec(query, &uuid, &statusStr)
	if err != nil {
		return err
	}
	return nil
}
