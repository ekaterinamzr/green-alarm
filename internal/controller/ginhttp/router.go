package ginhttp

import (
	"context"
	"time"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	SignUp(context.Context, dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(context.Context, dto.SignInRequest) (*dto.SignInResponse, error)

	ParseToken(context.Context, string) (id int, role int, err error)
}

type Incident interface {
	Create(context.Context, dto.CreateIncidentRequest) (*dto.CreateIncidentResponse, error)
	GetAll(context.Context) (*dto.GetIncidentsResponse, error)
	GetById(context.Context, dto.GetIncidentByIdRequest) (*dto.GetIncidentByIdResponse, error)
	Update(context.Context, dto.UpdateIncidentRequest) error
	Delete(context.Context, dto.DeleteIncidentRequest) error
	GetByType(context.Context, dto.GetIncidentsByTypeRequest) (*dto.GetIncidentsResponse, error)
}

type IncidentType interface {
	Create(context.Context, dto.CreateTypeRequest) (*dto.CreateTypeResponse, error)
	GetAll(context.Context) (*dto.GetAllTypesResponse, error)
	GetById(context.Context, dto.GetTypeByIdRequest) (*dto.GetTypeByIdResponse, error)
	Update(context.Context, dto.UpdateTypeRequest) error
	Delete(context.Context, dto.DeleteTypeRequest) error
}

type IncidentStatus interface {
	Create(context.Context, dto.CreateStatusRequest) (*dto.CreateStatusResponse, error)
	GetAll(context.Context) (*dto.GetAllStatusesResponse, error)
	GetById(context.Context, dto.GetStatusByIdRequest) (*dto.GetStatusByIdResponse, error)
	Update(context.Context, dto.UpdateStatusRequest) error
	Delete(context.Context, dto.DeleteStatusRequest) error
}

type User interface {
	GetAll(context.Context) (*dto.GetAllUsersResponse, error)
	GetById(context.Context, dto.GetUserByIdRequest) (*dto.GetUserByIdResponse, error)
	Update(context.Context, dto.UpdateUserRequest) error
	Delete(context.Context, dto.DeleteUserRequest) error
	ChangeRole(context.Context, dto.ChangeRoleRequest) error
}

type UserRole interface {
	Create(context.Context, dto.CreateRoleRequest) (*dto.CreateRoleResponse, error)
	GetAll(context.Context) (*dto.GetAllRolesResponse, error)
	GetById(context.Context, dto.GetRoleByIdRequest) (*dto.GetRoleByIdResponse, error)
	Update(context.Context, dto.UpdateRoleRequest) error
	Delete(context.Context, dto.DeleteRoleRequest) error
}

func NewRouter(handler *gin.Engine, l logger.Logger, a Auth, i Incident, t IncidentType, s IncidentStatus, u User, r UserRole) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// handler.Use(cors.Default())

	handler.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	m := newMiddleware(a.ParseToken)

	h := handler.Group("/api")
	{
		setAuthRoutes(h, a, l)
		setIncidentRoutes(h, m, i, l)

		setTypeRoutes(h, m, t, l)
		setStatusRoutes(h, m, s, l)

		setUserRoutes(h, m, u, l)
		setRoleRoutes(h, m, r, l)
	}

}
