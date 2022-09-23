package entity

import "github.com/ekaterinamzr/green-alarm/pkg/date"

// Type
const (
	Oil = iota + 1
	Radiation
	Chemical
	Bio
	Fire
	Dump
	Other
)

type IncidentType struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"type_name" db:"type_name"`
}

// Status
const (
	Confirmed = iota + 1
	Unconfirmed
)

type IncidentStatus struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"status_name" db:"status_name"`
}

// time_format:"2006-01-02"

type Incident struct {
	Id               string    `json:"id" bson:"_id" db:"id"`
	Name             string    `json:"incident_name" bson:"incident_name" db:"incident_name"`
	Date             date.CustomDate `json:"incident_date" bson:"incident_date" db:"incident_date" time_format:"2006-01-02"`
	Country          string    `json:"country" bson:"country" db:"country"`
	Latitude         float64   `json:"latitude" bson:"latitude" db:"latitude"`
	Longitude        float64   `json:"longitude" bson:"longitude" db:"longitude"`
	Publication_date date.CustomDate `json:"publication_date" bson:"publication_date" db:"publication_date" time_format:"2006-01-02"`
	Comment          string    `json:"comment" bson:"comment" db:"comment"`
	Status           int       `json:"incident_status" bson:"incident_status" db:"incident_status"`
	Type             int       `json:"incident_type" bson:"incident_type" db:"incident_type"`
	Author           int       `json:"author" bson:"author" db:"author"`
}
