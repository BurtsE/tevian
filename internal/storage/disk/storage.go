package disk

import (
	"bytes"
	"os"
)

type storage struct {
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) SaveImage(uuid, title string, img []byte) error {
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

	fileName.WriteByte('/')
	fileName.WriteString(title)
	err = os.WriteFile(fileName.String(), img, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) mkDir(dirName string) error {
	err := os.Mkdir(dirName, 0777)
	if !os.IsExist(err) {
		return err
	}
	return nil
}
