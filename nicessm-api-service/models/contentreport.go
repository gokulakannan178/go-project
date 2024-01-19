package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ContentReportFilter
type ContentReportFilter struct {
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
	KnowledgeDomain struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
	SubDomain struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"subDomain" bson:"subDomain,omitempty"`
	SubTopic struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"subTopic" bson:"subTopic,omitempty"`
	Topic struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"topic" bson:"topic,omitempty"`
	Classfication struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"classfication" bson:"classfication,omitempty"`
	Season struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"season" bson:"season,omitempty"`
	Market struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"market" bson:"market,omitempty"`
	Soil_type struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"soil_type" bson:"soil_type,omitempty"`
	Status struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"status" bson:"status,omitempty"`
	Commodity struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"commodity" bson:"commodity,omitempty"`
	Type struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"type"`
	Organisation struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"organisation" bson:"organisation,omitempty"`
	Project struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"project" bson:"project,omitempty"`
	CreatedFrom struct {
		Condition string     `json:"condition"  bson:"condition,omitempty"`
		Date      *time.Time `json:"date"`
	} `json:"createdFrom"`
	CreatedTo struct {
		Condition string     `json:"condition"  bson:"condition,omitempty"`
		Date      *time.Time `json:"date"`
	} `json:"createdTo"`
	Condition string `json:"condition"  bson:"condition,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	Regex     struct {
		Query string `json:"query" bson:"query,omitempty"`
	} `json:"regex" bson:"regex"`
}
