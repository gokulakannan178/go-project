package models

import (
	"time"
)

type VehicleType struct {
	UniqueID     string     `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Name         string     `json:"name"  bson:"name,omitempty"`
	FuelType     string     `json:"fuelType,omitempty" bson:"fuelType,omitempty"`
	VehicleType  string     `json:"vehicleType" bson:"vehicleType,omitempty"`
	FuelCapacity string     `json:"fuelCapacity,omitempty" bson:"fuelCapacity,omitempty"`
	Model        string     `json:"model,omitempty" bson:"model,omitempty"`
	Brand        string     `json:"brand,omitempty" bson:"brand,omitempty"`
	YearOfMaF    string     `json:"yearOfMaF,omitempty" bson:"yearOfMaF,omitempty"`
	Engine       string     `json:"engine,omitempty" bson:"engine,omitempty"`
	Payload      string     `json:"payload"  bson:"payload,omitempty"`
	Mileage      string     `json:"mileage"  bson:"mileage,omitempty"`
	GVW          string     `json:"gvw"  bson:"gvw,omitempty"`
	Power        string     `json:"power" bson:"power,omitempty"`
	Description  string     `json:"description" bson:"description,omitempty"`
	Date         *time.Time `json:"date" bson:"date,omitempty"`
	Status       string     `json:"status" bson:"status,omitempty"`
	Created      Created    `json:"created"  bson:"created,omitempty"`
	Updated      Updated    `json:"updated"  bson:"updated,omitempty"`
}
type VehicleTypeFilter struct {
	Status      []string `json:"status" bson:"status,omitempty"`
	UniqueID    []string `json:"uniqueId"  bson:"uniqueId,omitempty"`
	VehicleType []string `json:"vehicleType" bson:"vehicleType,omitempty"`

	Regex struct {
		Name     string `json:"name" bson:"name"`
		FuelType string `json:"fuelType" bson:"fuelType,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

//VehicleTypeMember:""
type VehicleTypeMember struct {
	Status        string  `json:"status" bson:"status,omitempty"`
	UniqueID      string  `json:"uniqueId"  bson:"uniqueId,omitempty"`
	VehicleTypeID string  `json:"VehicleTypeId" bson:"VehicleTypeId,omitempty"`
	UserName      string  `json:"userName" bson:"userName,omitempty"`
	Created       Created `json:"created"  bson:"created,omitempty"`
}
type RefVechileType struct {
	VehicleType `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
