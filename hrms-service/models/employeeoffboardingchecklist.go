package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeOffboardingCheckList struct {
	ID                           primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID                     string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId                   string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId               string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	OffboardingPolicyId          string             `json:"offboardingpolicyId,omitempty" bson:"offboardingpolicyId,omitempty"`
	OffboardingCheckListMasterId string             `json:"offboardingchecklistmasterId,omitempty" bson:"offboardingchecklistmasterId,omitempty"`
	IsChecked                    bool               `json:"isChecked,omitempty" bson:"isChecked,omitempty"`
	Created                      *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status                       string             `json:"status" bson:"status,omitempty"`
}

type RefEmployeeOffboardingCheckList struct {
	EmployeeOnboardingCheckList `bson:",inline"`
	Ref                         struct {
		OffboardingpolicyId          OffboardingPolicy          `json:"offboardingpolicyId" bson:"offboardingpolicyId,omitempty"`
		OffboardingchecklistmasterId OffboardingCheckListMaster `json:"offboardingchecklistmasterId" bson:"offboardingchecklistmasterId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefOffboardingPolicyV2 struct {
	OffboardingPolicy `bson:",inline"`
	CheckList         []RefOffboardingCheckListV2 `json:"checklist" bson:"checklist,omitempty"`
}

type RefOffboardingCheckListV2 struct {
	OffboardingCheckList `bson:",inline"`
	Ref                  struct {
		ChecklistName OffboardingCheckListMaster   `json:"checklistName" bson:"checklistName,omitempty"`
		IsChecked     EmployeeOffboardingCheckList `json:"isChecked" bson:"isChecked,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefEmployeeOffboardingCheckListv2 struct {
	Employee          `bson:",inline"`
	OffboardingPolicy RefOffboardingPolicyV2 `json:"offboardingpolicy" bson:"offboardingpolicy,omitempty"`
	Ref               struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeOffboardingCheckList struct {
	Status                       []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId               []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId                   []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OffboardingpolicyId          string   `json:"offboardingpolicyId,omitempty" bson:"offboardingpolicyId,omitempty"`
	OffboardingchecklistmasterId string   `json:"offboardingchecklistmasterId,omitempty" bson:"offboardingchecklistmasterId,omitempty"`
}
