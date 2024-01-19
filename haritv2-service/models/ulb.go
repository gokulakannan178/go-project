package models

import "time"

//ULB : ""
type ULB struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	DispName string `json:"dispname" bson:"dispname,omitempty"`
	Website  string `json:"website" bson:"website,omitempty"`
	Remarks  string `json:"remarks" bson:"remarks,omitempty"`
	//PrimaryContact Contact       `json:"primaryContact" bson:"primaryContact,omitempty"`
	//Contacts       []Contact     `json:"contacts" bson:"contacts,omitempty"`
	//Addresses []Address `json:"addresses" bson:"addresses,omitempty"`
	Address Address `json:"address" bson:"address,omitempty"`
	//Banks          []Bank        `json:"banks" bson:"banks,omitempty"`
	//Tax            Tax           `json:"tax" bson:"tax,omitempty"`
	Created           Created      `json:"createdOn" bson:"createdOn,omitempty"`
	Updated           Updated      `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog         []Updated    `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Status            string       `json:"status" bson:"status,omitempty"`
	TestCert          ULBTestCert  `json:"testcert" bson:"testcert,omitempty"`
	Logo              string       `json:"logo" bson:"logo,omitempty"`
	NodalOfficer      NodalOfficer `json:"nodalOfficer" bson:"nodalOfficer,omitempty"`
	CO                CO           `json:"co" bson:"co,omitempty"`
	ULBCode           string       `json:"ulbCode" bson:"ulbCode"`
	IsLocationUpdated bool         `json:"isLocationUpdated" bson:"isLocationUpdated,omitempty"`
	IsProfileUpdated  bool         `json:"isProfileUpdated" bson:"isProfileUpdated,omitempty"`
}

type ULBLessData struct {
	UniqueID     string       `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name         string       `json:"name" bson:"name,omitempty"`
	DispName     string       `json:"dispname" bson:"dispname,omitempty"`
	NodalOfficer NodalOfficer `json:"nodalOfficer" bson:"nodalOfficer,omitempty"`
}

//ULBFilter : ""
type ULBFilter struct {
	Name           string        `json:"name" bson:"name,omitempty"`
	UniqueID       []string      `json:"uniqueId" bson:"uniqueId,omitempty"`
	Contact        []string      `json:"contacts" bson:"contacts,omitempty"`
	OmitIDs        []string      `json:"omitIds" bson:"omitIds,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	TestCertStatus []string      `json:"testCertStatus" bson:"testCertStatus,omitempty"`
	MobileNo       []string      `json:"mobile" bson:"mobile,omitempty"`
	SortBy         string        `json:"sortBy"`
	SortOrder      int           `json:"sortOrder"`
	IsExpDate      string        `json:"isExpdate" bson:"isExpdate,omitempty"`
	ExcelType      string        `json:"excelType" bson:"excelType,omitempty"`

	Regex struct {
		Name     string `json:"name" bson:"name"`
		NoName   string `json:"noName" bson:"noName"`
		NoMobile string `json:"noMobile" bson:"noMobile"`
		CoName   string `json:"coName" bson:"coName"`
		CoMobile string `json:"coMobile" bson:"coMobile"`
	} `json:"regex" bson:"regex"`
}

//NodalOfficer : ""
type NodalOfficer struct {
	Name        string `json:"name" bson:"name,omitempty"`
	NonMobileNo string `json:"nonMobile" bson:"nonMobile,omitempty"`
	MobileNo    string `json:"mobile" bson:"mobile,omitempty"`
	Email       string `json:"email" bson:"email,omitempty"`
	UserName    string `json:"userName" bson:"userName,omitempty"`
	RollID      string `json:"rollId" bson:"rollId,omitempty"`
}

//RefULB : ""
type RefULB struct {
	ULB `bson:",inline"`
	Ref struct {
		Inventory ULBInventory `json:"inventory" bson:"inventory,omitempty"`
		Address   *RefAddress  `json:"address" bson:"address,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

//ULBTestCert : ""
type ULBTestCert struct {
	Doc                  string     `json:"doc" bson:"doc,omitempty"`
	AppliedDoc           string     `json:"appliedDoc" bson:"appliedDoc,omitempty"`
	AppliedDate          *time.Time `json:"appliedDate" bson:"appliedDate,omitempty"`
	RejectedDate         *time.Time `json:"rejectedDate" bson:"rejectedDate,omitempty"`
	RefDoc               string     `json:"refdoc" bson:"refdoc,omitempty"`
	ReportDate           *time.Time `json:"reportDate" bson:"reportDate,omitempty"`
	ExpDate              *time.Time `json:"expdate" bson:"expdate,omitempty"`
	IsExpDate            string     `json:"isExpdate" bson:"isExpdate,omitempty"`
	REGDate              *time.Time `json:"regDate" bson:"regDate,omitempty"`
	Status               string     `json:"status" bson:"status,omitempty"`
	Remarks              string     `json:"remarks" bson:"remarks,omitempty"`
	WasteGen             float64    `json:"wasteGen" bson:"wasteGen,omitempty"`
	TotalWasteGen        float64    `json:"totalWasteGen" bson:"totalWasteGen,omitempty"`
	WasteColl            float64    `json:"wasteColl" bson:"wasteColl,omitempty"`
	ApplyFor             string     `json:"applyFor" bson:"applyFor,omitempty"`
	TypeOfProcessing     string     `json:"typeOfProcessing" bson:"typeOfProcessing,omitempty"`
	SampleCheckedBy      string     `json:"sampleCheckedBy" bson:"sampleCheckedBy,omitempty"`
	SampleCheckedByID    string     `json:"sampleCheckedById" bson:"sampleCheckedById,omitempty"`
	GFCRating            int        `json:"gFCRating" bson:"gFCRating,omitempty"`
	StarRating           int        `json:"starRating" bson:"starRating,omitempty"`
	StarRatingDate       *time.Time `json:"starRatingDate" bson:"starRatingDate,omitempty"`
	StarRatingExpiryDate *time.Time `json:"starRatingExpiryDate" bson:"starRatingExpiryDate,omitempty"`
	Rejected             Rejected   `json:"rejected" bson:"rejected,omitempty"`
}
type RefULBTestCert struct {
	ULBTestCert `bson:",inline"`
	Ref         struct {
	} `json:"ref" bson:"ref,omitempty"`
}

type ULBInventoryUpdateMessageReport struct {
	ULBData []struct {
		ULBName  string `json:"ulbName" bson:"ulbName"`
		NoName   string `json:"noName" bson:"noName"`
		NoMobile string `json:"noMobile" bson:"noMobile"`
		ULBID    string `json:"ulbId" bson:"ulbId"`
	} `json:"ulbdata" bson:"ulb"`
}

type ULBInventoryUpdateMessageFilter struct {
	Date *time.Time `json:"date" bson:"date"`
}

type ULBInventoryUpdateMessageFilterV2 struct {
	Date       *DateRange `json:"date" bson:"date"`
	NotifyType string     `json:"notifyType" bson:"notifyType"`
	Month      string     `json:"month" bson:"month"`
	Year       string     `json:"year" bson:"year"`
	LastDate   string     `json:"lastDate" bson:"lastDate"`
}

type ULBMasterReportV2Filter struct {
	Status []string `json:"status" bson:"status,omitempty"`
	Months []string `json:"months" bson:"months,omitempty"`
	Year   int      `json:"year" bson:"year,omitempty"`
}

type RefULBMasterReportV2 struct {
	ULB    `bson:",inline"`
	Months []struct {
		Month            `bson:",inline"`
		StartDate        *time.Time `json:"startDate" bson:"startDate,omitempty"`
		EndDate          *time.Time `json:"endDate" bson:"endDate,omitempty"`
		CompostGenerated struct {
			Quantity float64 `json:"quantity" bson:"quantity,omitempty"`
		} `json:"compostGenerated" bson:"compostGenerated,omitempty"`
		Sale struct {
			Self struct {
				Quantity      float64 `json:"quantity" bson:"quantity,omitempty"`
				CustomerCount int64   `json:"customerCount" bson:"customerCount,omitempty"`
				Amount        float64 `json:"amount" bson:"amount,omitempty"`
			} `json:"self" bson:"Self,omitempty"`
			ULB struct {
				Quantity      float64 `json:"quantity" bson:"quantity,omitempty"`
				CustomerCount int64   `json:"customerCount" bson:"customerCount,omitempty"`
				Amount        float64 `json:"amount" bson:"amount,omitempty"`
			} `json:"ulb" bson:"ULB,omitempty"`
			Customer struct {
				Quantity      float64 `json:"quantity" bson:"quantity,omitempty"`
				CustomerCount int64   `json:"customerCount" bson:"customerCount,omitempty"`
				Amount        float64 `json:"amount" bson:"amount,omitempty"`
			} `json:"customer" bson:"Customer,omitempty"`
		} `json:"sale" bson:"sale,omitempty"`
	} `json:"months" bson:"months,omitempty"`
}
type Rejected struct {
	By     string `json:"by" bson:"by,omitempty"`
	ByType string `json:"byType,omitempty" form:"byType" bson:"byType,omitempty"`
}
type CO struct {
	Name     string `json:"name" bson:"name,omitempty"`
	MobileNo string `json:"mobile" bson:"mobile,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	UserName string `json:"userName" bson:"userName,omitempty"`
}
