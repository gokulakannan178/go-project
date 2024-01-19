package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseSubcategory struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	UniqueID       string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Desc           string             `json:"description,omitempty" bson:"description,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefExpenseSubcategory struct {
	ExpenseSubcategory `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterExpenseSubcategory struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
