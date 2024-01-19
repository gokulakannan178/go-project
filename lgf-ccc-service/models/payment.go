package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name               string             `json:"name,omitempty" bson:"name,omitempty"`
	Date               *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	EmployeeId         string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId     string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	BankId             string             `json:"bankId,omitempty" bson:"bankId,omitempty"`
	PaymentInformation string             `json:"paymentInformation,omitempty" bson:"paymentInformation,omitempty"`
	ModeOfPayment      ModeOfPayment      `json:"modeOfPayment,omitempty" bson:"modeOfPayment,omitempty"`
	Amount             int                `json:"amount,omitempty" bson:"amount,omitempty"`
	Created            *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated            Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	Status             string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefPayment struct {
	Payment `bson:",inline"`
	Ref     struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		//	BankId         BankInformation `json:"bankId,omitempty" bson:"bankId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterPayment struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
}

type ModeOfPayment struct {
	Cheque        string `json:"cheque,omitempty" bson:"cheque,omitempty"`
	DD            string `json:"dd,omitempty" bson:"dd,omitempty"`
	OnlinePayment string `json:"onlinePayment,omitempty" bson:"onlinePayment,omitempty"`
}
