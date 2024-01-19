package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type Wallet struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	UserID   string             `json:"userId" bson:"userId,omitempty"`
	UserType string             `json:"userType" bson:"userType,omitempty"`
	Created  *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Amount   float64            `json:"amount" bson:"amount,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefWallet struct {
	Wallet `bson:",inline"`
	Ref    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type WalletFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	Regex     struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`

		UserID   string `json:"userId" bson:"userId,omitempty"`
		UserType string `json:"userType" bson:"userType,omitempty"`
	} `json:"regex" bson:"regex"`
}
