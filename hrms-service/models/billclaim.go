package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillClaim struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Title          string             `json:"title" bson:"title,omitempty"`
	Desc           string             `json:"desc" bson:"desc,omitempty"`
	Category       string             `json:"Category" bson:"Category,omitempty"`
	EmployeeId     string             `json:"employeeId" bson:"employeeId,omitempty"`
	GradeId        string             `json:"gradeId" bson:"gradeId,omitempty"`
	Bills          []Bills            `json:"bills" bson:"bills,omitempty"`
	ApprovedBy     string             `json:"approvedBy" bson:"approvedBy,omitempty"`
	ApprovedDate   *time.Time         `json:"approvedDate" bson:"approvedDate,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time         `json:"date" bson:"date,omitempty"`
	RejectedBy     string             `json:"rejectedBy" bson:"rejectedBy,omitempty"`
	RejectedDate   *time.Time         `json:"rejectedDate" bson:"rejectedDate,omitempty"`
	Created        Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Remarks        string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
	TotalAmount    float64            `json:"totalAmount,omitempty" bson:"totalAmount,omitempty"`
	Updated        Updated            `json:"updated" bson:"updated,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefBillClaim struct {
	BillClaim `bson:",inline"`
	Ref       struct {
		Employee     *Employee `json:"employee" bson:"employee,omitempty"`
		ApprovedUser *User     `json:"approvedUser" bson:"approvedUser,omitempty"`
		RejectedBy   *User     `json:"rejectedUser" bson:"rejectedUser,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterBillClaim struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId     []string `json:"employeeId" bson:"employeeId,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type Bills struct {
	Title  string     `json:"title" bson:"title,omitempty"`
	Desc   string     `json:"desc" bson:"desc,omitempty"`
	File   string     `json:"file" bson:"file,omitempty"`
	Amount float64    `json:"amount,omitempty" bson:"amount,omitempty"`
	Date   *time.Time `json:"date" bson:"date,omitempty"`
}
type ReviewedBillClaim struct {
	BillClaim  string `json:"billClaim,omitempty" bson:"billClaim,omitempty"`
	ReviewedBy string `json:"reviewedBy" bson:"reviewedBy,omitempty"`
	Remarks    string `json:"remarks,omitempty" bson:"remarks,omitempty"`
}
