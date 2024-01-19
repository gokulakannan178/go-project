package models

import "time"

type PenalChargeFYRange struct {
	PenalCharge string     `json:"penalCharge" bson:"penalCharge,omitempty"`
	UniqueID    string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	From        *time.Time `json:"from" bson:"from,omitempty"`
	To          *time.Time `json:"to" bson:"to,omitempty"`
	Status      string     `json:"status" bson:"status,omitempty"`
}

type PenalChargeFYDetails struct {
	PenalChargeStatus string `json:"penalChargeStatus" bson:"penalChargeStatus,omitempty"`
	FyID              string `json:"fyId" bson:"fyId,omitempty"`
}
type RefPenalChargeFYRange struct {
	PenalChargeFYRange `bson:",inline"`
}
