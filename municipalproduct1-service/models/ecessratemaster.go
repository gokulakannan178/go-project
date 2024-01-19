package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//EcessRateMaster : ""
type EcessRateMaster struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"  bson:"name,omitempty"`
	UniqueID string             `json:"uniqueId,omitempty"  bson:"uniqueId,omitempty"`
	Rate     float64            `json:"rate,omitempty"  bson:"rate,omitempty"`
	DOE      *time.Time         `json:"doe,omitempty"  bson:"doe,omitempty"`
	ON       string             `json:"on,omitempty"  bson:"on,omitempty"`
}

//RefEcessRateMaster : ""
type RefEcessRateMaster struct {
	EcessRateMaster `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
