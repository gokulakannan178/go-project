package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Salary struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty"`
	Date            *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	EmployeeId      string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId  string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	PerMonthSalary  int                `json:"perMonthSalary,omitempty" bson:"perMonthSalary,omitempty"`
	PerAnnualSalary int                `json:"perAnnualSalary,omitempty" bson:"perAnnualSalary,omitempty"`
	Created         *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated         Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefSalary struct {
	Salary `bson:",inline"`
	Ref    struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterSalary struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
}
