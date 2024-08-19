package disk

import (
	"fmt"
	"os"
)

func (s *storage) DeleteImages(uuid string) error {
	folderName := fmt.Sprintf("images/%s", uuid)
	return os.RemoveAll(folderName)
}
