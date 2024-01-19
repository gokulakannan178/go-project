package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//PlanReqDocument : ""
type PlanReqDocument struct {
	ID           primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Desc         string             `json:"desc" bson:"desc,omitempty"`
	Address      Address            `json:"address" bson:"address,omitempty"`
	DepartmentId string             `json:"departmentId,omitempty"  bson:"departmentId,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
	Created      Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated      []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	OrgID        string             `json:"orgId,omitempty"  bson:"orgId,omitempty"`
	OrgType      string             `json:"orgType,omitempty"  bson:"orgType,omitempty"`
}

//RefPlanReqDocument : ""
type RefPlanReqDocument struct {
	PlanReqDocument `bson:",inline"`
	Ref             struct {
		ULB *ULB `json:"ulb,omitempty" bson:"ulb,omitempty"`
		//	Department *ULB        `json:"department,omitempty" bson:"department,omitempty"`
		Department *Department `json:"department,omitempty" bson:"department,omitempty"`
		Address    *RefAddress `json:"address,omitempty" bson:"address,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PlanReqDocumentFilter : ""
type PlanReqDocumentFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	Org       []string `json:"org,omitempty" org:"status,omitempty"`
	OrgType   []string `json:"orgType,omitempty" orgType:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
