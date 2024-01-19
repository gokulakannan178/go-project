package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CompositeTaxRateMaster : ""
type CompositeTaxRateMaster struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ConstructionTypeID string             `json:"constructionTypeId" bson:"constructionTypeId,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	MinBuildUpArea     float64            `json:"minBuildUpArea" bson:"minBuildUpArea,omitempty"`
	MaxBuildUpArea     float64            `json:"maxBuildUpArea" bson:"maxBuildUpArea,omitempty"`
	Rate               float64            `json:"rate" bson:"rate,omitempty"`
	RateType           string             `json:"rateType" bson:"rateType,omitempty"`
	DOE                *time.Time         `json:"doe" bson:"doe,omitempty"`
}

//RefCompositeTaxRateMaster :""
type RefCompositeTaxRateMaster struct {
	CompositeTaxRateMaster `bson:",inline"`
	Ref                    struct {
		ConstructionType *ConstructionType `json:"constructionType" bson:"constructionType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//CompositeTaxRateMasterFilter : ""
type CompositeTaxRateMasterFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

//PanelChargeRateMaster : ""
type PanelChargeRateMaster struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Rate     float64            `json:"rate" bson:"rate,omitempty"`
	RateType string             `json:"rateType" bson:"rateType,omitempty"`
	DOE      *time.Time         `json:"doe" bson:"doe,omitempty"`
}
