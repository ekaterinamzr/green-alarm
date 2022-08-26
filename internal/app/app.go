package app

import (
	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)
	l.Debug("App is running")
}
