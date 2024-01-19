package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeJob struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name               string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID           string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Office             string             `json:"office,omitempty" bson:"office,omitempty"`
	Title              string             `json:"title,omitempty" bson:"title,omitempty"`
	Date               *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	ProbationStartDate *time.Time         `json:"probationStartDate,omitempty" bson:"probationStartDate,omitempty"`
	ProbationEndDate   *time.Time         `json:"probationEndDate,omitempty" bson:"probationEndDate,omitempty"`
	ContractStartDate  *time.Time         `json:"contractStartDate,omitempty" bson:"contractStartDate,omitempty"`
	ContractEndDate    *time.Time         `json:"contractEndDate,omitempty" bson:"contractEndDate,omitempty"`
	Department         string             `json:"department,omitempty" bson:"department,omitempty"`
	OnBench            string             `json:"onBench,omitempty" bson:"onBench,omitempty"`
	Manager            string             `json:"manager,omitempty" bson:"manager,omitempty"`
	EmployeeType       string             `json:"employeeType,omitempty" bson:"employeeType,omitempty"`
	Engineer           string             `json:"engineer,omitempty" bson:"engineer,omitempty"`
	Organisation       string             `json:"organisation,omitempty" bson:"organisation,omitempty"`
	Created            Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status             string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeJob struct {
	EmployeeJob `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeJob struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
