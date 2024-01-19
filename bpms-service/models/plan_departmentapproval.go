package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//PlanDepartmentApproval : ""
type PlanDepartmentApproval struct {
	ID               primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PlanID           string             `json:"planId" bson:"planId,omitempty"`
	DepartmentID     string             `json:"departmentId" bson:"departmentId,omitempty"`
	DepartmentTypeID string             `json:"departmentTypeId" bson:"departmentTypeId,omitempty"`
	Status           string             `json:"status" bson:"status,omitempty"`
	Check            string             `json:"check" bson:"check,omitempty"`
	Created          Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated          []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefPlanDepartmentApproval : ""
type RefPlanDepartmentApproval struct {
	PlanDepartmentApproval `bson:",inline"`
	Ref                    struct {
		Organisation   *ULB            `json:"organisation,omitempty" bson:"organisation,omitempty"`
		DepartmentType *DepartmentType `json:"departmentType,omitempty" bson:"departmentType,omitempty"`
		Address        *RefAddress     `json:"address,omitempty" bson:"address,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PlanDepartmentApprovalFilter : ""
type PlanDepartmentApprovalFilter struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	Organisation   []string `json:"organisation,omitempty" organisation:"status,omitempty"`
	DepartmentType []string `json:"departmentType,omitempty" bson:"departmentType,omitempty"`
	SortBy         string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder      int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//GetAPlanDeptsApproval : ""
type GetAPlanDeptsApproval struct {
	Department            `bson:",inline"`
	PlanRegistrationTypes []struct {
		PlanRegistrationType   `bson:",inline"`
		PlanDepartmentApproval *PlanDepartmentApproval `json:"plandeptapproval" bson:"plandeptapproval,omitempty"`
	} `json:"planregtypes,omitempty" bson:"planregtypes,omitempty"`
}

//GetAPlanDeptsApprovalV2 : ""
type GetAPlanDeptsApprovalV2 struct {
	PlanRegistrationType `bson:",inline"`
	Departments          []struct {
		Department             `bson:",inline"`
		PlanDepartmentApproval *PlanDepartmentApproval `json:"plandeptapproval" bson:"plandeptapproval,omitempty"`
	} `json:"departments,omitempty" bson:"departments,omitempty"`
}

//GetAPlanDeptsApprovalV3 : ""
type GetAPlanDeptsApprovalV3 struct {
	PlanRegistrationType `bson:",inline"`
	Departments          []struct {
		DepartmentType         `bson:",inline"`
		PlanDepartmentApproval *PlanDepartmentApproval `json:"plandeptapproval" bson:"plandeptapproval,omitempty"`
	} `json:"departments,omitempty" bson:"departments,omitempty"`
}
