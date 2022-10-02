package dto

import (
	"time"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type Incident struct {
	Id          int       `json:"id"`
	Name        string    `json:"incident_name"`
	Date        time.Time `json:"incident_date"`
	Country     string    `json:"country"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Publication time.Time `json:"publication_date"`
	Comment     string    `json:"comment"`
	Status      int       `json:"incident_status"`
	Type        int       `json:"incident_type"`
	Author      int       `json:"author"`
}

func FromIncident(incidentEntity *entity.Incident) Incident {
	return Incident{
		Id:          incidentEntity.Id,
		Name:        incidentEntity.Name,
		Date:        incidentEntity.Date,
		Country:     incidentEntity.Country,
		Latitude:    incidentEntity.Latitude,
		Longitude:   incidentEntity.Longitude,
		Publication: incidentEntity.Publication,
		Comment:     incidentEntity.Comment,
		Status:      incidentEntity.Status,
		Type:        incidentEntity.Type,
		Author:      incidentEntity.Author,
	}
}

func FromIncidents(incidentEntities []entity.Incident) []Incident {
	dto := make([]Incident, len(incidentEntities))
	for i := range(dto) {
		dto[i] = FromIncident(&incidentEntities[i])
	}
	return dto
}

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

// type GetAllIncidentsResponse struct {
// 	Incidents []Incident
// }

type GetAllIncidentsResponse []Incident

type GetIncidentsByTypeRequest struct {
	IncidentType int `json:"incident_type"`
}

// type GetIncidentsByTypeResponse struct {
// 	Incidents []Incident
// }

type GetIncidentsByTypeResponse []Incident

type GetIncidentByIdRequest struct {
	Id int `json:"id"`
}

// type GetIncidentByIdResponse struct {
// 	Incident Incident
// }

type GetIncidentByIdResponse Incident

type UpdateIncidentRequest struct {
	Id        int       `json:"id"`
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

type DeleteIncidentRequest struct {
	Id int `json:"id"`
}
