package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//OnePageAdvisory : ""
type OnePageAdvisory struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Name         string             `json:"name,omitempty"  bson:"name,omitempty"`
	Html         string             `json:"html,omitempty"  bson:"html,omitempty"`
	Desc         string             `json:"desc,omitempty"  bson:"desc,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type OnePageAdvisoryFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefOnePageAdvisory struct {
	OnePageAdvisory `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
