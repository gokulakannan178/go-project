package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PropertyOtherDemand : ""
type PropertyOtherDemand struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID            string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID          string             `json:"propertyId" bson:"propertyId,omitempty"`
	FyID                string             `json:"fyId" bson:"fyId,omitempty"`
	Reason              string             `json:"reason" bson:"reason,omitempty"`
	OneTimePenalCharges string             `json:"oneTimePenalCharges" bson:"oneTimePenalCharges,omitempty"`
	Amount              float64            `json:"amount" bson:"amount,omitempty"`
	PaymentStatus       string             `json:"paymentStatus" bson:"paymentStatus,omitempty"`
	Status              string             `json:"status" bson:"status,omitempty"`
	Created             CreatedV2          `json:"created" bson:"created,omitempty"`
}

//PropertyOtherDemandFilter : ""
type PropertyOtherDemandFilter struct {
	UniqueIDs     []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	FyID          []string `json:"fyId" bson:"fyId,omitempty"`
	PropertyID    []string `json:"propertyId" bson:"propertyId,omitempty"`
	Status        []string `json:"status" bson:"status,omitempty"`
	PaymentStatus []string `json:"paymentStatus" bson:"paymentStatus,omitempty"`
	Regex         struct {
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefPropertyOtherDemand : ""
type RefPropertyOtherDemand struct {
	PropertyOtherDemand `bson:",inline"`
	Ref                 struct {
		FY FinancialYear `json:"fy" bson:"fy,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
