package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//DocumentList : "Holds single DocumentList data"
type PropertyDocuments struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Url          []string           `json:"url,omitempty"  bson:"url,omitempty"`
	Uri          string             `json:"uri,omitempty"  bson:"uri,omitempty"`
	DocumentType string             `json:"documentType,omitempty" bson:"documentType,omitempty"`
	PropertyID   string             `json:"propertyId,omitempty"  bson:"propertyId,omitempty"`
	DocumentID   string             `json:"documentId,omitempty"  bson:"documentId,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      Created            `json:"created,omitempty"  bson:"created,omitempty"`
}

//RefDocumentList : ""
type RefPropertyDocuments struct {
	PropertyDocuments `bson:",inline"`
	Ref               struct {
		DocumentID DocumentList `json:"documentId,omitempty"  bson:"documentId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DocumentListFilter : "Used for constructing filter query"
type PropertyDocumentsFilter struct {
	PropertyID []string `json:"propertyId,omitempty"  bson:"propertyId,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy     string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder  int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
