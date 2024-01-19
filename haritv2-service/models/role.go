package models

import "gopkg.in/mgo.v2/bson"

//Role : ""
type Role struct {
	ID       bson.ObjectId `json:"id,omitempty"  form:"id" bson:"_id,omitempty"`
	UniqueID string        `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string        `json:"name" bson:"name,omitempty"`
	Desc     string        `json:"desc" bson:"desc,omitempty"`
	Status   string        `json:"status" bson:"status,omitempty"`
	Company  struct {
		ID   string `json:"id" bson:"id,omitempty"`
		Type string `json:"type" bson:"type,omitempty"`
	} `json:"company" bson:"company,omitempty"`
	Created   *CreatedV2     `json:"createdOn" bson:"createdOn,omitempty"`
	Updated   Updated        `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog []Updated      `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Features  []RoleFeatures `json:"features" bson:"features,omitempty"`
}

//RoleFeatures : ""
type RoleFeatures struct {
	ID      string `json:"id" bson:"id,omitempty"`
	Visible bool   `json:"visible" bson:"visible,omitempty"`
}

type VisibilityFeature struct {
	ID       string  `json:"id" bson:"id,omitempty"`
	Visible  bool    `json:"visible" bson:"visible,omitempty"`
	Features Feature `json:"feature" bson:"feature,omitempty"`
}

//RoleFilter : ""
type RoleFilter struct {
	UniqueID         []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	CompanyID        []string `json:"companyId" bson:"companyId,omitempty"`
	Name             string   `json:"name" bson:"name,omitempty"`
	Status           []string `json:"status" bson:"status,omitempty"`
	ProjectionFields []string `json:"projectionFields" bson:"projectionFields,omitempty"`
	SortBy           string   `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder        int      `json:"sortOrder" bson:"sortOrder,omitempty"`
}

type RefRole struct {
	Role `bson:",inline"`
	Ref  struct {
		Features []VisibilityFeature `json:"-" bson:"features,omitempty"`
		Feature  interface{}         `json:"feature" bson:"-"`
	} `json:"ref" bson:"ref,omitempty"`
}
