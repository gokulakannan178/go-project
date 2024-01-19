package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeDocuments struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationID   string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeID       string             `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	DocumentPolicyID string             `json:"documentPolicyID,omitempty" bson:"documentPolicyID,omitempty"`
	DocumentMasterID string             `json:"documentMasterID,omitempty" bson:"documentMasterID,omitempty"`
	Uri              string             `json:"uri,omitempty" bson:"uri,omitempty"`
	Type             string             `json:"type,omitempty" bson:"type,omitempty"`
	Created          Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
}
type RefEmployeeDocuments struct {
	EmployeeDocuments `bson:",inline"`
	Ref               struct {
		Employee       Employee       `json:"employee,omitempty" bson:"employee,omitempty"`
		DocumentPolicy DocumentPolicy `json:"documentPolicy,omitempty" bson:"documentPolicy,omitempty"`
		DocumentMaster DocumentMaster `json:"documentMaster,omitempty" bson:"documentMaster,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeDocuments struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeID     []string `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type FilterEmployeeDocumentslist struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeID string   `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	Regex      struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type EmployeeDocumentsList struct {
	EmployeeID       string `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	DocumentPolicyID string `json:"documentPolicyID,omitempty" bson:"documentPolicyID,omitempty"`
	Data             []struct {
		Name     string `json:"name,omitempty" bson:"name,omitempty"`
		File     string `json:"file" bson:"file,omitempty"`
		UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	} `json:"data,omitempty" bson:"data,omitempty"`
}
