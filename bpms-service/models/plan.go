package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Plan : ""
type Plan struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Creator  struct {
		ID   string `json:"id" bson:"id,omitempty"`
		Type string `json:"type" bson:"type,omitempty"`
	} `json:"creator" bson:"creator,omitempty"`
	RegType        string           `json:"regType" bson:"regType,omitempty"`
	PlotArea       float64          `json:"plotArea" bson:"plotArea,omitempty"`
	RoadType       string           `json:"roadType" bson:"roadType,omitempty"`
	MOS            float64          `json:"mos" bson:"mos,omitempty"`
	GroundCoverage float64          `json:"groundCoverage" bson:"groundCoverage,omitempty"`
	FAR            float64          `json:"far" bson:"far,omitempty"`
	BuildingHeight float64          `json:"buildingHeight" bson:"buildingHeight,omitempty"`
	Floors         []PlanFLoor      `json:"floors" bson:"floors,omitempty"`
	OtherDetails   PlanOtherDetails `json:"otherDetails" bson:"otherDetails,omitempty"`
	Owner          PlanOwner        `json:"owner" bson:"owner,omitempty"`
	Address        Address          `json:"address" bson:"address,omitempty"`
	Status         string           `json:"status" bson:"status,omitempty"`
	Created        Created          `json:"created,omitempty"  bson:"created,omitempty"`
	Updated        []Updated        `json:"updated,omitempty"  bson:"updated,omitempty"`
	SaveType       string           `json:"saveType,omitempty"  bson:"-"`
	Log            []PlanTimeline   `json:"log,omitempty"  bson:"log,omitempty"`
	Remarks        string           `json:"remarks" bson:"remarks,omitempty"`
	URL            string           `json:"url" bson:"url,omitempty"`
}

//PlanOtherDetails : ""
type PlanOtherDetails struct {
	IsParking            string   `json:"isParking" bson:"uniqueId,omitempty"`
	ParkingArea          string   `json:"parkingArea" bson:"parkingArea,omitempty"`
	Amenities            []string `json:"amenities" bson:"amenities,omitempty"`
	Components           string   `json:"components" bson:"components,omitempty"`
	FireFightingReq      string   `json:"fireFightingReq" bson:"fireFightingReq,omitempty"`
	GreenBuildingNorms   string   `json:"greenBuildingNorms" bson:"greenBuildingNorms,omitempty"`
	RainWaterHarvesting  string   `json:"rainWaterHarvesting" bson:"rainWaterHarvesting,omitempty"`
	WaterSupply          string   `json:"waterSupply" bson:"waterSupply,omitempty"`
	WaterTreatment       string   `json:"waterTreatment" bson:"waterTreatment,omitempty"`
	SolidWasteManagement string   `json:"solidWasteManagement" bson:"solidWasteManagement,omitempty"`
	IsNOCRequired        string   `json:"isNOCRequired" bson:"isNOCRequired,omitempty"`
}

//PlanFLoor : ""
type PlanFLoor struct {
	FloorNo           string  `json:"floorNo" bson:"floorNo,omitempty"`
	FloorNoName       string  `json:"floorNoName" bson:"floorNoName,omitempty"`
	Area              float64 `json:"area" bson:"area,omitempty"`
	OccupancyTypeID   string  `json:"occupancyTypeId" bson:"occupancyTypeId,omitempty"`
	OccupancyTypeName string  `json:"occupancyTypeName" bson:"occupancyTypeName,omitempty"`
	RoofTypeID        string  `json:"roofTypeId" bson:"roofTypeId,omitempty"`
	RoofTypeName      string  `json:"roofTypeName" bson:"roofTypeName,omitempty"`
}

//PlanOwner : ""
type PlanOwner struct {
	Name   string     `json:"name" bson:"name,omitempty"`
	DOB    *time.Time `json:"age" bson:"age,omitempty"`
	Mobile string     `json:"mobile" bson:"mobile,omitempty"`
	Email  string     `json:"email" bson:"email,omitempty"`
	Gender string     `json:"gender" bson:"gender,omitempty"`
}

//RefPlan : ""
type RefPlan struct {
	Plan `bson:",inline"`
	Ref  struct {
		RegType *PlanRegistrationType `json:"regType,omitempty" bson:"regType,omitempty"`
		Address *RefAddress           `json:"address,omitempty" bson:"address,omitempty"`
		Creator *Applicant            `json:"creator,omitempty" bson:"creator,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PlanFilter : ""
type PlanFilter struct {
	Status       []string       `json:"status,omitempty" bson:"status,omitempty"`
	Address      *AddressSearch `json:"address"`
	Applicant    []string       `json:"applicant"`
	Organisation []string       `json:"organisation,omitempty" bson:"organisation,omitempty"`
	RegType      []string       `json:"regType,omitempty" bson:"regType,omitempty"`
	SortBy       string         `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder    int            `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//PlanTimeline : ""
type PlanTimeline struct {
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

//PlanFailScrutiny : ""
type PlanFailScrutiny struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanPassScrutiny : ""
type PlanPassScrutiny struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//ProceedPCP : ""
type ProceedPCP struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//MakePCPDefective : ""
type MakePCPDefective struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PCPAccept : ""
type PCPAccept struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//DeptApprovalAccept : ""
type DeptApprovalAccept struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//DeptApprovalReject : ""
type DeptApprovalReject struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//CCAccept : ""
type CCAccept struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//CCReject : ""
type CCReject struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//MakePayment : ""
type MakePayment struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//ReapplyDefective : ""
type ReapplyDefective struct {
	PlanID       string       `json:"planId,omitempty" bson:"planId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}
