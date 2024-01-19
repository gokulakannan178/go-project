package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FloorRatableArea : ""
type FloorRatableArea struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created" bson:"created,omitempty"`
	Updated  []Updated          `json:"updated" bson:"updated,omitempty"`
	Rate     float64            `json:"rate" bson:"rate,omitempty"`
	RateType string             `json:"rateType" bson:"rateType,omitempty"`
	DOE      *time.Time         `json:"doe" bson:"doe,omitempty"`
}

//RefFloorRatableArea :""
type RefFloorRatableArea struct {
	FloorRatableArea `bson:",inline"`
	Ref              struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//FloorRatableAreaFilter : ""
type FloorRatableAreaFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
