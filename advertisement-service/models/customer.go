package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type Customer struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status        string             `json:"status,omitempty" bson:"status,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Created       *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Gender        string             `json:"gender" bson:"gender,omitempty"`
	Mobile        string             `json:"mobile" bson:"mobile,omitempty"`
	Password      string             `json:"-" bson:"password,omitempty"`
	Photo         string             `json:"photo" bson:"photo,omitempty"`
	Email         string             `json:"email" bson:"email,omitempty"`
	BelongsTo     string             `json:"belongsTo" bson:"belongsTo,omitempty"`
	BelongsToType string             `json:"belongsToType" bson:"belongsToType,omitempty"`
	Address       Address            `json:"address" bson:"address,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefCustomer struct {
	Customer `bson:",inline"`
	Ref      struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type CustomerFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	Regex     struct {
		Name     string `json:"name" bson:"name"`
		Mobile   string `json:"mobile" bson:"mobile"`
		Email    string `json:"email" bson:"email,omitempty"`
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	} `json:"regex" bson:"regex"`
}
