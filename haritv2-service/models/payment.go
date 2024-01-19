package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTIONPAYMENT = "payment"
)

//Payment : ""
type Payment struct {
	ID       bson.ObjectId `json:"id,omitempty"  form:"id" bson:"_id,omitempty"`
	SaleID   string        `json:"saleId" bson:"saleId,omitempty"`
	UniqueID string        `json:"uniqueId" bson:"uniqueId,omitempty"`
	Type     string        `json:"type" bson:"type,omitempty"`
	Desc     string        `json:"desc" bson:"desc,omitempty"`
	Amount   float64       `json:"amount" bson:"amount,omitempty"`
	Name     string        `json:"name" bson:"name,omitempty"`
	//Data   IPayment `json:"data" bson:"data,omitempty"`
	//Check
	CheckNo string `json:"checkNo" bson:"checkNo,omitempty"`
	//Net Banking
	ReferenceID string     `json:"referenceId" bson:"referenceId,omitempty"`
	BankName    string     `json:"bankName" bson:"bankName,omitempty"`
	IFSCCode    string     `json:"ifscCode" bson:"ifscCode,omitempty"`
	Branch      string     `json:"branch" bson:"branch,omitempty"`
	Accno       string     `json:"accno" bson:"accno,omitempty"`
	Date        *time.Time `json:"date" bson:"date,omitempty"`
	Status      string     `json:"status" bson:"status,omitempty"`
	Created     CreatedV2  `json:"createdOn" bson:"createdOn,omitempty"`
}

//RefPayment : ""
type RefPayment struct {
	Payment `bson:",inline"`
	Ref     struct {
	} `json:"ref" bson:"ref,omitempty"`
}

//PaymentFilter : ""
type PaymentFilter struct {
	UniqueID  []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	SaleID    []string `json:"saleId" bson:"saleId,omitempty"`
	Status    []string `json:"status" bson:"status,omitempty"`
	Type      []string `json:"type" bson:"type,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

//PaymentsDone : ""
type PaymentsDone struct {
	PaymentDone                float64 `json:"paymentDone" bson:"paymentDone,omitempty"`
	PaymentVerificationPending float64 `json:"paymentVerificationPending" bson:"paymentVerificationPending,omitempty"`
	PaymentRemaining           float64 `json:"paymentRemaining" bson:"paymentRemaining,omitempty"`
}
