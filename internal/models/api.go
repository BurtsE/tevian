package models

type FaceServiceTask struct {
	Data []struct {
		Bbox struct {
			Height int `json:"height"`
			Width  int `json:"width"`
			X      int `json:"x"`
			Y      int `json:"y"`
		} `json:"bbox"`
		Demographics struct {
			Age struct {
				Mean float64 `json:"mean"`
			} `json:"age"`
			Gender string `json:"gender"`
		} `json:"demographics"`
	} `json:"data"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
}
type FaceServiceLogin struct {
	Data struct {
		Token string `json:"access_token"`
	} `json:"data"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code"`
}
