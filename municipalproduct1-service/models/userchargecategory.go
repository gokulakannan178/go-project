package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserChargeCategory : ""
type UserChargeCategory struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Rate     float64            `json:"rate" bson:"rate,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	DOE      *time.Time         `json:"doe" bson:"doe,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created" bson:"created,omitempty"`
	Updated  []Updated          `json:"updated" bson:"updated,omitempty"`
}

//RefUserChargeCategory :""
type RefUserChargeCategory struct {
	UserChargeCategory `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserChargeCategoryFilter : ""
type UserChargeCategoryFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

type UserChargeDemand struct {
	UniqueID string                  `json:"uniqueId" bson:"uniqueId,omitempty"`
	UCDemand UserChargeDemandSummary `json:"ucdemand" bson:"ucdemand,omitempty"`
	Ref      struct {
		Fy []UserChargeDemandFY `json:"fy" bson:"fy,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

type UserChargeDemandFY struct {
	FinancialYear     `bson:",inline"`
	FyMonth           []UserChargeDemandFyMonth `json:"fymonth" bson:"fymonth,omitempty"`
	Tax               float64                   `json:"Tax" bson:"Tax,omitempty"`
	CalculatedPenalty float64                   `json:"calculatedPenalty" bson:"calculatedPenalty,omitempty"`
	Total             float64                   `json:"total" bson:"total,omitempty"`
	ToBePaid          float64                   `json:"toBePaid" bson:"toBePaid,omitempty"`
	Alreadypaid       float64                   `json:"alreadypaid" bson:"alreadypaid,omitempty"`
	TotalTaxToBePaid  float64                   `json:"totalTaxToBePaid" bson:"totalTaxToBePaid,omitempty"`
}

type UserChargeDemandFyMonth struct {
	Month   `bson:",inline"`
	From    *time.Time `json:"from" bson:"from,omitempty"`
	To      *time.Time `json:"to" bson:"to,omitempty"`
	FYOrder int        `json:"fyOrder" bson:"fyOrder,omitempty"`
	Name    string     `json:"name" bson:"name,omitempty"`
	Rate    struct {
		Rate float64    `json:"rate" bson:"rate,omitempty"`
		DOE  *time.Time `json:"doe" bson:"doe,omitempty"`
	} `json:"rate" bson:"rate,omitempty"`
	CurrentPenalty      float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
	Tax                 float64 `json:"Tax" bson:"Tax,omitempty"`
	Penalty             float64 `json:"penalty" bson:"penalty,omitempty"`
	Total               float64 `json:"total" bson:"total,omitempty"`
	ToBePaid            float64 `json:"toBePaid" bson:"toBePaid,omitempty"`
	Alreadypaid         float64 `json:"alreadypaid" bson:"alreadypaid,omitempty"`
	TotalTaxToBePaid    float64 `json:"totalTaxToBePaid" bson:"totalTaxToBePaid,omitempty"`
	PaidTax             float64 `json:"paidTax" bson:"paidTax,omitempty"`
	PaidPenalty         float64 `json:"paidPenalty" bson:"paidPenalty,omitempty"`
	PaidTotalTaxPenalty float64 `json:"paidTotalTaxPenalty" bson:"paidTotalTaxPenalty,omitempty"`
}

type UserChargeDemandSummary struct {
	Actual struct {
		ArrearTax  float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
		CurrentTax float64 `json:"currentTax" bson:"currentTax,omitempty"`
		TotalTax   float64 `json:"totalTax" bson:"totalTax,omitempty"`
	}
	ArrearTax      float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
	ArrearPenalty  float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
	ArrearTotal    float64 `json:"arrearTotal" bson:"arrearTotal,omitempty"`
	CurrentPenalty float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
	CurrentTax     float64 `json:"currentTax" bson:"currentTax,omitempty"`
	CurrentTotal   float64 `json:"currentTotal" bson:"currentTotal,omitempty"`
	TotalTax       float64 `json:"totalTax" bson:"totalTax,omitempty"`
}
