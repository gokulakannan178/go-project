package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeePayroll struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	CTC            float64            `json:"ctc,omitempty" bson:"ctc,omitempty"`
	NetAmount      float64            `json:"netAmount,omitempty" bson:"netAmount,omitempty"`
	GrossAmount    float64            `json:"grossAmount,omitempty" bson:"grossAmount,omitempty"`
	Deduction      float64            `json:"deduction,omitempty" bson:"deduction,omitempty"`
	StartDate      *time.Time         `json:"startDate" bson:"startDate,omitempty"`
	EndDate        *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	Remarks        string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

type RefEmployeePayroll struct {
	EmployeePayroll `bson:",inline"`
	Ref             struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeePayroll struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type EmployeePayrollWithEmployee struct {
	EmployeeId string              `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Earning    []EmployeeEarning   `json:"earning,omitempty" bson:"earning,omitempty"`
	Deduction  []EmployeeDeduction `json:"deduction,omitempty" bson:"deduction,omitempty"`
	PayRoll    *EmployeePayroll    `json:"payRoll,omitempty" bson:"payRoll,omitempty"`
}
type EmployeePayrollWithEarningDeduction struct {
	OrganisationId string  `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	UniqueID       string  `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId     string  `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	NetAmount      float64 `json:"netAmount,omitempty" bson:"netAmount,omitempty"`
	CTC            float64 `json:"ctc,omitempty" bson:"ctc,omitempty"`
	Earning        []struct {
		EmployeeId string  `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
		EarningId  string  `json:"earningId,omitempty" bson:"earningId,omitempty"`
		Amount     float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	} `json:"earning,omitempty" bson:"earning,omitempty"`
	Deduction []struct {
		EmployeeId  string  `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
		DeductionId string  `json:"deductionId,omitempty" bson:"deductionId,omitempty"`
		Amount      float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	} `json:"deduction,omitempty" bson:"deduction,omitempty"`
}
