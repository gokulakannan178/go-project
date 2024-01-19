package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//PkgType : ""
type PkgType struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Desc      string             `json:"desc" bson:"desc,omitempty"`
	KGs       string             `json:"kgs" bson:"kgs,omitempty"`
	Created   *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	IsDefault bool               `json:"isDefault" bson:"isDefault,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
}

//PkgTypeFilter : ""
type PkgTypeFilter struct {
	Name             []string `json:"name" bson:"name,omitempty"`
	Status           []string `json:"status" bson:"status,omitempty"`
	ProjectionFields []string `json:"projectionFields" bson:"projectionFields,omitempty"`
	SortBy           string   `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder        int      `json:"sortOrder" bson:"sortOrder,omitempty"`
}

// RefPkgType : ""
type RefPkgType struct {
	PkgType `bson:",inline"`
	Ref     struct {
	} `json:"ref" bson:"ref,omitempty"`
}
