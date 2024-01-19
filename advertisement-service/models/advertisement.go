package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type Advertisement struct {
	ID         primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	CustomerId primitive.ObjectID `json:"customerId,omitempty" bson:"customerId,omitempty"`
	Project    primitive.ObjectID `json:"project,omitempty" bson:"project,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Size       string             `json:"size" bson:"size,omitempty"`
	Position   string             `json:"position" bson:"position,omitempty"`
	Created    *Created           `json:"createdOn" bson:"createdOn,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefAdvertisement struct {
	Advertisement `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type AdvertisementFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
