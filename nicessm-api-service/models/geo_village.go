package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Village : "Holds single state data"
type Village struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	GramPanchayat primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	ActiveStatus  bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created       CreatedV2          `json:"createdOn" bson:"createdOn,omitempty"`
	Location      struct {
		Longitude float64 `json:"longitude" bson:"longitude,omitempty"`
		Latitude  float64 `json:"latitude" bson:"latitude,omitempty"`
	} `json:"location" bson:"location,omitempty"`
	Population      int64  `json:"population"  bson:"population,omitempty"`
	VillageHead     string `json:"villageHead"  bson:"villageHead,omitempty"`
	CommiteeDetails string `json:"commiteeDetails"  bson:"commiteeDetails,omitempty"`
	School          string `json:"schoolName"  bson:"schoolName,omitempty"`
	FieldAgent      string `json:"fieldAgent"  bson:"fieldAgent,omitempty"`
	Version         int64  `json:"version"  bson:"version,omitempty"`
}

//RefVillage : "Village with refrence data such as language..."
type RefVillage struct {
	Village `bson:",inline"`
	Ref     struct {
		State         State         `json:"state,omitempty" bson:"state,omitempty"`
		District      District      `json:"district,omitempty" bson:"district,omitempty"`
		Block         Block         `json:"block,omitempty" bson:"block,omitempty"`
		GramPanchayat GramPanchayat `json:"gramPanchayat,omitempty" bson:"gramPanchayat,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//VillageFilter : "Used for constructing filter query"
type VillageFilter struct {
	ID            []primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	DataAccess    DataAccessRequest    `json:"dataAccess" bson:"dataAccess,omitempty"`
	GramPanchayat []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	ActiveStatus  []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status        []string             `json:"status" bson:"status,omitempty"`
	SortBy        string               `json:"sortBy"`
	SortOrder     int                  `json:"sortOrder"`
	Regex         struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type"  bson:"type"`
	} `json:"regex" bson:"regex"`
}
type NeWAddVillage struct {
	State             primitive.ObjectID `json:"state,omitempty" bson:"state,omitempty"`
	District          primitive.ObjectID `json:"district,omitempty" bson:"district,omitempty"`
	Block             primitive.ObjectID `json:"block,omitempty" bson:"block,omitempty"`
	GramPanchayat     primitive.ObjectID `json:"gramPanchayat,omitempty" bson:"gramPanchayat,omitempty"`
	Village           primitive.ObjectID `json:"village,omitempty" bson:"village,omitempty"`
	StateName         string             `json:"stateName,omitempty" bson:"stateName,omitempty"`
	DistrictName      string             `json:"districtName,omitempty" bson:"districtName,omitempty"`
	BlockName         string             `json:"blockName,omitempty" bson:"blockName,omitempty"`
	GramPanchayatName string             `json:"gramPanchayatName,omitempty" bson:"gramPanchayatName,omitempty"`
	VillageName       string             `json:"villageName" bson:"villageName,omitempty"`
}
