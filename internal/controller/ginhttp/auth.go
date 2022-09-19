package ginhttp

import (
	"net/http"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	uc Auth
	l  logger.Logger
}

func setAuthRoutes(handler *gin.RouterGroup, u Auth, l logger.Logger) {
	r := &authRoutes{u, l}

	h := handler.Group("/auth")
	{
		h.POST("sign-up", r.signUp)
		h.POST("sign-in", r.signIn)
	}
}

func (r *authRoutes) signUp(c *gin.Context) {
	var input dto.SignUpRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - auth - signUp")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	output, err := r.uc.SignUp(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - auth - signUp")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *authRoutes) signIn(c *gin.Context) {
	var input dto.SignInRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - auth - signUp")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	output, err := r.uc.SignIn(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - auth - signUp")
		errorResponse(c, http.StatusInternalServerError, "no rows in result set")
		return
	}

	c.JSON(http.StatusOK, output)
}
