package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Checklist struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID    string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	BoxAvilable string             `json:"boxAvilable" bson:"boxAvilable,omitempty"`
	FuelCheck   string             `json:"fuelCheck" bson:"fuelCheck,omitempty"`
	VehicleId   string             `json:"vehicleId" bson:"vehicleId,omitempty"`
	WheelCheck  string             `json:"wheelCheck" bson:"wheelCheck,omitempty"`
	ThinksCheck string             `json:"thinksCheck" bson:"thinksCheck,omitempty"`
	CheckBy     RegisterBy         `json:"checkBy" bson:"checkBy,omitempty"`
	Address     *Address           `json:"address" bson:"address,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Type        string             `json:"type" bson:"type,omitempty"`
	Date        *time.Time         `json:"date,omitempty"  bson:"date,omitempty"`
	Created     *Created           `json:"created"  bson:"created,omitempty"`
}

type RefChecklist struct {
	Checklist `bson:",inline"`
	Ref       struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
		VehicleId      Vechile      `json:"vehicle" bson:"vehicle,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterChecklist struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	VehicleId      []string `json:"vehicleId" bson:"vehicleId,omitempty"`
	CheckUserId    []string `json:"checkUserId" bson:"checkUserId,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	SortBy         string   `json:"sortBy"`
	SortOrder      int      `json:"sortOrder"`
	Regex          struct {
		Name      string `json:"name" bson:"name"`
		ContactNo string `json:"contactNo" bson:"contactNo"`
		Type      string `json:"type" bson:"type"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
}
