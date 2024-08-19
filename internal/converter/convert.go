package converter

import (
	"errors"
	"tevian/internal/models"
)

func ImageFromFaceApi(data models.FaceServiceApi) models.Image {
	img := models.Image{}
	img.Faces = make([]models.Face, 0)
	for _, face := range data.Data {
		img.Faces = append(img.Faces, models.Face{
			Bbox:   face.Bbox,
			Gender: face.Demographics.Gender,
			Age:    face.Demographics.Age.Mean,
		})
	}
	return img
}
func TaskStatusFromString(status string) (models.TaskStatus, error) {
	switch status {
	case "pending":
		return models.Pending, nil
	case "processed":
		return models.Pending, nil
	case "completed":
		return models.Completed, nil
	case "failed":
		return models.Failed, nil
	default:
		return nil, errors.New("unknown status")
	}
}
