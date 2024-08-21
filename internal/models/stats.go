package models

type Stats struct {
	FaceCount      uint32  `json:"face_count"`
	HumanCount     uint32  `json:"human_count"`
	MalesMeanAge   float64 `json:"males_mean_age"`
	FemalesMeanAge float64 `json:"females_mean_age"`
}
