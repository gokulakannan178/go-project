package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetPolicy struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	AssetMasterId  []string           `json:"assertmasterId,omitempty" bson:"-"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefAssetPolicy struct {
	AssetPolicy `bson:",inline"`
	Ref         struct {
		AssetPolicyAssetsId []AssetPolicyAssets `json:"assetPolicyAssetsId,omitempty" bson:"assetPolicyAssetsId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterAssetPolicy struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
