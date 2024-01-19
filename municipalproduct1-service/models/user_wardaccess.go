package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserWardAccess struct {
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	WardID   primitive.ObjectID `json:"wardId" bson:"wardId,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	UserName string             `json:"userName" bson:"userName,omitempty`
}

type UserWardAccessFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefUserWardAccess struct {
	UserWardAccess `bson:",inline"`
	RefWard        `bson:",inline"`
}
