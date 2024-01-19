package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//District : "Holds single district data"
type District struct {
	ID                  primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	State               primitive.ObjectID   `json:"state"  bson:"state,omitempty"`
	UniqueID            string               `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name                string               `json:"name" bson:"name,omitempty"`
	Status              string               `json:"status" bson:"status,omitempty"`
	ActiveStatus        bool                 `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created             Created              `json:"createdOn" bson:"createdOn,omitempty"`
	Languages           []primitive.ObjectID `json:"languages"  bson:"languages,omitempty"`
	AgroEcologicalZones []primitive.ObjectID `json:"agroEcologicalZones"  bson:"agroEcologicalZones"`
	Location            Location             `json:"location" bson:"location,omitempty"`
	SoilTypes           []primitive.ObjectID `json:"soilTypes"  bson:"soilTypes"`
	Version             int64                `json:"version"  bson:"version,omitempty"`
}

//RefDistrict : ""
type RefDistrict struct {
	District `bson:",inline"`
	Ref      struct {
		State               State                `json:"state,omitempty" bson:"state,omitempty"`
		AgroEcologicalZones []AgroEcologicalZone `json:"agroEcologicalZones"  bson:"agroEcologicalZones,omitempty"`
		SoilTypes           []SoilType           `json:"soilTypes"  bson:"soilTypes,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DistrictFilter : "Used for constructing filter query"
type DistrictFilter struct {
	ID           []primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	State        []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string             `json:"status" bson:"status,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
