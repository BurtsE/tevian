package disk

import (
	"os"
	def "tevian/internal/storage"
)

var _ def.DiskStorage = (*storage)(nil)

type storage struct {
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) mkDir(dirName string) error {
	err := os.Mkdir(dirName, 0777)
	if !os.IsExist(err) {
		return err
	}
	return nil
}
