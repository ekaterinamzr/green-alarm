package dto

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type CreateStatusRequest struct {
	Name string `json:"status_name"`
}

type CreateStatusResponse struct {
	Id int `json:"id"`
}

type GetAllStatusesResponse struct {
	Statuses []entity.IncidentStatus
}

type GetStatusByIdRequest struct {
	Id int `json:"id"`
}

type GetStatusByIdResponse struct {
	Id   int    `json:"id"`
	Name string `json:"status_name"`
}

type UpdateStatusRequest struct {
	Id   int    `json:"id"`
	Name string `json:"status_name"`
}

type DeleteStatusRequest struct {
	Id int `json:"id"`
}
