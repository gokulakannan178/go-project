package models

import "time"

type UserLocationLog struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName  string     `json:"userName" bson:"userName,omitempty"`
	UserType  string     `json:"userType" bson:"userType,omitempty"`
	Location  Location   `json:"location" bson:"location,omitempty"`
	TimeStamp *time.Time `json:"timeStamp" bson:"timeStamp,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	Created   *CreatedV2 `json:"created" bson:"created,omitempty"`
}

type UserLocationLogFilter struct {
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   []string `json:"status" bson:"status,omitempty"`
	UserName []string `json:"userName" bson:"userName,omitempty"`
	UserType []string `json:"userType" bson:"userType,omitempty"`
}
type RefUserLocationLog struct {
	UserLocationLog `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
