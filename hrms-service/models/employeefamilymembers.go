package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeFamilyMembers struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeID     string             `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Relationship   string             `json:"relationship,omitempty" bson:"relationship,omitempty"`
	DOB            *time.Time         `json:"dob,omitempty" bson:"dob,omitempty"`
	Phone          string             `json:"phone,omitempty" bson:"phone,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeFamilyMembers struct {
	EmployeeFamilyMembers `bson:",inline"`
	Ref                   struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeFamilyMembers struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeID     []string `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
