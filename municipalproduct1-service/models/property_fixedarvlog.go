package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PropertyFixedArvLog : ""
type PropertyFixedArvLog struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID string             `json:"propertyId" bson:"propertyId,omitempty"`
	FyID       string             `json:"fyId" bson:"fyId,omitempty"`
	ARV        float64            `json:"arv" bson:"arv,omitempty"`
	Tax        float64            `json:"tax" bson:"tax,omitempty"`
	Total      float64            `json:"total" bson:"total,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    CreatedV2          `json:"created" bson:"created,omitempty"`
}

//PropertyFixedArvLogFilter : ""
type PropertyFixedArvLogFilter struct {
	UniqueIDs   []string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID  []string   `json:"propertyId" bson:"propertyId,omitempty"`
	Status      []string   `json:"status" bson:"status,omitempty"`
	FyID        []string   `json:"fyId" bson:"fyId,omitempty"`
	ARV         []string   `json:"arv" bson:"arv,omitempty"`
	CreatedBy   []string   `json:"createdBy" bson:"createdBy,omitempty"`
	CreatedDate *DateRange `json:"createdDate"`

	Regex struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Name     string `json:"name" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefPropertyFixedArvLog : ""
type RefPropertyFixedArvLog struct {
	PropertyFixedArvLog `bson:",inline"`
	Ref                 struct {
	} `json:"ref" bson:"ref,omitempty"`
}
