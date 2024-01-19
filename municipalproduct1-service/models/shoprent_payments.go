package models

import "time"

type ShopRentPayments struct {
	TnxID              string                 `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID         string                 `json:"propertyId" bson:"propertyId,omitempty"`
	ShopRentID         string                 `json:"shopRentId" bson:"shopRentId,omitempty"`
	ReciptNo           string                 `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptNoOrder      string                 `json:"reciptNoOrder" bson:"reciptNoOrder,omitempty"`
	FinancialYear      FinancialYear          `json:"financialYear" bson:"financialYear,omitempty"`
	Details            ShopRentPaymentDetails `json:"details" bson:"details,omitempty"`
	Demand             ShopRentPaymentDemand  `json:"demand" bson:"demand,omitempty"`
	CompletionDate     *time.Time             `json:"completionDate" bson:"completionDate,omitempty"`
	Status             string                 `json:"status" bson:"status,omitempty"`
	Address            Address                `json:"address" bson:"address,omitempty"`
	ReciptURL          string                 `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark             string                 `json:"remark" bson:"remark,omitempty"`
	RejectedInfo       ShopRentPaymentsAction `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	VerifiedInfo       ShopRentPaymentsAction `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	NotVerifiedInfo    ShopRentPaymentsAction `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Created            CreatedV2              `json:"created" bson:"created,omitempty"`
	Scenario           string                 `json:"scenario" bson:"scenario,omitempty"`
	CollectionReceived CollectionReceived     `json:"collectionReceived" bson:"collectionReceived,omitempty"`
	BouncedInfo        struct {
		BouncedActionDate *time.Time `json:"bouncedActionDate" bson:"bouncedActionDate,omitempty"`
		BouncedDate       *time.Time `json:"bouncedDate" bson:"bouncedDate,omitempty"`
		Remark            string     `json:"remark" bson:"remark,omitempty"`
		By                string     `json:"by" bson:"by,omitempty"`
		ByType            string     `json:"byType" bson:"byType,omitempty"`
	} `json:"bouncedInfo" bson:"bouncedInfo,omitempty"`
}

type RefShopRentPayments struct {
	ShopRentPayments `bson:",inline"`
	FYs              []ShopRentPaymentsfY    `json:"fys" bson:"fys,omitempty"`
	Basic            *ShopRentPaymentsBasics `json:"basics" bson:"basics,omitempty"`
	Ref              struct {
		Address   RefAddress `json:"address" bson:"address,omitempty"`
		Collector User       `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

//PropertyPaymentDetails : ""
type ShopRentPaymentDetails struct {
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

type ShopRentPaymentDemand struct {
	FYs     []RefShopRentDemandFYLog      `json:"fys,omitempty" bson:"-"`
	Current ShopRentPaymentDemandSplitage `json:"current" bson:"current,omitempty"`
	Arrear  ShopRentPaymentDemandSplitage `json:"arrear" bson:"arrear,omitempty"`
	Total   ShopRentPaymentDemandSplitage `json:"total" bson:"total,omitempty"`
}

type ShopRentPaymentDemandSplitage struct {
	Tax     float64 `json:"tax" bson:"tax,omitempty"`
	Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
	Other   float64 `json:"other" bson:"other,omitempty"`
	Total   float64 `json:"total" bson:"total,omitempty"`
}
type ShopRentPaymentsfY struct {
	TnxID      string                 `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID string                 `json:"propertyId" bson:"propertyId,omitempty"`
	ShopRentID string                 `json:"shopRentId" bson:"shopRentId,omitempty"`
	FY         RefShopRentDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Status     string                 `json:"status" bson:"status,omitempty"`
	Remark     string                 `json:"remark" bson:"remark,omitempty"`
}
type ShopRentPaymentsBasics struct {
	TnxID      string      `json:"tnxId" bson:"tnxId,omitempty"`
	ShopRentID string      `json:"shopRentId" bson:"shopRentId,omitempty"`
	ShopRent   RefShopRent `json:"shopRent" bson:"shopRent,omitempty"`
	Remark     string      `json:"remark" bson:"remark,omitempty"`
	Status     string      `json:"status" bson:"status,omitempty"`
}

type InitiateShopRentPaymentReq struct {
	ShopRentID string   `json:"shopRentId" bson:"shopRentId,omitempty"`
	By         string   `json:"by" bson:"by,omitempty"`
	ByType     string   `json:"byType" bson:"byType,omitempty"`
	FYs        []string `json:"fys" bson:"fys,omitempty"`
}

type InitiateShopRentMonthlyPaymentReq struct {
	ShopRentID string                    `json:"shopRentId" bson:"shopRentId,omitempty"`
	By         string                    `json:"by" bson:"by,omitempty"`
	ByType     string                    `json:"byType" bson:"byType,omitempty"`
	Months     []SingleMonthIdentifierV2 `json:"months" bson:"months,omitempty"`
}

type MakeShopRentPaymentReq struct {
	TnxID          string                 `json:"tnxId" bson:"tnxId,omitempty"`
	Details        ShopRentPaymentDetails `json:"details" bson:"details,omitempty"`
	Creator        CreatedV2              `json:"creator" bson:"creator,omitempty"`
	Status         string                 `json:"status" bson:"status,omitempty"`
	CompletionDate *time.Time             `json:"completionDate" bson:"completionDate,omitempty"`
}

type ShopRentPaymentsFilter struct {
	ShopRentID     []string      `json:"shopRentId" bson:"shopRentId,omitempty"`
	ReceiptNO      []string      `json:"receiptNo" bson:"receiptNo,omitempty"`
	MOP            []string      `json:"mop" bson:"mop,omitempty"`
	MadeAT         []string      `json:"madeAt" bson:"madeAt,omitempty"`
	FY             []string      `json:"fy" bson:"fy,omitempty"`
	CompletionDate *DateRange    `json:"completionDate" bson:"completionDate,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	Scenario       []string      `json:"scenario" bson:"scenario,omitempty"`
	Regex          struct {
		ReceiptNO   string `json:"receiptNo" bson:"receiptNo,omitempty"`
		ShopRentID  string `json:"shoprentId" bson:"shoprentId"`
		OwnerName   string `json:"ownerName" bson:"ownerName"`
		OwnerMobile string `json:"ownerMobile" bson:"ownerMobile"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type ShopRentMonthlyPaymentsFilter struct {
	ShopRentID     []string      `json:"shopRentId" bson:"shopRentId,omitempty"`
	ReceiptNO      []string      `json:"receiptNo" bson:"receiptNo,omitempty"`
	MOP            []string      `json:"mop" bson:"mop,omitempty"`
	MadeAT         []string      `json:"madeAt" bson:"madeAt,omitempty"`
	FY             []string      `json:"fy" bson:"fy,omitempty"`
	CompletionDate DateRange     `json:"completionDate" bson:"completionDate,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Scenario       []string      `json:"scenario" bson:"scenario,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	SearchBox      struct {
		ShopRentID  string `json:"shopRentId" bson:"shopRentId"`
		OwnerName   string `json:"ownerName" bson:"ownerName"`
		OwnerMobile string `json:"ownerMobile" bson:"ownerMobile"`
	} `json:"searchBox" bson:"searchBox"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type ShopRentPaymentsAction struct {
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	ActualDate *time.Time `json:"actualDate" bson:"actualDate,omitempty"`
	Remark     string     `json:"remark" bson:"remark,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
}

type MakeShopRentPaymentsAction struct {
	TnxID                  string `json:"tnxId" bson:"tnxId,omitempty"`
	ShopRentPaymentsAction `bson:",inline"`
}

type ShopRentIDsWithOwnerNames struct {
	ShopRentIDs []string `json:"shopRentIds" bson:"shopRentIds,omitempty"`
}

type ShopRentIDsWithMobileNos struct {
	ShopRentIDs []string `json:"shopRentIds" bson:"shopRentIds,omitempty"`
}

// BasicShopRentUpdateLog : ""
type BasicShopRentUpdateLog struct {
	ShopRentID string      `json:"shopRentId,omitempty" bson:"shopRentId,omitempty"`
	UniqueID   string      `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous   RefShopRent `json:"previous,omitempty" bson:"previous,omitempty"`
	New        RefShopRent `json:"new,omitempty" bson:"new,omitempty"`
	UserName   string      `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType   string      `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester  Updated     `json:"requester" bson:"requester,omitempty"`
	Action     Updated     `json:"action" bson:"action,omitempty"`
	Proof      []string    `json:"proof,omitempty" bson:"proof,omitempty"`
	Status     string      `json:"status,omitempty" bson:"status,omitempty"`
}

// RefBasicShopRentUpdateLog : ""
type RefBasicShopRentUpdateLog struct {
	BasicShopRentUpdateLog `bson:",inline"`
	Ref                    struct {
		RequestedBy     User        `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType    `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User        `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType    `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		Previous        RefShopRent `json:"previous,omitempty" bson:"previous,omitempty"`
		New             RefShopRent `json:"new,omitempty" bson:"new,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
