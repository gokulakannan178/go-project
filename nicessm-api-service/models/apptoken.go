package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apptoken struct {
	ID                primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID          string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserID            primitive.ObjectID `json:"userid" bson:"userid,omitempty"`
	Usertype          string             `json:"usertype" bson:"usertype,omitempty"`
	Status            string             `json:"status" bson:"status,omitempty"`
	RegistrationToken string             `json:"registrationtoken" bson:"registrationtoken"`
	Currenttime       *time.Time         `json:"currenttime" bson:"currenttime,omitempty"`
}
type ApptokenFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}
type RefApptoken struct {
	Apptoken `bson:",inline"`
}

type AppTokenNotification struct {
	Notification string   `json:"notification" bson:"notification,omitempty"`
	Types        []string `json:"types" bson:"types,omitempty"`
}

type RefAppTokenNotification struct {
	RegistrationToken []string `json:"registrationtoken" bson:"registrationtoken,omitempty"`
}
