package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ApplicantType : ""
type ApplicantType struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefApplicantType : ""
type RefApplicantType struct {
	ApplicantType `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ApplicantTypeFilter : ""
type ApplicantTypeFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
