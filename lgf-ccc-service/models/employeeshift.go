package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeShift struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID    string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"desc" bson:"desc,omitempty"`
	StartTime   *time.Time         `json:"startTime" bson:"startTime,omitempty"`
	EndTime     *time.Time         `json:"endTime" bson:"endTime,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Created     *Created           `json:"created"  bson:"created,omitempty"`
}

type RefEmployeeShift struct {
	EmployeeShift `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeShift struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
