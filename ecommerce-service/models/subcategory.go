package models

import "time"

// SubCategory : ""
type SubCategory struct {
	UniqueID   string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	CategoryID string     `json:"categoryId" bson:"categoryId,omitempty"`
	Name       string     `json:"name" bson:"name,omitempty"`
	Desc       string     `json:"desc,omitempty"  bson:"desc,omitempty"`
	Status     string     `json:"status" bson:"status,omitempty"`
	IMG        string     `json:"img" bson:"img,omitempty"`
	Created    *CreatedV2 `json:"created" bson:"created,omitempty"`
}

// SubCategoryFilter : ""
type SubCategoryFilter struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	UniqueID   []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	CategoryID []string `json:"categoryId" bson:"categoryId,omitempty"`
	SearchText struct {
		Name     string `json:"name" bson:"name,omitempty"`
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	} `json:"searchText" bson:"searchText"`
	DateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
}

type RefSubCategory struct {
	SubCategory `bson:",inline"`
	Ref         struct {
		Category *Category `json:"category,omitempty" bson:"category,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
