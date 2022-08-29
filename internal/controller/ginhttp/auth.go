package ginhttp

import (
	"net/http"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/usecase"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	uc usecase.Auth
	l  logger.Logger
}

func setAuthRoutes(handler *gin.RouterGroup, u usecase.Auth, l logger.Logger) {
	r := &authRoutes{u, l}

	h := handler.Group("/user")
	{
		h.POST("signup", r.signUp)
	}
}

func (r *authRoutes) signUp(c *gin.Context) {
	var input dto.SignUpRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - auth - signUp")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := r.uc.SignUp(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - auth - signUp")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, dto.SignUpResponse{Id: id})
}
