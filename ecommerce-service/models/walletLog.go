package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type WalletLog struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	WalletID     string             `json:"walletID" bson:"walletID,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
	Scenario     string             `json:"scenario,omitempty" bson:"scenario,omitempty"`
	UserID       string             `json:"userId" bson:"userId,omitempty"`
	UserType     string             `json:"userType" bson:"userType,omitempty"`
	Created      *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	PrevAmount   float64            `json:"prevAmount" bson:"prevAmount,omitempty"`
	AdjustAmount float64            `json:"adjustAmount" bson:"adjustAmount,omitempty"`
	NewAmount    float64            `json:"newAmount" bson:"newAmount,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefWalletLog struct {
	WalletLog `bson:",inline"`
	Ref       struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type WalletLogFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	Regex     struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
		WalletID string `json:"walletID" bson:"walletID,omitempty"`
		UserID   string `json:"userId" bson:"userId,omitempty"`
		UserType string `json:"userType" bson:"userType,omitempty"`
	} `json:"regex" bson:"regex"`
}
