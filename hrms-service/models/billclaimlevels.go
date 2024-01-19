package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BillclaimLevels : "Holds single billclaimlevels data"
type BillclaimLevels struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueID" bson:"uniqueID,omitempty"`
	Grade          string             `json:"grade"  bson:"grade,omitempty"`
	EmployeeId     string             `json:"employeeId" bson:"employeeId,omitempty"`
	Level          int64              `json:"level"  bson:"level,omitempty"`
	NoOfLevel      int64              `json:"noOflevel"  bson:"noOflevel,omitempty"`
	Organisation   string             `json:"organisation"  bson:"organisation,omitempty"`
	Bill           string             `json:"bill"  bson:"bill,omitempty"`
	Status         string             `json:"status"  bson:"status,omitempty"`
	AssignedBy     string             `json:"assignedBy" bson:"assignedBy,omitempty"`
	ApprovedBy     string             `json:"approvedBy" bson:"approvedBy,omitempty"`
	ApprovedDate   *time.Time         `json:"approvedDate" bson:"approvedDate,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time         `json:"date" bson:"date,omitempty"`
	RejectedBy     string             `json:"rejectedBy" bson:"rejectedBy,omitempty"`
	RejectedDate   *time.Time         `json:"rejectedDate" bson:"rejectedDate,omitempty"`
	Remarks        string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Created        Created            `json:"created"  bson:"created,omitempty"`
	Updated        []Updated          `json:"updated"  bson:"updated,omitempty"`
}

//RefBillclaimLevels : "BillclaimLevels with refrence data such as language..."
type RefBillclaimLevels struct {
	BillclaimLevels `bson:",inline"`
	Ref             struct {
		OrganisationID *Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		Employee       *Employee     `json:"employee" bson:"employee,omitempty"`
		AssignedBy     *Employee     `json:"assignedBy" bson:"assignedBy,omitempty"`
		ApprovedBy     *Employee     `json:"approvedBy" bson:"approvedBy,omitempty"`
		RejectedBy     *Employee     `json:"rejectedBy" bson:"rejectedBy,omitempty"`
		Bill           *BillClaim    `json:"bill" bson:"bill,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//BillclaimLevelsFilter : "Used for constructing filter query"
type BillclaimLevelsFilter struct {
	Codes          []string `json:"codes,omitempty" bson:"codes,omitempty"`
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OmiStatus      []string `json:"omiStatus,omitempty" bson:"omiStatus,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Employee       []string `json:"employee,omitempty" bson:"employee,omitempty"`
	RejectedBy     []string `json:"rejectedBy" bson:"rejectedBy,omitempty"`
	AssignedBy     []string `json:"assignedBy,omitempty" bson:"assignedBy,omitempty"`
	SortBy         string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder      int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
