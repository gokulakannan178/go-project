package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//DocumentList : "Holds single DocumentList data"
type DocumentList struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty"  bson:"name,omitempty"`
	UniqueID     string             `json:"uniqueId,omitempty"  bson:"uniqueId,omitempty"`
	DocumentType string             `json:"documentType,omitempty"  bson:"documentType,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated      []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefDocumentList : ""
type RefDocumentList struct {
	DocumentList `bson:",inline"`
	Ref          struct {
		Name string `json:"name,omitempty"  bson:"name,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DocumentListFilter : "Used for constructing filter query"
type DocumentListFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
