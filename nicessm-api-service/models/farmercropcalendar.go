package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FarmerCropCalendar : ""
type FarmerCropCalendar struct {
	ID         primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Stage      primitive.ObjectID `json:"stage" bson:"stage,omitempty"`
	Version    int                `json:"version" bson:"version,omitempty"`
	FarmerCrop primitive.ObjectID `json:"farmerCrop,omitempty"  bson:"farmerCrop,omitempty"`
	Status     string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created    *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	EndDate    *time.Time         `json:"endDate,omitempty"  bson:"endDate,omitempty"`
	StartDate  *time.Time         `json:"startDate,omitempty"  bson:"startDate,omitempty"`
}

type FarmerCropCalendarFilter struct {
	Status     []string           `json:"status,omitempty"  bson:"status,omitempty"`
	Stage      primitive.ObjectID `json:"stage" bson:"stage,omitempty"`
	FarmerCrop primitive.ObjectID `json:"farmerCrop,omitempty"  bson:"farmerCrop,omitempty"`
	StartDate  *time.Time         `json:"startDate,omitempty"  bson:"startDate,omitempty"`
}

type RefFarmerCropCalendar struct {
	FarmerCropCalendar `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
