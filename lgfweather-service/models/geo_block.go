package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Block : "Holds single state data"
type Block struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID        string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty"`
	ImdStateName    string             `json:"imdStateName,omitempty"  bson:"imdStateName,omitempty"`
	ImdDistrictName string             `json:"imdDistrictName,omitempty"  bson:"imdDistrictName,omitempty"`
	ImdBlockName    string             `json:"imdBlockName,omitempty"  bson:"imdBlockName,omitempty"`
	ImdFileName     string             `json:"imdFileName" bson:"imdFileName,omitempty"`
	Code            string             `json:"code"  bson:"code,omitempty"`
	DistrictCode    string             `json:"districtCode,omitempty"  bson:"districtCode,omitempty"`
	Status          string             `json:"status"  bson:"status,omitempty"`
	Created         *CreatedV2         `json:"created"  bson:"created,omitempty"`
	Languages       []string           `json:"languages"  bson:"languages,omitempty"`
	Updated         []Updated          `json:"updated"  bson:"updated,omitempty"`
	Location        Location           `json:"location" bson:"location,omitempty"`
}

//RefBlock : "Village with refrence data such as language..."
type RefBlock struct {
	Village `bson:",inline"`
	Ref     struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//BlockFilter : "Used for constructing filter query"
type BlockFilter struct {
	UniqueID      []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Codes         []string `json:"codes,omitempty" bson:"codes,omitempty"`
	StateName     string   `json:"stateName,omitempty"  bson:"stateName,omitempty"`
	DistrictName  string   `json:"districtName,omitempty"  bson:"districtName,omitempty"`
	DistrictCodes []string `json:"districtCodes,omitempty" bson:"districtCodes,omitempty"`
	Status        []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy        string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder     int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

type BlockName struct {
	StateName    string   `json:"stateName" bson:"stateName,omitempty"`
	DistrcitName string   `json:"districtName" bson:"districtName,omitempty"`
	BlockName    []string `json:"blockName" bson:"blockName,omitempty"`
}
