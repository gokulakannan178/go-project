package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//BankInformation : ""
type BankInformation struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	EmployeeID    string             `json:"employeeId" bson:"employeeId,omitempty"`
	BankName      string             `json:"bankName" bson:"bankName,omitempty"`
	Branch        string             `json:"branch" bson:"branch,omitempty"`
	AccountName   string             `json:"accountName" bson:"accountName,omitempty"`
	SWIFT         string             `json:"swift" bson:"swift,omitempty"`
	IFSC          string             `json:"ifsc" bson:"ifsc,omitempty"`
	IBAN          string             `json:"iban" bson:"iban,omitempty"`
	AccountNumber string             `json:"accountNumber" bson:"accountNumber,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	Created       *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated       Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
}

//RefBankInformation :""
type RefBankInformation struct {
	BankInformation `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//BankInformationFilter : ""
type BankInformationFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	UniqueID  string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Regex     struct {
		BankName string `json:"bankName" bson:"bankName,omitempty"`
	} `json:"regex" bson:"regex"`
}
