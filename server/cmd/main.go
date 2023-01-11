package main

import (
	"log"

	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/internal/app"
)

// @title Green Alarm
// @version 1.0
// @description Swagger API for Golang Project Green Alarm.
// @termsOfService http://swagger.io/terms/

// @contact.email ekaterinaamzr@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in Bearer token
// @name Authorization
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
