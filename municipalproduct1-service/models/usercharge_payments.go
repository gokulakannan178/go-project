package models

import "time"

type UserChargePayments struct {
	TnxID          string                    `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID     string                    `json:"propertyId" bson:"propertyId,omitempty"`
	UserChargeID   string                    `json:"userchargeId" bson:"userchargeId,omitempty"`
	ReciptNo       string                    `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptNoOrder  string                    `json:"reciptNoOrder" bson:"reciptNoOrder,omitempty"`
	FinancialYear  FinancialYear             `json:"financialYear" bson:"financialYear,omitempty"`
	Details        *UserChargePaymentDetails `json:"details" bson:"details,omitempty"`
	Demand         *UserChargePaymentDemand  `json:"demand" bson:"demand,omitempty"`
	CompletionDate *time.Time                `json:"completionDate" bson:"completionDate,omitempty"`
	Status         string                    `json:"status" bson:"status,omitempty"`
	Address        Address                   `json:"address" bson:"address,omitempty"`
	ReciptURL      string                    `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark         string                    `json:"remark" bson:"remark,omitempty"`
	RejectedInfo   UserChargePaymentsAction  `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	VerifiedInfo   UserChargePaymentsAction  `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	BouncedInfo    struct {
		BouncedActionDate *time.Time `json:"bouncedActionDate" bson:"bouncedActionDate,omitempty"`
		BouncedDate       *time.Time `json:"bouncedDate" bson:"bouncedDate,omitempty"`
		Remark            string     `json:"remark" bson:"remark,omitempty"`
		By                string     `json:"by" bson:"by,omitempty"`
		ByType            string     `json:"byType" bson:"byType,omitempty"`
	} `json:"bouncedInfo" bson:"bouncedInfo,omitempty"`
	NotVerifiedInfo    UserChargePaymentsAction `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Created            CreatedV2                `json:"created" bson:"created,omitempty"`
	Scenario           string                   `json:"scenario" bson:"scenario,omitempty"`
	CollectionReceived CollectionReceived       `json:"collectionReceived" bson:"collectionReceived,omitempty"`
}

type RefUserChargePayments struct {
	UserChargePayments `bson:",inline"`
	FYs                []UserChargePaymentsfY    `json:"fys" bson:"fys,omitempty"`
	Basic              *UserChargePaymentsBasics `json:"basics" bson:"basics,omitempty"`
	Ref                struct {
		Address            RefAddress `json:"address" bson:"address,omitempty"`
		Collector          User       `json:"collector" bson:"collector,omitempty"`
		CollectionReceived User       `json:"collectionReceivedBy" bson:"collectionReceivedBy,omitempty"`
		RejectedBy         User       `json:"rejectedBy" bson:"rejectedBy,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

//PropertyPaymentDetails : ""
type UserChargePaymentDetails struct {
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

type UserChargePaymentDemand struct {
	//	FYs     []RefUserChargeDemandFYLog      `json:"fys,omitempty" bson:"-"`
	Current UserChargePaymentDemandSplitage `json:"current" bson:"current,omitempty"`
	Arrear  UserChargePaymentDemandSplitage `json:"arrear" bson:"arrear,omitempty"`
	Total   UserChargePaymentDemandSplitage `json:"total" bson:"total,omitempty"`
}

type UserChargePaymentDemandSplitage struct {
	Tax     float64 `json:"tax" bson:"tax,omitempty"`
	Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
	Other   float64 `json:"other" bson:"other,omitempty"`
	Total   float64 `json:"total" bson:"total,omitempty"`
}
type UserChargePaymentsfY struct {
	TnxID        string `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID   string `json:"propertyId" bson:"propertyId,omitempty"`
	UserChargeID string `json:"userchargeId" bson:"userchargeId,omitempty"`
	//	FY           RefUserChargeDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Status  string `json:"status" bson:"status,omitempty"`
	Remarks string `json:"remarks" bson:"remarks,omitempty"`
}
type UserChargePaymentsBasics struct {
	TnxID        string     `json:"tnxId" bson:"tnxId,omitempty"`
	UserChargeID string     `json:"userchargeId" bson:"userchargeId,omitempty"`
	UserCharge   UserCharge `json:"userCharge" bson:"userCharge,omitempty"`
	Created      CreatedV2  `json:"created" bson:"created,omitempty"`
	Status       string     `json:"status" bson:"status,omitempty"`
	Remarks      string     `json:"remarks" bson:"remarks,omitempty"`
}

type InitiateUserChargePaymentReq struct {
	UserChargeID string   `json:"UserChargeId" bson:"UserChargeId,omitempty"`
	By           string   `json:"by" bson:"by,omitempty"`
	ByType       string   `json:"byType" bson:"byType,omitempty"`
	FYs          []string `json:"fys" bson:"fys,omitempty"`
}

type InitiateUserChargeMonthlyPaymentReq struct {
	UserChargeID string                  `json:"userchargeId" bson:"userchargeId,omitempty"`
	By           string                  `json:"by" bson:"by,omitempty"`
	ByType       string                  `json:"byType" bson:"byType,omitempty"`
	Months       []SingleMonthIdentifier `json:"months" bson:"months,omitempty"`
}

type MakeUserChargePaymentReq struct {
	TnxID          string                   `json:"tnxId" bson:"tnxId,omitempty"`
	Details        UserChargePaymentDetails `json:"details" bson:"details,omitempty"`
	Creator        CreatedV2                `json:"creator" bson:"creator,omitempty"`
	Status         string                   `json:"status" bson:"status,omitempty"`
	CompletionDate *time.Time               `json:"completionDate" bson:"completionDate,omitempty"`
}

type UserChargePaymentsFilter struct {
	UserChargeID   []string      `json:"UserChargeId" bson:"UserChargeId,omitempty"`
	ReceiptNO      []string      `json:"receiptNo" bson:"receiptNo,omitempty"`
	MOP            []string      `json:"mop" bson:"mop,omitempty"`
	MadeAT         []string      `json:"madeAt" bson:"madeAt,omitempty"`
	FY             []string      `json:"fy" bson:"fy,omitempty"`
	CompletionDate *DateRange    `json:"completionDate" bson:"completionDate,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	Scenario       []string      `json:"scenario" bson:"scenario,omitempty"`
	Regex          struct {
		ReceiptNO    string `json:"receiptNo" bson:"receiptNo,omitempty"`
		UserChargeID string `json:"UserChargeId" bson:"UserChargeId"`
		OwnerName    string `json:"ownerName" bson:"ownerName"`
		OwnerMobile  string `json:"ownerMobile" bson:"ownerMobile"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type UserChargeMonthlyPaymentsFilter struct {
	UserChargeID   []string      `json:"userchargeId" bson:"userchargeId,omitempty"`
	ReceiptNO      []string      `json:"receiptNo" bson:"receiptNo,omitempty"`
	MOP            []string      `json:"mop" bson:"mop,omitempty"`
	MadeAT         []string      `json:"madeAt" bson:"madeAt,omitempty"`
	Collector      []string      `json:"collector" bson:"collector,omitempty"`
	FY             []string      `json:"fy" bson:"fy,omitempty"`
	CompletionDate DateRange     `json:"completionDate" bson:"completionDate,omitempty"`
	DateRange      *DateRange    `json:"dateRange" bson:"dateRange,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Scenario       []string      `json:"scenario" bson:"scenario,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	SearchBox      struct {
		UserChargeID string `json:"UserChargeId" bson:"UserChargeId"`
		OwnerName    string `json:"ownerName" bson:"ownerName"`
		OwnerMobile  string `json:"ownerMobile" bson:"ownerMobile"`
	} `json:"searchBox" bson:"searchBox"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type UserChargePaymentsAction struct {
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	ActualDate *time.Time `json:"actualDate" bson:"actualDate,omitempty"`
	Remark     string     `json:"remark" bson:"remark,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
}

type MakeUserChargePaymentsAction struct {
	TnxID                    string `json:"tnxId" bson:"tnxId,omitempty"`
	UserChargePaymentsAction `bson:",inline"`
}

type UserChargeIDsWithOwnerNames struct {
	UserChargeIDs []string `json:"UserChargeIds" bson:"UserChargeIds,omitempty"`
}

type UserChargeIDsWithMobileNos struct {
	UserChargeIDs []string `json:"UserChargeIds" bson:"UserChargeIds,omitempty"`
}

// BasicUserChargeUpdateLog : ""
type BasicUserChargeUpdateLog struct {
	UserChargeID string `json:"UserChargeId,omitempty" bson:"UserChargeId,omitempty"`
	UniqueID     string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	//	Previous     RefUserCharge `json:"previous,omitempty" bson:"previous,omitempty"`
	//	New          RefUserCharge `json:"new,omitempty" bson:"new,omitempty"`
	UserName  string   `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType  string   `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester Updated  `json:"requester" bson:"requester,omitempty"`
	Action    Updated  `json:"action" bson:"action,omitempty"`
	Proof     []string `json:"proof,omitempty" bson:"proof,omitempty"`
	Status    string   `json:"status,omitempty" bson:"status,omitempty"`
}

// RefBasicUserChargeUpdateLog : ""
type RefBasicUserChargeUpdateLog struct {
	BasicUserChargeUpdateLog `bson:",inline"`
	Ref                      struct {
		RequestedBy     User     `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User     `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		//	Previous        RefUserCharge `json:"previous,omitempty" bson:"previous,omitempty"`
		//	New             RefUserCharge `json:"new,omitempty" bson:"new,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type DateWiseUserchargeReportFilter struct {
	Date *time.Time `json:"date" bson:"date,omitempty"`
}

type UserwiseUsercharge struct {
	User               `bson:",inline"`
	TotalAmount        float64    `json:"totalAmount" bson:"totalAmount,omitempty"`
	Date               *time.Time `json:"date" bson:"date,omitempty"`
	UserChargePayments struct {
		Cash struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"cash" bson:"cash,omitempty"`
		Cheque struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"cheque" bson:"cheque,omitempty"`
		NetBanking struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"netbanking" bson:"netbanking,omitempty"`
		DD struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"dd" bson:"dd,omitempty"`
	} `json:"userchargepayments" bson:"userchargepayments,omitempty"`
}

type UserwiseTradeLicense struct {
	User                 `bson:",inline"`
	TotalAmount          float64    `json:"totalAmount" bson:"totalAmount,omitempty"`
	Date                 *time.Time `json:"date" bson:"date,omitempty"`
	TradeLicensePayments struct {
		Cash struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"cash" bson:"cash,omitempty"`
		Cheque struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"cheque" bson:"cheque,omitempty"`
		NetBanking struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"netbanking" bson:"netbanking,omitempty"`
		DD struct {
			Count       float64 `json:"count" bson:"count,omitempty"`
			TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		} `json:"dd" bson:"dd,omitempty"`
	} `json:"tradelicensepayments" bson:"tradelicensepayments,omitempty"`
}
