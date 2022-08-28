package ginhttp

import (
	"github.com/ekaterinamzr/green-alarm/internal/usecase"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	u usecase.User
}

func setUserRoutes(handler *gin.RouterGroup, u usecase.User, l logger.Logger) {
	r := &userRoutes{u}

	h := handler.Group("/user")
	{
		h.POST("register", r.register)
	}
}

func (r *userRoutes) register(c *gin.Context) {

}
