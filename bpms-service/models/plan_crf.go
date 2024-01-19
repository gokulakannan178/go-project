package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CRF : ""
type CRF struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PlanID        string             `json:"planId" bson:"planId,omitempty"`
	PlanRegTypeID string             `json:"planRegTypeId" bson:"planRegTypeId,omitempty"`
	DeptID        string             `json:"deptId" bson:"deptId,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	Created       Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Log           []PlanCRFTimeline  `json:"log,omitempty"  bson:"log,omitempty"`
	Updated       []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefCRF : ""
type RefCRF struct {
	CRF `bson:",inline"`
	Ref struct {
		Plan        *RefPlan                 `json:"plan" bson:"plan,omitempty"`
		Department  *RefDepartment           `json:"department" bson:"department,omitempty"`
		ULB         RefULB                   `json:"ulb" bson:"ulb,omitempty"`
		PlanRegType *RefPlanRegistrationType `json:"planRegType" bson:"planRegType,omitempty"`
		ULBAddress  *RefAddress              `json:"ulbAddress" bson:"ulbAddress,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//CRFFilter : ""
type CRFFilter struct {
	Plan        []string `json:"plan,omitempty" bson:"plan,omitempty"`
	Status      []string `json:"status,omitempty" bson:"status,omitempty"`
	Dept        []string `json:"dept,omitempty" bson:"dept,omitempty"`
	PlanRegType []string `json:"planRegType,omitempty" bson:"planRegType,omitempty"`
	SortBy      string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder   int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//PlanCRFTimeline : ""
type PlanCRFTimeline struct {
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

//CRFInspection : ""
type CRFInspection struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PlanID        string             `json:"planId" bson:"planId,omitempty"`
	PlanRegTypeID string             `json:"planRegTypeId" bson:"planRegTypeId,omitempty"`
	DeptID        string             `json:"deptId" bson:"deptId,omitempty"`
	CheckListID   string             `json:"checkListId" bson:"checkListId,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	Time          *time.Time         `json:"time" bson:"time,omitempty"`
	Location      Location           `json:"location" bson:"location,omitempty"`
	Created       Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated       []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	Photo         string             `json:"photo"  bson:"photo,omitempty"`
	Remarks       string             `json:"remarks"  bson:"remarks,omitempty"`
}

//RefCRFInspection : ""
type RefCRFInspection struct {
	CRFInspection `bson:",inline"`
	Ref           struct {
		Checklist *RefDeptChecklist `json:"checklist" bson:"checklist,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//CRFInspectionFilter : ""
type CRFInspectionFilter struct {
	Status      []string `json:"status,omitempty" bson:"status,omitempty"`
	Dept        []string `json:"dept,omitempty" bson:"dept,omitempty"`
	PlanRegType []string `json:"planRegType,omitempty" bson:"planRegType,omitempty"`
	SortBy      string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder   int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//PlanCRFAccept : ""
type PlanCRFAccept struct {
	CRFID        string       `json:"crfId,omitempty" bson:"crfId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanCRFReapply : ""
type PlanCRFReapply struct {
	CRFID        string       `json:"crfId,omitempty" bson:"crfId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanCRFReject : ""
type PlanCRFReject struct {
	CRFID        string       `json:"crfId,omitempty" bson:"crfId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanCRFPostInspectionAccept : ""
type PlanCRFPostInspectionAccept struct {
	CRFID        string       `json:"crfId,omitempty" bson:"crfId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanCRFPostInspectionReject : ""
type PlanCRFPostInspectionReject struct {
	CRFID        string       `json:"crfId,omitempty" bson:"crfId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanCRFCertificateComplete : ""
type PlanCRFCertificateComplete struct {
	CRFID        string       `json:"crfId,omitempty" bson:"crfId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanCRFStartInspection : ""
type PlanCRFStartInspection struct {
	CRFID        string       `json:"crfId,omitempty" bson:"crfId,omitempty"`
	PlanTimeline PlanTimeline `json:"info,omitempty" bson:"info,omitempty"`
	Scenario     string       `json:"scenario" bson:"scenario,omitempty"`
}

//PlanCRFStartInspectionReqPayload : ""
type PlanCRFStartInspectionReqPayload struct {
	CRF  `bson:",inline"`
	Flow PlanCRFStartInspection `json:"flow,omitempty" bson:"flow,omitempty"`
}

//PlanCRFEndInspectionReqPayload : ""
type PlanCRFEndInspectionReqPayload struct {
	CRF  `bson:",inline"`
	Flow PlanCRFStartInspection `json:"flow,omitempty" bson:"flow,omitempty"`
}
