package main

import (
	"log"
	"login-user/config"
	"login-user/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Hubo un error en la configuracion:(. ", err)
	}
	app.RunApp(cfg)
}
