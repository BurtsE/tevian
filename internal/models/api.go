package models

type FaceServiceApi struct {
	Data []struct {
		Bbox         string `json:"bbox"`
		Demographics struct {
			Age struct {
				Mean int `json:"mean"`
			} `json:"age"`
			Gender string `json:"gender"`
		} `json:"demographics"`
	} `json:"data"`
}
