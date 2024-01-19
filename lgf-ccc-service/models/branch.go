package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Branch struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Desc           string             `json:"description" bson:"description,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	Spoc           string             `json:"spoc" bson:"spoc,omitempty"`
	ContactNo      string             `json:"contactNo" bson:"contactNo,omitempty"`
	Address        *Address           `json:"address" bson:"address,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Type           string             `json:"type" bson:"type,omitempty"`
	Created        *Created           `json:"created"  bson:"created,omitempty"`
}

type RefBranch struct {
	Branch `bson:",inline"`
	Ref    struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterBranch struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	SortBy         string   `json:"sortBy"`
	SortOrder      int      `json:"sortOrder"`
	Regex          struct {
		Name      string `json:"name" bson:"name"`
		ContactNo string `json:"contactNo" bson:"contactNo"`
		Type      string `json:"type" bson:"type"`
	} `json:"regex" bson:"regex"`
}
