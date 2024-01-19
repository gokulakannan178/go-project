package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//OrderPayments : ""
type OrderPayment struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	PayeeName     string             `json:"payeeName"  bson:"payeeName,omitempty"`
	OrderID       primitive.ObjectID `json:"orderId" bson:"orderId,omitempty"`
	UniqueID      primitive.ObjectID `json:"uniqueId" bson:"uniqueId,omitempty"`
	Amount        float64            `json:"amount" bson:"amount,omitempty"`
	Status        string             `json:"status"  bson:"status,omitempty"`
	AccountNumber float64            `json:"accountNumber"  bson:"accountNumber,omitempty"`
	TaxID         primitive.ObjectID `json:"taxId" bson:"taxId,omitempty"`
	Bank          string             `json:"bank"  bson:"bank,omitempty"`
	Branch        string             `json:"branch"  bson:"branch,omitempty"`
	IFSCCode      string             `json:"ifscCode"  bson:"ifscCode,omitempty"`
	MOP           string             `json:"mop"  bson:"mop,omitempty"`
	Date          *time.Time         `json:"date"  bson:"date,omitempty"`
	Created       *Created           `json:"created"  bson:"created,omitempty"`
}

type OrderPaymentFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	SearchBox struct {
		PayeeName string `json:"payeeName"  bson:"payeeName,omitempty"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefOrderPayment struct {
	OrderPayment `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
