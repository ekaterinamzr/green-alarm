package ginhttp

import (
	"context"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	SignUp(context.Context, dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(context.Context, dto.SignInRequest) (*dto.SignInResponse, error)

	ParseToken(context.Context, string) (id int, role int, err error)
}

type Incident interface {
	Create(context.Context, dto.CreateIncidentRequest) (*dto.CreateIncidentResponse, error)
	GetAll(context.Context) (*dto.GetAllIncidentsResponse, error)
	GetById(context.Context, dto.GetIncidentByIdRequest) (*dto.GetIncidentByIdResponse, error)
	Update(context.Context, dto.UpdateIncidentRequest) error
	Delete(context.Context, dto.DeleteIncidentRequest) error
}

func NewRouter(handler *gin.Engine, l logger.Logger, a Auth, i Incident) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	
	m := newMiddleware(a.ParseToken)

	h := handler.Group("/api")
	{
		setAuthRoutes(h, a, l)
		setIncidentRoutes(h, m, i, l)
	}

}
