package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeEarning struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	EarningId      string             `json:"earningId,omitempty" bson:"earningId,omitempty"`
	Amount         float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	StartDate      *time.Time         `json:"startDate" bson:"startDate,omitempty"`
	EndDate        *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeEarning struct {
	EmployeeEarning `bson:",inline"`
	Ref             struct {
		OrganisationId Organisation          `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		EarningId      EmployeeEarningMaster `json:"earningId,omitempty" bson:"earningId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeEarning struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
