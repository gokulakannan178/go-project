package models

import "time"

type SurveyAndTax struct {
	Survey struct {
		Today struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		}
		Month struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		}
		FinancialYear struct {
			FinancialYearID string `json:"financialYearId" bson:"financialYearId,omitempty"`
			Value           int64  `json:"value" bson:"value"`
		} `json:"survey" bson:"survey,omitempty"`
	}
	Tax struct {
		Today struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		}
		Month struct {
			Date  *time.Time `json:"date" bson:"date,omitempty"`
			Value int64      `json:"value" bson:"value"`
		}
		FinancialYear struct {
			FinancialYearID string `json:"financialYearId" bson:"financialYearId,omitempty"`
			Value           int64  `json:"value" bson:"value"`
		}
	} `json:"tax" bson:"tax,omitempty"`
	Created Created `json:"created" bson:"created,omitempty"`
}
