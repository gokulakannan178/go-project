package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentPolicy struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Description      string             `json:"description,omitempty" bson:"description,omitempty"`
	DocumentMasterId []string           `json:"documentmasterId,omitempty" bson:"-"`
	OrganisationID   string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created          Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefDocumentPolicy struct {
	DocumentPolicy `bson:",inline"`
	Ref            struct {
		DocumentMaster []DocumentMaster `json:"documentMaster,omitempty" bson:"documentMaster,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDocumentPolicy struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
