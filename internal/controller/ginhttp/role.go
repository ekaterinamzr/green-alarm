package ginhttp

import (
	"net/http"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type roleRoutes struct {
	uc UserRole
	l  logger.Logger
}

func setRoleRoutes(handler *gin.RouterGroup, m *middleware, uc UserRole, l logger.Logger) {
	r := &roleRoutes{uc, l}

	h := handler.Group("/roles")
	{
		h.GET("/", r.getAll)
		h.POST("/", r.create)
		h.GET(":id", r.getById)
		h.PUT(":id", r.update)
		h.DELETE(":id", r.delete)
	}
}

func (r *roleRoutes) create(c *gin.Context) {
	var input dto.CreateRoleRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - role - create")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	output, err := r.uc.Create(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - role - create")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *roleRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - role - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *roleRoutes) getById(c *gin.Context) {
	var input dto.GetRoleByIdRequest

	input.Id = c.Param("id")

	output, err := r.uc.GetById(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - role - getById")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *roleRoutes) update(c *gin.Context) {
	var input dto.UpdateRoleRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - role - update")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input.Id = c.Param("id")

	err := r.uc.Update(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - role - update")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

func (r *roleRoutes) delete(c *gin.Context) {
	var input dto.DeleteRoleRequest

	input.Id = c.Param("id")

	err := r.uc.Delete(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - role - delete")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}
