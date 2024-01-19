package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SolidWasteUserChargeSubCategory : ""
type SolidWasteUserChargeSubCategory struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	CategoryID string             `json:"categoryId" bson:"categoryId,omitempty"`
	Rate       float64            `json:"rate" bson:"rate,omitempty"`
	DOE        *time.Time         `json:"doe" bson:"doe,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    CreatedV2          `json:"created" bson:"created,omitempty"`
}

//SolidWasteUserChargeSubCategoryFilter : ""
type SolidWasteUserChargeSubCategoryFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	UniqueIDs []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Regex     struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Name     string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefSolidWasteUserChargeSubCategory : ""
type RefSolidWasteUserChargeSubCategory struct {
	SolidWasteUserChargeSubCategory `bson:",inline"`
	Ref                             struct {
		Category SolidWasteUserChargeCategory `json:"category" bson:"category,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
