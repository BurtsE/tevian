package converter

import "tevian/internal/models"

func ImageFromFaceApi(data models.FaceServiceApi) models.Image {
	img := models.Image{}
	img.Faces = make([]models.Face, 0)
	for _, face := range data.Data {
		img.Faces = append(img.Faces, models.Face{
			Bbox: face.Bbox,
			Gender: face.Demographics.Gender,
			Age: face.Demographics.Age.Mean,
		})
	}
	return img
}
