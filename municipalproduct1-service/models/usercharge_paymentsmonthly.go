package models

import "time"

type UserChargeMonthlyPayments struct {
	TnxID           string                   `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID      string                   `json:"propertyId" bson:"propertyId,omitempty"`
	UserChargeID    string                   `json:"UserChargeId" bson:"UserChargeId,omitempty"`
	ReciptNo        string                   `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptNoOrder   string                   `json:"reciptNoOrder" bson:"reciptNoOrder,omitempty"`
	FinancialYear   FinancialYear            `json:"financialYear" bson:"financialYear,omitempty"`
	Details         UserChargePaymentDetails `json:"details" bson:"details,omitempty"`
	Demand          UserChargeDemandSummary  `json:"demand" bson:"demand,omitempty"`
	CompletionDate  *time.Time               `json:"completionDate" bson:"completionDate,omitempty"`
	Status          string                   `json:"status" bson:"status,omitempty"`
	Address         Address                  `json:"address" bson:"address,omitempty"`
	ReciptURL       string                   `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark          string                   `json:"remark" bson:"remark,omitempty"`
	RejectedInfo    UserChargePaymentsAction `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	VerifiedInfo    UserChargePaymentsAction `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	NotVerifiedInfo UserChargePaymentsAction `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Created         CreatedV2                `json:"created" bson:"created,omitempty"`
	Scenario        string                   `json:"scenario" bson:"scenario,omitempty"`
}

type RefUserChargeMonthlyPayments struct {
	UserChargeMonthlyPayments `bson:",inline"`
	FYs                       []UserChargetMonthlyPaymentsfYForReceipt `json:"fys" bson:"fys,omitempty"`
	Basic                     *UserChargePaymentsBasics                `json:"basics" bson:"basics,omitempty"`
	Ref                       struct {
		Address       RefAddress    `json:"address" bson:"address,omitempty"`
		Collector     User          `json:"collector" bson:"collector,omitempty"`
		Property      Property      `json:"property" bson:"property,omitempty"`
		PropertyOwner PropertyOwner `json:"propertyowner" bson:"propertyowner,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

type UserChargetMonthlyPaymentsfY struct {
	TnxID        string                  `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID   string                  `json:"propertyId" bson:"propertyId,omitempty"`
	UserChargeID string                  `json:"userchargeId" bson:"userchargeId,omitempty"`
	Fy           FinancialYear           `json:"fy" bson:"fy,omitempty"`
	Month        UserChargeDemandFyMonth `json:"month" bson:"month,omitempty"`
	Status       string                  `json:"status" bson:"status,omitempty"`
	Created      CreatedV2               `json:"created" bson:"created,omitempty"`
}
type UserChargetMonthlyPaymentsfYForReceipt struct {
	Fy struct {
		Name string `json:"name" bson:"name,omitempty"`
		FyId string `json:"fyId" bson:"fyId,omitempty"`
	} `json:"fy" bson:"fy,omitempty"`
	Month    []UserChargeDemandFyMonth `json:"months" bson:"months,omitempty"`
	TotalTax float64                   `json:"totalTax" bson:"totalTax,omitempty"`
}
