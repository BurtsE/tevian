package postgres

import (
	"errors"
	"tevian/internal/converter"
	"tevian/internal/models"
)

func (s *Storage) DeleteTask(task models.Task) error {
	var status string
	query := `
		SELECT progress
		FROM tasks
		where uuid = $1
	`
	row := s.db.QueryRow(query)
	err := row.Scan(&status)
	if err != nil {
		return err
	}

	task.Status, err = converter.TaskStatusFromString(status)
	if err != nil {
		return err
	}
	if task.Status == models.Pending {
		return errors.New("permission denied")
	}

	query = `
		DELETE FROM tasks
		WHERE uuid = $1 
	`
	_, err = s.db.Exec(query, &task.UUID)
	if err != nil {
		return err
	}
	return nil
}
