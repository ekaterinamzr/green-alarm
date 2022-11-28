package ginhttp

import (
	"net/http"
	"strconv"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/logger"
	"github.com/gin-gonic/gin"
)

type statusRoutes struct {
	uc IncidentStatus
	l  logger.Logger
}

func setStatusRoutes(handler *gin.RouterGroup, m *middleware, uc IncidentStatus, l logger.Logger) {
	r := &statusRoutes{uc, l}

	h := handler.Group("/statuses")
	h.Use(m.userIdentity(), m.checkRole(entity.Admin))
	{
		h.GET("", r.getAll)
		h.POST("", r.create)
		h.GET("/:id", r.getById)
		h.PUT("/:id", r.update)
		h.DELETE("/:id", r.delete)
	}
}

// @Summary Create
// @Security ApiKeyAuth
// @Tags Statuses
// @Description Create status
// @Accept json
// @Produce json
// @Param input body dto.CreateStatusRequest true "New status data"
// @Success 200 {object} dto.CreateStatusResponse
// @Failure 400,500 {object} response
// @Router /statuses [post]
func (r *statusRoutes) create(c *gin.Context) {
	var input dto.CreateStatusRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - status - create")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	output, err := r.uc.Create(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - status - create")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get all
// @Security ApiKeyAuth
// @Tags Statuses
// @Description Get list of statuses
// @Produce json
// @Success 200 {object} dto.GetAllStatusesResponse
// @Failure 400,500 {object} response
// @Router /statuses [get]
func (r *statusRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - status - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get by id
// @Security ApiKeyAuth
// @Tags Statuses
// @Description Get status by id
// @Produce json
// @Success 200 {object} dto.GetStatusByIdResponse
// @Failure 400,500 {object} response
// @Router /statuses/{id} [get]
func (r *statusRoutes) getById(c *gin.Context) {
	var input dto.GetStatusByIdRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - status - getById")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	output, err := r.uc.GetById(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - status - getById")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Update
// @Security ApiKeyAuth
// @Tags Statuses
// @Description Update status
// @Accept json
// @Param id path int true "id"
// @Param input body dto.UpdateStatusRequest true "Updated status data"
// @Success 200
// @Failure 400,500 {object} response
// @Router /statuses/{id} [put]
func (r *statusRoutes) update(c *gin.Context) {
	var input dto.UpdateStatusRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - status - update")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - status - update")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Update(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - status - update")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete
// @Security ApiKeyAuth
// @Tags Statuses
// @Description Delete status
// @Success 200
// @Failure 400,500 {object} response
// @Router /statuses/{id} [delete]
func (r *statusRoutes) delete(c *gin.Context) {
	var input dto.DeleteStatusRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - status - delete")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Delete(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - status - delete")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}
