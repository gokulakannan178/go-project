package models

import "time"

type OrderPayment struct {
	OrderID    string     `json:"orderId" bson:"orderId,omitempty"`
	UniqueID   string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Mop        string     `json:"mop" bson:"mop,omitempty"`
	Created    Created    `json:"created" bson:"created,omitempty"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
	TnxID      string     `json:"tnxId" bson:"tnxId,omitempty"`
	BankName   string     `json:"bankName" bson:"bankName,omitempty"`
	BranchName string     `json:"branchName" bson:"branchName,omitempty"`
	Amount     float64    `json:"amount" bson:"amount,omitempty"`
	Ifsc       string     `json:"ifsc" bson:"ifsc,omitempty"`
	Upi        string     `json:"upi" bson:"upi,omitempty"`
	Status     string     `json:"status" bson:"status,omitempty"`
	No         string     `json:"no" bson:"no,omitempty"`
	RecordDate *time.Time `json:"recordDate" bson:"recordDate,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	Payee      string     `json:"payee" bson:"payee,omitempty"`
}

//OrderPaymentFilter : ""
type OrderPaymentFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	SearchText struct {
		Payee string `json:"payee" bson:"payee,omitempty"`
	} `json:"searchText" bson:"searchText"`
	DateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
	AmountRange *struct {
		From float64 `json:"from"`
		To   float64 `json:"to"`
	} `json:"amountRange"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

//RefOrderPayment : ""
type RefOrderPayment struct {
	OrderPayment `bson:",inline"`
	// Ref     struct {
	// } `json:"ref" bson:"ref,omitempty"`
}
