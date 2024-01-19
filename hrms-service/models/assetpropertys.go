package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetPropertys struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationID  string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty"`
	Description     string             `json:"description,omitempty" bson:"description,omitempty"`
	AssetTypeID     string             `json:"assetTypeId,omitempty" bson:"assetTypeId,omitempty"`
	AssetID         string             `json:"assetId,omitempty" bson:"assetId,omitempty"`
	AssetPropertyId string             `json:"assetPropertyId,omitempty" bson:"assetPropertyId,omitempty"`
	Value           string             `json:"value,omitempty" bson:"value,omitempty"`
	Created         Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefAssetPropertys struct {
	AssetPropertys `bson:",inline"`
	Ref            struct {
		AssetTypeId AssetType `json:"assetTypeId,omitempty" bson:"assetTypeId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterAssetPropertys struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
