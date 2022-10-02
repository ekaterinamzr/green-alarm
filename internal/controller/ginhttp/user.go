package ginhttp

import (
	"net/http"
	"strconv"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	uc User
	l  logger.Logger
}

func setUserRoutes(handler *gin.RouterGroup, m *middleware, uc User, l logger.Logger) {
	r := &userRoutes{uc, l}

	h := handler.Group("/users")
	{
		h.GET("", r.getAll)
		h.GET("/:id", r.getById)
		h.PUT("/:id", r.update)
		h.PATCH("/:id", r.changeRole)
		h.DELETE("/:id", r.delete)
	}
}

func (r *userRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - user - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *userRoutes) getById(c *gin.Context) {
	var input dto.GetUserByIdRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - user - getById")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	output, err := r.uc.GetById(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - user - getById")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *userRoutes) update(c *gin.Context) {
	var input dto.UpdateUserRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - user - update")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - user - update")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Update(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - user - update")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

func (r *userRoutes) changeRole(c *gin.Context) {
	var input dto.ChangeRoleRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - user - updateRole")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - user - updateRole")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.ChangeRole(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - user - makeDefault")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

func (r *userRoutes) delete(c *gin.Context) {
	var input dto.DeleteUserRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - user - delete")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Delete(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - user - delete")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}
