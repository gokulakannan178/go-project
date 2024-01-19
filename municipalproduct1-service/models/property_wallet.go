package models

import "time"

//PropertyWallet : ""
type PropertyWallet struct {
	OwnerName     string     `json:"ownerName" bson:"ownerName,omitempty"`
	UniqueID      string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID    string     `json:"propertyId" bson:"propertyId,omitempty"`
	Amount        float64    `json:"amount" bson:"amount,omitempty"`
	BalanceAmount float64    `json:"balanceAmount" bson:"balanceAmount,omitempty"`
	Status        string     `json:"status" bson:"status,omitempty"`
	Created       *CreatedV2 `json:"created,omitempty"  bson:"created,omitempty"`
	MobileNo      string     `json:"mobileNo" bson:"mobileNo,omitempty"`
}

//PropertyWalletFilter : ""
type PropertyWalletFilter struct {
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	PropertyID []string `json:"propertyId" bson:"propertyId,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	SearchText struct {
		OwnerName string `json:"ownerName" bson:"ownerName,omitempty"`
		MobileNo  string `json:"mobileNo" bson:"mobileNo,omitempty"`
		UniqueID  string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	} `json:"searchText"`
	Amount *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"amount"`
}

//RefPropertyWallet : ""
type RefPropertyWallet struct {
	PropertyWallet `bson:",inline"`
	Ref            struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
