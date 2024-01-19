package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeAssets struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationID  string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId      string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	AssetPropertyId string             `json:"assetPropertyId,omitempty" bson:"assetPropertyId,omitempty"`
	Created         *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeAssets struct {
	EmployeeAssets `bson:",inline"`
	Ref            struct {
		OrganisationId  Organisation   `json:"organisationId" bson:"organisationId,omitempty"`
		AssetPropertyId AssetPropertys `json:"assetPropertyId,omitempty" bson:"assetPropertyId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeAssets struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
