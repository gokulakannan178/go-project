package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Cluster : "Holds single KnowlegdeDomain data"
type Cluster struct {
	ID            primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	Name          string               `json:"name" bson:"name,omitempty"`
	Status        string               `json:"status" bson:"status,omitempty"`
	ActiveStatus  bool                 `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created       Created              `json:"createdOn" bson:"createdOn,omitempty"`
	GramPanchayat primitive.ObjectID   `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	State         primitive.ObjectID   `json:"state"  bson:"state,omitempty"`
	Block         primitive.ObjectID   `json:"block"  bson:"block,omitempty"`
	District      primitive.ObjectID   `json:"district"  bson:"district,omitempty"`
	//Version       string               `json:"version"  bson:"version,omitempty"`
	Description string `json:"description"  bson:"description,omitempty"`
}

//RefCluster : "RefCluster with refrence data such as language..."
type RefCluster struct {
	Cluster `bson:",inline"`
	Ref     struct {
		GramPanchayat GramPanchayat `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		Village       []Village     `json:"village"  bson:"village,omitempty"`
		State         State         `json:"state"  bson:"state,omitempty"`
		Block         Block         `json:"block"  bson:"block,omitempty"`
		District      District      `json:"district"  bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ClusterFilter : "Used for constructing filter query"
type ClusterFilter struct {
	GramPanchayat []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	ActiveStatus  []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status        []string             `json:"status" bson:"status,omitempty"`
	SortBy        string               `json:"sortBy"`
	SortOrder     int                  `json:"sortOrder"`
	Regex         struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
