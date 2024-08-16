package postgres

import (
	"database/sql"
	"fmt"
	"tevian/internal/config"
	"tevian/internal/models"
	def "tevian/internal/storage"

	_ "github.com/lib/pq"
)

var _ def.Storage = (*Storage)(nil)

type Storage struct {
	db *sql.DB
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	DSN := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		cfg.Postgres.DB,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Sslmode,
	)
	db, _ := sql.Open("postgres", DSN)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Storage{
		db: db,
	}, nil
}

// AddImage implements storage.Storage.
func (s *Storage) AddImage(string, models.Image) error {
	panic("unimplemented")
}

// CreateTask implements storage.Storage.
func (s *Storage) CreateTask(task models.Task) error {
	query := `
		INSERT INTO tasks(uuid, progress)
		VALUES($1,$2)
	`
	_, err := s.db.Exec(query, task.UUID, task.Status.String())
	return err
}

// DeleteTask implements storage.Storage.
func (s *Storage) DeleteTask(models.Task) error {
	panic("unimplemented")
}

// FinishTask implements storage.Storage.
func (s *Storage) FinishTask(models.Task) error {
	panic("unimplemented")
}
