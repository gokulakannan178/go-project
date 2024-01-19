package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AVRRange : ""
type AVRRange struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	From     float64            `json:"from" bson:"from,omitempty"`
	To       float64            `json:"to" bson:"to,omitempty"`
	Rate     float64            `json:"rate" bson:"rate,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	DOE      *time.Time         `json:"doe" bson:"doe,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created" bson:"created,omitempty"`
	Updated  []Updated          `json:"updated" bson:"updated,omitempty"`
}

//RefAVRRange :""
type RefAVRRange struct {
	AVRRange `bson:",inline"`
	Ref      struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//AVRRangeFilter : ""
type AVRRangeFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
