package converter

import (
	"errors"
	"tevian/internal/models"
)

func ImageFromFaceApi(data models.FaceServiceTask) models.Image {
	img := models.Image{}
	img.Faces = make([]models.Face, 0)
	for _, face := range data.Data {
		img.Faces = append(img.Faces, models.Face{
			Height: face.Bbox.Height,
			Width:  face.Bbox.Width,
			X:      face.Bbox.X,
			Y:      face.Bbox.Y,
			Gender: face.Demographics.Gender,
			Age:    int(face.Demographics.Age.Mean),
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
