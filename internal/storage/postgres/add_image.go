package postgres

import (
	"context"
	"tevian/internal/models"
)

func (s *Storage) AddImage(uuid string, img models.Image) (uint64, error) {
	var imageId uint64
	query := `
		INSERT INTO IMAGES(task_id, title)
		VALUES($1,$2)
		RETURNING id
	`
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	row := tx.QueryRow(query, &uuid, &img.Name)
	err = row.Scan(&imageId)
	if err != nil {
		return 0, err
	}
	query = `
		INSERT INTO FACES(bbox, gender, age, image_id)
		VALUES($1,$2,$3,$4)
	`
	for _, face := range img.Faces {
		_, err = tx.Exec(query, &face.Bbox, &face.Gender, &face.Age, &imageId)
		if err != nil {
			return 0, err
		}
	}
	tx.Commit()
	return imageId, nil
}
