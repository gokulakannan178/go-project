package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VehicleLocation struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	VehicleId string             `json:"vehicleId" bson:"vehicleId,omitempty"`
	BeatId    string             `json:"beatId" bson:"beatId,omitempty"`
	Date      string             `json:"date" bson:"date,omitempty"`
	Location  Location           `json:"location" bson:"location,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
	Created   *Created           `json:"created"  bson:"created,omitempty"`
}

type RefVehicleLocation struct {
	VehicleLocation `bson:",inline"`
	Ref             struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
		BeatId         Beat         `json:"beatId" bson:"beatId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterVehicleLocation struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	VehicleId []string `json:"vehicleId" bson:"vehicleId,omitempty"`
	Regex     struct {
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
