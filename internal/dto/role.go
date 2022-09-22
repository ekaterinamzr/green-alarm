package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type CreateRoleRequest struct {
	Name string `json:"role_name"`
}

type CreateRoleResponse struct {
	Id string `json:"id"`
}

type GetAllRolesResponse struct {
	Roles []entity.UserRole
}

type GetRoleByIdRequest struct {
	Id string `json:"id"`
}

type GetRoleByIdResponse struct {
	Id string    `json:"id"`
	Name string `json:"role_name"`
}

type UpdateRoleRequest struct {
	Id string    `json:"id"`
	Name string `json:"role_name"`
}

type DeleteRoleRequest struct {
	Id string `json:"id"`
}
