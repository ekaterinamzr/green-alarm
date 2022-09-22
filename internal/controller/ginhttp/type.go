package ginhttp

import (
	"net/http"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type typeRoutes struct {
	uc IncidentType
	l  logger.Logger
}

func setTypeRoutes(handler *gin.RouterGroup, m *middleware, uc IncidentType, l logger.Logger) {
	r := &typeRoutes{uc, l}

	h := handler.Group("/types")
	{
		h.GET("/", r.getAll)
		h.POST("/", r.create)
		h.GET(":id", r.getById)
		h.PUT(":id", r.update)
		h.DELETE(":id", r.delete)
	}
}

func (r *typeRoutes) create(c *gin.Context) {
	var input dto.CreateTypeRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - type - create")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	output, err := r.uc.Create(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - type - create")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *typeRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - type - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *typeRoutes) getById(c *gin.Context) {
	var input dto.GetTypeByIdRequest

	input.Id = c.Param("id")

	output, err := r.uc.GetById(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - type - getById")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *typeRoutes) update(c *gin.Context) {
	var input dto.UpdateTypeRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - type - update")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input.Id = c.Param("id")

	err := r.uc.Update(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - type - update")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

func (r *typeRoutes) delete(c *gin.Context) {
	var input dto.DeleteTypeRequest

	input.Id = c.Param("id")

	err := r.uc.Delete(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - type - delete")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}
