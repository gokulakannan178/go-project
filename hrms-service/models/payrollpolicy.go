package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PayrollPolicy struct {
	ID              primitive.ObjectID      `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string                  `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name            string                  `json:"name,omitempty" bson:"name,omitempty"`
	Description     string                  `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationID  string                  `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EarningMaster   []LeavematserWithpolicy `json:"earningMaster,omitempty" bson:"earningMaster,omitempty"`
	DetectionMaster []LeavematserWithpolicy `json:"detectionMaster,omitempty" bson:"detectionMaster,omitempty"`
	CTC             float64                 `json:"ctc,omitempty" bson:"ctc,omitempty"`
	TakeHome        float64                 `json:"takeHome,omitempty" bson:"takeHome,omitempty"`
	GrossAmount     float64                 `json:"grossAmount,omitempty" bson:"grossAmount,omitempty"`
	Deduction       float64                 `json:"deduction,omitempty" bson:"deduction,omitempty"`
	Created         *Created                `json:"createdOn" bson:"createdOn,omitempty"`
	Status          string                  `json:"status,omitempty" bson:"status,omitempty"`
}
type SalaryCalc struct {
	GrossAmount    float64 `json:"grossAmount,omitempty" bson:"grossAmount,omitempty"`
	TotalDeduction float64 `json:"totalDeduction,omitempty" bson:"totalDeduction,omitempty"`
	NetSalary      float64 `json:"netSalary,omitempty" bson:"netSalary,omitempty"`
	SalaryConfigId string  `json:"salaryConfigId,omitempty" bson:"salaryConfigId,omitempty"`
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
	} `json:"detections,omitempty" bson:"detections,omitempty"`
}
type RefPayrollPolicy struct {
	PayrollPolicy `bson:",inline"`
	Ref           struct {
		OrganisationID  Organisation                `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		EarningMaster   []RefPayrollPolicyEarning   `json:"earningMaster,omitempty" bson:"earningMaster,omitempty"`
		DetectionMaster []RefPayrollPolicyDetection `json:"detectionMaster,omitempty" bson:"detectionMaster,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterPayrollPolicy struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
