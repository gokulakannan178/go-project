package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetPolicyAssets struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	AssetPolicyID string             `json:"assetPolicyId,omitempty" bson:"assetPolicyId,omitempty"`
	AssetMasterID string             `json:"assetMasterId,omitempty" bson:"assetMasterId,omitempty"`
	Message       string             `json:"message,omitempty" bson:"message,omitempty"`
	Remarks       string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Created       Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status        string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefAssetPolicyAssets struct {
	AssetPolicyAssets `bson:",inline"`
	Ref               struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterAssetPolicyAssets struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
