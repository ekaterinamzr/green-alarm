package ginhttp

import (
	"github.com/ekaterinamzr/green-alarm/internal/usecase"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Logger, a usecase.Auth) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	setAuthRoutes(&handler.RouterGroup, a, l)

}
