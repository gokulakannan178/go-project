package models

import (
	"time"
)

type TradeLicense struct {
	UniqueID           string                       `json:"uniqueId" bson:"uniqueId,omitempty"`
	Desc               string                       `json:"desc" bson:"desc,omitempty"`
	Created            CreatedV2                    `json:"created" bson:"created,omitempty"`
	BuiltUpArea        float64                      `json:"builtUpArea" bson:"builtUpArea,omitempty"`
	Address            Address                      `json:"address" bson:"address,omitempty"`
	TLBTID             string                       `json:"tlbtId" bson:"tlbtId,omitempty"`
	TLCTID             string                       `json:"tlctId" bson:"tlctId,omitempty"`
	MobileNo           string                       `json:"mobileNo" bson:"mobileNo,omitempty"`
	OwnerName          string                       `json:"ownerName" bson:"ownerName,omitempty"`
	GuardianName       string                       `json:"guardianName" bson:"guardianName,omitempty"`
	BusinessName       string                       `json:"businessName" bson:"businessName,omitempty"`
	Photo              string                       `json:"photo" bson:"photo,omitempty"`
	Status             string                       `json:"status" bson:"status,omitempty"`
	ESign              string                       `json:"esign" bson:"esign,omitempty"`
	LicenseExpiryDate  *time.Time                   `json:"licenseExpiryDate" bson:"licenseExpiryDate,omitempty"`
	LicenseDate        *time.Time                   `json:"licenseDate" bson:"licenseDate,omitempty"`
	StartDate          *time.Time                   `json:"startDate" bson:"startDate,omitempty"`
	DateFrom           *time.Time                   `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo             *time.Time                   `json:"dateTo" bson:"dateTo,omitempty"`
	LicenseAmount      float64                      `json:"licenseAmount" bson:"licenseAmount,omitempty"`
	Demand             TradeLicenseTotalDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        TradeLicenseTotalCollection  `json:"collection" bson:"collection,omitempty"`
	PendingCollections TradeLicenseTotalCollection  `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        TradeLicenseTotalOutStanding `json:"outstanding" bson:"outstanding,omitempty"`
	Documents          TLDocuments                  `json:"documents" bson:"documents,omitempty"`
	Approved           Action                       `json:"approved,omitempty" bson:"approved,omitempty"`
	NotApproved        Action                       `json:"notApproved,omitempty" bson:"notApproved,omitempty"`
	Verify             Action                       `json:"verified,omitempty" bson:"verified,omitempty"`
}

type TLDocuments struct {
	Proof string `json:"proof,omitempty"  bson:"proof,omitempty"`
}

type TradeLicenseFilter struct {
	Address       *AddressSearch `json:"address"`
	UniqueID      []string       `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status        []string       `json:"status" bson:"status,omitempty"`
	TLBTID        []string       `json:"tlbtId" bson:"tlbtId,omitempty"`
	TLCTID        []string       `json:"tlctId" bson:"tlctId,omitempty"`
	CreatedBy     []string       `json:"createdBy" bson:"createdBy,omitempty"`
	IsExpired     bool           `json:"isExpired" bson:"isExpired,omitempty"`
	ApprovedBy    []string       `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
	NotApprovedBy []string       `json:"notApprovedBy,omitempty" bson:"notApprovedBy,omitempty"`
	VerifiedBy    []string       `json:"verifiedBy,omitempty" bson:"verifiedBy,omitempty"`
	SearchText    struct {
		MobileNo     string `json:"mobileNo" bson:"mobileNo,omitempty"`
		OwnerName    string `json:"ownerName" bson:"ownerName,omitempty"`
		GuardianName string `json:"guardianName" bson:"guardianName,omitempty"`
		UniqueID     string `json:"uniqueId" bson:"uniqueId,omitempty"`
		LisenceNo    string `json:"lisenceNo" bson:"lisenceNo,omitempty"`
	} `json:"searchText"`
	LicenseExpiryDate *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"licenseExpiryDate"`
	LicenseDate *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"licenseDate"`
	CreatedDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"createdDateRange"`
	SortBy    string `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder int    `json:"sortOrder" bson:"sortOrder,omitempty"`
}

type RefTradeLicense struct {
	TradeLicense `bson:",inline"`

	Ref struct {
		Address                  RefAddress                `json:"address" bson:"address,omitempty"`
		TradeLicenseBusinessType TradeLicenseBusinessType  `json:"tradeLicenseBusinessType" bson:"tradeLicenseBusinessType,omitempty"`
		TradeLicenseCategoryType TradeLicenseCategoryType  `json:"tradeLicenseCategoryType" bson:"tradeLicenseCategoryType,omitempty"`
		Payments                 []RefTradeLicensePayments `json:"payments" bson:"payments,omitempty"`
		MarketRate               TradeLicenseRateMaster    `json:"marketRate" bson:"marketRate,omitempty"`
		CategoryNames            []string                  `json:"categoryNames" bson:"categoryNames,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (res *RefTradeLicense) Inc(a int) int {
	return a + 1
}

type TradeLicensePaymentsAction struct {
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	ActualDate *time.Time `json:"actualDate" bson:"actualDate,omitempty"`
	Remark     string     `json:"remark" bson:"remark,omitempty"`
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
}
type MakeTradeLicensePaymentsAction struct {
	TnxID                      string `json:"tnxId" bson:"tnxId,omitempty"`
	TradeLicensePaymentsAction `bson:",inline"`
}

// TradeLicenseOverallDemandReport
type TradeLicenseOverallDemandReport struct {
	TradeLicenses []RefTradeLicense `json:"tradeLicenses" bson:"tradeLicenses,omitempty"`
	CFY           RefFinancialYear  `json:"cfy" bson:"cfy,omitempty"`
}

type TradeLicenseTotalDemand struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
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

type TradeLicenseTotalCollection struct {
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

type TradeLicenseTotalOutStanding struct {
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

// BasicPropertyUpdate : ""
type BasicTradeLicenseUpdate struct {
	TradeLicenseID string                 `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
	UpdateData     TradeLicenseUpdateData `json:"updateData,omitempty" bson:"updateData,omitempty"`
	UserName       string                 `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType       string                 `json:"userType,omitempty" bson:"userType,omitempty"`
	Proof          []string               `json:"proof,omitempty" bson:"proof,omitempty"`
	Remarks        string                 `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

// BasicTradeLicenseUpdateLog : ""
type BasicTradeLicenseUpdateLog struct {
	TradeLicenseID string                 `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
	UniqueID       string                 `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous       TradeLicenseUpdateData `json:"previous,omitempty" bson:"previous,omitempty"`
	New            TradeLicenseUpdateData `json:"new,omitempty" bson:"new,omitempty"`
	UserName       string                 `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType       string                 `json:"userType,omitempty" bson:"userType,omitempty"`
	// Requester      Updated                `json:"requester" bson:"requester,omitempty"`
	Action  Updated  `json:"action" bson:"action,omitempty"`
	Proof   []string `json:"proof,omitempty" bson:"proof,omitempty"`
	Status  string   `json:"status,omitempty" bson:"status,omitempty"`
	Remarks string   `json:"remarks,omitempty" bson:"remarks,omitempty"`
	// ApproverName string   `json:"approverName,omitempty" bson:"approverName,omitempty"`
	// ApproverType string   `json:"approverType,omitempty" bson:"approverType,omitempty"`
}

// FilterBasicTradeLicenseUpdateLog : ""
type FilterBasicTradeLicenseUpdateLog struct {
	TradeLicenseID []string `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
	UniqueID       []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
}

//  AcceptBasicTradeLicenseUpdate : ""
type AcceptBasicTradeLicenseUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectBasicTradeLicenseUpdate : ""
type RejectBasicTradeLicenseUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

type RefBasicTradeLicenseUpdateLog struct {
	BasicTradeLicenseUpdateLog `bson:",inline"`
	Ref                        struct {
		RequestedBy     User     `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ActionBy        User     `json:"actionBy,omitempty" bson:"actionBy,omitempty"`
		ActionByType    UserType `json:"actionByType,omitempty" bson:"actionByType,omitempty"`
		Previous        struct {
			Address RefAddress `json:"address,omitempty" bson:"address,omitempty"`
		} `json:"previous,omitempty" bson:"previous,omitempty"`
		New struct {
			Address RefAddress `json:"address,omitempty" bson:"address,omitempty"`
		} `json:"new,omitempty" bson:"new,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type TradeLicenseUpdateData struct {
	Desc         string  `json:"desc" bson:"desc,omitempty"`
	Address      Address `json:"address" bson:"address,omitempty"`
	TLBTID       string  `json:"tlbtId" bson:"tlbtId,omitempty"`
	TLCTID       string  `json:"tlctId" bson:"tlctId,omitempty"`
	MobileNo     string  `json:"mobileNo" bson:"mobileNo,omitempty"`
	OwnerName    string  `json:"ownerName" bson:"ownerName,omitempty"`
	GuardianName string  `json:"guardianName" bson:"guardianName,omitempty"`
	BusinessName string  `json:"businessName" bson:"businessName,omitempty"`
	Photo        string  `json:"photo" bson:"photo,omitempty"`
}

//  AcceptPropertyPaymentModeChange : ""
type ApproveTradeLicense struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ESign    string     `json:"esign" bson:"esign,omitempty"`
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectPropertyPaymentModeChangeRequest : ""
type NotApproveTradeLicense struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}

type TradeLicenseSAFDashboard struct {
	Init struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"init,omitempty" bson:"init,omitempty"`
	Pending struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"pending,omitempty" bson:"pending,omitempty"`
	Active struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"active,omitempty" bson:"active,omitempty"`
	Rejected struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"rejected,omitempty" bson:"rejected,omitempty"`
	Expired struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"expired,omitempty" bson:"expired,omitempty"`
}

type GetTradeLicenseSAFDashboardFilter struct {
}
