package models

import "time"

type TradeLicensePayments struct {
	TnxID              string                     `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID         string                     `json:"propertyId" bson:"propertyId,omitempty"`
	TradeLicenseID     string                     `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	ReciptNo           string                     `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptNoOrder      string                     `json:"reciptNoOrder" bson:"reciptNoOrder,omitempty"`
	FinancialYear      FinancialYear              `json:"financialYear" bson:"financialYear,omitempty"`
	Details            TradeLicensePaymentDetails `json:"details" bson:"details,omitempty"`
	Demand             TradeLicensePaymentDemand  `json:"demand" bson:"demand,omitempty"`
	CompletionDate     *time.Time                 `json:"completionDate" bson:"completionDate,omitempty"`
	Status             string                     `json:"status" bson:"status,omitempty"`
	Address            Address                    `json:"address" bson:"address,omitempty"`
	ReciptURL          string                     `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark             string                     `json:"remark" bson:"remark,omitempty"`
	RejectedInfo       TradeLicensePaymentsAction `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	VerifiedInfo       TradeLicensePaymentsAction `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	NotVerifiedInfo    TradeLicensePaymentsAction `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Created            CreatedV2                  `json:"created" bson:"created,omitempty"`
	CollectionReceived CollectionReceived         `json:"collectionReceived" bson:"collectionReceived,omitempty"`
	BouncedInfo        struct {
		BouncedActionDate *time.Time `json:"bouncedActionDate" bson:"bouncedActionDate,omitempty"`
		BouncedDate       *time.Time `json:"bouncedDate" bson:"bouncedDate,omitempty"`
		Remark            string     `json:"remark" bson:"remark,omitempty"`
		By                string     `json:"by" bson:"by,omitempty"`
		ByType            string     `json:"byType" bson:"byType,omitempty"`
	} `json:"bouncedInfo" bson:"bouncedInfo,omitempty"`
}

// RefTradeLicensePayments : ""
type RefTradeLicensePayments struct {
	TradeLicensePayments `bson:",inline"`
	FYs                  []TradeLicensePaymentsfY    `json:"fys" bson:"fys,omitempty"`
	Basic                *TradeLicensePaymentsBasics `json:"basics" bson:"basics,omitempty"`
	Ref                  struct {
		Address   RefAddress `json:"address" bson:"address,omitempty"`
		Collector User       `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

// TradeLicensePaymentDetails : ""
type TradeLicensePaymentDetails struct {
	PayeeName     string              `json:"payeeName" bson:"payeeName,omitempty"`
	Amount        float64             `json:"amount" bson:"amount,omitempty"`
	AmountInWords string              `json:"amountInWords" bson:"amountInWords,omitempty"`
	MadeAt        *TradeLicenceMadeAt `json:"madeAt" bson:"madeAt,omitempty"`

	MOP struct {
		CardRNet   CardRNet   `json:"cardRNet" bson:"cardRNet,omitempty"`
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
		VendorInfo struct {
			Paytm struct {
				BankName     string     `json:"bankName" bson:"bankName,omitempty"`
				BankTnxID    string     `json:"bankTnxId" bson:"bankTnxId,omitempty"`
				CheckSumHash string     `json:"checkSumHash" bson:"checkSumHash,omitempty"`
				Currency     string     `json:"currency" bson:"currency,omitempty"`
				GateWayName  string     `json:"gateWayName" bson:"gateWayName,omitempty"`
				MID          string     `json:"mId" bson:"mId,omitempty"`
				OrderId      string     `json:"orderId" bson:"orderId,omitempty"`
				PaymentMode  string     `json:"paymentMode" bson:"paymentMode,omitempty"`
				RespCode     string     `json:"respCode" bson:"respCode,omitempty"`
				RespMsg      string     `json:"respMsg" bson:"respMsg,omitempty"`
				Status       string     `json:"status" bson:"status,omitempty"`
				TxnAmount    string     `json:"txnAmount" bson:"txnAmount,omitempty"`
				TxnID        string     `json:"txnId" bson:"txnId,omitempty"`
				TxnDate      *time.Time `json:"TxnDate" bson:"TxnDate,omitempty"`
			} `json:"paytm" bson:"paytm,omitempty"`
			HDFC HDFCPaymentGatewayCheckPaymentStatusResponse
		} `json:"vendorInfo" bson:"vendorInfo,omitempty"`
	} `json:"mop" bson:"mop,omitempty"`
	Collector CreatedV2 `json:"collector" bson:"collector,omitempty"`
	PayedVia  string    `json:"payedVia" bson:"payedVia,omitempty"`
}

type TradeLicenceMadeAt struct {
	At       string     `json:"at" bson:"at,omitempty"`
	ID       string     `json:"id" bson:"id,omitempty"`
	Name     string     `json:"name" bson:"name,omitempty"`
	Branch   string     `json:"branch" bson:"branch,omitempty"`
	TxnID    string     `json:"txnId" bson:"txnId,omitempty"`
	Location string     `json:"location" bson:"location,omitempty"`
	DOP      *time.Time `json:"dop" bson:"dop,omitempty"`
}

type TradeLicensePaymentDemand struct {
	FYs     []RefTradeLicenseDemandFYLog      `json:"fys,omitempty" bson:"-"`
	Current TradeLicensePaymentDemandSplitage `json:"current" bson:"current,omitempty"`
	Arrear  TradeLicensePaymentDemandSplitage `json:"arrear" bson:"arrear,omitempty"`
	Total   TradeLicensePaymentDemandSplitage `json:"total" bson:"total,omitempty"`
}

type TradeLicensePaymentDemandSplitage struct {
	Tax     float64 `json:"tax" bson:"tax,omitempty"`
	Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
	Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
	Other   float64 `json:"other" bson:"other,omitempty"`
	Total   float64 `json:"total" bson:"total,omitempty"`
}
type TradeLicensePaymentsfY struct {
	TnxID          string                     `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID     string                     `json:"propertyId" bson:"propertyId,omitempty"`
	TradeLicenseID string                     `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	FY             RefTradeLicenseDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Status         string                     `json:"status" bson:"status,omitempty"`
	Remark         string                     `json:"remark" bson:"remark,omitempty"`
}
type TradeLicensePaymentsBasics struct {
	TnxID          string          `json:"tnxId" bson:"tnxId,omitempty"`
	TradeLicenseID string          `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	TradeLicense   RefTradeLicense `json:"tradeLicense" bson:"tradeLicense,omitempty"`
	Remark         string          `json:"remark" bson:"remark,omitempty"`
	Status         string          `json:"status" bson:"status,omitempty"`
}

type InitiateTradeLicensePaymentReq struct {
	TradeLicenseID string   `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	By             string   `json:"by" bson:"by,omitempty"`
	ByType         string   `json:"byType" bson:"byType,omitempty"`
	FYs            []string `json:"fys" bson:"fys,omitempty"`
}

type MakeTradeLicensePaymentReq struct {
	TnxID          string                     `json:"tnxId" bson:"tnxId,omitempty"`
	Details        TradeLicensePaymentDetails `json:"details" bson:"details,omitempty"`
	Creator        CreatedV2                  `json:"creator" bson:"creator,omitempty"`
	Status         string                     `json:"status" bson:"status,omitempty"`
	CompletionDate *time.Time                 `json:"completionDate" bson:"completionDate,omitempty"`
}

type TradeLicensePaymentsFilter struct {
	TradeLicenseID []string      `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	ReceiptNO      []string      `json:"receiptNo" bson:"receiptNo,omitempty"`
	MOP            []string      `json:"mop" bson:"mop,omitempty"`
	MadeAT         []string      `json:"madeAt" bson:"madeAt,omitempty"`
	FY             []string      `json:"fy" bson:"fy,omitempty"`
	CompletionDate DateRange     `json:"completionDate" bson:"completionDate,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	Regex          struct {
		TradeLicenseID string `json:"shoprentId" bson:"shoprentId"`
		OwnerName      string `json:"ownerName" bson:"ownerName"`
		OwnerMobile    string `json:"ownerMobile" bson:"ownerMobile"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type TradeLicenseIDsWithOwnerNames struct {
	TradeLicenseIDs []string `json:"tradeLicenseIds" bson:"tradeLicenseIds,omitempty"`
}

// BasicTradeLicenseUpdateLog : ""
type BasicTradeLicenseUpdateLogV2 struct {
	TradeLicenseID string          `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
	UniqueID       string          `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous       RefTradeLicense `json:"previous,omitempty" bson:"previous,omitempty"`
	New            RefTradeLicense `json:"new,omitempty" bson:"new,omitempty"`
	UserName       string          `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType       string          `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester      Updated         `json:"requester" bson:"requester,omitempty"`
	Action         Updated         `json:"action" bson:"action,omitempty"`
	Proof          []string        `json:"proof,omitempty" bson:"proof,omitempty"`
	Status         string          `json:"status,omitempty" bson:"status,omitempty"`
}

// RefBasicTradeLicenseUpdateLogV2 : ""
type RefBasicTradeLicenseUpdateLogV2 struct {
	BasicTradeLicenseUpdateLogV2 `bson:",inline"`
	Ref                          struct {
		RequestedBy     User            `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType        `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User            `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType        `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		Previous        RefTradeLicense `json:"previous,omitempty" bson:"previous,omitempty"`
		New             RefTradeLicense `json:"new,omitempty" bson:"new,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type DateWiseTradeLicenseReportFilter struct {
	Date *time.Time `json:"date" bson:"date,omitempty"`
}
type RefDateWiseTradeLicensePaymentReport struct {
	Overall struct {
		ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
		CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
	} `json:"overall" bson:"overall,omitempty"`
	Year struct {
		ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
		CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
		FyName            string  `json:"fyName" bson:"fyName,omitempty"`
	} `json:"year" bson:"year,omitempty"`
	Month struct {
		ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
		CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
		FyMonth           string  `json:"fyMonth" bson:"fyMonth,omitempty"`
	} `json:"month" bson:"month,omitempty"`
	Week struct {
		ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
		CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
		FyWeek            string  `json:"fyWeek" bson:"fyWeek,omitempty"`
	} `json:"week" bson:"week,omitempty"`
	Day struct {
		ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
		CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		PropertyCount     float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
		FyDay             string  `json:"fyDay" bson:"fyDay,omitempty"`
	} `json:"day" bson:"day,omitempty"`
}
