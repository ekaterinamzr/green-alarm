package main

import (
	"log"

	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/internal/app"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
