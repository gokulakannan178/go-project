package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//State : "Holds single state data"
type State struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	ImdStateName     string             `json:"imdStateName" bson:"imdStateName,omitempty"`
	ImdFileName      string             `json:"imdFileName" bson:"imdFileName,omitempty"`
	ImdBlockFileName string             `json:"imdBlockFileName" bson:"imdBlockFileName,omitempty"`
	Code             string             `json:"code"  bson:"code,omitempty"`
	Status           string             `json:"status"  bson:"status,omitempty"`
	Created          Created            `json:"created"  bson:"created,omitempty"`
	Languages        []string           `json:"languages"  bson:"languages,omitempty"`
	Updated          []Updated          `json:"updated"  bson:"updated,omitempty"`
	Location         Location           `json:"location" bson:"location,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefState struct {
	State `bson:",inline"`
	Ref   struct {
		Languages []Language `json:"languages,omitempty" bson:"languages,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type Statename struct {
	Name string `json:"name" bson:"name,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type StateFilter struct {
	UniqueID  string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Codes     []string `json:"codes,omitempty" bson:"codes,omitempty"`
	State     []string `json:"state,omitempty" bson:"state,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
