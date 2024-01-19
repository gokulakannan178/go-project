package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PolicyRule struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	LeaveMasterID  string             `json:"leavemasterId,omitempty" bson:"leavemasterId,omitempty"`
	LeavePolicyID  string             `json:"leavepolicyId,omitempty" bson:"leavepolicyId,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Value          float64            `json:"value,omitempty" bson:"value,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefPolicyRule struct {
	PolicyRule `bson:",inline"`
	Ref        struct {
		LeaveMasterID  LeaveMaster  `json:"leavemasterId" bson:"leavemasterId,omitempty"`
		LeavePolicyID  LeavePolicy  `json:"leavepolicyId" bson:"leavepolicyId,omitempty"`
		OrganisationID Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterPolicyRule struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	Regex  struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"regex," bson:"regex"`
}
