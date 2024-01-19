package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WardWiseDumpHistory struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	WardCode       string             `json:"wardCode"  bson:"wardCode,omitempty"`
	CircleCode     string             `json:"circleCode,omitempty"  bson:"circleCode,omitempty"`
	Quantity       string             `json:"quantity" bson:"quantity,omitempty"`
	Date           string             `json:"date" bson:"date,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	VehicleId      string             `json:"vehicleId" bson:"vehicleId,omitempty"`
	Created        *Created           `json:"created"  bson:"created,omitempty"`
}

type RefWardWiseDumpHistory struct {
	WardWiseDumpHistory `bson:",inline"`
	Ref                 struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterWardWiseDumpHistory struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	BranchId       []string `json:"branchId" bson:"branchId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
