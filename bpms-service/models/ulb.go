package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ULB : ""
type ULB struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Address  Address            `json:"address" bson:"address,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefULB : ""
type RefULB struct {
	ULB `bson:",inline"`
	Ref struct {
		Address RefAddress `json:"address,omitempty" bson:"address,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ULBFilter : ""
type ULBFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	State     []string `json:"state,omitempty" bson:"state,omitempty"`
	District  []string `json:"district,omitempty" bson:"district,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
