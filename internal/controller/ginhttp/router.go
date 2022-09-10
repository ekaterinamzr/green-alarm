package ginhttp

import (
	"context"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

// type TokenParserFunc func(context.Context, string) (id int, role int, err error)

type Auth interface {
	SignUp(context.Context, dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(context.Context, dto.SignInRequest) (*dto.SignInResponse, error)

	//TokenParser() TokenParserFunc
	ParseToken(context.Context, string) (id int, role int, err error)
}

func NewRouter(handler *gin.Engine, l logger.Logger, a Auth) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/api")
	{
		setAuthRoutes(h, a, l)
	}

}
