package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payroll struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	Method         string             `json:"method,omitempty" bson:"method,omitempty"`
	SalaryConfigId string             `json:"salaryConfigId,omitempty" bson:"salaryConfigId,omitempty"`
	CTC            float64            `json:"ctc,omitempty" bson:"ctc,omitempty"`
	GrossAmount    float64            `json:"grossAmount,omitempty" bson:"grossAmount,omitempty"`
	TotalDeduction float64            `json:"totalDeduction,omitempty" bson:"totalDeduction,omitempty"`
	NetSalary      float64            `json:"netSalary,omitempty" bson:"netSalary,omitempty"`
	Earnings       struct {
		BasicSalary          float64 `json:"basicSalary,omitempty" bson:"basicSalary,omitempty"`
		Hra                  float64 `json:"hra,omitempty" bson:"hra,omitempty"`
		ConveyanceAllowances float64 `json:"conveyanceAllowances,omitempty" bson:"conveyanceAllowances,omitempty"`
		EducationAllowance   float64 `json:"educationAllowance,omitempty" bson:"educationAllowance,omitempty"`
		PerformanceAllowance float64 `json:"performanceAllowance,omitempty" bson:"performanceAllowance,omitempty"`
	} `json:"earnings,omitempty" bson:"earnings,omitempty"`
	Detections struct {
		PfContribution   float64 `json:"pfContribution,omitempty" bson:"pfContribution,omitempty"`
		ESICContribution float64 `json:"eSICContribution,omitempty" bson:"eSICContribution,omitempty"`
		Lop              float64 `json:"lop,omitempty" bson:"lop,omitempty"`
	} `json:"detections,omitempty" bson:"detections,omitempty"`
	Date    *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Created *Created   `json:"createdOn" bson:"createdOn,omitempty"`
	Status  string     `json:"status,omitempty" bson:"status,omitempty"`
}

type RefPayroll struct {
	Payroll `bson:",inline"`
	Ref     struct {
		OrganisationID Organisation    `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		Bank           BankInformation `json:"bank" bson:"bank,omitempty"`
		DesignationID  Designation     `json:"designationId,omitempty" bson:"designationId,omitempty"`
		EmployeeId     Employee        `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterPayroll struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
