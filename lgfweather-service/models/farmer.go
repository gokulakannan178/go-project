package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Farmer struct {
	ID primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`

	Name     string     `json:"name" bson:"name,omitempty"`
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	COName   string     `json:"cOName" bson:"cOName,omitempty"`
	Gender   string     `json:"gender" bson:"gender,omitempty"`
	DOB      *time.Time `json:"dob" bson:"dob,omitempty"`
	MobileNo string     `json:"mobileNo" bson:"mobileNo,omitempty"`
	Email    string     `json:"email" bson:"email,omitempty"`
	Address  Address    `json:"address" bson:"address,omitempty"`
	Token    string     `json:"token" bson:"-"`

	Created *CreatedV2 `json:"created" bson:"created,omitempty"`
}

type FarmerFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

type RefFarmer struct {
	Farmer `bson:",inline"`
	Ref    struct {
		Address RefAddress `json:"address" bson:"address,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
