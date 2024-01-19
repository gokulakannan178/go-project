package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//OccupancyType : ""
type OccupancyType struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created" bson:"created,omitempty"`
	Updated  []Updated          `json:"updated" bson:"updated,omitempty"`
	Factor   float64            `json:"factor" bson:"factor,omitempty"`
	DOA      *time.Time         `json:"doa" bson:"doa,omitempty"`
}

//RefOccupancyType :""
type RefOccupancyType struct {
	OccupancyType `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//OccupancyTypeFilter : ""
type OccupancyTypeFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}
