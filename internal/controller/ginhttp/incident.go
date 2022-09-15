package ginhttp

import (
	"net/http"
	"strconv"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type incidentRoutes struct {
	uc Incident
	l  logger.Logger
}

func setIncidentRoutes(handler *gin.RouterGroup, m *middleware, uc Incident, l logger.Logger) {
	r := &incidentRoutes{uc, l}

	h := handler.Group("/incidents")
	{
		h.GET("/", r.getAll)
		h.POST("/", r.create)
		h.GET(":id", r.getById)
		h.PUT(":id", r.update)
		h.DELETE(":id", r.delete)
	}
}

func (r *incidentRoutes) create(c *gin.Context) {
	var input dto.CreateIncidentRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - incident - create")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	// input.Author = c.GetInt("userId")
	input.Author = 1

	output, err := r.uc.Create(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - incident - create")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *incidentRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - incident - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *incidentRoutes) getById(c *gin.Context) {
	var input dto.GetIncidentByIdRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - incident - getById")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	output, err := r.uc.GetById(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - incident - getById")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.JSON(http.StatusOK, output)
}

func (r *incidentRoutes) update(c *gin.Context) {
	var input dto.UpdateIncidentRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - incident - update")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - incident - update")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Update(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - incident - update")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

func (r *incidentRoutes) delete(c *gin.Context) {
	var input dto.DeleteIncidentRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - incident - delete")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Delete(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - incident - delete")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}
