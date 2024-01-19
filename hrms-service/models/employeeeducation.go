package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeEducation struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeID     string             `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	InstituteName  string             `json:"instituteName,omitempty" bson:"instituteName,omitempty"`
	YearOfPassout  int64              `json:"yearOfPassout,omitempty" bson:"yearOfPassout,omitempty"`
	Percentage     float64            `json:"percentage,omitempty" bson:"percentage,omitempty"`
	Department     string             `json:"department,omitempty" bson:"department,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeEducation struct {
	EmployeeEducation `bson:",inline"`
	Ref               struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeEducation struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeID     []string `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		InstituteName string `json:"instituteName" bson:"instituteName"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
