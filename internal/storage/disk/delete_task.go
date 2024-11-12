package disk

import (
	"context"
	"fmt"
	"os"
)

func (s *storage) DeleteImages(ctx context.Context, uuid string) error {
	folderName := fmt.Sprintf("images/%s", uuid)
	return os.RemoveAll(folderName)
}
