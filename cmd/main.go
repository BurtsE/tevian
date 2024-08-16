package main

import (
	"log"
	"tevian/internal/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
}
