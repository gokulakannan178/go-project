package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//MonthSeason : ""
type MonthSeason struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Season       primitive.ObjectID `json:"season" form:"season," bson:"season,omitempty"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Name         string             `json:"name,omitempty"  bson:"name,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type MonthSeasonFilter struct {
	Status       []string             `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Season       []primitive.ObjectID `json:"season" form:"season," bson:"season,omitempty"`

	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefMonthSeason struct {
	MonthSeason `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
