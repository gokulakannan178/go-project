package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//OffenceType : ""
type OffenceType struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID    string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Desc        string             `json:"desc" bson:"desc,omitempty"`
	VehicleType string             `json:"vehicleType" bson:"vehicleType,omitempty"`
	Penalty     float64            `json:"penaty" bson:"penaty,omitempty"`
	IsVideo     string             `json:"isVideo" bson:"isVideo,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Created     Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated     Updated            `json:"updated"  bson:"updated,omitempty"`
	UpdateLog   []Updated          `json:"updatedLog" bson:"updatedLog,omitempty"`
}

//RefOffenceType :""
type RefOffenceType struct {
	OffenceType `bson:"inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//OffenceTypeFilter : ""
type OffenceTypeFilter struct {
	Status      []string `json:"status"`
	OmitID      []string `json:"omitId"`
	VehicleType []string `json:"vehicleType" bson:"vehicleType,omitempty"`
	SortBy      string   `json:"sortBy"`
	SortOrder   int      `json:"sortOrder"`
}
