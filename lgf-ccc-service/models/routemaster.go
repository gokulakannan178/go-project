package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RouteMaster struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name               string             `json:"name" bson:"name,omitempty"`
	Area               Address            `json:"area" bson:"area,omitempty"`
	Roadtype           string             `json:"roadtype,omitempty" bson:"roadtype,omitempty"`
	NumberOfProperties int64              `json:"numberOfProperties,omitempty" bson:"numberOfProperties,omitempty"`
	StartTime          *time.Time         `json:"startTime" bson:"startTime,omitempty"`
	EndTime            *time.Time         `json:"EndTime" bson:"EndTime,omitempty"`
	Location           []Location         `json:"location" bson:"location,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            *Created           `json:"created"  bson:"created,omitempty"`
}

type RefRouteMaster struct {
	RouteMaster `bson:",inline"`
	Ref         struct {
		EmployeeId Beat      `json:"organisationId" bson:"organisationId,omitempty"`
		Roadtype   RoadType  `json:"roadtype,omitempty" bson:"roadtype,omitempty"`
		State      *State    `json:"state,omitempty" bson:"state,omitempty"`
		District   *District `json:"district,omitempty" bson:"district,omitempty"`
		Village    *District `json:"village,omitempty" bson:"village,omitempty"`
		Zone       *Zone     `json:"zone,omitempty" bson:"zone,omitempty"`
		Ward       *Ward     `json:"ward,omitempty" bson:"ward,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterRouteMaster struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId []string `json:"employeeID" bson:"employeeID,omitempty"`
	WardCode   []string `json:"wardCode" bson:"wardCode,omitempty"`
	ZoneCode   []string `json:"zoneCode" bson:"zoneCode,omitempty"`
	Roadtype   []string `json:"roadtype,omitempty" bson:"roadtype,omitempty"`

	Regex struct {
		Name      string `json:"name" bson:"name"`
		Type      string `json:"type" bson:"type"`
		ContactNo string `json:"ContactNo" bson:"ContactNo"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
