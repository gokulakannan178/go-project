package models

import (
	"time"
)

type VehicleLog struct {
	UniqueID    string       `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Vehicle     Vechile      `json:"vehicle"  bson:"vehicle,omitempty"`
	VechileType *VehicleType `json:"vechileType"  bson:"vechileType,omitempty"`
	Driver      User         `json:"driver"  bson:"driver,omitempty"`
	Description string       `json:"description" bson:"description,omitempty"`
	StartDate   *time.Time   `json:"startDate" bson:"startDate,omitempty"`
	EndDate     *time.Time   `json:"endDate" bson:"endDate,omitempty"`
	Status      string       `json:"status" bson:"status,omitempty"`
	Created     Created      `json:"created"  bson:"created,omitempty"`
	Updated     Updated      `json:"updated"  bson:"updated,omitempty"`
}
type VehicleLogFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	UniqueID  []string `json:"uniqueId"  bson:"uniqueId,omitempty"`
	VehicleId []string `json:"vehicleId"  bson:"vehicleId,omitempty"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type RefVehicleLog struct {
	VehicleLog `bson:",inline"`
	Ref        struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
