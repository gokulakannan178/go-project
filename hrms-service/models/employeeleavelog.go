package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeLeaveLog struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId     string             `json:"employeeId" bson:"employeeId,omitempty"`
	LeaveType      string             `json:"leaveType" bson:"leaveType,omitempty"`
	Value          int64              `json:"value,omitempty" bson:"value,omitempty"`
	Revert         bool               `json:"revert,omitempty" bson:"revert,omitempty"`
	CreateBy       string             `json:"createBy" bson:"createBy,omitempty"`
	CreateDate     *time.Time         `json:"createDate" bson:"createDate,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Date           *time.Time         `json:"date" bson:"date,omitempty"`
	Remarks        string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Created        *Created           `json:"created"  bson:"created,omitempty"`
}

type RefEmployeeLeaveLog struct {
	EmployeeLeaveLog `bson:",inline"`
	Ref              struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeLeaveLog struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId     []string `json:"employeeId" bson:"employeeId,omitempty"`
	Regex          struct {
		Name      string `json:"name" bson:"name"`
		LeaveType string `json:"leaveType" bson:"leaveType"`
	} `json:"regex" bson:"regex"`
}

type EmployeeLeaveLogCount struct {
	OrganisationId string `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId     string `json:"employeeId" bson:"employeeId,omitempty"`
	LeaveType      string `json:"leaveType" bson:"leaveType,omitempty"`
}

type RefEmployeeLeaveLogCount struct {
	TotalLeave int64 `json:"totalLeave" bson:"totalLeave,omitempty"`
}
