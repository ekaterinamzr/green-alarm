package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type UserRole struct {
	Id int `json:"id"`
	Name string `json:"role_name"`
}

func FromRole(roleEntity *entity.UserRole) UserRole {
	return UserRole{
		Id:          roleEntity.Id,
		Name:        roleEntity.Name,
	}
}

func FromRoles(roleEntities []entity.UserRole) []UserRole {
	dto := make([]UserRole, len(roleEntities))
	for i := range(dto) {
		dto[i] = FromRole(&roleEntities[i])
	}
	return dto
}

type CreateRoleRequest struct {
	Name string `json:"role_name"`
}

type CreateRoleResponse struct {
	Id int `json:"id"`
}

type GetAllRolesResponse []UserRole

type GetRoleByIdRequest struct {
	Id int `json:"id"`
}

type GetRoleByIdResponse UserRole

type UpdateRoleRequest UserRole

type DeleteRoleRequest struct {
	Id int `json:"id"`
}
