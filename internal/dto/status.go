package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type IncidentStatus struct {
	Id int `json:"id"`
	Name string `json:"status_name"`
}

func FromStatus(statusEntity *entity.IncidentStatus) IncidentStatus {
	return IncidentStatus{
		Id:          statusEntity.Id,
		Name:        statusEntity.Name,
	}
}

func FromStatuses(statusEntities []entity.IncidentStatus) []IncidentStatus {
	dto := make([]IncidentStatus, len(statusEntities))
	for i := range(dto) {
		dto[i] = FromStatus(&statusEntities[i])
	}
	return dto
}

type CreateStatusRequest struct {
	Name string `json:"status_name"`
}

type CreateStatusResponse struct {
	Id int `json:"id"`
}

type GetAllStatusesResponse []IncidentStatus

type GetStatusByIdRequest struct {
	Id int `json:"id"`
}

type GetStatusByIdResponse IncidentStatus

type UpdateStatusRequest IncidentStatus

type DeleteStatusRequest struct {
	Id int `json:"id"`
}
