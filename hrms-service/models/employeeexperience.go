package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeExperience struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeID     string             `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	CompanyName    string             `json:"companyName,omitempty" bson:"companyName,omitempty"`
	YearOfWorking  int64              `json:"yearOfWorking,omitempty" bson:"yearOfWorking,omitempty"`
	Experience     string             `json:"experience,omitempty" bson:"experience,omitempty"`
	Role           string             `json:"role,omitempty" bson:"role,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeExperience struct {
	EmployeeExperience `bson:",inline"`
	Ref                struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeExperience struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeID     []string `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		CompanyName string `json:"companyName" bson:"companyName"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
