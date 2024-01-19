package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FarmerCrop : ""
type FarmerCrop struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus  bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Area          string             `json:"area,omitempty"  bson:"area,omitempty"`
	Crop          primitive.ObjectID `json:"crop" bson:"crop,omitempty"`
	InterCrop     primitive.ObjectID `json:"interCrop" bson:"interCrop,omitempty"`
	Farmer        primitive.ObjectID `json:"farmer,omitempty"  bson:"farmer,omitempty"`
	Status        string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created       *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	Irrigation    string             `json:"irrigation,omitempty"  bson:"irrigation,omitempty"`
	Season        primitive.ObjectID `json:"season,omitempty"  bson:"season,omitempty"`
	StartDate     *time.Time         `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	Unit          string             `json:"unit,omitempty"  bson:"unit,omitempty"`
	Version       int                `json:"version,omitempty"  bson:"version,omitempty"`
	Year          int                `json:"year,omitempty"  bson:"year,omitempty"`
	Variety       primitive.ObjectID `json:"variety,omitempty"  bson:"variety,omitempty"`
	Yeild         float64            `json:"yeild"  bson:"yeild,omitempty"`
	YieldUnit     string             `json:"yieldUnit,omitempty"  bson:"yieldUnit,omitempty"`
	InputCost     int                `json:"inputCost,omitempty"  inputCost:"year,omitempty"`
	YieldValue    int                `json:"yieldValue,omitempty"  bson:"yieldValue,omitempty"`
	Consumption   string             `json:"consumption,omitempty"  consumption:"year,omitempty"`
	CompletedDate *time.Time         `json:"completedDate,omitempty"  bson:"completedDate,omitempty"`
}

type FarmerCropFilter struct {
	Status    []string             `json:"status,omitempty"  bson:"status,omitempty"`
	Crop      []primitive.ObjectID `json:"crop" bson:"crop,omitempty"`
	Category  []primitive.ObjectID `json:"category" bson:"category,omitempty"`
	Farmer    primitive.ObjectID   `json:"farmer,omitempty"  bson:"farmer,omitempty"`
	Season    primitive.ObjectID   `json:"season,omitempty"  bson:"season,omitempty"`
	StartDate *time.Time           `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	Veriety   primitive.ObjectID   `json:"veriety,omitempty"  bson:"veriety,omitempty"`
}

type RefFarmerCrop struct {
	FarmerCrop `bson:",inline"`
	Ref        struct {
		InterCrop Commodity         `json:"interCrop" bson:"interCrop,omitempty"`
		Crop      Commodity         `json:"crop" bson:"crop,omitempty"`
		Farmer    Farmer            `json:"farmer,omitempty"  bson:"farmer,omitempty"`
		Season    Cropseason        `json:"season,omitempty"  bson:"season,omitempty"`
		Variety   CommodityVariety  `json:"variety,omitempty"  bson:"variety,omitempty"`
		Category  CommodityCategory `json:"category" bson:"category,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type FarmerCropCount struct {
	Crop float64 `json:"crop" bson:"crop,omitempty"`
}
