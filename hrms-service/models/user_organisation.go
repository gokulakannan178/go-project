package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Organisation : ""
type Organisation struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"description" bson:"description,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated  Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
}

//RefOrganisation :""
type RefOrganisation struct {
	Organisation `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//OrganisationFilter : ""
type OrganisationFilter struct {
	Desc     []string `json:"description" bson:"description,omitempty"`
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	//UniqueID  []string `json:"uniqueId"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
type DashboardOrganisationCountReport struct {
	Active   int `json:"active" bson:"active,omitempty"`
	InActive int `json:"inActive" bson:"inActive,omitempty"`
}
