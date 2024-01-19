package models

import (
	"time"
)

type VehicleInsurance struct {
	UniqueID        string     `json:"uniqueId"  bson:"uniqueId,omitempty"`
	OwnerName       string     `json:"ownerName"  bson:"ownerName,omitempty"`
	VehicleName     string     `json:"vehicleName"  bson:"vehicleName,omitempty"`
	InsuranceValid  *time.Time `json:"insuranceValid,omitempty" bson:"insuranceValid,omitempty"`
	InsuranceExpiry *time.Time `json:"insuranceExpiry,omitempty" bson:"insuranceExpiry,omitempty"`
	PermitNo        string     `json:"permitNo,omitempty" bson:"permitNo,omitempty"`
	RcNO            string     `json:"rcNO,omitempty" bson:"rcNO,omitempty"`
	RcNOTill        *time.Time `json:"rcNOTill,omitempty" bson:"rcNOTill,omitempty"`
	PUCNo           string     `json:"pUCNo,omitempty" bson:"pUCNo,omitempty"`
	PUCValidDate    *time.Time `json:"pUCValidDate,omitempty" bson:"pUCValidDate,omitempty"`
	Date            *time.Time `json:"date" bson:"date,omitempty"`
	Status          string     `json:"status" bson:"status,omitempty"`
	Created         Created    `json:"created"  bson:"created,omitempty"`
	Updated         Updated    `json:"updated"  bson:"updated,omitempty"`
}
type VehicleInsuranceFilter struct {
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

type RefVehicleInsurance struct {
	VehicleInsurance `bson:",inline"`
	Ref              struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
