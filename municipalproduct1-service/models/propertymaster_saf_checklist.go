package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PropertyChecklist : ""
type PropertyChecklist struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Label      string             `json:"label" bson:"label,omitempty"`
	IsDiscount string             `json:"isDiscount" bson:"isDiscount,omitempty"`
	Value      float64            `json:"value" bson:"value,omitempty"`
	Type       string             `json:"type" bson:"type,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created" bson:"created,omitempty"`
	Updated    []Updated          `json:"updated" bson:"updated,omitempty"`
}

//RefPropertyChecklist :""
type RefPropertyChecklist struct {
	PropertyChecklist `bson:",inline"`
	Ref               struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PropertyChecklistFilter : ""
type PropertyChecklistFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

//PropertyCheckedChecklist : ""
type PropertyCheckedChecklist struct {
	UniqueID   string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Label      string    `json:"label" bson:"label,omitempty"`
	IsDiscount string    `json:"isDiscount" bson:"isDiscount,omitempty"`
	Value      float64   `json:"value" bson:"value,omitempty"`
	Type       string    `json:"type" bson:"type,omitempty"`
	Status     string    `json:"status" bson:"status,omitempty"`
	Created    Created   `json:"created" bson:"created,omitempty"`
	Updated    []Updated `json:"updated" bson:"updated,omitempty"`
}
