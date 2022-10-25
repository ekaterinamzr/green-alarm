package main

import (
	"log"

	"github.com/ekaterinamzr/green-alarm/config"
	app "github.com/ekaterinamzr/green-alarm/internal/app/console"
)

func main() {
	cfg, err := config.Load("console")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
