package models

import "time"

//TaxCollector Request
type CollectionSubmissionRequest struct {
	UniqueID        string     `json:"uniqueId,omitempty"  bson:"uniqueId,omitempty"`
	Status          string     `json:"status,omitempty" bson:"status,omitempty"`
	Date            *time.Time `json:"date,omitempty"  bson:"date,omitempty"`
	Amount          float64    `json:"amount,omitempty"  bson:"amount,omitempty"`
	Actioner        Action     `json:"actioner,omitempty"  bson:"actioner,omitempty"`
	Requestor       Action     `json:"requestor,omitempty"  bson:"requestor,omitempty"`
	PaymentReceipts []struct {
		TxnId     string `json:"txnId,omitempty"  bson:"txnId,omitempty"`
		ReceiptNo string `json:"receiptNo,omitempty"  bson:"receiptNo,omitempty"`
	} `json:"paymentReceipts,omitempty"  bson:"paymentReceipts,omitempty"`
}

type RefCollectionSubmissionRequest struct {
	CollectionSubmissionRequest `bson:",inline"`
	Ref                         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//TaxCollector Request Filter
type CollectionSubmissionRequestFilter struct {
	Status    []string  `json:"status,omitempty" bson:"status,omitempty"`
	DateRange DateRange `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	Actioner  []Action  `json:"actioner,omitempty"  bson:"actioner,omitempty"`
	Requestor []Action  `json:"requestor,omitempty"  bson:"requestor,omitempty"`
}
