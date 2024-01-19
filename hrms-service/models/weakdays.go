package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DayOfWeek struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	DayOfWeek int64              `json:"dayOfWeek" bson:"dayOfWeek,omitempty"`
	Created   Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefDayOfWeek struct {
	DayOfWeek `bson:",inline"`
	Ref       struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDayOfWeek struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
