package models

import "time"

// PropertyOtherDemandPayment : ""
type PropertyOtherDemandPayment struct {
	TnxID           string                  `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID      string                  `json:"propertyId" bson:"propertyId,omitempty"`
	RecordID        string                  `json:"recordId" bson:"recordId,omitempty"`
	TotalTax        float64                 `json:"totalTax" bson:"totalTax,omitempty"`
	ReciptNo        string                  `json:"reciptNo" bson:"reciptNo,omitempty"`
	FinancialYear   FinancialYear           `json:"financialYear" bson:"financialYear,omitempty"`
	Type            string                  `json:"type" bson:"type,omitempty"`
	Details         *PropertyPaymentDetails `json:"details" bson:"details,omitempty"`
	Demand          *PropertyPaymentDemand  `json:"demand" bson:"demand,omitempty"`
	RemainingAmount float64                 `json:"remainingAmount" bson:"remainingAmount,omitempty"`
	CompletionDate  *time.Time              `json:"completionDate" bson:"completionDate,omitempty"`
	Status          string                  `json:"status" bson:"status,omitempty"`
	PendingAmount   float64                 `json:"pendingAmount" bson:"pendingAmount,omitempty"`
	Address         Address                 `json:"address" bson:"address,omitempty"`
	ReciptURL       string                  `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark          string                  `json:"remark" bson:"remark,omitempty"`
	RejectedInfo    struct {
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
	CollectionReceived CollectionReceived `json:"collectionReceived" bson:"collectionReceived,omitempty"`
}

// PropertyOtherDemandPaymentDemandBasic : ""
type PropertyOtherDemandPaymentDemandBasic struct {
	TnxID    string             `json:"tnxId" bson:"tnxId,omitempty"`
	Property Property           `json:"property,omitempty" bson:"property,omitempty"`
	Owners   []RefPropertyOwner `json:"owners,omitempty" bson:"owners,omitempty"`
	Floors   []RefPropertyFloor `json:"floors,omitempty" bson:"floors,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
}

// PropertyOtherDemandPaymentDemandFy : ""
type PropertyOtherDemandPaymentDemandFy struct {
	TnxID      string              `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID string              `json:"propertyId" bson:"propertyId,omitempty"`
	FY         FinancialYearDemand `json:"fy,omitempty" bson:"fy,omitempty"`
	Status     string              `json:"status" bson:"status,omitempty"`
}

func (prop *PropertyOtherDemandPaymentDemandFy) Inc(a int) int {
	return a + 1
}

type PropertyOtherDemandPaymentDemand struct {
	PercentAreaBuildup        float64               `json:"percentAreaBuildup,omitempty" bson:"percentAreaBuildup,omitempty"`
	TaxableVacantLand         float64               `json:"taxableVacantLand,omitempty" bson:"taxableVacantLand,omitempty"`
	FYs                       []FinancialYearDemand `json:"fys,omitempty" bson:"-"`
	Property                  Property              `json:"property,omitempty" bson:"-"`
	ServiceCharge             float64               `json:"serviceCharge" bson:"serviceCharge"`
	IsServiceChargeApplicable bool                  `json:"isServiceChargeApplicable" bson:"isServiceChargeApplicable"`
	PropertyConfig            PropertyConfiguration `json:"propertyConfig" bson:"propertyConfig"`
	FYTax                     float64               `json:"fyTax" bson:"fyTax"`
	CompositeTax              float64               `json:"compositeTax" bson:"compositeTax"`
	EducationChess            float64               `json:"educationChess" bson:"educationChess"`
	SWUC                      float64               `json:"swuc" bson:"swuc"`
	PenalCharge               float64               `json:"penalCharge" bson:"penalCharge"`
	OtherAmount               float64               `json:"otherAmount" bson:"otherAmount"`
	FormFee                   float64               `json:"formFee" bson:"formFee"`
	TotalTax                  float64               `json:"totalTax" bson:"totalTax"`
	BoreCharge                float64               `json:"boreCharge" bson:"boreCharge"`
	Arrear                    float64               `json:"arrear" bson:"arrear"`
	Current                   float64               `json:"current" bson:"current"`
	PreviousCollection        PreviousCollection    `json:"previousCollection" bson:"previousCollection"`
}

// RefPropertyOtherDemandPayment : ""
type RefPropertyOtherDemandPayment struct {
	PropertyOtherDemandPayment `bson:",inline"`
	Basic                      PropertyOtherDemandPaymentDemandBasic `json:"basic" bson:"basic,omitempty"`
	Fys                        []PropertyOtherDemandPaymentDemandFy  `json:"fys" bson:"fys,omitempty"`
	Ref                        struct {
		PropertyOwner       []RefPropertyOwner                  `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		Address             RefAddress                          `json:"address" bson:"address,omitempty"`
		Collector           User                                `json:"collector" bson:"collector,omitempty"`
		PartPayments        []RefPropertyOtherDemandPartPayment `json:"partPayments" bson:"partPayments,omitempty"`
		PartAmountCollected float64                             `json:"partAmountCollected" bson:"partAmountCollected,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (prop *RefPropertyOtherDemandPayment) Inc(a int) int {
	return a + 1
}

// InitiatePropertyOtherDemandFilter : ""
type InitiatePropertyOtherDemandFilter struct {
	PropertyID string   `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	Status     []string `json:"status"`
	RecordID   []string `json:"recordId,omitempty" bson:"recordId,omitempty"`
}

// PropertyOtherDemandMakePayment : ""
type PropertyOtherDemandMakePayment struct {
	TnxID     string                 `json:"tnxId" bson:"tnxId,omitempty"`
	Details   PropertyPaymentDetails `json:"details" bson:"details,omitempty"`
	ReciptURL string                 `json:"reciptURL" bson:"reciptURL,omitempty"`
}

// PropertyOtherDemandVerifyPayment : ""
type PropertyOtherDemandVerifyPayment struct {
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
	TnxID      string     `json:"tnxId" bson:"tnxId,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
}

// PropertyOtherDemandNotVerifiedPayment : ""
type PropertyOtherDemandNotVerifiedPayment struct {
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	TnxID      string     `json:"tnxId" bson:"tnxId,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

// PropertyOtherDemandRejectPayment : ""
type PropertyOtherDemandRejectPayment struct {
	Date       *time.Time `json:"date" bson:"date,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	UniqueID   string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	TnxID      string     `json:"tnxId,omitempty" bson:"tnxId,omitempty"`
	Status     string     `json:"status" bson:"status,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
}
