package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//WaterTaxConnectionType : ""
type WaterTaxConnectionType struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"description" bson:"description,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  CreatedV2          `json:"created" bson:"created,omitempty"`
}

//WaterTaxConnectionTypeFilter : ""
type WaterTaxConnectionTypeFilter struct {
	UniqueIDs []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status    []string `json:"status" bson:"status,omitempty"`
	Regex     struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Name     string `json:"name" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefWaterTaxConnectionType : ""
type RefWaterTaxConnectionType struct {
	WaterTaxConnectionType `bson:",inline"`
	Ref                    struct {
	} `json:"ref" bson:"ref,omitempty"`
}
