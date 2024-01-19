package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//WaterTaxArv : ""
type WaterTaxArv struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Desc      string             `json:"description" bson:"description,omitempty"`
	DOE       *time.Time         `json:"doe" bson:"doe,omitempty"`
	ARV       float64            `json:"arv" bson:"arv,omitempty"`
	RangeFrom float64            `json:"rangeFrom" bson:"rangeFrom,omitempty"`
	RangeTo   float64            `json:"rangeTo" bson:"rangeTo,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
	Created   CreatedV2          `json:"created" bson:"created,omitempty"`
}

//WaterTaxArvFilter : ""
type WaterTaxArvFilter struct {
	UniqueIDs []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status    []string `json:"status" bson:"status,omitempty"`
	Regex     struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Name     string `json:"name" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefWaterTaxArv : ""
type RefWaterTaxArv struct {
	WaterTaxArv `bson:",inline"`
	Ref         struct {
	} `json:"ref" bson:"ref,omitempty"`
}
