package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryConfig struct {
	ID                  primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID            string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeType        string             `json:"employeeType,omitempty" bson:"employeeType,omitempty"`
	GrossPercentage     float64            `json:"grossPercentage,omitempty" bson:"grossPercentage,omitempty"`
	EarningsPercentage  float64            `json:"earningsPercentage,omitempty" bson:"earningsPercentage,omitempty"`
	DeductionPercentage float64            `json:"deductionPercentage,omitempty" bson:"deductionPercentage,omitempty"`
	NetPercentage       float64            `json:"netPercentage,omitempty" bson:"netPercentage,omitempty"`
	Earnings            struct {
		BasicSalary          float64 `json:"basicSalary,omitempty" bson:"basicSalary,omitempty"`
		Hra                  float64 `json:"hra,omitempty" bson:"hra,omitempty"`
		ConveyanceAllowances float64 `json:"conveyanceAllowances,omitempty" bson:"conveyanceAllowances,omitempty"`
		EducationAllowance   float64 `json:"educationAllowance,omitempty" bson:"educationAllowance,omitempty"`
		PerformanceAllowance float64 `json:"performanceAllowance,omitempty" bson:"performanceAllowance,omitempty"`
		Others               float64 `json:"others,omitempty" bson:"others,omitempty"`
	} `json:"earnings,omitempty" bson:"earnings,omitempty"`
	Detections struct {
		PfContribution   float64 `json:"pfContribution,omitempty" bson:"pfContribution,omitempty"`
		ESICContribution float64 `json:"eSICContribution,omitempty" bson:"eSICContribution,omitempty"`
		Others           float64 `json:"others,omitempty" bson:"others,omitempty"`
	} `json:"detections,omitempty" bson:"detections,omitempty"`
	OrganisationID string     `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate        *time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
	Created        *Created   `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string     `json:"status,omitempty" bson:"status,omitempty"`
}

type RefSalaryConfig struct {
	SalaryConfig `bson:",inline"`
	Ref          struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterSalaryConfig struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeID     []string `json:"employeeID,omitempty" bson:"employeeID,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
