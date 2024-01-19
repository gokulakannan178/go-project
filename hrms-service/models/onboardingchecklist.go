package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OnboardingCheckList struct {
	ID                          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name                        string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID                    string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	OnboardingchecklistmasterID string             `json:"onboardingchecklistmasterId,omitempty" bson:"onboardingchecklistmasterId,omitempty"`
	OrganisationID              string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	OnboardingpolicyID          string             `json:"onboardingpolicyId,omitempty" bson:"onboardingpolicyId,omitempty"`
	Created                     *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status                      string             `json:"status,omitempty" bson:"status,omitempty"`
}

// type OnboardingCheckListUpsert struct {
// 	OnboardingpolicyID string   `json:"onboardingpolicyId,omitempty" bson:"onboardingpolicyId,omitempty"`
// 	ArrayValue         []string `json:"arrayValue,omitempty" bson:"arrayValue,omitempty"`
// }

type RefOnboardingCheckList struct {
	OnboardingCheckList `bson:",inline"`
	Ref                 struct {
		OnboardingchecklistmasterID OnboardingCheckListMaster `json:"onboardingchecklistmasterId" bson:"onboardingchecklistmasterId,omitempty"`
		OrganisationID              Organisation              `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		OnboardingpolicyID          OnboardingPolicy          `json:"onboardingpolicyId" bson:"onboardingpolicyId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterOnboardingCheckList struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
