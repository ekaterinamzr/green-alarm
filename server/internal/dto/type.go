package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type IncidentType struct {
	Id   int    `json:"id"`
	Name string `json:"type_name"`
}

func FromType(typeEntity *entity.IncidentType) IncidentType {
	return IncidentType{
		Id:   typeEntity.Id,
		Name: typeEntity.Name,
	}
}

func FromTypes(typeEntities []entity.IncidentType) []IncidentType {
	dto := make([]IncidentType, len(typeEntities))
	for i := range dto {
		dto[i] = FromType(&typeEntities[i])
	}
	return dto
}

type CreateTypeRequest struct {
	Name string `json:"type_name"`
}

type CreateTypeResponse struct {
	Id int `json:"id"`
}

type GetAllTypesResponse []IncidentType

type GetTypeByIdRequest struct {
	Id int `json:"id"`
}

type GetTypeByIdResponse IncidentType

type UpdateTypeRequest IncidentType

type DeleteTypeRequest struct {
	Id int `json:"id"`
}
