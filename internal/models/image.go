package models

type Image struct {
	Id    int64
	Name  string `json:"name"`
	Data  []byte `json:"-"`
	Faces []Face `json:"faces"`
}
