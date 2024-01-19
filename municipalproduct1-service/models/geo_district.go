package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//District : "Holds single district data"
type District struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"  bson:"name,omitempty"`
	Code      string             `json:"code,omitempty"  bson:"code,omitempty"`
	StateCode string             `json:"stateCode,omitempty"  bson:"stateCode,omitempty"`
	Status    string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created   Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Languages []string           `json:"languages,omitempty" bson:"languages,omitempty"`
	Updated   []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//RefDistrict : ""
type RefDistrict struct {
	District `bson:",inline"`
	Ref      struct {
		State *State `json:"state,omitempty" bson:"state,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DistrictFilter : "Used for constructing filter query"
type DistrictFilter struct {
	Codes      []string `json:"codes,omitempty" bson:"codes,omitempty"`
	StateCodes []string `json:"stateCodes,omitempty" bson:"stateCodes,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy     string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder  int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
