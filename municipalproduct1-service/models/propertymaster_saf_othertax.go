package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PropertyOtherTax : ""
type PropertyOtherTax struct {
	ID                  primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID            string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ResidentialDiscount struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
	} `json:"residentialDiscount" bson:"residentialDiscount,omitempty"`
	PropertyTax struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
	} `json:"propertyTax" bson:"propertyTax,omitempty"`
	SelfResidentialDiscount struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
	} `json:"selfResidentialDiscount" bson:"selfResidentialDiscount,omitempty"`
	EducationTax struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
	} `json:"educationTax" bson:"educationTax,omitempty"`
	CompositeTax struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
	} `json:"compositeTax" bson:"compositeTax,omitempty"`
	UserCharges struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
	} `json:"userCharges" bson:"userCharges,omitempty"`
	FormCharges struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
	} `json:"formCharges" bson:"formCharges,omitempty"`
	Penalty struct {
		Value float64 `json:"value" bson:"value,omitempty"`
		Type  string  `json:"type" bson:"type,omitempty"`
		On    string  `json:"on" bson:"on,omitempty"`
	} `json:"penalty" bson:"penalty,omitempty"`
	DOE     *time.Time `json:"doe" bson:"doe,omitempty"`
	Status  string     `json:"status" bson:"status,omitempty"`
	Created Created    `json:"created" bson:"created,omitempty"`
	Updated []Updated  `json:"updated" bson:"updated,omitempty"`
}

//RefPropertyOtherTax :""
type RefPropertyOtherTax struct {
	PropertyOtherTax `bson:",inline"`
	Ref              struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PropertyOtherTaxFilter : ""
type PropertyOtherTaxFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
