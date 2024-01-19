package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserReportFilter
type UserReportFilter struct {
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
	Language struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"language" bson:"language,omitempty"`
	Role struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"role" bson:"role,omitempty"`
	UserName struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"userName" bson:"userName,omitempty"`
	Experience struct {
		Condition string  `json:"condition"  bson:"condition,omitempty"`
		ID        float64 `json:"id"  bson:"id,omitempty"`
	} `json:"experience" bson:"experience,omitempty"`
	Gender struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"gender" bson:"gender,omitempty"`
	AccessLevel struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"accessLevel" bson:"accessLevel,omitempty"`
	Organisation struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"organisation" bson:"organisation,omitempty"`
	Project struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"project" bson:"project,omitempty"`
	Condition   string `json:"condition"  bson:"condition,omitempty"`
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
