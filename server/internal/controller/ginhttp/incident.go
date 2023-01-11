package ginhttp

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
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
		h.GET("", r.getAll)
		h.POST("", m.userIdentity(), r.create)
		h.GET("/:id", r.getById)
		h.PUT("/:id", m.userIdentity(), m.checkRole(entity.Moderator), r.update)
		h.DELETE("/:id", m.userIdentity(), m.checkRole(entity.Moderator), r.delete)
	}

}

// @Summary Create
// @Security ApiKeyAuth
// @Tags Incidents
// @Description Report an incident
// @Accept json
// @Produce json
// @Param input body dto.CreateIncidentRequest true "New incident data"
// @Success 200 {object} dto.CreateIncidentResponse
// @Failure 400,500 {object} response
// @Router /incidents [post]
func (r *incidentRoutes) create(c *gin.Context) {
	var input dto.CreateIncidentRequest

	fmt.Println(c.Request)

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - incident - create")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	input.Author = c.GetInt("userId")
	// input.Author = 1

	output, err := r.uc.Create(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - incident - create")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get all
// @Tags Incidents
// @Description Get list of incidents
// @Produce json
// @Success 200 {object} dto.GetIncidentsResponse
// @Failure 400,500 {object} response
// @Router /incidents [get]
func (r *incidentRoutes) getAll(c *gin.Context) {
	if c.Query("type") == "" {
		output, err := r.uc.GetAll(c.Request.Context())
		if err != nil {
			r.l.Error(err, "ginhttp - incident - getAll")
			errorResponse(c, http.StatusInternalServerError, "invalid request body")
			return
		}
		c.JSON(http.StatusOK, output)
	} else {
		incidentType, err := strconv.Atoi(c.Query("type"))
		if err != nil {
			r.l.Error(err, "ginhttp - incident - getAll")
			errorResponse(c, http.StatusBadRequest, "invalid type qurey value")
			return
		}
		output, err := r.uc.GetByType(c.Request.Context(), dto.GetIncidentsByTypeRequest{IncidentType: incidentType})
		if err != nil {
			r.l.Error(err, "ginhttp - incident - getAll")
			errorResponse(c, http.StatusInternalServerError, "invalid request body")
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

// @Summary Get by id
// @Tags Incidents
// @Description Get incident by id
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.GetIncidentByIdResponse
// @Failure 400,500 {object} response
// @Router /incidents/{id} [get]
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

// @Summary Update
// @Security ApiKeyAuth
// @Tags Incidents
// @Description Update incident
// @Accept json
// @Param        id   path      int  true  "id"
// @Param input body dto.UpdateIncidentRequest true "Updated incident data"
// @Success 200
// @Failure 400,500 {object} response
// @Router /incidents/{id} [put]
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

// @Summary Delete
// @Security ApiKeyAuth
// @Tags Incidents
// @Param id path int true "id"
// @Description Delete incident
// @Success 200
// @Failure 400,500 {object} response
// @Router /incidents/{id} [delete]
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
