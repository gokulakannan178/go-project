package models

import "time"

// SolidWasteUserChargeMonthlyPayments : ""
type SolidWasteUserChargeMonthlyPayments struct {
	TnxID                  string                          `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID             string                          `json:"propertyId" bson:"propertyId,omitempty"`
	SolidWasteUserChargeID string                          `json:"solidWasteUserChargeId" bson:"solidWasteUserChargeId,omitempty"`
	ReciptNo               string                          `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptNoOrder          string                          `json:"reciptNoOrder" bson:"reciptNoOrder,omitempty"`
	FinancialYear          FinancialYear                   `json:"financialYear" bson:"financialYear,omitempty"`
	Details                SolidWasteUserPaymentDetails    `json:"details" bson:"details,omitempty"`
	Demand                 SolidWasteUserChargeTotalDemand `json:"demand" bson:"demand,omitempty"`
	CompletionDate         *time.Time                      `json:"completionDate" bson:"completionDate,omitempty"`
	Status                 string                          `json:"status" bson:"status,omitempty"`
	Address                Address                         `json:"address" bson:"address,omitempty"`
	ReciptURL              string                          `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark                 string                          `json:"remark" bson:"remark,omitempty"`
	// RejectedInfo    SolidWasteUserChargePaymentsAction `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	// VerifiedInfo    SolidWasteUserChargePaymentsAction `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	// NotVerifiedInfo SolidWasteUserChargePaymentsAction `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Created  CreatedV2 `json:"created" bson:"created,omitempty"`
	Scenario string    `json:"scenario" bson:"scenario,omitempty"`
}

// SolidWasteUserPaymentDetails : ""
type SolidWasteUserPaymentDetails struct {
	PayeeName     string  `json:"payeeName" bson:"payeeName,omitempty"`
	Amount        float64 `json:"amount" bson:"amount,omitempty"`
	AmountInWords string  `json:"amountInWords" bson:"amountInWords,omitempty"`
	MadeAt        *struct {
		At       string     `json:"at" bson:"at,omitempty"`
		ID       string     `json:"id" bson:"id,omitempty"`
		Name     string     `json:"name" bson:"name,omitempty"`
		Branch   string     `json:"branch" bson:"branch,omitempty"`
		TxnID    string     `json:"txnId" bson:"txnId,omitempty"`
		Location string     `json:"location" bson:"location,omitempty"`
		DOP      *time.Time `json:"dop" bson:"dop,omitempty"`
	} `json:"madeAt" bson:"madeAt,omitempty"`
	MOP struct {
		Mode       string     `json:"mode" bson:"mode,omitempty"`
		No         string     `json:"no" bson:"no,omitempty"`
		Date       *time.Time `json:"date" bson:"date,omitempty"`
		Bank       string     `json:"bank" bson:"bank,omitempty"`
		Branch     string     `json:"branch" bson:"branch,omitempty"`
		BounceDate *time.Time `json:"bounceDate" bson:"bounceDate,omitempty"`
		TxnID      string     `json:"txnId" bson:"txnId,omitempty"`
		DOP        *time.Time `json:"dop" bson:"dop,omitempty"`
		VendorType string     `json:"vendorType" bson:"vendorType,omitempty"`
		Vendor     string     `json:"vendor" bson:"vendor,omitempty"`
		Proof      string     `json:"proof" bson:"proof,omitempty"`
	} `json:"mop" bson:"mop,omitempty"`
	Collector CreatedV2 `json:"collector" bson:"collector,omitempty"`
	PayedVia  string    `json:"payedVia" bson:"payedVia,omitempty"`
}

// InitiateSolidWasteChargeMonthlyPaymentReq : ""
type InitiateSolidWasteChargeMonthlyPaymentReq struct {
	SolidWasteChargeID string                  `json:"solidWasteUserChargeId" bson:"solidWasteUserChargeId,omitempty"`
	By                 string                  `json:"by" bson:"by,omitempty"`
	ByType             string                  `json:"byType" bson:"byType,omitempty"`
	Months             []SingleMonthIdentifier `json:"months" bson:"months,omitempty"`
}

// SolidWasteChargeMonthlyPaymentsfY : ""
type SolidWasteChargeMonthlyPaymentsfY struct {
	TnxID                  string                             `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID             string                             `json:"propertyId" bson:"propertyId,omitempty"`
	SolidWasteUserChargeID string                             `json:"SolidWasteUserChargeId" bson:"SolidWasteUserChargeId,omitempty"`
	FY                     RefSolidWasteUserChargeDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Status                 string                             `json:"status" bson:"status,omitempty"`
}

// SolidWasteChargeMonthlyPaymentsBasics : ""
type SolidWasteChargeMonthlyPaymentsBasics struct {
	TnxID                  string                  `json:"tnxId" bson:"tnxId,omitempty"`
	SolidWasteUserChargeID string                  `json:"solidWasteUserChargeId" bson:"solidWasteUserChargeId,omitempty"`
	SolidWasteUserCharge   RefSolidWasteUserCharge `json:"solidWasteUserCharge" bson:"solidWasteUserCharge,omitempty"`
	Status                 string                  `json:"status" bson:"status,omitempty"`
}

// RefSolidWasteChargeMonthlyPayments : ""
type RefSolidWasteChargeMonthlyPayments struct {
	SolidWasteUserChargeMonthlyPayments `bson:",inline"`
	FYs                                 []SolidWasteChargeMonthlyPaymentsfY    `json:"fys" bson:"fys,omitempty"`
	Basic                               *SolidWasteChargeMonthlyPaymentsBasics `json:"basics" bson:"basics,omitempty"`
	Ref                                 struct {
		Address   RefAddress `json:"address" bson:"address,omitempty"`
		Collector User       `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

// MakeSolidWasteUserChargePaymentReq : ""
type MakeSolidWasteUserChargePaymentReq struct {
	TnxID          string                       `json:"tnxId" bson:"tnxId,omitempty"`
	Details        SolidWasteUserPaymentDetails `json:"details" bson:"details,omitempty"`
	Creator        CreatedV2                    `json:"creator" bson:"creator,omitempty"`
	Status         string                       `json:"status" bson:"status,omitempty"`
	CompletionDate *time.Time                   `json:"completionDate" bson:"completionDate,omitempty"`
}

// SolidWasteUserChargePaymentsFilter : ""
type SolidWasteUserChargePaymentsFilter struct {
	SolidWasteUserChargeID []string      `json:"solidWasteUserChargeId" bson:"solidWasteUserChargeId,omitempty"`
	ReceiptNo              []string      `json:"receiptNo" bson:"receiptNo,omitempty"`
	MOP                    []string      `json:"mop" bson:"mop,omitempty"`
	MadeAT                 []string      `json:"madeAt" bson:"madeAt,omitempty"`
	FY                     []string      `json:"fy" bson:"fy,omitempty"`
	CompletionDate         *DateRange    `json:"completionDate" bson:"completionDate,omitempty"`
	Status                 []string      `json:"status" bson:"status,omitempty"`
	Address                AddressSearch `json:"address" bson:"address,omitempty"`
	Scenario               []string      `json:"scenario" bson:"scenario,omitempty"`
	Regex                  struct {
		ReceiptNO              string `json:"receiptNo" bson:"receiptNo,omitempty"`
		SolidWasteUserChargeID string `json:"solidWasteUserChargeId" bson:"solidWasteUserChargeId"`
		OwnerName              string `json:"ownerName" bson:"ownerName"`
		OwnerMobile            string `json:"ownerMobile" bson:"ownerMobile"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// SolidWasteUserChargeIDsWithOwnerNames : ""
type SolidWasteUserChargeIDsWithOwnerNames struct {
	SolidWasteUserChargeIDs []string `json:"solidWasteUserChargeIds" bson:"solidWasteUserChargeIds,omitempty"`
}

// SolidWasteUserChargeIDsWithMobileNos : ""
type SolidWasteUserChargeIDsWithMobileNos struct {
	SolidWasteUserChargeIDs []string `json:"solidWasteUserChargeIds" bson:"solidWasteUserChargeIds,omitempty"`
}

// SolidWasteUserChargePaymentsAction : ""
type SolidWasteUserChargePaymentsAction struct {
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	ActualDate *time.Time `json:"actualDate" bson:"actualDate,omitempty"`
	Remark     string     `json:"remark" bson:"remark,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
}

// MakeSolidWasteUserChargePaymentsAction : ""
type MakeSolidWasteUserChargePaymentsAction struct {
	TnxID                              string `json:"tnxId" bson:"tnxId,omitempty"`
	SolidWasteUserChargePaymentsAction `bson:",inline"`
}

// BasicSolidWasteUpdateLog : ""
type BasicSolidWasteUpdateLog struct {
	SolidWasteID string      `json:"solidWasteId,omitempty" bson:"solidWasteId,omitempty"`
	UniqueID     string      `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous     RefShopRent `json:"previous,omitempty" bson:"previous,omitempty"`
	New          RefShopRent `json:"new,omitempty" bson:"new,omitempty"`
	UserName     string      `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType     string      `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester    Updated     `json:"requester" bson:"requester,omitempty"`
	Action       Updated     `json:"action" bson:"action,omitempty"`
	Proof        []string    `json:"proof,omitempty" bson:"proof,omitempty"`
	Status       string      `json:"status,omitempty" bson:"status,omitempty"`
}

// RefBasicSolidWasteUpdateLog : ""
type RefBasicSolidWasteUpdateLog struct {
	BasicSolidWasteUpdateLog `bson:",inline"`
	Ref                      struct {
		RequestedBy     User                    `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType                `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User                    `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType                `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		Previous        RefSolidWasteUserCharge `json:"previous,omitempty" bson:"previous,omitempty"`
		New             RefSolidWasteUserCharge `json:"new,omitempty" bson:"new,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
