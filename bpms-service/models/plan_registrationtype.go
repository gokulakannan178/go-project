package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//PlanRegistrationType : ""
type PlanRegistrationType struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefPlanRegistrationType : ""
type RefPlanRegistrationType struct {
	PlanRegistrationType `bson:",inline"`
	Ref                  struct {
		Organisation *ULB        `json:"organisation,omitempty" bson:"organisation,omitempty"`
		Address      *RefAddress `json:"address,omitempty" bson:"address,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PlanRegistrationTypeFilter : ""
type PlanRegistrationTypeFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	Organisation []string `json:"organisation,omitempty" organisation:"status,omitempty"`
	SortBy       string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder    int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
