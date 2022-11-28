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
		h.POST("/sign-up", r.signUp)
		h.POST("/sign-in", r.signIn)
	}
}

// @Summary SignUp
// @Tags auth
// @Description Create account
// @Accept json
// @Produce json
// @Param input body dto.SignUpRequest true "New user info"
// @Success 200 {object} dto.SignUpResponse
// @Failure 400,500 {object} response
// @Router /auth/sign-up [post]
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

// @Summary SignIn
// @Tags auth
// @Description Sign in
// @Accept json
// @Produce json
// @Param input body dto.SignInRequest true "Login, password"
// @Success 200 {object} dto.SignInResponse
// @Failure 400,500 {object} response
// @Router /auth/sign-in [post]
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
