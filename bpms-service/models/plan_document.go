package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PlanDocument : ""
type PlanDocument struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	OrgID      string             `json:"orgId,omitempty"  bson:"orgId,omitempty"`
	OrgType    string             `json:"orgType,omitempty"  bson:"orgType,omitempty"`
	PlanID     string             `json:"planId,omitempty"  bson:"planId,omitempty"`
	DocID      string             `json:"docId,omitempty"  bson:"docId,omitempty"`
	URL        string             `json:"url"  bson:"url,omitempty"`
	IssuedDate *time.Time         `json:"issuedDate"  bson:"issuedDate,omitempty"`
}

//RefPlanDocument : ""
type RefPlanDocument struct {
	PlanDocument `bson:",inline"`
	Ref          struct {
		ULB        *ULB             `json:"ulb,omitempty" bson:"ulb,omitempty"`
		Department *ULB             `json:"department,omitempty" bson:"department,omitempty"`
		Address    *RefAddress      `json:"address,omitempty" bson:"address,omitempty"`
		Plan       *Plan            `json:"plan,omitempty" bson:"plan,omitempty"`
		Doc        *PlanReqDocument `json:"doc,omitempty" bson:"doc,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PlanDocumentFilter : ""
type PlanDocumentFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	Plan      []string `json:"plan,omitempty" bson:"plan,omitempty"`
	Org       []string `json:"org,omitempty" org:"status,omitempty"`
	OrgType   []string `json:"orgType,omitempty" orgType:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//PlanDocumentFilter : ""
type GetPendingPlanDocumentFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	Plan      string   `json:"plan,omitempty" bson:"plan,omitempty"`
	Org       string   `json:"org,omitempty" org:"status,omitempty"`
	OrgType   string   `json:"orgType,omitempty" orgType:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
