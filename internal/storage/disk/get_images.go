package disk

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"tevian/internal/models"
)

// GetImages implements storage.DiskStorage.
func (s *storage) Images(uuid string) ([]models.Image, error) {
	dirName := fmt.Sprintf("images/%s/", uuid)
	entries, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	images := make([]models.Image, 0)
	for _, entry := range entries {
		imageBytes, err := os.ReadFile(dirName + entry.Name())
		if err != nil {
			return nil, err
		}
		imageId, err := strconv.ParseInt(strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())), 10, 64)
		if err != nil {
			return nil, err
		}
		images = append(images, models.Image{
			Id:   imageId,
			Data: imageBytes,
		})
	}
	return images, nil
}
