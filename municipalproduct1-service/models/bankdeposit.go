package models

import "time"

type BankDeposit struct {
	Status   string `json:"status,omitempty" bson:"status,omitempty"`
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Usertype string `json:"usertype,omitempty" bson:"usertype,omitempty"`
	Bank     struct {
		Name   string `json:"name,omitempty" bson:"name,omitempty"`
		IFSC   string `json:"ifsc,omitempty" bson:"ifsc,omitempty"`
		Branch string `json:"branch,omitempty" bson:"branch,omitempty"`
	}
	RefTnxID string              `json:"refTnxId,omitempty" bson:"refTnxId,omitempty"`
	Proof    string              `json:"proof,omitempty" bson:"proof,omitempty"`
	TnxIDs   []string            `json:"tnxids,omitempty" bson:"tnxids,omitempty"`
	Amount   float64             `json:"amount,omitempty" bson:"amount,omitempty"`
	Verifier BankDepositVerifier `json:"verifier,omitempty" bson:"verifier,omitempty"`
	On       *time.Time          `json:"on,omitempty" bson:"on,omitempty"`
}
type BankDepositVerifier struct {
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	UserName string     `json:"userName,omitempty"  bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty"  bson:"userType,omitempty"`
	Remarks  string     `json:"remarks" bson:"remarks,omitempty"`
}
type NotBankDepositVerifier struct {
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	UserName string     `json:"userName,omitempty"  bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty"  bson:"userType,omitempty"`
	Remarks  string     `json:"remarks" bson:"remarks,omitempty"`
}
type BankDepositFilter struct {
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	UserName []string `json:"userName,omitempty"  bson:"userName,omitempty"`
	UserType []string `json:"userType,omitempty"  bson:"userType,omitempty"`
	Status   []string `json:"status,omitempty" bson:"status,omitempty"`
}
