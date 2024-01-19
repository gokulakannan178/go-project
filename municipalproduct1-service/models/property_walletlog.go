package models

import "time"

//PropertyWalletLog : ""
type PropertyWalletLog struct {
	UniqueID      string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	WalletID      string     `json:"walletId" bson:"walletId,omitempty"`
	Message       string     `json:"message" bson:"message,omitempty"`
	PropertyID    string     `json:"propertyId" bson:"propertyId,omitempty"`
	Scenario      string     `json:"scenario" bson:"scenario,omitempty"`
	OwnerName     string     `json:"ownerName" bson:"ownerName,omitempty"`
	Amount        float64    `json:"amount" bson:"amount,omitempty"`
	PostTnxAmount float64    `json:"postTnxAmount" bson:"postTnxAmount,omitempty"`
	PreTnxAmount  float64    `json:"preTnxAmount" bson:"preTnxAmount,omitempty"`
	Date          *time.Time `json:"date" bson:"date,omitempty"`
	Status        string     `json:"status" bson:"status,omitempty"`
	Created       *CreatedV2 `json:"created,omitempty"  bson:"created,omitempty"`
	MobileNo      string     `json:"mobileNo" bson:"mobileNo,omitempty"`
}

//PropertyWalletLogFilter : ""
type PropertyWalletLogFilter struct {
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	Scenario   []string `json:"scenario,omitempty" bson:"scenario,omitempty"`
	CreatedBy  []string `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	PropertyID []string `json:"propertyId" bson:"propertyId,omitempty"`
	SearchText struct {
		UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
		WalletID string `json:"walletId,omitempty" bson:"walletId,omitempty"`
		MobileNo string `json:"mobileNo" bson:"mobileNo,omitempty"`
	} `json:"searchText"`
	PreTnxAmountDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"preTnxAmountDateRange"`
	PostTnxAmountDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"postTnxAmountDateRange"`
	DateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
}

//RefPropertyWalletLog : ""
type RefPropertyWalletLog struct {
	PropertyWalletLog `bson:",inline"`
}
