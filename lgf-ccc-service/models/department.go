package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Department struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Desc           string             `json:"description" bson:"description,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	Branch         string             `json:"branch" bson:"branch,omitempty"`
	Type           string             `json:"type" bson:"type,omitempty"`
	Created        *Created           `json:"created"  bson:"created,omitempty"`
}

type RefDepartment struct {
	Department `bson:",inline"`
	Ref        struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDepartment struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	BranchId       []string `json:"branchId" bson:"branchId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
