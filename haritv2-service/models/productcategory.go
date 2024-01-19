package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ProductCategory : ""
type ProductCategory struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Desc      string             `json:"desc" bson:"desc,omitempty"`
	IsDefault bool               `json:"isDefault" bson:"isDefault,omitempty"`
	Created   *CreatedV2         `json:"created" bson:"created,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
}

//RefProductCategory : ""
type RefProductCategory struct {
	ProductCategory `bson:",inline"`
	// Ref             struct {
	// } `json:"ref" bson:"ref,omitempty"`
}

//ProductCategoryFilter : ""
type ProductCategoryFilter struct {
	UniqueID         []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name             []string `json:"name" bson:"name,omitempty"`
	Status           []string `json:"status" bson:"status,omitempty"`
	ProjectionFields []string `json:"projectionFields" bson:"projectionFields,omitempty"`
	SortBy           string   `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder        int      `json:"sortOrder" bson:"sortOrder,omitempty"`
}
