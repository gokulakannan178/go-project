package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//NutrientValue : ""
type NutrientValue struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus   bool               `json:"activeStatus " bson:"activeStatus,omitempty"`
	Status         string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created        *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	Nutrient       primitive.ObjectID `json:"nutrient,omitempty"  bson:"nutrient,omitempty"`
	Val            float64            `json:"val,omitempty"  bson:"val,omitempty"`
	FarmerSoilData primitive.ObjectID `json:"farmerSoilData" bson:"farmerSoilData,omitempty"`
}

type NutrientValueFilter struct {
	Status []string `json:"status,omitempty"  bson:"status,omitempty"`
	// SearchBox struct {
	// 	Name string `json:"name" bson:"name"`
	// } `json:"searchbox" bson:"searchbox"`
}

type RefNutrientValue struct {
	NutrientValue `bson:",inline"`
	Ref           struct {
		FarmerSoilData FarmerSoilData `json:"farmerSoilData" bson:"farmerSoilData,omitempty"`
		Nutrient       Nutrients      `json:"nutrient,omitempty"  bson:"nutrient,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
