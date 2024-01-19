package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeHistory struct {
	ID         primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name       string             `json:"name" bson:"name"`
	FromDate   *time.Time         `json:"fromDate,omitempty" bson:"fromDate,omitempty"`
	ToDate     *time.Time         `json:"toDate,omitempty" bson:"toDate,omitempty"`
	Message    string             `json:"message,omitempty" bson:"message,omitempty"`
	EmployeeId primitive.ObjectID `json:"EmployeeId" form:"EmployeeId," bson:"EmployeeId,omitempty"`
	Created    Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeHistory struct {
	EmployeeHistory `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeHistory struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
