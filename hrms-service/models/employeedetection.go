package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeDeduction struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	DeductionId    string             `json:"deductionId,omitempty" bson:"deductionId,omitempty"`
	Amount         float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	StartDate      *time.Time         `json:"startDate" bson:"startDate,omitempty"`
	EndDate        *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeDeduction struct {
	EmployeeDeduction `bson:",inline"`
	Ref               struct {
		OrganisationId Organisation            `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		DeductionId    EmployeeDeductionMaster `json:"deductionId,omitempty" bson:"deductionId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeDeduction struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
