package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeDeductionMaster struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Title          string             `json:"title,omitempty" bson:"title,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeDeductionMaster struct {
	EmployeeDeductionMaster `bson:",inline"`
	Ref                     struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeDeductionMaster struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Title string `json:"title,omitempty" bson:"title,omitempty"`
	} `json:"regex" bson:"regex"`
}
