package models

import "time"

type SurveyAndTax struct {
	Survey struct {
		Today struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		} `json:"today" bson:"today,omitempty"`
		Month struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		} `json:"month" bson:"month,omitempty"`
		FinancialYear struct {
			FinancialYearID string `json:"financialYearId" bson:"financialYearId,omitempty"`
			Value           int64  `json:"value" bson:"value"`
		} `json:"financialYear" bson:"financialYear,omitempty"`
	} `json:"survey" bson:"survey,omitempty"`
	Tax struct {
		Today struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		} `json:"today" bson:"today,omitempty"`
		Month struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		} `json:"month" bson:"month,omitempty"`
		FinancialYear struct {
			FinancialYearID string `json:"financialYearId" bson:"financialYearId,omitempty"`
			Value           int64  `json:"value" bson:"value"`
		} `json:"financialYear" bson:"financialYear,omitempty"`
	} `json:"tax" bson:"tax,omitempty"`
	Created  Created `json:"created" bson:"created,omitempty"`
	Status   string  `json:"status" bson:"status,omitempty"`
	UniqueID string  `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
}

type RefSurveyAndTax struct {
	SurveyAndTax `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type SurveyAndTaxFilter struct {
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
}
