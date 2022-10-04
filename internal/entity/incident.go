package entity

import "time"

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
	Id   int    `json:"id" db:"id"`
	Name string `json:"type_name" db:"type_name"`
}

// Status
const (
	Confirmed = iota + 1
	Unconfirmed
)

type IncidentStatus struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"status_name" db:"status_name"`
}

type Incident struct {
	Id          int       `json:"id" db:"id" bson:"_id"`
	Name        string    `json:"incident_name" db:"incident_name" bson:"incident_name"`
	Date        time.Time `json:"incident_date" db:"incident_date" bson:"incident_date"`
	Country     string    `json:"country" db:"country" bson:"country"`
	Latitude    float64   `json:"latitude" db:"latitude" bson:"latitude"`
	Longitude   float64   `json:"longitude" db:"longitude" bson:"longitude"`
	Publication time.Time `json:"publication_date" db:"publication_date" bson:"publication_date"`
	Comment     string    `json:"comment" db:"comment" bson:"comment"`
	Status      int       `json:"incident_status" db:"incident_status" bson:"incident_status"`
	Type        int       `json:"incident_type" db:"incident_type" bson:"incident_type"`
	Author      int       `json:"author" db:"author" bson:"author"`
}
