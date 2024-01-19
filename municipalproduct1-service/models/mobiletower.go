package models

import (
	"time"
)

// PropertyMobileTower : ""
type PropertyMobileTower struct {
	UniqueID        string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID      string    `json:"propertyId" bson:"propertyId,omitempty"`
	Provider        string    `json:"provider" bson:"provider,omitempty"`
	Status          string    `json:"status" bson:"status,omitempty"`
	Created         CreatedV2 `json:"created,omitempty" bson:"created,omitempty"`
	PaymentScenario string    `json:"paymentScenario" bson:"paymentScenario,omitempty"`
	OwnerName       string    `json:"ownerName" bson:"ownerName,omitempty"`
	MobileNo        string    `json:"mobileNo" bson:"mobileNo,omitempty"`
	Address         Address   `json:"address" bson:"address,omitempty"`
	// Date of Construction
	DateFrom           *time.Time                  `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo             *time.Time                  `json:"dateTo" bson:"dateTo,omitempty"`
	BuiltUpArea        float64                     `json:"builtUpArea" bson:"builtUpArea,omitempty"`
	Demand             MobileTowerTotalDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        MobileTowerTotalCollection  `json:"collection" bson:"collection,omitempty"`
	PendingCollections MobileTowerTotalCollection  `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        MobileTowerTotalOutStanding `json:"outstanding" bson:"outstanding,omitempty"`
	NewPropertyID      string                      `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID      string                      `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

type MobileTowerTotalDemand struct {
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

type MobileTowerTotalCollection struct {
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

type MobileTowerTotalOutStanding struct {
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

// RefPropertyMobileTower : ""
type RefPropertyMobileTower struct {
	PropertyMobileTower `bson:",inline"`
	Ref                 struct {
		Address                    RefAddress                    `json:"address" bson:"address,omitempty"`
		MobileTowerRegistrationTax RefMobileTowerRegistrationTax `json:"mobileTowerRegistrationTax" bson:"mobileTowerRegistrationTax,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (res *RefPropertyMobileTower) Inc(a int) int {
	return a + 1
}

// PropertyMobileTowerFilter : ""
type PropertyMobileTowerFilter struct {
	Status     []string       `json:"status" bson:"status,omitempty"`
	Address    *AddressSearch `json:"address"`
	SearchText struct {
		UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
		Mobile     string `json:"mobile" bson:"mobile,omitempty"`
		PropertyNo string `json:"propertyNo" bson:"propertyNo"`
		OwnerName  string `json:"ownerName" bson:"ownerName"`
	} `json:"searchText" bson:"searchText"`
	SortBy    string `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int    `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

type MobileTowerWithMobileNoReq struct {
	// Status   string `json:"status" bson:"status,omitempty"`
	MobileNo string `json:"mobileNo" bson:"mobileNo,omitempty"`
}

type MobileTowerWithMobileNoRes struct {
	MobileTowers []PropertyMobileTower `json:"mobileTowers" bson:"mobiletowers,omitempty"`
}

type MobileTowerPenaltyUpdate struct {
	UniqueID           string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	MobileTowerPenalty float64 `json:"mobileTowerPenalty" bson:"mobileTowerPenalty,omitempty"`
}

type MobileTowerUpdateData struct {
	// UniqueID    string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID  string     `json:"propertyId" bson:"propertyId,omitempty"`
	Address     Address    `json:"address" bson:"address,omitempty"`
	DateFrom    *time.Time `json:"dateFrom" bson:"dateFrom,omitempty"`
	BuiltUpArea float64    `json:"builtUpArea" bson:"builtUpArea,omitempty"`
}
type BasicMobileTowerUpdateData struct {
	MobileTowerID string                `json:"mobileTowerId,omitempty" bson:"mobileTowerId,omitempty"`
	UpdateData    MobileTowerUpdateData `json:"updateData,omitempty" bson:"updateData,omitempty"`
	UserName      string                `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType      string                `json:"userType,omitempty" bson:"userType,omitempty"`
	Proof         []string              `json:"proof,omitempty" bson:"proof,omitempty"`
	Remarks       string                `json:"remarks,omitempty" bson:"remarks,omitempty"`
}
type BasicMobileTowerUpdateLog struct {
	MobileTowerID string                `json:"mobileTowerId,omitempty" bson:"mobileTowerId,omitempty"`
	UniqueID      string                `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous      MobileTowerUpdateData `json:"previous,omitempty" bson:"previous,omitempty"`
	New           MobileTowerUpdateData `json:"new,omitempty" bson:"new,omitempty"`
	UserName      string                `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType      string                `json:"userType,omitempty" bson:"userType,omitempty"`
	Action        Updated               `json:"action" bson:"action,omitempty"`
	Proof         []string              `json:"proof,omitempty" bson:"proof,omitempty"`
	Status        string                `json:"status,omitempty" bson:"status,omitempty"`
	Remarks       string                `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

// FilterBasicMobileTowerUpdateLog : ""
type FilterBasicMobileTowerUpdateLog struct {
	MobileTowerID []string `json:"mobileTowerId,omitempty" bson:"mobileTowerId,omitempty"`
	UniqueID      []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status        []string `json:"status,omitempty" bson:"status,omitempty"`
}

//  AcceptBasicMobileTowerUpdate : ""
type AcceptBasicMobileTowerUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectBasicMobileTowerUpdate : ""
type RejectBasicMobileTowerUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

type RefBasicMobileTowerUpdateLog struct {
	BasicMobileTowerUpdateLog `bson:",inline"`
	Ref                       struct {
		RequestedBy     User     `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ActionBy        User     `json:"action,omitempty" bson:"actionBy,omitempty"`
		ActionByType    UserType `json:"actionByType,omitempty" bson:"actionByType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type MobileTowerOverallDemandReport struct {
	MobileTowers []RefPropertyMobileTower `json:"mobileTowers" bson:"mobileTowers,omitempty"`
	CFY          RefFinancialYear         `json:"cfy" bson:"cfy,omitempty"`
}
