package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env/v11"
)

func InitConfig() (Config, error) {
	cfg := Config{}
	data, err := os.ReadFile("./configs/config.json")
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}
	err = env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
