package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FarmerReportFilter
type FarmerReportFilter struct {
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
	Education struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"education" bson:"education,omitempty"`
	Asset struct {
		Condition string             `json:"condition"  bson:"condition,omitempty"`
		ID        primitive.ObjectID `json:"id"  bson:"id,omitempty"`
	} `json:"asset" bson:"asset,omitempty"`
	YearlyIncome struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"yearlyIncome" bson:"yearlyIncome,omitempty"`
	Age struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        int    `json:"id"  bson:"id,omitempty"`
	} `json:"age" bson:"age,omitempty"`
	CreatedDate struct {
		Condition string     `json:"condition"  bson:"condition,omitempty"`
		ID        *time.Time `json:"id"  bson:"id,omitempty"`
	} `json:"createdDate" bson:"createdDate,omitempty"`
	VoiceSmsStatus struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        bool   `json:"id"  bson:"id,omitempty"`
	} `json:"voiceSmsStatus" bson:"voiceSmsStatus,omitempty"`
	Gender struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        string `json:"id"  bson:"id,omitempty"`
	} `json:"gender" bson:"gender,omitempty"`
	SmsStatus struct {
		Condition string `json:"condition"  bson:"condition,omitempty"`
		ID        bool   `json:"id"  bson:"id,omitempty"`
	} `json:"smsStatus" bson:"smsStatus,omitempty"`
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
