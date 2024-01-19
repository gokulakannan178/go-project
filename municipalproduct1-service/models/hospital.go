package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Hospital : ""
type Hospital struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Address  Address            `json:"address" bson:"address,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  CreatedV2          `json:"created" bson:"created,omitempty"`
}

//HospitalFilter : ""
type HospitalFilter struct {
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   []string `json:"status" bson:"status,omitempty"`
	Regex    struct {
		Name string `json:"name" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefHospital : ""
type RefHospital struct {
	Hospital `bson:",inline"`
	Ref      struct {
	} `json:"ref" bson:"ref,omitempty"`
}
