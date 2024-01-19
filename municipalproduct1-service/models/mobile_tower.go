package models

import (
	"time"
)

// MobileTowerDemandFYLog : ""
type MobileTowerDemandFYLog struct {
	FinancialYear `bson:",inline"`
	PropertyID    string `json:"propertyId" bson:"propertyId,omitempty"`
	MobileTowerID string `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	Status        string `json:"status" bson:"status,omitempty"`
	Details       struct {
		Tax            float64 `json:"tax" bson:"tax,omitempty"`
		Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
		Other          float64 `json:"other" bson:"other,omitempty"`
		TotalTaxAmount float64 `json:"totalTaxAmount,omitempty" bson:"totalTaxAmount,omitempty"`
	} `json:"details,omitempty" bson:"details,omitempty"`
	Ref struct {
		MobileTowerTax MobileTowerTax `json:"mobileTowerTax,omitempty" bson:"mobileTowerTax,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

// MobileTowerTax : ""
type MobileTowerTax struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string     `json:"name" bson:"name,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Value    float64    `json:"value" bson:"value,omitempty"`
	DOE      *time.Time `json:"doe" bson:"doe,omitempty"`
	Created  CreatedV2  `json:"created,omitempty" bson:"created,omitempty"`
}

type RefMobileTowerDemandFYLog struct {
	MobileTowerDemandFYLog `bson:",inline"`
}

// MobileTowerTaxFilter : ""
type MobileTowerTaxFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefMobileTowerTax struct {
	MobileTowerTax `bson:",inline"`
}
