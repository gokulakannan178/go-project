package models

import (
	"time"
)

//FinancialYear : ""
type FinancialYear struct {
	// ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID              string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name                  string     `json:"name" bson:"name,omitempty"`
	Desc                  string     `json:"desc" bson:"desc,omitempty"`
	IsCurrent             bool       `json:"isCurrent" bson:"isCurrent"`
	From                  *time.Time `json:"from" bson:"from,omitempty"`
	To                    *time.Time `json:"to" bson:"to,omitempty"`
	Status                string     `json:"status" bson:"status,omitempty"`
	CommonVLR             string     `json:"commonVlr" bson:"commonVlr,omitempty"`
	Created               Created    `json:"created" bson:"created,omitempty"`
	LastDate              *time.Time `json:"lastDate" bson:"lastDate,omitempty"`
	RebateLastDate        *time.Time `json:"rebateLastDate" bson:"rebateLastDate,omitempty"`
	Updated               []Updated  `json:"updated" bson:"updated,omitempty"`
	Order                 int        `json:"order" bson:"order,omitempty"`
	PenaltyRate           float64    `json:"penaltyRate" bson:"penaltyRate,omitempty"`
	TLPenaltyRate         float64    `json:"tlPenaltyRate" bson:"tlPenaltyRate,omitempty"`
	TLPenaltyRateType     string     `json:"tlPenaltyRateType" bson:"tlPenaltyRateType,omitempty"`
	MobileTowerPenalty    float64    `json:"mobileTowerPenalty" bson:"mobileTowerPenalty,omitempty"`
	OperatingFinacialYear bool       `json:"operatingFinacialYear" bson:"operatingFinacialYear,omitempty"`
}

// SingleMonthIdentifier : ""
type SingleMonthIdentifier struct {
	FYID   string `json:"fyId" bson:"fyId,omitempty"`
	Months []int  `json:"months" bson:"months,omitempty"`
}

type SingleMonthIdentifierV2 struct {
	FYID   string   `json:"fyId" bson:"fyId,omitempty"`
	Months []string `json:"months" bson:"months,omitempty"`
}

// FinancialYearWithMonths : ""
type FinancialYearWithMonths struct {
	FinancialYear `bson:",inline"`
	Months        []Month `json:"months" bson:"months,omitempty"`
}

//RefFinancialYear :""
type RefFinancialYear struct {
	FinancialYear `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (res *RefFinancialYear) Inc(a int) int {
	return a + 1
}

//FinancialYearFilter : ""
type FinancialYearFilter struct {
	Status                []string   `json:"status"`
	UniqueID              []string   `json:"uniqueId"`
	SortBy                string     `json:"sortBy"`
	SortOrder             int        `json:"sortOrder"`
	OperatingFinacialYear []bool     `json:"operatingFinacialYear"`
	DateRange             *DateRange `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
}

type DateWiseFilter struct {
	Date *time.Time `json:"date"`
}
