package facecloud

func (s *service) AddImageToTask(uuid, title string, img []byte) error {
	imageId, err := s.storage.AddImage(uuid, title)
	if err != nil {
		return err
	}

	err = s.diskStorage.SaveImage(uuid, title, imageId, img)
	if err != nil {
		return err
	}
	return nil
}
