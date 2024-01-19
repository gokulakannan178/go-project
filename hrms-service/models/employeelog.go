package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeLog struct {
	ID             primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	DepartmentId   string             `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	BranchId       string             `json:"branchId,omitempty" bson:"branchId,omitempty"`
	DesignationId  string             `json:"designationId,omitempty" bson:"designationId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Desc           string             `json:"description" bson:"description,omitempty"`
	Action         struct {
		UserName string    `json:"userName,omitempty" bson:"username,omitempty"`
		UserType string    `json:"userType,omitempty" bson:"usertype,omitempty"`
		UserID   string    `json:"userId,omitempty" bson:"userId,omitempty"`
		Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
	} `json:"action,omitempty" bson:"action,omitempty"`
	Remark  string   `json:"remark,omitempty" bson:"remark,omitempty"`
	Created *Created `json:"createdOn" bson:"createdOn,omitempty"`
	Status  string   `json:"status" bson:"status,omitempty"`
}

type RefEmployeeLog struct {
	EmployeeLog `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeLog struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex," bson:"regex"`
}
