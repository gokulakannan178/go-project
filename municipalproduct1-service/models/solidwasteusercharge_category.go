package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SolidWasteUserChargeCategory : ""
type SolidWasteUserChargeCategory struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Rate     float64            `json:"rate" bson:"rate,omitempty"`
	DOE      *time.Time         `json:"doe" bson:"doe,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  CreatedV2          `json:"created" bson:"created,omitempty"`
}

//SolidWasteUserChargeCategoryFilter : ""
type SolidWasteUserChargeCategoryFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	UniqueIDs []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Regex     struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Name     string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefSolidWasteUserChargeCategory : ""
type RefSolidWasteUserChargeCategory struct {
	SolidWasteUserChargeCategory `bson:",inline"`
	Ref                          struct {
	} `json:"ref" bson:"ref,omitempty"`
}
