package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//DeptChecklist : ""
type DeptChecklist struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	DeptID   string             `json:"deptId" bson:"deptId,omitempty"`
	IsPhoto  string             `json:"isPhoto" bson:"isPhoto,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefDeptChecklist : ""
type RefDeptChecklist struct {
	DeptChecklist `bson:",inline"`
	Ref           struct {
		Department *Department `json:"department,omitempty" bson:"department,omitempty"`
		ULB        *ULB        `json:"ulb,omitempty" bson:"ulb,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DeptChecklistFilter : ""
type DeptChecklistFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	Dept      []string `json:"dept,omitempty" bson:"dept,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
