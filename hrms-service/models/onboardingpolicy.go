package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OnboardingPolicy struct {
	ID                       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name                     string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID                 string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Desc                     string             `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationID           string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	OnboardCheckListMasterId []string           `json:"onboardchecklistmasterId,omitempty" bson:"-"`
	Created                  *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status                   string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefOnboardingPolicy struct {
	OnboardingPolicy `bson:",inline"`
	Ref              struct {
		OrganisationID      Organisation          `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		OnboardingCheckList []OnboardingCheckList `json:"onboardingchecklist,omitempty" bson:"onboardingchecklist,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterOnboardingPolicy struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
