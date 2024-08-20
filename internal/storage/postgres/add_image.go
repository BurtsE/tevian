package postgres

func (s *Storage) AddImage(uuid, title string) (uint64, error) {
	var imageId uint64
	query := `
		INSERT INTO IMAGES(task_id, title)
		VALUES($1,$2)
		RETURNING id
	`
	row := s.db.QueryRow(query, &uuid, &title)
	err := row.Scan(&imageId)
	if err != nil {
		return 0, err
	}
	return imageId, nil
}
