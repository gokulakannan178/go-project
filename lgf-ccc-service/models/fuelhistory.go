package models

import (
	"time"
)

type FuelHistory struct {
	UniqueID string   `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Vehicle  *MinUser `json:"vehicle"  bson:"vehicle,omitempty"`
	Driver   *Driver  `json:"driver"  bson:"driver,omitempty"`
	OdoMeter string   `json:"odoMeter" bson:"odoMeter,omitempty"`
	// StartReading string     `json:"startReading" bson:"startReading,omitempty"`
	// EndReading   string     `json:"endReading" bson:"endReading,omitempty"`
	Quantity string     `json:"quantity" bson:"quantity,omitempty"`
	FuelType string     `json:"fuelType" bson:"fuelType,omitempty"`
	Price    string     `json:"price" bson:"price,omitempty"`
	Mileage  string     `json:"mileage" bson:"mileage,omitempty"`
	Date     *time.Time `json:"date" bson:"date,omitempty"`
	Time     *time.Time `json:"time" bson:"time,omitempty"`
	EndDate  *time.Time `json:"endDate" bson:"endDate,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Created  Created    `json:"created"  bson:"created,omitempty"`
	Updated  Updated    `json:"updated"  bson:"updated,omitempty"`
}
type FuelHistoryFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Regex    struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type RefFuelHistory struct {
	FuelHistory `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
