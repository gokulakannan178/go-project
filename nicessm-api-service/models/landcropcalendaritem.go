package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//LandCropCalendar : ""
type LandCropCalendar struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Stage           primitive.ObjectID `json:"stage" bson:"stage,omitempty"`
	Version         int                `json:"version" bson:"version,omitempty"`
	IrrigationCount int                `json:"irrigationCount" bson:"irrigationCount,omitempty"`
	LandCrop        primitive.ObjectID `json:"landCrop,omitempty"  bson:"landCrop,omitempty"`
	Status          string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created         *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	EndDate         *time.Time         `json:"endDate,omitempty"  bson:"endDate,omitempty"`
	StartDate       *time.Time         `json:"startDate,omitempty"  bson:"startDate,omitempty"`
}

type LandCropCalendarFilter struct {
	Status    []string           `json:"status,omitempty"  bson:"status,omitempty"`
	Stage     primitive.ObjectID `json:"stage" bson:"stage,omitempty"`
	LandCrop  primitive.ObjectID `json:"landCrop,omitempty"  bson:"landCrop,omitempty"`
	StartDate *time.Time         `json:"startDate,omitempty"  bson:"startDate,omitempty"`
}

type RefLandCropCalendar struct {
	LandCropCalendar `bson:",inline"`
	Ref              struct {
		Farmer Farmer    `json:"farmer" bson:"farmer,omitempty"`
		Stage  Commodity `json:"stage" bson:"stage,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
