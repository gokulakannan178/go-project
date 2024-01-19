package models

import "time"

type ShopRentMonthlyPayments struct {
	TnxID           string                 `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID      string                 `json:"propertyId" bson:"propertyId,omitempty"`
	ShopRentID      string                 `json:"shopRentId" bson:"shopRentId,omitempty"`
	ReciptNo        string                 `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptNoOrder   string                 `json:"reciptNoOrder" bson:"reciptNoOrder,omitempty"`
	FinancialYear   FinancialYear          `json:"financialYear" bson:"financialYear,omitempty"`
	Details         ShopRentPaymentDetails `json:"details" bson:"details,omitempty"`
	Demand          ShopRentPaymentDemand  `json:"demand" bson:"demand,omitempty"`
	CompletionDate  *time.Time             `json:"completionDate" bson:"completionDate,omitempty"`
	Status          string                 `json:"status" bson:"status,omitempty"`
	Address         Address                `json:"address" bson:"address,omitempty"`
	ReciptURL       string                 `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark          string                 `json:"remark" bson:"remark,omitempty"`
	RejectedInfo    ShopRentPaymentsAction `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	VerifiedInfo    ShopRentPaymentsAction `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	NotVerifiedInfo ShopRentPaymentsAction `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Created         CreatedV2              `json:"created" bson:"created,omitempty"`
	Scenario        string                 `json:"scenario" bson:"scenario,omitempty"`
}

type RefShopRentMonthlyPayments struct {
	ShopRentMonthlyPayments `bson:",inline"`
	FYs                     []ShopRenttMonthlyPaymentsfY `json:"fys" bson:"fys,omitempty"`
	Basic                   *ShopRentPaymentsBasics      `json:"basics" bson:"basics,omitempty"`
	Ref                     struct {
		Address   RefAddress `json:"address" bson:"address,omitempty"`
		Collector User       `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

type ShopRenttMonthlyPaymentsfY struct {
	TnxID      string                        `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID string                        `json:"propertyId" bson:"propertyId,omitempty"`
	ShopRentID string                        `json:"shopRentId" bson:"shopRentId,omitempty"`
	FY         RefShopRentMonthlyDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Status     string                        `json:"status" bson:"status,omitempty"`
}
