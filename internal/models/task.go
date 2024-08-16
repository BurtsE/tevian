package models

type Task struct {
	UUID   string
	Status TaskStatus
	Stats  Stats
	Images []Image
}

func (t *Task) CalcStats() {
	var (
		humanCounter uint32
		malesMeanAge float64
		femalesMeanAge float64
		males int
		females int
	)
	for _, img := range t.Images {
		for _, face := range img.Faces {
			switch face.Sex {
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
	t.Stats.FemalesMeanAge = femalesMeanAge/float64(females)
	t.Stats.MalesMeanAge = malesMeanAge/float64(males)
	t.Stats.FaceCount = humanCounter
	t.Stats.

}
