package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CommodityCategory : ""
type CommodityCategory struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	Icon           string             `json:"icon,omitempty" bson:"icon,omitempty"`
	Classification string             `json:"classification,omitempty" bson:"classification,omitempty"`
	ActiveStatus   bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	//Version        string             `json:"version,omitempty" bson:"version,omitempty"`
	Created *CreatedV2 `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommodityCategoryFilter struct {
	Classification []string `json:"classification,omitempty" bson:"classification,omitempty"`
	Status         []string `json:"status" form:"status" bson:"status,omitempty"`
	SortBy         string   `json:"sortBy"`
	SortOrder      int      `json:"sortOrder"`
	SearchBox      struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefCommodityCategory struct {
	CommodityCategory `bson:",inline"`
	// Ref               struct {
	// } `json:"ref,omitempty" bson:"ref,omitempty"`
}
