package ginhttp

import (
	"net/http"
	"strconv"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
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
		h.GET("", r.getAll)
	}

	hAdmin := handler.Group("/roles")
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
// @Tags Roles
// @Description Create role
// @Accept json
// @Produce json
// @Param input body dto.CreateRoleRequest true "New role data"
// @Success 200 {object} dto.CreateRoleResponse
// @Failure 400,500 {object} response
// @Router /roles [post]
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

// @Summary Get all
// @Tags Roles
// @Description Get list of roles
// @Produce json
// @Success 200 {object} dto.GetAllRolesResponse
// @Failure 400,500 {object} response
// @Router /roles [get]
func (r *roleRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - role - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get by id
// @Tags Roles
// @Description Get role by id
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.GetRoleByIdResponse
// @Failure 400,500 {object} response
// @Router /roles/{id} [get]
func (r *roleRoutes) getById(c *gin.Context) {
	var input dto.GetRoleByIdRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - role - getById")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	output, err := r.uc.GetById(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - role - getById")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Update
// @Security ApiKeyAuth
// @Tags Roles
// @Description Update role
// @Accept json
// @Param id path int true "id"
// @Param input body dto.UpdateRoleRequest true "Updated role data"
// @Success 200
// @Failure 400,500 {object} response
// @Router /roles/{id} [put]
func (r *roleRoutes) update(c *gin.Context) {
	var input dto.UpdateRoleRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - role - update")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - role - update")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Update(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - role - update")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete
// @Security ApiKeyAuth
// @Tags Roles
// @Description Delete role
// @Param id path int true "id"
// @Success 200
// @Failure 400,500 {object} response
// @Router /roles/{id} [delete]
func (r *roleRoutes) delete(c *gin.Context) {
	var input dto.DeleteRoleRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - role - delete")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	err = r.uc.Delete(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - role - delete")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}
