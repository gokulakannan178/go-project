package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetType struct {
	ID                   primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID             string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name                 string             `json:"name" bson:"name,omitempty"`
	Description          string             `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationID       string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	AssetTypePropertysId []string           `json:"assetTypePropertysId,omitempty" bson:"-"`
	Status               string             `json:"status,omitempty" bson:"status,omitempty"`
	Created              Created            `json:"createdOn" bson:"createdOn,omitempty"`
}

type RefAssetType struct {
	AssetType `bson:",inline"`
	Ref       struct {
		Organisation         Organisation         `json:"organisation,omitempty" bson:"organisation,omitempty"`
		AssetTypePepropertys []AssetTypePropertys `json:"assetTypePepropertys,omitempty" bson:"assetTypePepropertys,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterAssetType struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
type UpdateAssetType struct {
	AssetType          `bson:",inline"`
	AssetTypePropertys []AssetTypePropertys `json:"assetTypePropertys,omitempty" bson:"assetTypePropertys,omitempty"`
}
