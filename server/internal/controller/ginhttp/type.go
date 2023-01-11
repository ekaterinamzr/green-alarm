package ginhttp

import (
	"net/http"
	"strconv"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
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
		h.GET("", r.getAll)
	}

	hAdmin := handler.Group("/types")
	hAdmin.Use(m.userIdentity(), m.checkRole(entity.Admin))
	{
		hAdmin.POST("", r.create)
		hAdmin.GET("/:id", r.getById)
		hAdmin.PUT("/:id", r.update)
		hAdmin.DELETE("/:id", r.delete)
	}
}

// @Summary Create
// @Security ApiKeyAuth
// @Tags Types
// @Description Create type
// @Accept json
// @Produce json
// @Param input body dto.CreateTypeRequest true "New type data"
// @Success 200 {object} dto.CreateTypeResponse
// @Failure 400,500 {object} response
// @Router /types [post]
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

// @Summary Get all
// @Tags Types
// @Description Get list of types
// @Produce json
// @Success 200 {object} dto.GetAllTypesResponse
// @Failure 400,500 {object} response
// @Router /types [get]
func (r *typeRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - type - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get by id
// @Tags Types
// @Description Get type by id
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.GetTypeByIdResponse
// @Failure 400,500 {object} response
// @Router /types/{id} [get]
func (r *typeRoutes) getById(c *gin.Context) {
	var input dto.GetTypeByIdRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - type - getById")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	output, err := r.uc.GetById(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - type - getById")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Update
// @Security ApiKeyAuth
// @Tags Types
// @Description Update type
// @Accept json
// @Param id path int true "id"
// @Param input body dto.UpdateTypeRequest true "Updated type data"
// @Success 200
// @Failure 400,500 {object} response
// @Router /types/{id} [put]
func (r *typeRoutes) update(c *gin.Context) {
	var input dto.UpdateTypeRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - type - update")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - type - update")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Update(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - type - update")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete
// @Security ApiKeyAuth
// @Tags Types
// @Description Delete type
// @Param id path int true "id"
// @Success 200
// @Failure 400,500 {object} response
// @Router /types/{id} [delete]
func (r *typeRoutes) delete(c *gin.Context) {
	var input dto.DeleteTypeRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - type - delete")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Delete(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - type - delete")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}
