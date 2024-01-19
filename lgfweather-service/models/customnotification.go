package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomNotification struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Title    string             `json:"title" bson:"title,omitempty"`
	Body     string             `json:"body" bson:"body,omitempty"`
	Image    string             `json:"image" bson:"image,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  *CreatedV2         `json:"created" bson:"created,omitempty"`
}

type CustomNotificationFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

type RefCustomNotification struct {
	CustomNotification `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
