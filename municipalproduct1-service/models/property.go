package models

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Property : ""
type Property struct {
	ID                    primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	MunicipalityID        string             `json:"municipalityId" bson:"municipalityId,omitempty"`
	ApplicationNo         string             `json:"applicationNo" bson:"applicationNo,omitempty"`
	YOA                   string             `json:"yoa" bson:"yoa,omitempty"`
	Mobile                string             `json:"mobile" bson:"mobile,omitempty"`
	NewHoldingNumber      string             `json:"newHoldingNumber" bson:"newHoldingNumber,omitempty"`
	OldHoldingNumber      string             `json:"oldHoldingNumber" bson:"oldHoldingNumber,omitempty"`
	UniqueID              string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	RoadName              string             `json:"roadName" bson:"roadName,omitempty"`
	PropertyTypeID        string             `json:"propertyTypeId" bson:"propertyTypeId,omitempty"`
	RoadTypeID            string             `json:"roadTypeId" bson:"roadTypeId,omitempty"`
	DOA                   *time.Time         `json:"doa" bson:"doa,omitempty"`
	EndDate               *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	AreaOfPlot            float64            `json:"areaOfPlot" bson:"areaOfPlot,omitempty"`
	BuiltUpArea           float64            `json:"builtUpArea" bson:"builtUpArea,omitempty"`
	IsRainWaterHarvesting string             `json:"isRainWaterHarvesting" bson:"isRainWaterHarvesting,omitempty"`
	IsGovtProperty        string             `json:"isGovtProperty" bson:"isGovtProperty,omitempty"`
	Address               Address            `json:"address" bson:"address,omitempty"`
	IsPrincipleMainRoad   string             `json:"isPrincipleMainRoad" bson:"isPrincipleMainRoad,omitempty"`
	BPLCategory           string             `json:"bplCategory" bson:"bplCategory,omitempty"`
	GeoTagged             string             `json:"geoTagged" bson:"geoTagged,omitempty"`
	OwnerID               []string           `json:"ownerId" bson:"ownerIds,omitempty"` // multiple owners
	Status                string             `json:"status" bson:"status,omitempty"`
	IsMatched             string             `json:"isMatched" bson:"isMatched,omitempty"`
	From                  *time.Time         `json:"from" bson:"from,omitempty"`
	Created               Created            `json:"created,omitempty" bson:"created,omitempty"`
	Updated               []Updated          `json:"updated" bson:"updated,omitempty"`
	Checklist             []Checklist        `json:"checklist" bson:"checklist,omitempty"`
	Owner                 []PropertyOwner    `json:"owner" bson:"-"`
	Floors                []PropertyFloor    `json:"floors" bson:"-"`
	Legacy                struct {
		IsLegacy       bool               `json:"isLegacy" bson:"isLegacy,omitempty"`
		LegacyProperty *RegLegacyProperty `json:"legacyProperty" bson:"legacyProperty,omitempty"`
	} `json:"legacy" bson:"-"`
	MobileTower struct {
		IsMobileTower       bool                 `json:"isMobileTower" bson:"isMobileTower,omitempty"`
		PropertyMobileTower *PropertyMobileTower `json:"propertyMobileTower" bson:"propertyMobileTower,omitempty"`
	} `json:"mobileTower" bson:"-"`
	Log                         []PropertyTimeline          `json:"log" bson:"log,omitempty"`
	MunicipalityWaterConnection string                      `json:"municipalityWaterConnection" bson:"municipalityWaterConnection,omitempty"`
	IsBoringChargePayed         bool                        `json:"isBoringChargePayed" bson:"isBoringChargePayed,omitempty"`
	IsFormFeePayed              bool                        `json:"isFormFeePayed" bson:"isFormFeePayed,omitempty"`
	Demand                      UpdateDemand                `json:"demand" bson:"demand"`
	Collection                  UpdateCollection            `json:"collection" bson:"collection"`
	Penalty                     UpdatePenalty               `json:"penalty" bson:"penalty"`
	Rebate                      UpdateRebate                `json:"rebate" bson:"rebate"`
	Advance                     float64                     `json:"advance" bson:"advance"`
	PreviousCollection          PreviousCollection          `json:"previousCollection" bson:"previousCollection"`
	GIS                         PropertyGISTagging          `json:"gis" bson:"gis"`
	SortOrder                   int64                       `json:"sortOrder" bson:"sortOrder,omitempty"`
	NDemand                     PropertyTaxTotalDemand      `json:"ndemand" bson:"ndemand,omitempty"`
	NCollection                 PropertyTaxTotalCollection  `json:"ncollection" bson:"ncollection,omitempty"`
	NOutstanding                PropertyTaxTotalOutStanding `json:"noutstanding" bson:"noutstanding,omitempty"`
	NPending                    PropertyTaxTotalPending     `json:"npending" bson:"npending,omitempty"`
	ParkPenalty                 bool                        `json:"parkPenalty" bson:"parkPenalty,omitempty"`
	PropertyDocument            []PropertyDocuments         `json:"propertyDocument" bson:"-"`
	Images                      []string                    `json:"images" bson:"images,omitempty"`
	Documents                   []string                    `json:"documents" bson:"documents,omitempty"`
	SPD                         StoredPropertyDemand        `json:"spd" bson:"spd,omitempty"`
	Picture                     struct {
		Location Location `json:"location" bson:"location,omitempty"`
		Image    string   `json:"image" bson:"image,omitempty"`
	} `json:"picture" bson:"picture,omitempty"`
	OldPID                  string                   `json:"oldPID" bson:"oldPID,omitempty"`
	HoldingStatus           string                   `json:"holdingStatus" bson:"holdingStatus,omitempty"`
	Reason                  string                   `json:"reason" bson:"reason,omitempty"`
	Environment             string                   `json:"environment" bson:"environment,omitempty"`
	NewUniqueID             string                   `json:"newUniqueId" bson:"newUniqueId,omitempty"`
	OldUniqueID             string                   `json:"oldUniqueId" bson:"oldUniqueId,omitempty"`
	UserCharge              UserCharge               `json:"userCharge" bson:"userCharge,omitempty"`
	UserChargeRejectedInfo  UserChargePaymentsAction `json:"userchargerejectedInfo" bson:"userchargerejectedInfo,omitempty"`
	UserChargeVerifiedInfo  UserChargePaymentsAction `json:"userchargeverifiedInfo" bson:"userchargeverifiedInfo,omitempty"`
	IsViewProperty          string                   `json:"isViewProperty" bson:"isViewProperty,omitempty"`
	PlsContactAdministrator bool                     `json:"plsContactAdministrator" bson:"plsContactAdministrator"`
	WalletBalanceAmount     float64                  `json:"walletBalanceAmount" bson:"walletBalanceAmount,omitempty"`
	Sumary                  PropertyDemandSummary    `json:"summary" bson:"summary,omitempty"`
}
type UserCharge struct {
	CategoryID   string     `json:"categoryId" bson:"categoryId,omitempty"`
	DOA          *time.Time `json:"doa" bson:"doa,omitempty"`
	IsUserCharge string     `json:"isUserCharge" bson:"isUserCharge,omitempty"`
	Status       string     `json:"status" bson:"status,omitempty"`
	Createdby    CreatedV2  `json:"createdBy" bson:"createdBy,omitempty"`
}

type PropertyGISTagging struct {
	Images       []string   `json:"images" bson:"images,omitempty"`
	Location     Location   `json:"location" bson:"location,omitempty"`
	Time         *time.Time `json:"time" bson:"time,omitempty"`
	PropertyID   string     `json:"propertyId" bson:"propertyId,omitempty"`
	PropertyType string     `json:"propertyType" bson:"propertyType,omitempty"`
}
type PreviousCollection struct {
	Amount       float64 `json:"amount" bson:"amount"`
	IsCalculated bool    `json:"isCalculated" bson:"isCalculated"`
}

// PropertyHoldingStatus : ""
type PropertyHoldingStatus struct {
	PropertyID    string `json:"propertyId" bson:"propertyId,omitempty"`
	HoldingStatus string `json:"holdingStatus" bson:"holdingStatus,omitempty"`
	Reason        string `json:"reason" bson:"reason,omitempty"`
}

//PropertyTimeline : ""
type PropertyTimeline struct {
	On *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	By struct {
		ID   string `json:"id,omitempty" bson:"id,omitempty"`
		Type string `json:"type,omitempty" bson:"type,omitempty"`
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"by,omitempty" bson:"by,omitempty"`
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	TypeLabel string `json:"typeLabel,omitempty" bson:"typeLabel,omitempty"`
	Remarks   string `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

//ActivateProperty : ""
type ActivateProperty struct {
	PropertyID       string           `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	PropertyTimeline PropertyTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario         string           `json:"scenario" bson:"scenario,omitempty"`
}

//RejectProperty : ""
type RejectProperty struct {
	PropertyID       string           `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	PropertyTimeline PropertyTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario         string           `json:"scenario" bson:"scenario,omitempty"`
}

//RefProperty :""
type RefProperty struct {
	Property `bson:",inline"`
	Ref      struct {
		ReassessmentFloors    []RefPropertyFloor         `json:"reassessmentFloors" bson:"reassessmentFloors,omitempty"`
		ReassessmentOwners    []RefPropertyOwner         `json:"reassessmentOwners" bson:"reassessmentOwners,omitempty"`
		PropertyOwner         []RefPropertyOwner         `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		Floors                []RefPropertyFloor         `json:"floors" bson:"floors,omitempty"`
		ReassessmentDocuments []RefPropertyDocuments     `json:"reassessmentDocuments" bson:"reassessmentDocuments,omitempty"`
		Documents             []RefPropertyDocuments     `json:"documents" bson:"documents,omitempty"`
		Address               RefAddress                 `json:"address" bson:"address,omitempty"`
		Wallet                RefPropertyWallet          `json:"wallet" bson:"wallet,omitempty"`
		PropertyType          *RefPropertyType           `json:"propertyType" bson:"propertyType,omitempty"`
		RoadType              *RefRoadType               `json:"roadType" bson:"roadType,omitempty"`
		YOA                   *RefFinancialYear          `json:"yoa" bson:"yoa,omitempty"`
		User                  *RefUser                   `json:"user" bson:"user,omitempty"`
		MunicipalType         *RefMunicipalType          `json:"municipalType" bson:"municipalType,omitempty"`
		Demand                OverallPropertyDemand      `json:"demand" bson:"demand,omitempty"`
		Activator             User                       `json:"activator" bson:"activator,omitempty"`
		UserChargeActivator   User                       `json:"userchargeactivator" bson:"userchargeactivator,omitempty"`
		UserChargeRejector    User                       `json:"userchargerejector" bson:"userchargerejector,omitempty"`
		UserChargeCreator     User                       `json:"userChargeCreator" bson:"userChargeCreator,omitempty"`
		Ratemaster            UserChargeRateMaster       `json:"userchargeratemaster" bson:"userchargeratemaster,omitempty"`
		CategoryID            UserChargeCategory         `json:"userchargecategory" bson:"userchargecategory,omitempty"`
		Basic                 PropertyPaymentDemandBasic `json:"basic" bson:"basic,omitempty"`
		Collections           struct {
			ArrearTax      float64 `json:"arrearTax" bson:"arrearTax,omitempty"`
			ArrearPenalty  float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
			ArrearRebate   float64 `json:"arrearRebate" bson:"arrearRebate,omitempty"`
			CurrentTax     float64 `json:"currentTax" bson:"currentTax,omitempty"`
			CurrentPenalty float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
			CurrentRebate  float64 `json:"currentRebate" bson:"currentRebate,omitempty"`
			TotalTax       float64 `json:"totalTax" bson:"totalTax,omitempty"`
			OtherDemand    float64 `json:"otherDemand" bson:"otherDemand,omitempty"`
		} `json:"collections" bson:"collections,omitempty"`
		PropertyPayments struct {
			Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
			FormFee float64 `json:"formFee" bson:"formFee,omitempty"`
		} `json:"propertyPayments" bson:"propertyPayments,omitempty"`
		Payments struct {
			Payments float64 `json:"payments" bson:"payments,omitempty"`
			Amount   float64 `json:"amount" bson:"amount,omitempty"`
		} `json:"payments" bson:"payments,omitempty"`
		// Inc           func(int) int
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

// DayWisePropertyDemandChartFilter : ""
type DayWisePropertyDemandChartFilter struct {
	PropertyFilter `bson:",inline"`
	Status         []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

// DayWisePropertyDeamndChart : ""
type DayWisePropertyDemandChart struct {
	Records []struct {
		ID            int     `json:"day" bson:"_id,omitempty"`
		PropertyCount int64   `json:"propertyCount" bson:"propertyCount,omitempty"`
		Amount        float64 `json:"amount" bson:"amount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
	Total float64 `json:"total" bson:"total,omitempty"`
}

type PropertyOverallDemandReport struct {
	Properties []RefProperty    `json:"properties" bson:"properties,omitempty"`
	CFY        RefFinancialYear `json:"cfy" bson:"cfy,omitempty"`
}

func (prop *RefProperty) Inc(a int) int {
	return a + 1
}

func (prop *RefProperty) Breakstring(a string) string {
	// return strings.Replace(a, ",", "<br>", -1)
	data := strings.Split(a, ",")
	if len(data) > 0 {
		return data[0]
	}
	return "NA"

}

//PropertyFilter : ""
type PropertyFilter struct {
	Status           []string       `json:"status"`
	UserChargeStatus []string       `json:"userchargestatus"`
	Type             []string       `json:"type"`
	Address          *AddressSearch `json:"address"`
	DemandCalc       bool           `json:"demandCalc"`
	UniqueID         []string       `json:"uniqueId"`
	IsLocation       bool           `json:"isLocation"`
	IsGeoTagged      string         `json:"isGeoTagged"`
	NeedRecordCount  bool           `json:"needRecordCount"`
	OmitZeroDemand   bool           `json:"omitZeroDemand"`
	Date             *time.Time     `json:"date"`
	SortBy           string         `json:"sortBy"`
	SortOrder        int            `json:"sortOrder"`
	DemandSortBy     string         `json:"demandSortBy"`
	DemandSortOrder  int            `json:"demandSortOrder"`
	IsUserCharge     []string       `json:"isUserCharge" bson:"isUserCharge,omitempty"`
	IsMatched        []string       `json:"isMatched" bson:"isMatched,omitempty"`
	OldHoldingNumber []string       `json:"oldHoldingNumber" bson:"oldHoldingNumber,omitempty"`
	Regex            struct {
		Mobile        string `json:"mobile" bson:"mobile,omitempty"`
		PropertyNo    string `json:"propertyNo" bson:"propertyNo"`
		ApplicationNo string `json:"applicationNo" bson:"applicationNo"`
		OwnerName     string `json:"ownerName" bson:"ownerName"`
	} `json:"regex" bson:"regex"`
	AppliedRange *PropertyAppliedRange `json:"appliedRange" bson:"appliedRange"`
	RemoveLookup struct {
		PropertyOwner bool `json:"propertysOwner"`
		State         bool `json:"state"`
		District      bool `json:"district"`
		Village       bool `json:"village"`
		Zone          bool `json:"zone"`
		Ward          bool `json:"ward"`
		PropertyFloor bool `json:"propertyFloor"`
		PropertyType  bool `json:"propertyType"`
		RoadType      bool `json:"roadType"`
		MunicipalType bool `json:"municipalType"`
		Wallet        bool `json:"wallet"`
		User          bool `json:"user"`
		Activator     bool `json:"activator"`
	} `json:"removeLookup" bson:"removeLookup"`
}

type PropertyAppliedRange struct {
	From *time.Time `json:"from" bson:"from"`
	To   *time.Time `json:"to" bson:"to"`
}

// PropertyWardwiseDemandFilter : ""
type PropertyWardwiseDemandFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	Zone      []string `json:"zone" bson:"zone,omitempty"`
	Ward      []string `json:"ward" bson:"ward,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

// WardwiseDemand : ""
type WardwiseDemandandCollection struct {
	Ward       `bson:",inline"`
	Properties struct {
		TotalDemandArrear      float64 `json:"totalDemandArrear" bson:"totalDemandArrear"`
		TotalDemandCurrent     float64 `json:"totalDemandCurrent" bson:"totalDemandCurrent"`
		TotalDemandTax         float64 `json:"totalDemandTax" bson:"totalDemandTax"`
		TotalCollectionArrear  float64 `json:"totalCollectionArrear" bson:"totalCollectionArrear"`
		TotalCollectionCurrent float64 `json:"totalCollectionCurrent" bson:"totalCollectionCurrent"`
		TotalCollectionTax     float64 `json:"totalCollectionTax" bson:"totalCollectionTax"`
	} `json:"properties" bson:"properties"`
}

//Checklist : ""
type Checklist struct {
	UniqueID   string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	Label      string  `json:"label" bson:"label,omitempty"`
	IsDiscount string  `json:"isDiscount" bson:"isDiscount,omitempty"`
	Value      float64 `json:"value" bson:"value,omitempty"`
	Type       string  `json:"type" bson:"type,omitempty"`
}

//DashboardPropertyStatusFilter : ""
type DashboardPropertyStatusFilter struct {
	Status  []string       `json:"status"`
	Address *AddressSearch `json:"address"`
}

//DashboardPropertyStatus : ""
type DashboardPropertyStatus struct {
	Data struct {
		Init struct {
			Count int64 `json:"count" bson:"count,omitempty"`
		} `json:"init" bson:"Init,omitempty"`
		Rejected struct {
			Count int64 `json:"count" bson:"count,omitempty"`
		} `json:"rejected" bson:"Rejected,omitempty"`
		Active struct {
			Count int64 `json:"count" bson:"count,omitempty"`
		} `json:"active" bson:"Active,omitempty"`
	} `json:"data" bson:"data,omitempty"`
}

type PropertyPreviousYrCollection struct {
	UniqueID string  `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Amount   float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	Status   string  `json:"status" bson:"status,omitempty"`
}

// BasicPropertyUpdate : ""
type BasicPropertyUpdate struct {
	PropertyID string        `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	Address    Address       `json:"address,omitempty" bson:"address,omitempty"`
	Owner      PropertyOwner `json:"owner,omitempty" bson:"owner,omitempty"`
	UserName   string        `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType   string        `json:"userType,omitempty" bson:"userType,omitempty"`
	Proof      []string      `json:"proof,omitempty" bson:"proof,omitempty"`
	Remarks    string        `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

// BasicPropertyUpdateLog : ""
type BasicPropertyUpdateLog struct {
	PropertyID string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	UniqueID   string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous   struct {
		Address Address          `json:"address,omitempty" bson:"address,omitempty"`
		Owner   RefPropertyOwner `json:"owner,omitempty" bson:"owner,omitempty"`
	} `json:"previous,omitempty" bson:"previous,omitempty"`
	New struct {
		Address Address       `json:"address,omitempty" bson:"address,omitempty"`
		Owner   PropertyOwner `json:"owner,omitempty" bson:"owner,omitempty"`
	} `json:"new,omitempty" bson:"new,omitempty"`
	UserName      string   `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType      string   `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester     Updated  `json:"requester" bson:"requester,omitempty"`
	Action        Updated  `json:"action" bson:"action,omitempty"`
	Proof         []string `json:"proof,omitempty" bson:"proof,omitempty"`
	Status        string   `json:"status,omitempty" bson:"status,omitempty"`
	NewPropertyID string   `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string   `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

//  AcceptBasicPropertyUpdate : ""
type AcceptBasicPropertyUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectBasicPropertyUpdate : ""
type RejectBasicPropertyUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// FilterBasicPropertyUpdate : ""
type FilterBasicPropertyUpdate struct {
	PropertyID   []string `json:"propertyId"`
	UniqueId     []string `json:"uniqueId"`
	UserName     []string `json:"userName"`
	UserType     []string `json:"userType"`
	Approver     []string `json:"approver"`
	ApproverType []string `json:"approverType"`
	Status       []string `json:"status"`
	SearchText   struct {
		PropertyID string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	} `json:"searchText"`
}

type RefBasicPropertyUpdateLog struct {
	BasicPropertyUpdateLog `bson:",inline"`
	Ref                    struct {
		RequestedBy     User     `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User     `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		Previous        struct {
			Address RefAddress       `json:"address,omitempty" bson:"address,omitempty"`
			Owner   RefPropertyOwner `json:"owner,omitempty" bson:"owner,omitempty"`
		} `json:"previous,omitempty" bson:"previous,omitempty"`
		New struct {
			Address RefAddress    `json:"address,omitempty" bson:"address,omitempty"`
			Owner   PropertyOwner `json:"owner,omitempty" bson:"owner,omitempty"`
		} `json:"new,omitempty" bson:"new,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type PropertyTaxTotalDemand struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
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
type PropertyTaxTotalCollection struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
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
type PropertyTaxTotalOutStanding struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
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
type PropertyTaxTotalPending struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate  float64 `json:"rebate" bson:"rebate,omitempty"`
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
type PropertyLocation struct {
	PropertyID string   `json:"propertyId" bson:"propertyId,omitempty"`
	Location   Location `json:"location" bson:"location,omitempty"`
	GeoTagged  string   `json:"geoTagged" bson:"geoTagged,omitempty"`
}

//  PropertyUpdatePicture : ""
type PropertyPicture struct {
	PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
	Picture    struct {
		Location Location `json:"location" bson:"location,omitempty"`
		Image    string   `json:"image" bson:"image,omitempty"`
	} `json:"picture" bson:"picture,omitempty"`
}

// YearWisePropertyDemandChartFilter : ""
type YearWisePropertyDemandChartFilter struct {
	PropertyFilter `bson:",inline"`
	Status         []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

// YearWisePropertyDemandReportFilter : ""
type YearWisePropertyDemandReportFilter struct {
	PropertyFilter `bson:",inline"`
	FYID           string     `json:"fyId" bson:"fyId,omitempty"`
	Status         []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

// YearWisePropertyDemandReport : ""
type YearWisePropertyDemandReport struct {
	NoOfProperties float64 `json:"noOfProperties" bson:"noOfProperties,omitempty"`
	Demand         float64 `json:"demand" bson:"demand,omitempty"`
	Month          int     `json:"month" bson:"month,omitempty"`
}

// YearWisePropertyDemandReportFilter : ""
type YearWisePropertyCollectionReportFilter struct {
	PropertyFilter `bson:",inline"`
	FYID           string     `json:"fyId" bson:"fyId,omitempty"`
	Status         []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

// YearWisePropertyDemandReport : ""
type YearWisePropertyCollectionReport struct {
	PropertyCount float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
	TotalTax      float64 `json:"totalTax" bson:"totalTax,omitempty"`
	Month         int     `json:"month" bson:"month,omitempty"`
}

// PropertyUpdateLocationFilter :""
type PropertyUpdateLocationFilter struct {
	PropertyFilter `bson:",inline"`
}
type WardWisePropertyDemandAndCollectionReport struct {
	Report []struct {
		TotalAmount            float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
		PaidAmount             float64 `json:"paidAmount" bson:"paidAmount,omitempty"`
		UnPaidAmount           float64 `json:"unPaidAmount" bson:"unPaidAmount,omitempty"`
		OutstandingDemand      float64 `json:"outstandingDemand" bson:"outstandingDemand,omitempty"`
		CurrentDemand          float64 `json:"currentDemand" bson:"currentDemand,omitempty"`
		TotalDemand            float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		ArrearCollection       float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
		CurrentCollection      float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
		Penalty                float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate                 float64 `json:"rebate" bson:"rebate,omitempty"`
		AdvanceAmount          float64 `json:"advanceAmount" bson:"advanceAmount,omitempty"`
		TotalCollection        float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		TotalOutstandingDemand float64 `json:"totalOutstandingDemand" bson:"totalOutstandingDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

type UpdatePropertyTotalDemand struct {
	PropertyID  string  `json:"propertyId" bson:"propertyId,omitempty"`
	TotalAmount float64 `json:"totalAmount" bson:"totalAmount,omitempty"`
}

type UpdatePropertyUniqueID struct {
	UniqueIDs   []string `json:"uniqueIds" bson:"uniqueIds,omitempty"`
	UniqueID    string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	NewUniqueID string   `json:"newUniqueId" bson:"newUniqueId,omitempty"`
	OldUniqueID string   `json:"oldUniqueId" bson:"oldUniqueId,omitempty"`
}

type UpdatePropertyPayeeName struct {
	TnxID string `json:"tnxId" bson:"tnxId,omitempty"`
	Name  string `json:"name" bson:"name,omitempty"`
}
