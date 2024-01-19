package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeLeave struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId     string             `json:"employeeId" bson:"employeeId,omitempty"`
	LeaveType      string             `json:"leaveType" bson:"leaveType,omitempty"`
	Value          int64              `json:"value,omitempty" bson:"value,omitempty"`
	Date           *time.Time         `json:"date" bson:"date,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Created        *Created           `json:"created"  bson:"created,omitempty"`
}

type RefEmployeeLeave struct {
	EmployeeLeave `bson:",inline"`
	Ref           struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeLeave struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	LeaveType      []string `json:"leaveType" bson:"leaveType,omitempty"`
	EmployeeId     []string `json:"employeeId" bson:"employeeId,omitempty"`
	Regex          struct {
		Name      string `json:"name" bson:"name"`
		LeaveType string `json:"leaveType" bson:"leaveType"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type FilterEmployeeLeaveList struct {
	EmployeeId string `json:"employeeId" bson:"employeeId,omitempty"`
	IsZero     string `json:"isZero" bson:"isZero,omitempty"`
}
type UpdateEmployeeLeave struct {
	EmployeeId string `json:"employeeId" bson:"employeeId,omitempty"`
	LeaveType  string `json:"leaveType" bson:"leaveType,omitempty"`
	Value      int64  `json:"value,omitempty" bson:"value,omitempty"`
	Remarks    string `json:"remarks,omitempty" bson:"remarks,omitempty"`
}
type EmployeeLeaveCount struct {
	OrganisationId string `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId     string `json:"employeeId" bson:"employeeId,omitempty"`
	LeaveType      string `json:"leaveType" bson:"leaveType,omitempty"`
}

type RefEmployeeLeaveCount struct {
	TotalLeave int64 `json:"totalLeave" bson:"totalLeave,omitempty"`
}
type EmployeeLeaveList struct {
	Employee      string `json:"employee" bson:"employee,omitempty"`
	EmployeeLeave []struct {
		//Leave struct {
		UniqueID     string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
		Name         string `json:"name,omitempty" bson:"name,omitempty"`
		Value        int64  `json:"value" bson:"value,omitempty"`
		NumberOfDays int64  `json:"numberOfDays,omitempty" bson:"numberOfDays,omitempty"`
		LeaveType    string `json:"leaveType" bson:"leaveType,omitempty"`
		//	} `json:"leave,omitempty" bson:"leave,omitempty"`
	} `json:"employeeLeave,omitempty" bson:"employeeLeave,omitempty"`
}
type EmployeeLeaveListV2 struct {
	Employee      string `json:"employee" bson:"employee,omitempty"`
	EmployeeLeave struct {
		Annual          *EmployeeLeaveValue `json:"Annual,omitempty"  bson:"Annual,omitempty"`
		Engagement      *EmployeeLeaveValue `json:"Engagement,omitempty"  bson:"Engagement,omitempty"`
		RelativeFuneral *EmployeeLeaveValue `json:"RelativeFuneral,omitempty"  bson:"RelativeFuneral,omitempty"`
		Paternity       *EmployeeLeaveValue `json:"Paternity,omitempty"  bson:"Paternity,omitempty"`
		SickLeave       *EmployeeLeaveValue `json:"SickLeave,omitempty"  bson:"SickLeave,omitempty"`
		Wedding         *EmployeeLeaveValue `json:"Wedding,omitempty"  bson:"Wedding,omitempty"`
		UnpaidTimeOff   *EmployeeLeaveValue `json:"UnpaidTimeOff,omitempty"  bson:"UnpaidTimeOff,omitempty"`
	} `json:"employeeLeave,omitempty" bson:"employeeLeave,omitempty"`
}
type EmployeeLeaveValue struct {
	UniqueID  string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	LeaveType string `json:"leaveType" bson:"leaveType,omitempty"`
	Value     int64  `json:"value" bson:"value,omitempty"`
}
