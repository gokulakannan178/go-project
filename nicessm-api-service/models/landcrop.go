package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AidCategory : ""
type LandCrop struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	AreaInAcre   float64            `json:"areaInAcre,omitempty"  bson:"areaInAcre,omitempty"`
	Crop         primitive.ObjectID `json:"crop" bson:"crop,omitempty"`
	Farmer       primitive.ObjectID `json:"farmer" bson:"farmer,omitempty"`
	FarmerLand   primitive.ObjectID `json:"farmerLand,omitempty"  bson:"farmerLand,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	Season       primitive.ObjectID `json:"season,omitempty"  bson:"season,omitempty"`
	StartDate    *time.Time         `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	Unit         string             `json:"unit,omitempty"  bson:"unit,omitempty"`
	Version      int                `json:"version,omitempty"  bson:"version,omitempty"`
	Year         int                `json:"year,omitempty"  bson:"year,omitempty"`
	Veriety      primitive.ObjectID `json:"veriety,omitempty"  bson:"veriety,omitempty"`
	Yeild        float64            `json:"yeild,omitempty"  bson:"yeild,omitempty"`
}

type LandCropFilter struct {
	Status     []string           `json:"status,omitempty"  bson:"status,omitempty"`
	Crop       primitive.ObjectID `json:"crop" bson:"crop,omitempty"`
	FarmerLand primitive.ObjectID `json:"farmerLand,omitempty"  bson:"farmerLand,omitempty"`
	Farmer     primitive.ObjectID `json:"farmer" bson:"farmer,omitempty"`
	Season     primitive.ObjectID `json:"season,omitempty"  bson:"season,omitempty"`
}

type RefLandCrop struct {
	LandCrop `bson:",inline"`
	Ref      struct {
		Farmer     Farmer           `json:"farmer" bson:"farmer,omitempty"`
		Crop       Commodity        `json:"crop" bson:"crop,omitempty"`
		FarmerLand FarmerLand       `json:"farmerLand,omitempty"  bson:"farmerLand,omitempty"`
		Season     Cropseason       `json:"season,omitempty"  bson:"season,omitempty"`
		Veriety    CommodityVariety `json:"veriety,omitempty"  bson:"veriety,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
