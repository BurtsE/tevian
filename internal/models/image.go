package models

type Image struct {
	Name  string
	Data  []byte
	Faces []Face
}
