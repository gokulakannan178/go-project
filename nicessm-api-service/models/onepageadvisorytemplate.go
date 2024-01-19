package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//OnePageAdvisoryTemplate : ""
type OnePageAdvisoryTemplate struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Name         string             `json:"name,omitempty"  bson:"name,omitempty"`
	Html         string             `json:"html,omitempty"  bson:"html,omitempty"`
	Desc         string             `json:"desc,omitempty"  bson:"desc,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type OnePageAdvisoryTemplateFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefOnePageAdvisoryTemplate struct {
	OnePageAdvisoryTemplate `bson:",inline"`
	Ref                     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
