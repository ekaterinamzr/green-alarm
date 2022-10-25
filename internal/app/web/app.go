package app

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/internal/controller/ginhttp"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/hash"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/mgrepo"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/pgrepo"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/token"
	"github.com/ekaterinamzr/green-alarm/internal/usecase"
	"github.com/ekaterinamzr/green-alarm/pkg/httpserver"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/ekaterinamzr/green-alarm/pkg/mongo"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)
	l.Debug("App is running!")

	var (
		userRepo     usecase.UserRepository
		authRepo     usecase.AuthRepository
		incidentRepo usecase.IncidentRepository
		typeRepo     usecase.TypeRepository
		statusRepo   usecase.StatusRepository
		roleRepo     usecase.RoleRepository
	)

	if cfg.DBType == "pg" {
		pg, err := postgres.New(cfg.Database.URI)
		if err != nil {
			l.Fatal(err, "app - Run - postgres.New")
		}
		l.Debug("Connected to db!")

		userRepo = pgrepo.NewUserRepository(pg)
		authRepo = pgrepo.NewUserRepository(pg)
		incidentRepo = pgrepo.NewIncidentRepository(pg)
		typeRepo = pgrepo.NewTypeRepository(pg)
		statusRepo = pgrepo.NewStatusRepository(pg)
		roleRepo = pgrepo.NewRoleRepository(pg)

		defer pg.Close()
	} else if cfg.DBType == "mg" {
		mongo, err := mongo.New(cfg.MongoDB.URI)
		if err != nil {
			l.Fatal(err, "app - Run - mongo.New")
		}
		l.Debug("Connected to MongoDB!")

		authRepo = mgrepo.NewUserRepository(mongo)
		incidentRepo = mgrepo.NewIncidentRepository(mongo)

		defer mongo.Close()
	} else {
		l.Fatal(errors.New("Wrong DBType config"), "app - Run")
	}

	token := token.NewTokenService(cfg.Auth.TokenTTL, cfg.Auth.SigningKey)

	authUseCase := usecase.NewAuthUseCase(authRepo, token, hash.GenerateHash, cfg.Auth.Salt)
	incidentUseCase := usecase.NewIncidentUseCase(incidentRepo)

	statusUseCase := usecase.NewStatusUseCase(statusRepo)
	typeUseCase := usecase.NewTypeUseCase(typeRepo)

	userUseCase := usecase.NewUserUseCase(userRepo)
	roleUseCase := usecase.NewRoleUseCase(roleRepo)

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
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err), "")
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err), "")
	}
}
