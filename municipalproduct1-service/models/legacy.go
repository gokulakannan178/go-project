package models

import "time"

type LegacyProperty struct {
	UniqueID     string     `json:"uniqueId" bson:"uniqueId"` // auto generate
	LegacyID     string     `json:"legacyId,omitempty" bson:"legacyId,omitempty"`
	PropertyID   string     `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	ProofDoc     string     `json:"proofDoc,omitempty" bson:"proofDoc,omitempty"`
	Created      CreatedV2  `json:"created,omitempty" bson:"created,omitempty"`
	TotalTaxPaid float64    `json:"totalTaxPaid,omitempty" bson:"totalTaxPaid,omitempty"`
	Status       string     `json:"status,omitempty" bson:"status,omitempty"`
	UpdatedLog   []Updated  `json:"updatedLog,omitempty" bson:"updatedLog,omitempty"`
	ReceiptNo    string     `json:"receiptNo,omitempty" bson:"receiptNo,omitempty"`
	PaymentDate  *time.Time `json:"paymentDate,omitempty" bson:"paymentDate,omitempty"`
}
type LegacyPropertyFy struct {
	UniqueID   string  `json:"uniqueId" bson:"uniqueId"` // auto generate
	LegacyID   string  `json:"legacyId,omitempty" bson:"legacyId,omitempty"`
	PropertyID string  `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	FyID       string  `json:"fyId,omitempty" bson:"fyId"`
	TaxAmount  float64 `json:"taxAmount,omitempty" bson:"taxAmount,omitempty"`
	Status     string  `json:"status,omitempty" bson:"status,omitempty"`
}
type RegLegacyProperty struct {
	LegacyProperty   LegacyProperty     `json:"legacyProperty,omitempty" bson:"legacyProperty,omitempty"`
	LegacyPropertyFy []LegacyPropertyFy `json:"legacyPropertyFy,omitempty" bson:"legacyPropertyFy,omitempty"`
	CreatedBy        string             `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedType      string             `json:"createdType,omitempty" bson:"createdType,omitempty"`
	Updated          string             `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	UpdatedType      string             `json:"updatedType,omitempty" bson:"updatedType,omitempty"`
}

type LegacyPropertyFilter struct {
	Status      []string `json:"status,omitempty" bson:"status,omitempty"`
	LegacyID    []string `json:"legacyId,omitempty" bson:"legacyId,omitempty"`
	PropertyID  string   `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	PropertyIDs []string `json:"propertyIds,omitempty" bson:"propertyIds,omitempty"`
	UniqueID    []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	DateRange   *struct {
		From *time.Time `json:"from" bson:"from"`
		To   *time.Time `json:"to" bson:"to"`
	} `json:"dateRange" bson:"dateRange"`
	SortBy    string `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder int    `json:"sortOrder" bson:"sortOrder,omitempty"`
}

type RefLegacyPropertyPayment struct {
	IsAvailable bool `json:"isAvailable" bson:"isAvailable,omitempty"`

	RefLegacyProperty `bson:",inline"`
	LegacyPropertyFy  []RefLegacyPropertyFy `json:"legacyPropertyFy,omitempty" bson:"legacyPropertyFy,omitempty"`
}

type RefLegacyProperty struct {
	LegacyProperty `bson:",inline"`
	Ref            struct {
		User User `json:"user,omitempty" bson:"user,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefLegacyPropertyFy struct {
	LegacyPropertyFy `bson:",inline"`
	Ref              struct {
		Fy *FinancialYear `json:"fy,omitempty" bson:"fy,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefV2LegacyPropertyFy struct {
	FinancialYear `bson:",inline"`
	Legacyyear    LegacyPropertyFy `json:"legacyyear,omitempty" bson:"legacyyear,omitempty"`
}

type GetReqFinancialYear struct {
	Doa *time.Time `json:"doa,omitempty" bson:"doa,omitempty"`
}
