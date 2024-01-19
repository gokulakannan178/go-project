package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PayrollPolicyEarning struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty"`
	Description     string             `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationID  string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	PayRollPolicyId string             `json:"payRollPolicyId,omitempty" bson:"payRollPolicyId,omitempty"`
	EarningMasterId string             `json:"earningMasterId,omitempty" bson:"earningMasterId,omitempty"`
	Amount          float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	Created         *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefPayrollPolicyEarning struct {
	PayrollPolicyEarning `bson:",inline"`
	Ref                  struct {
		OrganisationID  Organisation          `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		EarningMasterId EmployeeEarningMaster `json:"earningMasterId,omitempty" bson:"earningMasterId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterPayrollPolicyEarning struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
