package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//OtherCharges : ""
type OtherCharges struct {
	ID                                                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID                                           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	OneTimeBoringChargeWithWaterConnection             float64            `json:"oneTimeBoringChargeWithWaterConnection" bson:"oneTimeBoringChargeWithWaterConnection,omitempty"`
	OneTimeBoringChargeWithoutWaterConnection          float64            `json:"oneTimeBoringChargeWithoutWaterConnection" bson:"oneTimeBoringChargeWithoutWaterConnection,omitempty"`
	PenaltyParking                                     string             `json:"penaltyParking" bson:"penaltyParking,omitempty"`
	BoringChargeParking                                string             `json:"boringChargeParking" bson:"boringChargeParking,omitempty"`
	FormFeeParking                                     string             `json:"formFeeParking" bson:"formFeeParking,omitempty"`
	FormFeeCharges                                     float64            `json:"formFeeCharges" bson:"formFeeCharges,omitempty"`
	Status                                             string             `json:"status" bson:"status,omitempty"`
	Created                                            Created            `json:"created" bson:"created,omitempty"`
	Updated                                            []Updated          `json:"updated" bson:"updated,omitempty"`
	OneTimeBoringChargeWithWaterConnectionSupplyAndOwn float64            `json:"oneTimeBoringChargeWithWaterConnectionSupplyAndOwn" bson:"oneTimeBoringChargeWithWaterConnectionSupplyAndOwn,omitempty"`
	OneTimeBoringChargeNotApplicable                   float64            `json:"oneTimeBoringChargeNotApplicable" bson:"oneTimeBoringChargeNotApplicable,omitempty"`
	OneTimeBoringChargeAlreadyPaied                    float64            `json:"oneTimeBoringChargeAlreadyPaied" bson:"oneTimeBoringChargeAlreadyPaied,omitempty"`
}

//RefOtherCharges :""
type RefOtherCharges struct {
	OtherCharges `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//OtherChargesFilter : ""
type OtherChargesFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
