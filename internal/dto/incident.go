package dto

import (
	"time"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type CreateIncidentRequest struct {
	Name      string    `json:"incident_name"`
	Date      time.Time `json:"incident_date"`
	Country   string    `json:"country"`
	Latitude  float64   `json:"latitude,string"`
	Longitude float64   `json:"longitude,string"`
	Comment   string    `json:"comment"`
	Status    int       `json:"incident_status,string"`
	Type      int       `json:"incident_type,string"`
	Author    int       `json:"author,string"`
}

type CreateIncidentResponse struct {
	Id int `json:"id"`
}

type GetAllIncidentsResponse struct {
	Incidents []entity.Incident
}

type GetIncidentByIdRequest struct {
	Id int `json:"id"`
}

type GetIncidentByIdResponse struct {
	Id               int       `json:"id"`
	Name             string    `json:"incident_name"`
	Date             time.Time `json:"incident_date"`
	Country          string    `json:"country"`
	Latitude         float64   `json:"latitude,string"`
	Longitude        float64   `json:"longitude,string"`
	Publication_date time.Time `json:"publication_date"`
	Comment          string    `json:"comment"`
	Status           int       `json:"incident_status,string"`
	Type             int       `json:"incident_type,string"`
	Author           int       `json:"author,string"`
}

type UpdateIncidentRequest struct {
	Id               int       `json:"id"`
	Name             string    `json:"incident_name"`
	Date             time.Time `json:"incident_date"`
	Country          string    `json:"country"`
	Latitude         float64   `json:"latitude,string"`
	Longitude        float64   `json:"longitude,string"`
	Comment          string    `json:"comment"`
	Status           int       `json:"incident_status,string"`
	Type             int       `json:"incident_type,string"`
	Author           int       `json:"author,string"`
}

type DeleteIncidentRequest struct {
	Id int `json:"id"`
}


