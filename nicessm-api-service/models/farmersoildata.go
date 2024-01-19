package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FarmerSoilData : ""
type FarmerSoilData struct {
	ID              primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus    bool                 `json:"activeStatus" bson:"activeStatus,omitempty"`
	FarmerLand      primitive.ObjectID   `json:"farmerLand" bson:"farmerLand,omitempty"`
	Farmer          primitive.ObjectID   `json:"farmer,omitempty"  bson:"farmer,omitempty"`
	Status          string               `json:"status,omitempty"  bson:"status,omitempty"`
	Created         *Created             `json:"created,omitempty"  bson:"created,omitempty"`
	LabName         primitive.ObjectID   `json:"labName,omitempty"  bson:"labName,omitempty"`
	EcValue         float64              `json:"ecValue,omitempty"  bson:"ecValue,omitempty"`
	Humidity        float64              `json:"humidity,omitempty"  bson:"humidity,omitempty"`
	PH              float64              `json:"pH,omitempty"  bson:"pH,omitempty"`
	SoilCollectedOn *time.Time           `json:"soilCollectedOn,omitempty"  bson:"soilCollectedOn,omitempty"`
	SoilTestedOn    *time.Time           `json:"soilTestedOn,omitempty"  bson:"soilTestedOn,omitempty"`
	ValidFrom       *time.Time           `json:"validFrom,omitempty"  bson:"validFrom,omitempty"`
	ValidTo         *time.Time           `json:"validTo,omitempty"  bson:"validTo,omitempty"`
	SoilSampleNo    float64              `json:"soilSampleNo,omitempty"  bson:"soilSampleNo,omitempty"`
	MicroNutrients  []primitive.ObjectID `json:"microNutrients" bson:"-"`
	MacroNutrients  []primitive.ObjectID `json:"macroNutrients,omitempty"  bson:"-"`
	Location        LocationV2           `json:"Location,omitempty"  bson:"Location,omitempty"`
	OrganicCarbon   float64              `json:"organicCarbon,omitempty"  bson:"organicCarbon,omitempty"`
}

type FarmerSoilDataFilter struct {
	Status []string `json:"status,omitempty"  bson:"status,omitempty"`
}

type RefFarmerSoilData struct {
	FarmerSoilData `bson:",inline"`
	Ref            struct {
		FarmerLand     FarmerLand  `json:"farmerLand" bson:"farmerLand,omitempty"`
		Farmer         Farmer      `json:"farmer,omitempty"  bson:"farmer,omitempty"`
		MicroNutrients []Nutrients `json:"microNutrients" bson:"microNutrients,omitempty"`
		MacroNutrients []Nutrients `json:"macroNutrients,omitempty"  bson:"macroNutrients,omitempty"`
		LabName        Aidlocation `json:"labName,omitempty"  bson:"labName,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type LocationV2 struct {
	Longitude int `json:"longitude,omitempty"  bson:"longitude,omitempty"`
	Latitude  int `json:"latitude,omitempty"  bson:"latitude,omitempty"`
}
