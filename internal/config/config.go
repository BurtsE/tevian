package config

type Config struct {
	// Postgres  `json:"postgres"`
	// Server    `json:"server"`
	FaceCloud `json:"face_cloud"`
}
type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}
type Server struct {
	Port int
}

type FaceCloud struct {
	Login    string `json:"login",env:"home"`
	Password string `json:"password"`
	Url      string `json:"url"`
}
