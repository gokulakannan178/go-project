package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmergencyContact struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID     string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	FullName     string             `json:"fullName" bson:"fullName,omitempty"`
	Relationship string             `json:"relationship,omitempty" bson:"relationship,omitempty"`
	PhoneNumber  string             `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Created      Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmergencyContact struct {
	EmergencyContact `bson:",inline"`
	Ref              struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmergencyContact struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		FullName string `json:"fullName" bson:"fullName"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
