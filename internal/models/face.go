package models

type Face struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
