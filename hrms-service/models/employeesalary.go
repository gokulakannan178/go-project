package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeSalary struct {
	ID                    primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID              string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationId        string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId            string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Earning               []Earning          `json:"earning,omitempty" bson:"earning,omitempty"`
	Deduction             []Deduction        `json:"deduction,omitempty" bson:"deduction,omitempty"`
	TotalEaringAmount     float64            `json:"totalEaringAmount,omitempty" bson:"totalEaringAmount,omitempty"`
	TotaldeductiongAmount float64            `json:"totaldeductiongAmount,omitempty" bson:"totaldeductiongAmount,omitempty"`
	GrossAmount           float64            `json:"grossAmount,omitempty" bson:"grossAmount,omitempty"`
	NoOfDaysFullyPaid     float64            `json:"noOfDaysFullyPaid,omitempty" bson:"noOfDaysFullyPaid,omitempty"`
	NoOfDaysParticalPaid  float64            `json:"noOfDaysParticalPaid,omitempty" bson:"noOfDaysParticalPaid,omitempty"`
	NoOfDaysLop           float64            `json:"noOfDaysLop,omitempty" bson:"noOfDaysLop,omitempty"`
	Date                  *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	Created               *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Updated               Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	Status                string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefEmployeeSalary struct {
	EmployeeSalary `bson:",inline"`
	Ref            struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeSalary struct {
	Status         []string   `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string   `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId     []string   `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Employee       string     `json:"employee,omitempty" bson:"employee,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	Date           *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type Earning struct {
	EarningType string  `json:"earningType,omitempty" bson:"earningType,omitempty"`
	Amount      float64 `json:"amount,omitempty" bson:"amount,omitempty"`
}
type Deduction struct {
	DeductionType string  `json:"deductionType,omitempty" bson:"deductionType,omitempty"`
	Amount        float64 `json:"amount,omitempty" bson:"amount,omitempty"`
}
type SalaryError struct {
	EmployeeId   string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	EmployeeName string `json:"employeeName,omitempty" bson:"employeeName,omitempty"`
	Message      string `json:"message,omitempty" bson:"message,omitempty"`
}
