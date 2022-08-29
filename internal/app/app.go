package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/internal/controller/ginhttp"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/pgrepo"
	"github.com/ekaterinamzr/green-alarm/internal/usecase"
	"github.com/ekaterinamzr/green-alarm/pkg/httpserver"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
	"github.com/gin-gonic/gin"
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

	userUseCase := usecase.NewAuthUseCase(pgrepo.NewUserRepository(pg), cfg.Auth.Salt)

	handler := gin.New()
	ginhttp.NewRouter(handler, l, userUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Server.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err), "")
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err), "")
	}
}
