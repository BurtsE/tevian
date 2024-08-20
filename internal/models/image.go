package models

type Image struct {
	Id    int64
	Name  string
	Data  []byte
	Faces []Face
}
