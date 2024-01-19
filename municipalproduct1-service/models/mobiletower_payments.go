package models

import "time"

type MobileTowerPayments struct {
	TnxID              string                    `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID         string                    `json:"propertyId" bson:"propertyId,omitempty"`
	MobileTowerID      string                    `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	ReciptNo           string                    `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptNoOrder      string                    `json:"reciptNoOrder" bson:"reciptNoOrder,omitempty"`
	FinancialYear      FinancialYear             `json:"financialYear" bson:"financialYear,omitempty"`
	Details            MobileTowerPaymentDetails `json:"details" bson:"details,omitempty"`
	Demand             MobileTowerPaymentDemand  `json:"demand" bson:"demand,omitempty"`
	CompletionDate     *time.Time                `json:"completionDate" bson:"completionDate,omitempty"`
	Status             string                    `json:"status" bson:"status,omitempty"`
	Address            Address                   `json:"address" bson:"address,omitempty"`
	ReciptURL          string                    `json:"reciptURL" bson:"reciptURL,omitempty"`
	Remark             string                    `json:"remark" bson:"remark,omitempty"`
	RejectedInfo       MobileTowerPaymentsAction `json:"rejectedInfo" bson:"rejectedInfo,omitempty"`
	VerifiedInfo       MobileTowerPaymentsAction `json:"verifiedInfo" bson:"verifiedInfo,omitempty"`
	NotVerifiedInfo    MobileTowerPaymentsAction `json:"notVerifiedInfo" bson:"notVerifiedInfo,omitempty"`
	Created            CreatedV2                 `json:"created" bson:"created,omitempty"`
	Scenario           string                    `json:"scenario" bson:"scenario,omitempty"`
	CollectionReceived CollectionReceived        `json:"collectionReceived" bson:"collectionReceived,omitempty"`
}
type RefMobileTowerPayments struct {
	MobileTowerPayments `bson:",inline"`
	FYs                 []MobileTowerPaymentsfY    `json:"fys" bson:"fys,omitempty"`
	Basic               *MobileTowerPaymentsBasics `json:"basics" bson:"basics,omitempty"`
	Ref                 struct {
		Address   RefAddress `json:"address" bson:"address,omitempty"`
		Collector User       `json:"collector" bson:"collector,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

//PropertyPaymentDetails : ""
type MobileTowerPaymentDetails struct {
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
}

type MobileTowerPaymentDemand struct {
	Current MobileTowerPaymentDemandSplitage `json:"current" bson:"current,omitempty"`
	Arrear  MobileTowerPaymentDemandSplitage `json:"arrear" bson:"arrear,omitempty"`
	Total   MobileTowerPaymentDemandSplitage `json:"total" bson:"total,omitempty"`
}

type MobileTowerPaymentDemandSplitage struct {
	Tax     float64 `json:"tax" bson:"tax,omitempty"`
	Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
	Other   float64 `json:"other" bson:"other,omitempty"`
	Total   float64 `json:"total" bson:"total,omitempty"`
}
type MobileTowerPaymentsfY struct {
	TnxID         string                    `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID    string                    `json:"propertyId" bson:"propertyId,omitempty"`
	MobileTowerID string                    `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	FY            RefMobileTowerDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Status        string                    `json:"status" bson:"status,omitempty"`
}

type MobileTowerPaymentsBasics struct {
	TnxID         string                 `json:"tnxId" bson:"tnxId,omitempty"`
	PropertyID    string                 `json:"propertyId" bson:"propertyId,omitempty"`
	MobileTowerID string                 `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	MobileTower   RefPropertyMobileTower `json:"mobileTower" bson:"mobileTower,omitempty"`
	Property      *RefProperty           `json:"property" bson:"property,omitempty"`
	Owner         []RefPropertyOwner     `json:"owners" bson:"owners,omitempty"`
	Status        string                 `json:"status" bson:"status,omitempty"`
}

type InitiateMobileTowerPaymentReq struct {
	MobileTowerID string   `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	By            string   `json:"by" bson:"by,omitempty"`
	ByType        string   `json:"byType" bson:"byType,omitempty"`
	FYs           []string `json:"fys" bson:"fys,omitempty"`
}

type MakeMobileTowerPaymentReq struct {
	TnxID          string                    `json:"tnxId" bson:"tnxId,omitempty"`
	Details        MobileTowerPaymentDetails `json:"details" bson:"details,omitempty"`
	Creator        CreatedV2                 `json:"creator" bson:"creator,omitempty"`
	Status         string                    `json:"status" bson:"status,omitempty"`
	CompletionDate *time.Time                `json:"completionDate" bson:"completionDate,omitempty"`
}

type MobileTowerPaymentsFilter struct {
	MobileTowerID []string `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	PropertyID    []string `json:"propertyId" bson:"propertyId,omitempty"`
	ReceiptNO     []string `json:"receiptNo" bson:"receiptNo,omitempty"`
	FY            []string `json:"fy" bson:"fy,omitempty"`
	Scenario      []string `json:"scenario" bson:"scenario,omitempty"`
	MOP           []string `json:"mop" bson:"mop,omitempty"`
	Collector     []string `json:"collector" bson:"collector,omitempty"`
	CollectorType []string `json:"collectorType" bson:"collectorType,omitempty"`

	CompletionDate DateRange     `json:"completionDate" bson:"completionDate,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	Regex          struct {
		PropertyID    string `json:"propertyId" bson:"propertyId"`
		OwnerName     string `json:"ownerName" bson:"ownerName"`
		OwnerMobile   string `json:"ownerMobile" bson:"ownerMobile"`
		MobileTowerID string `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
		ReciptNo      string `json:"reciptNo" bson:"reciptNo,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type MobileTowerPaymentsAction struct {
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	ActualDate *time.Time `json:"actualDate" bson:"actualDate,omitempty"`
	Remark     string     `json:"remark" bson:"remark,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
}

type MakeMobileTowerPaymentsAction struct {
	TnxID                     string `json:"tnxId" bson:"tnxId,omitempty"`
	MobileTowerPaymentsAction `bson:",inline"`
}

// BasicMobileTowerUpdateLogV2 : ""
type BasicMobileTowerUpdateLogV2 struct {
	MobileTowerID string                 `json:"mobileTowerId,omitempty" bson:"mobileTowerId,omitempty"`
	UniqueID      string                 `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous      RefPropertyMobileTower `json:"previous,omitempty" bson:"previous,omitempty"`
	New           RefPropertyMobileTower `json:"new,omitempty" bson:"new,omitempty"`
	UserName      string                 `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType      string                 `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester     Updated                `json:"requester" bson:"requester,omitempty"`
	Action        Updated                `json:"action" bson:"action,omitempty"`
	Proof         []string               `json:"proof,omitempty" bson:"proof,omitempty"`
	Status        string                 `json:"status,omitempty" bson:"status,omitempty"`
}

// RefBasicMobileTowerUpdateLogV2 : ""
type RefBasicMobileTowerUpdateLogV2 struct {
	BasicMobileTowerUpdateLogV2 `bson:",inline"`
	Ref                         struct {
		RequestedBy     User                   `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType               `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User                   `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType               `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		Previous        RefPropertyMobileTower `json:"previous,omitempty" bson:"previous,omitempty"`
		New             RefPropertyMobileTower `json:"new,omitempty" bson:"new,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
