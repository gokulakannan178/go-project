package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ConstructionType : ""
type ConstructionType struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Label    string             `json:"label" bson:"label,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created" bson:"created,omitempty"`
	Updated  []Updated          `json:"updated" bson:"updated,omitempty"`
}

//RefConstructionType :""
type RefConstructionType struct {
	ConstructionType `bson:",inline"`
	Ref              struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ConstructionTypeFilter : ""
type ConstructionTypeFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
