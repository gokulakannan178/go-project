package models

import "time"

// PropertyOtherDemandPartPayment : ""
type PropertyOtherDemandPartPayment struct {
	UniqueID              string                 `json:"uniqueId" bson:"uniqueId,omitempty"`
	TnxID                 string                 `json:"tnxId" bson:"tnxId,omitempty"`
	Address               Address                `json:"address" bson:"address,omitempty"`
	PropertyPartPaymentID string                 `json:"propertyPartPaymentId" bson:"propertyPartPaymentId,omitempty"`
	PropertyID            string                 `json:"propertyId" bson:"propertyId,omitempty"`
	ReciptNo              string                 `json:"reciptNo" bson:"reciptNo,omitempty"`
	Details               PropertyPaymentDetails `json:"details" bson:"details,omitempty"`
	Status                string                 `json:"status" bson:"status,omitempty"`
	PaymentDate           *time.Time             `json:"paymentDate" bson:"paymentDate,omitempty"`
	CompletionDate        *time.Time             `json:"completionDate" bson:"completionDate,omitempty"`
	RejectedInfo          struct {
		RejectedActionDate *time.Time `json:"rejectedActionDate" bson:"rejectedActionDate,omitempty"`
		RejectedDate       *time.Time `json:"rejectedDate" bson:"rejectedDate,omitempty"`
		Remark             string     `json:"remark" bson:"remark,omitempty"`
		By                 string     `json:"by" bson:"by,omitempty"`
		ByType             string     `json:"byType" bson:"byType,omitempty"`
	} `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	VerifiedInfo struct {
		VerifiedActionDate *time.Time `json:"verifiedActionDate" bson:"verifiedActionDate,omitempty"`
		VerifiedDate       *time.Time `json:"verifiedDate" bson:"verifiedDate,omitempty"`
		Remark             string     `json:"remark" bson:"remark,omitempty"`
		By                 string     `json:"by" bson:"by,omitempty"`
		ByType             string     `json:"byType" bson:"byType,omitempty"`
	} `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	NotVerifiedInfo struct {
		NotVerifiedActionDate *time.Time `json:"notVerifiedActionDate" bson:"notVerifiedActionDate,omitempty"`
		NotVerifiedDate       *time.Time `json:"notVerifiedDate" bson:"notVerifiedDate,omitempty"`
		Remark                string     `json:"remark" bson:"remark,omitempty"`
		By                    string     `json:"by" bson:"by,omitempty"`
		ByType                string     `json:"byType" bson:"byType,omitempty"`
	} `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
}

type PropertyOtherDemandPartPaymentFilter struct {
	Status      []string       `json:"status,omitempty" bson:"status,omitempty"`
	TnxID       []string       `json:"tnxId" bson:"tnxId,omitempty"`
	UniqueID    []string       `json:"uniqueId" bson:"uniqueId,omitempty"`
	ReciptNo    []string       `json:"reciptNo" bson:"reciptNo,omitempty"`
	PropertyID  []string       `json:"propertyId" bson:"propertyId,omitempty"`
	PayeeName   []string       `json:"payeeName" bson:"payeeName,omitempty"`
	CollectorID []string       `json:"collectorId" bson:"collectorId,omitempty"`
	MOP         []string       `json:"mop" bson:"mop,omitempty"`
	MadeAt      []string       `json:"madeAt" bson:"madeAt,omitempty"`
	Address     *AddressSearch `json:"address"`
	Regex       struct {
		ReciptNo   string `json:"reciptNo" bson:"reciptNo,omitempty"`
		PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
		PayeeName  string `json:"payeeName" bson:"payeeName,omitempty"`
	} `json:"regex" bson:"regex"`
	PaymentDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"paymentDateRange"`
	CompletionDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"completionDateRange"`
}

//RefPropertyWallet : ""
type RefPropertyOtherDemandPartPayment struct {
	PropertyOtherDemandPartPayment `bson:",inline"`
	Ref                            struct {
		Collector User `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
