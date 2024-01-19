package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Village : "Holds single state data"
type Village struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Code             string             `json:"code"  bson:"code,omitempty"`
	GramPanjayatCode string             `json:"gramPanjayatCode,omitempty"  bson:"gramPanjayatCode,omitempty"`
	Status           string             `json:"status"  bson:"status,omitempty"`
	Created          *CreatedV2         `json:"created"  bson:"created,omitempty"`
	Languages        []string           `json:"languages"  bson:"languages,omitempty"`
	Updated          []Updated          `json:"updated"  bson:"updated,omitempty"`
	Location         Location           `json:"location" bson:"location,omitempty"`
}

//RefVillage : "Village with refrence data such as language..."
type RefVillage struct {
	Village `bson:",inline"`
	Ref     struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//VillageFilter : "Used for constructing filter query"
type VillageFilter struct {
	Codes    []string `json:"codes,omitempty" bson:"codes,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`

	GramPanjayatCode []string `json:"gramPanjayatCodes,omitempty" bson:"gramPanjayatCodes,omitempty"`
	Status           []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy           string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder        int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
