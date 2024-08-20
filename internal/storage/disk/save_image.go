package disk

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func (s *storage) SaveImage(uuid, title string, imageId uint64, img []byte) error {
	fileName := bytes.NewBuffer([]byte("images"))

	err := s.mkDir(fileName.String())
	if err != nil {
		return err
	}

	fileName.WriteByte('/')
	fileName.WriteString(uuid)
	err = s.mkDir(fileName.String())
	if err != nil {
		return err
	}
	format := filepath.Ext(title)
	fileName.WriteByte('/')
	fileName.WriteString(fmt.Sprintf("%d", imageId, format))
	err = os.WriteFile(fileName.String(), img, 0777)
	if err != nil {
		return err
	}
	return nil
}
