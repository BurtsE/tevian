package config

type Config struct {
	Postgres  `json:"postgres"`
	Server    `json:"server"`
	FaceCloud `json:"face_cloud"`
}
type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `env:"POSTGRES_USER,notEmpty"`
	Password string `env:"POSTGRES_PASSWORD,notEmpty"`
	Sslmode  string `json:"sslmode"`
	DB       string `json:"db"`
}
type Server struct {
	Port string
	Host string
}

type FaceCloud struct {
	Login    string `env:"FACE_CLOUD_LOGIN"`
	Password string `json:"FACE_CLOUD_PASSWORD"`
	Url      string `json:"FACE_CLOUD_LOGIN_URL"`
}
