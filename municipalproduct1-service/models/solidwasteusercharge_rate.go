package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SolidWasteUserChargeRate : ""
type SolidWasteUserChargeRate struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	CategoryID    string             `json:"categoryId" bson:"categoryId,omitempty"`
	SubCategoryID string             `json:"subCategoryId" bson:"subCategoryId,omitempty"`
	Rate          float64            `json:"rate" bson:"rate,omitempty"`
	DOE           *time.Time         `json:"doe" bson:"doe,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
}

//SolidWasteUserChargeRateFilter : ""
type SolidWasteUserChargeRateFilter struct {
	UniqueIDs []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Regex     struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Name     string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefSolidWasteUserChargeRate : ""
type RefSolidWasteUserChargeRate struct {
	SolidWasteUserChargeRate `bson:",inline"`
	Ref                      struct {
		Category    SolidWasteUserChargeCategory    `json:"category" bson:"category,omitempty"`
		SubCategory SolidWasteUserChargeSubCategory `json:"subCategory" bson:"subCategory,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
