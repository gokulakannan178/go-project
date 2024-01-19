package models

import "time"

type ParkingPenalty struct {
	PropertyID string     `json:"propertyId" bson:"propertyIds,omitempty"`
	From       *time.Time `json:"from" bson:"from,omitempty"`
	To         *time.Time `json:"to" bson:"to,omitempty"`
	Status     string     `json:"status" bson:"status,omitempty"`
	UniqueID   string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Created    *CreatedV2 `json:"created" bson:"created,omitempty"`
	Updated    Updated    `json:"updated" bson:"updated,omitempty"`
}

type ParkingPenaltyFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefParkingPenalty struct {
	ParkingPenalty `bson:",inline"`
}
