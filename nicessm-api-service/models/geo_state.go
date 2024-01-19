package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//State : "Holds single state data"
type State struct {
	ID           primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID     string               `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name         string               `json:"name" bson:"name,omitempty"`
	Status       string               `json:"status" bson:"status,omitempty"`
	ActiveStatus bool                 `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      Created              `json:"createdOn" bson:"createdOn,omitempty"`
	Location     Location             `json:"location" bson:"location,omitempty"`
	Languages    []primitive.ObjectID `json:"languages"  bson:"languages,omitempty"`
	Version      int64                `json:"version"  bson:"version,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefState struct {
	State `bson:",inline"`
	Ref   struct {
		Languages []Languages       `json:"languages,omitempty" bson:"languages,omitempty"`
		Projects  []RefProjectState `json:"projects" bson:"projects,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type StateFilter struct {
	ID           []primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string             `json:"status" bson:"status,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	// OmitProjectState primitive.ObjectID `json:"omitProjectState"`
	OmitProjectState struct {
		Is      bool               `json:"is"`
		Project primitive.ObjectID `json:"project"`
	} `json:"omitProjectState"`
	Regex struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
type GeoDetailsReport struct {
	Data []struct {
		State             string `json:"state" bson:"states,omitempty"`
		StateCode         string `json:"stateCode" bson:"stateCode,omitempty"`
		Distric           string `json:"distric" bson:"districs,omitempty"`
		DistricCode       string `json:"districCode" bson:"districCode,omitempty"`
		Block             string `json:"block" bson:"blocks,omitempty"`
		BlockCode         string `json:"blockCode" bson:"blockCode,omitempty"`
		GramPanchayat     string `json:"grampanchayat" bson:"grampanchayats,omitempty"`
		GramPanchayatCode string `json:"grampanchayatCode" bson:"grampanchayatCode,omitempty"`
		Village           string `json:"village" bson:"villages,omitempty"`
		VillageCode       string `json:"villageCode" bson:"villageCode,omitempty"`
	} `json:"data,omitempty" bson:"data,omitempty"`
}
type GeoDetailsReport2 struct {
	//Date []struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	UniqueID  string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Districts []struct {
		ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Name     string             `json:"name" bson:"name,omitempty"`
		UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
		Block    []struct {
			ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
			Name          string             `json:"name" bson:"name,omitempty"`
			UniqueID      string             `json:"uniqueId" bson:"uniqueId,omitempty"`
			GramPanchayat []struct {
				ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
				Name     string             `json:"name" bson:"name,omitempty"`
				UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
				Village  []struct {
					ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
					Name     string             `json:"name" bson:"name,omitempty"`
					UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
				} `json:"villages" bson:"villages,omitempty"`
			} `json:"gramPanchayats" bson:"gramPanchayats,omitempty"`
		} `json:"blocks" bson:"blocks,omitempty"`
	} `json:"districts" bson:"districts,omitempty"`
	//	} `json:"data" bson:"data,omitempty"`
}
