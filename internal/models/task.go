package models

import "encoding/json"

type Task struct {
	UUID   string     `json:"uuid"`
	Status TaskStatus `json:"status"`
	Stats  Stats      `json:"statistics"`
	Images []Image    `json:"images"`
}

func (t *Task) CalcStats() {
	var (
		humanCounter   uint32
		malesMeanAge   float64
		femalesMeanAge float64
		males          int
		females        int
	)
	for _, img := range t.Images {
		for _, face := range img.Faces {
			switch face.Gender {
			case "male":
				males++
				malesMeanAge += float64(face.Age)
			case "female":
				females++
				femalesMeanAge += float64(face.Age)
			}
			humanCounter++
		}
	}
	t.Stats.FemalesMeanAge = femalesMeanAge / float64(females)
	t.Stats.MalesMeanAge = malesMeanAge / float64(males)
	t.Stats.FaceCount = humanCounter
	t.Stats.HumanCount = uint32(males) + uint32(females)
}

func (t *Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		UUID   string  `json:"uuid"`
		Status string  `json:"status"`
		Stats  Stats   `json:"statistics"`
		Images []Image `json:"images"`
	}{
		UUID:   t.UUID,
		Status: t.Status.String(),
		Stats:  t.Stats,
		Images: t.Images,
	})
}
