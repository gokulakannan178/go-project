package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetTypePropertys struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	AssetTypeID    string             `json:"assetTypeId,omitempty" bson:"assetTypeId,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefAssetTypePropertys struct {
	AssetTypePropertys `bson:",inline"`
	Ref                struct {
		Organisation Organisation `json:"organisation,omitempty" bson:"organisation,omitempty"`
		AssetTypeId  AssetType    `json:"assetTypeId,omitempty" bson:"assetTypeId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterAssetTypePropertys struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	AssetTypeIds   []string `json:"assetTypeIds,omitempty" bson:"assetTypeIds,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
