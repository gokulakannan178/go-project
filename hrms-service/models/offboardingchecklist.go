package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OffboardingCheckList struct {
	ID                           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name                         string             `json:"name,omitempty" bson:"name,omitempty"`
	OrganisationID               string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	UniqueID                     string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OffboardingchecklistmasterID string             `json:"offboardingchecklistmasterId,omitempty" bson:"offboardingchecklistmasterId,omitempty"`
	OffboardingpolicyID          string             `json:"offboardingpolicyId,omitempty" bson:"offboardingpolicyId,omitempty"`
	Created                      *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status                       string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefOffboardingCheckList struct {
	OffboardingCheckList `bson:",inline"`
	Ref                  struct {
		OrganisationID               Organisation               `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		OffboardingchecklistmasterID OffboardingCheckListMaster `json:"offboardingchecklistmasterId" bson:"offboardingchecklistmasterId,omitempty"`
		OffboardingpolicyID          OffboardingPolicy          `json:"offboardingpolicyId" bson:"offboardingpolicyId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterOffboardingCheckList struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
