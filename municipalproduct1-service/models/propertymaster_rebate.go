package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Rebate : ""
type Rebate struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created" bson:"created,omitempty"`
	Updated  []Updated          `json:"updated" bson:"updated,omitempty"`
	Rate     float64            `json:"rate" bson:"rate,omitempty"`
}

//RefRebate :""
type RefRebate struct {
	Rebate `bson:",inline"`
	Ref    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//RebateFilter : ""
type RebateFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
