package postgres

import (
	"context"
	"tevian/internal/models"
)

func (s *Storage) AddFaces(img models.Image) error {
	query := `
		INSERT INTO FACES(width, height, x, y, gender, age, image_id)
		VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, face := range img.Faces {
		_, err := tx.Exec(query, &face.Width, &face.Height, &face.X, &face.Y, &face.Gender, &face.Age, &img.Id)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}
