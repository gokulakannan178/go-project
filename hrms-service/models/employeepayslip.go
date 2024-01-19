package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeePayslip struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	ResponseUrl    string             `json:"responseUrl,omitempty" bson:"responseUrl,omitempty"`
	PayslipId      string             `json:"payslipId,omitempty" bson:"payslipId,omitempty"`
	YearOfMonth    string             `json:"yearOfMonth,omitempty" bson:"yearOfMonth,omitempty"`
	FileUrl        string             `json:"fileUrl,omitempty" bson:"fileUrl,omitempty"`
	Date           *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	Remarks        string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

type RefEmployeePayslip struct {
	EmployeePayslip `bson:",inline"`
	Ref             struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		Employee       Employee     `json:"employee,omitempty" bson:"employee,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeePayslip struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
