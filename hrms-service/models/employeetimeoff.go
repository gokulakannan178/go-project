package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeTimeOff struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId     string             `json:"employeeId" bson:"employeeId,omitempty"`
	LeaveType      string             `json:"leaveType" bson:"leaveType,omitempty"`
	PaidType       string             `json:"paidType" bson:"paidType,omitempty"`
	StartDate      *time.Time         `json:"startDate" bson:"startDate,omitempty"`
	EndDate        *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	NumberOfDays   int64              `json:"numberOfDays,omitempty" bson:"numberOfDays"`
	RequestDate    *time.Time         `json:"requestDate,omitempty" bson:"requestDate,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	Attachment     string             `json:"attachment,omitempty" bson:"attachment,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Created        *Created           `json:"created"  bson:"created,omitempty"`
	ApprovedBy     string             `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
	ApprovedDate   *time.Time         `json:"approvedDate,omitempty" bson:"approvedDate,omitempty"`
	RejectedBy     string             `json:"rejectedBy,omitempty" bson:"rejectedBy,omitempty"`
	RejectedDate   *time.Time         `json:"rejectedDate,omitempty" bson:"rejectedDate,omitempty"`
	Remarks        string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Revoke         string             `json:"revoke,omitempty" bson:"revoke,omitempty"`
	Branch         *Branch            `json:"branch,omitempty" bson:"branch,omitempty"`
	Designation    *Designation       `json:"designation,omitempty" bson:"designation,omitempty"`
}

type RefEmployeeTimeOff struct {
	EmployeeTimeOff `bson:",inline"`
	Ref             struct {
		OrganisationId Organisation   `json:"organisationId" bson:"organisationId,omitempty"`
		ApprovedUser   *User          `json:"approvedUser" bson:"approvedUser,omitempty"`
		LeaveType      *LeaveMaster   `json:"leaveType" bson:"leaveType,omitempty"`
		Employee       *Employee      `json:"employee" bson:"employee,omitempty"`
		EmployeeLeave  *EmployeeLeave `json:"employeeLeave" bson:"employeeLeave,omitempty"`
		RejectedBy     *User          `json:"rejectedUser" bson:"rejectedUser,omitempty"`
		Branch         *Branch        `json:"branch" bson:"branch,omitempty"`
		Designation    *Designation   `json:"designation" bson:"designation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeTimeOff struct {
	Status            []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId    []string `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId        []string `json:"employeeId" bson:"employeeId,omitempty"`
	RevokeRequest     []string `json:"revokeRequest" bson:"revokeRequest,omitempty"`
	OmitStatus        []string `json:"omitStatus" bson:"omitStatus,omitempty"`
	OmitRevokeRequest []string `json:"omitRevokeRequest" bson:"omitRevokeRequest,omitempty"`
	LeaveType         []string `json:"leaveType" bson:"leaveType,omitempty"`
	Manager           string   `json:"manager" bson:"manager,omitempty"`
	Regex             struct {
		EmployeeName string `json:"employeeName,omitempty" bson:"employeeName,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
type EmployeeTimeOffCount struct {
	OrganisationId string `json:"organisationId" bson:"organisationId,omitempty"`
	EmployeeId     string `json:"employeeId" bson:"employeeId,omitempty"`
	TimeOffType    string `json:"timeoffType" bson:"timeoffType,omitempty"`
}

type RefEmployeeTimeOffCount struct {
	TotalTimeOff int64 `json:"totalTimeOff" bson:"totalTimeOff,omitempty"`
}
type ReviewedEmployeeTimeOff struct {
	EmployeeTimeOff string `json:"employeeTimeOff,omitempty" bson:"employeeTimeOff,omitempty"`
	ReviewedBy      string `json:"reviewedBy" bson:"reviewedBy,omitempty"`
	Remarks         string `json:"remarks,omitempty" bson:"remarks,omitempty"`
}
