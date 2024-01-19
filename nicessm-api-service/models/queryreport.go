package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//QueryReportFilter
type QueryReportFilter struct {
	State struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"state"  bson:"state,omitempty"`
	District struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"district"  bson:"district,omitempty"`
	Block struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"block"  bson:"block,omitempty"`
	GramPanchayat struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"village"  bson:"village,omitempty"`
	Condition string `json:"condition"  bson:"condition,omitempty"`
	// CreatedFrom   *time.Time         `json:"createdFrom" bson:"createdFrom,omitempty"`
	// CreatedTo     *time.Time         `json:"createdTo" bson:"createdTo,omitempty"`
	CreatedFrom struct {
		Condition string     `json:"condition"  bson:"condition,omitempty"`
		Date      *time.Time `json:"date"`
	} `json:"createdFrom"`
	CreatedTo struct {
		Condition string     `json:"condition"  bson:"condition,omitempty"`
		Date      *time.Time `json:"date"`
	} `json:"createdTo"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	Regex     struct {
		Query string `json:"query" bson:"query,omitempty"`
	} `json:"regex" bson:"regex"`
}
