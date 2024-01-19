package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeaveMaster struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	LeaveType      string             `json:"leaveType,omitempty" bson:"leaveType,omitempty"`
	Value          int64              `json:"value,omitempty" bson:"value,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefLeaveMaster struct {
	LeaveMaster `bson:",inline"`
	Ref         struct {
		OrganisationID Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterLeaveMaster struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`

	Regex struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
