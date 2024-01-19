package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//BannedItem : ""
type BannedItem struct {
	ID      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"  bson:"name,omitempty"`
	Type    string             `json:"type,omitempty"  bson:"type,omitempty"`
	Desc    string             `json:"description" bson:"description,omitempty"`
	Version int                `json:"version,omitempty"  bson:"version,omitempty"`
	Status  string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type BannedItemFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefBannedItem struct {
	BannedItem `bson:",inline"`
	Ref        struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
