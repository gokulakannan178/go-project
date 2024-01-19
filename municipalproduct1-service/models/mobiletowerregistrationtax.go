package models

import "time"

type MobileTowerRegistrationTax struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name      string     `json:"name" bson:"name,omitempty"`
	IsDefault bool       `json:"isDefault" bson:"isDefault,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	Value     float64    `json:"value" bson:"value,omitempty"`
	DOE       *time.Time `json:"doe" bson:"doe,omitempty"`
	Created   CreatedV2  `json:"created,omitempty" bson:"created,omitempty"`
}

// MobileTowerTaxFilter : ""
type MobileTowerRegistrationTaxFilter struct {
	UniqueID []string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	Status   []string `json:"status" bson:"status,omitempty"`
}

// RefMobileTowerRegistrationTax : ""
type RefMobileTowerRegistrationTax struct {
	MobileTowerRegistrationTax `bson:",inline"`
	Ref                        struct {
	} `json:"ref" bson:"ref"`
}
