package app

import (
	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)
	l.Debug("App is running!")

	pg, err := postgres.New(cfg.Database.URL)
	if err != nil {
		l.Fatal(err, "app - Run - postgres.New")
	}
	l.Debug("Connected to db!")

	defer pg.Close()
}
