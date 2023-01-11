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

type userRoutes struct {
	uc User
	l  logger.Logger
}

func setUserRoutes(handler *gin.RouterGroup, m *middleware, uc User, l logger.Logger) {
	r := &userRoutes{uc, l}

	hAdmin := handler.Group("/users")
	hAdmin.Use(m.userIdentity(), m.checkRole(entity.Admin))
	{
		hAdmin.GET("", r.getAll)
		hAdmin.GET("/:id", r.getById)
		hAdmin.PUT("/:id", r.update)
		hAdmin.PATCH("/:id", r.changeRole)
		hAdmin.DELETE("/:id", r.delete)
	}
}

// @Summary Get all
// @Security ApiKeyAuth
// @Tags Users
// @Description Get list of users
// @Produce json
// @Success 200 {object} dto.GetAllUsersResponse
// @Failure 400,500 {object} response
// @Router /users [get]
func (r *userRoutes) getAll(c *gin.Context) {
	output, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "ginhttp - user - getAll")
		errorResponse(c, http.StatusInternalServerError, "invalid request body")
		return
	}

	c.JSON(http.StatusOK, output)
}

// @Summary Get by id
// @Security ApiKeyAuth
// @Tags Users
// @Description Get user by id
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.GetUserByIdResponse
// @Failure 400,500 {object} response
// @Router /users/{id} [get]
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

// @Summary Update
// @Security ApiKeyAuth
// @Tags Users
// @Description Update user
// @Accept json
// @Param id path int true "id"
// @Param input body dto.UpdateUserRequest true "Updated user data"
// @Success 200
// @Failure 400,500 {object} response
// @Router /users/{id} [put]
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

// @Summary Change role
// @Security ApiKeyAuth
// @Tags Users
// @Description Change user role
// @Accept json
// @Param id path int true "id"
// @Param input body dto.ChangeRoleRequest true "Updated user role"
// @Success 200
// @Failure 400,500 {object} response
// @Router /users/{id} [patch]
func (r *userRoutes) changeRole(c *gin.Context) {
	var input dto.ChangeRoleRequest

	if err := c.BindJSON(&input); err != nil {
		r.l.Error(err, "ginhttp - user - changeUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.l.Error(err, "ginhttp - user - changeUser")
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input.Id = id

	fmt.Println(input)

	err = r.uc.ChangeRole(c.Request.Context(), input)
	if err != nil {
		r.l.Error(err, "ginhttp - user - changeUser")
		errorResponse(c, http.StatusInternalServerError, "invalid request")
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete
// @Security ApiKeyAuth
// @Tags Users
// @Description Delete user
// @Param id path int true "id"
// @Success 200
// @Failure 400,500 {object} response
// @Router /users/{id} [delete]
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
