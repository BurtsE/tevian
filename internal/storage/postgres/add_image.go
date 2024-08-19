package postgres

import "tevian/internal/models"

func (s *Storage) AddImage(uuid string, img models.Image) (uint64, error) {
	var imageId uint64
	query := `
		INSERT INTO IMAGES(task_id, title)
		VALUES($1, $2)
	`
	row := s.db.QueryRow(query, uuid, img.Name)
	err := row.Scan(&imageId)
	if err != nil {
		return 0, err
	}
	return imageId, nil
}
