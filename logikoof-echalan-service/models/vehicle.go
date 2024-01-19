package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Vehicle : ""
type Vehicle struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RegAuthority    string             `json:"regAuthority" bson:"regAuthority,omitempty"`
	RegNo           string             `json:"regNo" bson:"regNo,omitempty"`
	RegDate         *time.Time         `json:"regDate" bson:"regDate,omitempty"`
	EngineNo        string             `json:"engineNo" bson:"engineNo,omitempty"`
	ChassisNo       string             `json:"chassisNo" bson:"chassisNo,omitempty"`
	OwnerName       string             `json:"ownerName" bson:"ownerName,omitempty"`
	VehicleClass    string             `json:"vehicleClass" bson:"vehicleClass,omitempty"`
	Type            string             `json:"type" bson:"type,omitempty"`
	FuelType        string             `json:"fuelType" bson:"fuelType,omitempty"`
	InsuranceUpTo   *time.Time         `json:"insuranceUpTo" bson:"insuranceUpTo,omitempty"`
	RoadTaxPaidUpto *time.Time         `json:"roadTaxPaidUpto" bson:"roadTaxPaidUpto,omitempty"`
	Model           string             `json:"model" bson:"model,omitempty"`
	FitnessUpTo     *time.Time         `json:"fitnessUpTo" bson:"fitnessUpTo,omitempty"`
	FuelNorms       string             `json:"fuelNorms" bson:"fuelNorms,omitempty"`
	Mobile          string             `json:"mobile" bson:"mobile,omitempty"`
	Status          string             `json:"status" bson:"status,omitempty"`
	Created         Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated         Updated            `json:"updated"  bson:"updated,omitempty"`
	UpdateLog       []Updated          `json:"updatedLog" bson:"updatedLog,omitempty"`
}

//RefVehicle :""
type RefVehicle struct {
	Vehicle `bson:",inline"`
	Ref     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//VehicleFilter : ""
type VehicleFilter struct {
	Status    []string `json:"status"`
	RegNo     []string `json:"regNo"`
	OmitID    []string `json:"omitId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

//VehiclePayment : ""
type VehiclePayment struct {
	Status string `json:"status" bson:"status,omitempty"`
}
