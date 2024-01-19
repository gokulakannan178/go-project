package models

import "time"

type PmAchievement struct {
	PmId              string     `json:"pmId" bson:"pmId,omitempty"`
	TargetAmount      float64    `json:"targetAmount" bson:"targetAmount,omitempty"`
	AchievementAmount float64    `json:"achievementAmount" bson:"achievementAmount,omitempty"`
	Status            string     `json:"status" bson:"status,omitempty"`
	UniqueID          string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Date              *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Created           *CreatedV2 `json:"created" bson:"created,omitempty"`
	FyID              string     `json:"fyId" bson:"fyId,omitempty"`
	Month             int        `json:"month" bson:"month,omitempty"`
}

type PmAchievementFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefPmAchievement struct {
	PmAchievement `bson:",inline"`
	Ref           struct {
		FyYear FinancialYear `json:"fyYear" bson:"fyYear,omitempty"`
		Month  Month         `json:"month" bson:"month,omitempty"`
	} `json:"status" bson:"status,omitempty"`
}

type PmAchievementMonthWiseFilter struct {
	FYID string `json:"fyId" bson:"fyId,omitempty"`
	PMID string `json:"pmId" bson:"pmId,omitempty"`
}

type PmAchievementMonthWise struct {
}
