package models

import (
	"time"
)

type MajorUpdate struct {
	Title         string     `json:"title" bson:"title,omitempty"`
	Desc          string     `json:"desc" bson:"desc,omitempty"`
	RequestedBy   []string   `json:"requestedBy" bson:"requestedBy,omitempty"`
	DateOfRequest *time.Time `json:"dateOfRequest" bson:"dateOfRequest,omitempty"`
	Documents     []string   `json:"documents" bson:"documents,omitempty"`
	Status        string     `json:"status" bson:"status,omitempty"`
	UniqueID      string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Created       *CreatedV2 `json:"created" bson:"created,omitempty"`
	Updated       Updated    `json:"updated" bson:"updated,omitempty"`
}

type MajorUpdateFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefMajorUpdate struct {
	MajorUpdate `bson:",inline"`
	Ref         struct {
	} `json:"ref" bson:"ref,omitempty"`
}
