package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CircleWiseHouseVisited struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty"`
	CircleCode      string             `json:"circleCode"  bson:"circleCode,omitempty"`
	TodayCollection int                `json:"todayCollection" bson:"todayCollection,omitempty"`
	TotalProperties int                `json:"totalProperties" bson:"totalProperties,omitempty"`
	Datestr         string             `json:"dateStr" bson:"dateStr,omitempty"`
	Date            *time.Time         `json:"date" bson:"date,omitempty"`
	Status          string             `json:"status" bson:"status,omitempty"`
	OrganisationId  string             `json:"organisationId" bson:"organisationId,omitempty"`
	VehicleId       string             `json:"vehicleId" bson:"vehicleId,omitempty"`
	Created         *Created           `json:"created"  bson:"created,omitempty"`
}

type RefCircleWiseHouseVisited struct {
	CircleWiseHouseVisited `bson:",inline"`
	Ref                    struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterCircleWiseHouseVisited struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	CircleCode []string `json:"circleCode"  bson:"circleCode,omitempty"`
	Datestr    []string `json:"dateStr" bson:"dateStr,omitempty"`
	Regex      struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
