package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeOnboardingCheckList struct {
	ID                          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID                    string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId                  string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OnboardingpolicyId          string             `json:"onboardingpolicyId,omitempty" bson:"onboardingpolicyId,omitempty"`
	OrganisationId              string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	OnboardingchecklistMasterId string             `json:"onboardingchecklistmasterId,omitempty" bson:"onboardingchecklistmasterId,omitempty"`
	IsChecked                   bool               `json:"ischecked,omitempty" bson:"ischecked,omitempty"`
	Created                     *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status                      string             `json:"status" bson:"status,omitempty"`
}

type RefEmployeeOnboardingCheckList struct {
	EmployeeOnboardingCheckList `bson:",inline"`
	Ref                         struct {
		OnboardingpolicyId          OnboardingPolicy          `json:"onboardingpolicyId,omitempty" bson:"onboardingpolicyId,omitempty"`
		OnboardingchecklistmasterId OnboardingCheckListMaster `json:"onboardingchecklistmasterId,omitempty" bson:"onboardingchecklistmasterId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefOnboardingPolicyV2 struct {
	OnboardingPolicy `bson:",inline"`
	CheckList        []RefOnboardingCheckListV2 `json:"checklist,omitempty" bson:"checklist,omitempty"`
}

type RefOnboardingCheckListV2 struct {
	OnboardingCheckList `bson:",inline"`
	Ref                 struct {
		ChecklistName OnboardingCheckListMaster   `json:"checklistName,omitempty" bson:"checklistName,omitempty"`
		IsChecked     EmployeeOnboardingCheckList `json:"isChecked,omitempty" bson:"isChecked,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefEmployeeOnboardingCheckListv2 struct {
	Employee         `bson:",inline"`
	OnboardingPolicy RefOnboardingPolicyV2 `json:"onboardingpolicy,omitempty" bson:"onboardingpolicy,omitempty"`
	Ref              struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterEmployeeOnboardingCheckList struct {
	Status                      []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId              []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	EmployeeId                  []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OnboardingpolicyId          string   `json:"onboardingpolicyId,omitempty" bson:"onboardingpolicyId,omitempty"`
	OnboardingchecklistmasterId string   `json:"onboardingchecklistmasterId,omitempty" bson:"onboardingchecklistmasterId,omitempty"`
}
