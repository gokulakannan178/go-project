package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeAttendanceCalendar struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	StartDate      *time.Time         `json:"startDate" bson:"startDate,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Created        *Created           `json:"created"  bson:"created,omitempty"`
}

type RefEmployeeAttendanceCalendar struct {
	EmployeeAttendanceCalendar `bson:",inline"`
	Ref                        struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeAttendanceCalendar struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
