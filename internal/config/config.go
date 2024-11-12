package config

type Config struct {
	Postgres    `json:"postgres"`
	Server      `json:"server"`
	FaceCloud   `json:"face_cloud"`
	Credentials `json:"credentials"`
	TGBot       `json:"tg_bot"`
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
	Email    string `env:"FACE_CLOUD_LOGIN"`
	Password string `env:"FACE_CLOUD_PASSWORD"`
	Url      string `env:"FACE_CLOUD_URL"`
}

type Credentials struct {
	Login    string `env:"LOGIN,notEmpty"`
	Password string `env:"PASSWORD,notEmpty"`
}

type TGBot struct {
	Token string `env:"TELEGRAM_TOKEN,notEmpty"`
}
