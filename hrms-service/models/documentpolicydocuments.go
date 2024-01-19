package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentPolicyDocuments struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	DocumentPolicyID string             `json:"documentPolicyID,omitempty" bson:"documentPolicyID,omitempty"`
	DocumentMasterID string             `json:"documentMasterID,omitempty" bson:"documentMasterID,omitempty"`
	Message          string             `json:"message,omitempty" bson:"message,omitempty"`
	Remarks          string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Created          Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefDocumentPolicyDocuments struct {
	DocumentPolicyDocuments `bson:",inline"`
	Ref                     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDocumentPolicyDocuments struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
