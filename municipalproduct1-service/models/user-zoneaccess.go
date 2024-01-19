package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserZoneAccess struct {
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ZoneID   primitive.ObjectID `json:"zoneId" bson:"zoneId,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	UserName string             `json:"userName" bson:"userName,omitempty`
}

type UserZoneAccessFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefUserZoneAccess struct {
	UserZoneAccess `bson:",inline"`
	RefZone        `bson:",inline"`
}
