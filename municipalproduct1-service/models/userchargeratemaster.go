package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserChargeRateMaster : ""
type UserChargeRateMaster struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	CategoryID string             `json:"categoryId" bson:"categoryId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Rate       float64            `json:"rate" bson:"rate,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	DOE        *time.Time         `json:"doe" bson:"doe,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created" bson:"created,omitempty"`
	Updated    []Updated          `json:"updated" bson:"updated,omitempty"`
}

//RefUserChargeRateMaster :""
type RefUserChargeRateMaster struct {
	UserChargeRateMaster `bson:",inline"`
	Ref                  struct {
		Category UserChargeCategory `json:"category" bson:"category,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserChargeRateMasterFilter : ""
type UserChargeRateMasterFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
