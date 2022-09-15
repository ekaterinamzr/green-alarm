package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type CreateTypeRequest struct {
	Name string `json:"type_name"`
}

type CreateTypeResponse struct {
	Id int `json:"id"`
}

type GetAllTypesResponse struct {
	Types []entity.IncidentType
}

type GetTypeByIdRequest struct {
	Id int `json:"id"`
}

type GetTypeByIdResponse struct {
	Id   int    `json:"id"`
	Name string `json:"type_name"`
}

type UpdateTypeRequest struct {
	Id   int    `json:"id"`
	Name string `json:"type_name"`
}

type DeleteTypeRequest struct {
	Id int `json:"id"`
}
