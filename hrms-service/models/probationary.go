package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Probationary struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name               string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID           string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationId     string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Desc               string             `json:"description,omitempty" bson:"description,omitempty"`
	ProbationaryPeroid float64            `json:"probationaryPeroid,omitempty" bson:"probationaryPeroid,omitempty"`
	ProbationaryDays   int                `json:"probationaryDays,omitempty" bson:"probationaryDays,omitempty"`
	Created            *Created           `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	Status             string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefProbationary struct {
	Probationary `bson:",inline"`
	Ref          struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterProbationary struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
