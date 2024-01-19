package models

import "time"

type PropertyPartPayment struct {
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

//PropertyWalletFilter : ""
type PropertyPartPaymentFilter struct {
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
type RefPropertyPartPayment struct {
	PropertyPartPayment `bson:",inline"`
	Ref                 struct {
		Collector User `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

type RefPropertyPartPaymentDemandFYLog struct {
	PropertyPartPaymentDemandFYLog `bson:",inline"`
}

type PropertyPartPaymentRateMaster struct {
	UniqueID          string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	ShopCategoryID    string     `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID string     `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Rate              float64    `json:"rate" bson:"rate,omitempty"`
	DOE               *time.Time `json:"doe" bson:"doe,omitempty"`
	Status            string     `json:"status" bson:"status,omitempty"`
	Created           *Created   `json:"created,omitempty"  bson:"created,omitempty"`
	Updated           []Updated  `json:"updated,omitempty"  bson:"updated,omitempty"`
}

type PropertyPartPaymentDemandFYLog struct {
	FinancialYear `bson:",inline"`
	PropertyID    string `json:"propertyId" bson:"propertyId,omitempty"`
	PartPaymentID string `json:"partPaymentId" bson:"partPaymentId,omitempty"`
	Status        string `json:"status" bson:"status,omitempty"`
	Details       struct {
		Tax            float64 `json:"tax" bson:"tax,omitempty"`
		Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
		Other          float64 `json:"other" bson:"other,omitempty"`
		TotalTaxAmount float64 `json:"totalTaxAmount,omitempty" bson:"totalTaxAmount,omitempty"`
	} `json:"details,omitempty" bson:"details,omitempty"`
	Ref struct {
		PropertyPartPaymentTax PropertyPartPaymentRateMaster `json:"propertyPartPaymentTax,omitempty" bson:"propertyPartPaymentTax,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type PropertyPartPaymentPaymentsfY struct {
	TnxID                 string                            `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID            string                            `json:"propertyId" bson:"propertyId,omitempty"`
	PropertyPartPaymentID string                            `json:"propertyPartPaymentId" bson:"propertyPartPaymentId,omitempty"`
	FY                    RefPropertyPartPaymentDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Status                string                            `json:"status" bson:"status,omitempty"`
}

type PropertyPartPaymentsBasics struct {
	TnxID                 string                 `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyPartPaymentID string                 `json:"propertyPartPaymentId" bson:"propertyPartPaymentId,omitempty"`
	PropertyPartPayment   RefPropertyPartPayment `json:"propertyPartPayment" bson:"propertyPartPayment,omitempty"`

	Status string `json:"status" bson:"status,omitempty"`
}

type PropertyPartPaymentPaymentsBasics struct {
	TnxID                 string                 `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyPartPaymentID string                 `json:"propertyPartPaymentId" bson:"propertyPartPaymentId,omitempty"`
	PropertyPartPayment   RefPropertyPartPayment `json:"propertyPartPayment" bson:"propertyPartPayment,omitempty"`

	Status string `json:"status" bson:"status,omitempty"`
}
type RefPropertyPartPayments struct {
	PropertyPartPayment `bson:",inline"`
	FYs                 []PropertyPartPaymentPaymentsfY    `json:"fys" bson:"fys,omitempty"`
	Basic               *PropertyPartPaymentPaymentsBasics `json:"basics" bson:"basics,omitempty"`
	Ref                 struct {
		Address   RefAddress `json:"address" bson:"address,omitempty"`
		Collector User       `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
type PropertyPartPaymentsAction struct {
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	ActualDate *time.Time `json:"actualDate" bson:"actualDate,omitempty"`
	Remark     string     `json:"remark" bson:"remark,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
}

type MakePropertyPartPaymentsAction struct {
	TnxID                      string `json:"tnxId" bson:"tnxId,omitempty"`
	UniqueID                   string `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyPartPaymentsAction `bson:",inline"`
}

type PropertyPartTotalCollection struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type PropertyPartTotalOutStanding struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type PropertyPartPaymentDemand struct {
	RefPropertyPartPayment `bson:",inline"`
	FY                     []RefPropertyPartPaymentDemandFYLog `json:"fy" bson:"fy,omitempty"`
	ProductConfiguration   *RefProductConfiguration            `json:"-" bson:"productConfiguration,omitempty"`
}

type PropertyPartPaymentCalcQueryFilter struct {
	PropertyPartPaymentID string   `json:"propertyPartPaymentId" bson:"propertyPartPaymentId,omitempty"`
	OmitFy                []string `json:"omitFy" bson:"omitFy,omitempty"`
	AddFy                 []string `json:"addFy" bson:"addFy,omitempty"`
	OmitPayedYears        bool     `json:"omitPayedYears" bson:"omitPayedYears4,omitempty"`
}

type RefMOPPartPayment struct {
	RefPropertyPayment `bson:",inline"`
	Ref2               struct {
		Completed     []RefPropertyPartPayment `json:"completed" bson:"completed,omitempty"`
		ChequePending []RefPropertyPartPayment `json:"chequePending" bson:"chequePending,omitempty"`
		ChequeBounced []RefPropertyPartPayment `json:"chequeBounced" bson:"chequeBounced,omitempty"`
		DDNBPending   []RefPropertyPartPayment `json:"ddnbPending" bson:"ddnbPending,omitempty"`
		DDNBBounced   []RefPropertyPartPayment `json:"ddnbBounced" bson:"ddnbBounced,omitempty"`
		Rejected      []RefPropertyPartPayment `json:"rejected" bson:"rejected,omitempty"`
		//Address       RefAddress           `json:"address" bson:"address,omitempty"`
		// DDNBRejected  []RefPropertyPayment `json:"ddnbRejected" bson:"ddnbRejected,omitempty"`
	} `json:"ref2" bson:"ref2,omitempty"`
}
