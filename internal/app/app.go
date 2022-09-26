package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/internal/controller/ginhttp"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/pgrepo"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/token"
	"github.com/ekaterinamzr/green-alarm/internal/usecase"
	"github.com/ekaterinamzr/green-alarm/pkg/httpserver"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)
	l.Debug("App is running!")

	pg, err := postgres.New(cfg.Database.URI)
	if err != nil {
		l.Fatal(err, "app - Run - postgres.New")
	}
	l.Debug("Connected to db!")

	defer pg.Close()

	// mongo, err := mongo.New(cfg.MongoDB.URI)
	// if err != nil {
	// 	l.Fatal(err, "app - Run - mongo.New")
	// }
	// l.Debug("Connected to MongoDB!")

	// defer mongo.Close()

	token := token.NewTokenService(cfg.Auth.TokenTTL, cfg.Auth.SigningKey)

	authUseCase := usecase.NewAuthUseCase(pgrepo.NewUserRepository(pg), token, cfg.Auth.Salt)
	incidentUseCase := usecase.NewIncidentUseCase(pgrepo.NewIncidentRepository(pg))
	// incidentUseCase := usecase.NewIncidentUseCase(mongorepo.NewIncidentRepository(mongo))

	statusUseCase := usecase.NewStatusUseCase(pgrepo.NewStatusRepository(pg))
	typeUseCase := usecase.NewTypeUseCase(pgrepo.NewTypeRepository(pg))

	userUseCase := usecase.NewUserUseCase(pgrepo.NewUserRepository(pg))
	roleUseCase := usecase.NewRoleUseCase(pgrepo.NewRoleRepository(pg))

	handler := gin.New()
	ginhttp.NewRouter(handler, l,
		authUseCase,
		incidentUseCase,
		typeUseCase,
		statusUseCase,
		userUseCase,
		roleUseCase,
	)

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
