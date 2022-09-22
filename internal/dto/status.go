package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type CreateStatusRequest struct {
	Name string `json:"status_name"`
}

type CreateStatusResponse struct {
	Id string `json:"id"`
}

type GetAllStatusesResponse struct {
	Statuss []entity.IncidentStatus
}

type GetStatusByIdRequest struct {
	Id string `json:"id"`
}

type GetStatusByIdResponse struct {
	Id   string    `json:"id"`
	Name string `json:"status_name"`
}

type UpdateStatusRequest struct {
	Id   string    `json:"id"`
	Name string `json:"status_name"`
}

type DeleteStatusRequest struct {
	Id string `json:"id"`
}
