package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentMuxMaster struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	DocumentTypeID     string             `json:"documentTypeID,omitempty" bson:"documentTypeID,omitempty"`
	DocumentScenarioID string             `json:"documentScenarioID,omitempty" bson:"documentScenarioID,omitempty"`
	Name               string             `json:"name" bson:"name,omitempty"`
	Remarks            string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Message            string             `json:"message,omitempty" bson:"message,omitempty"`
	Created            Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status             string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefDocumentMuxMaster struct {
	DocumentMuxMaster `bson:",inline"`
	Ref               struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDocumentMuxMaster struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
