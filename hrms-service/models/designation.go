package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Designation struct {
	ID             primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Desc           string             `json:"description,omitempty" bson:"description,omitempty"`
	Designation    string             `json:"designation,omitempty" bson:"designation,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefDesignation struct {
	Designation `bson:",inline"`
	Ref         struct {
		Organisation Organisation `json:"organisation,omitempty" bson:"organisation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDesignation struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	UniqueID       []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
