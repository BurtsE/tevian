package postgres

import (
	"database/sql"
	"fmt"
	"tevian/internal/config"
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
