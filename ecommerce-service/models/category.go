package models

import "time"

// Category : ""
type Category struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Created  *CreatedV2 `json:"created" bson:"created,omitempty"`
	Name     string     `json:"name" bson:"name,omitempty"`
	Desc     string     `json:"desc,omitempty"  bson:"desc,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	IMG      string     `json:"img" bson:"img,omitempty"`
}

// CategoryFilter : ""
type CategoryFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	UniqueID   []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	SearchText struct {
		Name     string `json:"name" bson:"name,omitempty"`
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	} `json:"searchText" bson:"searchText"`
	DateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
}

type RefCategory struct {
	Category `bson:",inline"`
}
