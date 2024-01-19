package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserChargeLog : ""
type UserChargeLog struct {
	ID               primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyId       string             `json:"propertyId" bson:"propertyId,omitempty"`
	BeforeUserCharge UserCharge         `json:"beforeUserCharge" bson:"beforeUserCharge,omitempty"`
	AfterUserCharge  UserCharge         `json:"afterUserCharge" bson:"afterUserCharge,omitempty"`
	Date             *time.Time         `json:"date" bson:"date,omitempty"`
	Status           string             `json:"status" bson:"status,omitempty"`
	Created          Created            `json:"created" bson:"created,omitempty"`
	Updated          []Updated          `json:"updated" bson:"updated,omitempty"`
}

//RefUserChargeLog :""
type RefUserChargeLog struct {
	UserChargeLog `bson:",inline"`
	Ref           struct {
		Category UserChargeCategory `json:"category" bson:"category,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserChargeLogFilter : ""
type UserChargeLogFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
