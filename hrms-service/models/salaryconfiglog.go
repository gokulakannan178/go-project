package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryConfigLog struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	PreSalaryConfig SalaryConfig       `json:"preSalaryConfig,omitempty" bson:"preSalaryConfig,omitempty"`
	NewSalaryConfig SalaryConfig       `json:"NewSalaryConfig,omitempty" bson:"NewSalaryConfig,omitempty"`
	Date            *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	Created         *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefSalaryConfigLog struct {
	SalaryConfigLog `bson:",inline"`
	Ref             struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterSalaryConfigLog struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeID     []string `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
