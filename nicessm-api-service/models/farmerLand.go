package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//FarmerLand : "Holds single FarmerLand data"
type FarmerLand struct {
	ID                  primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	FarmerID            string             `json:"farmerId" bson:"-"`
	Status              string             `json:"status" bson:"status,omitempty"`
	ActiveStatus        bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created             Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Farmer              primitive.ObjectID `json:"farmer"  bson:"farmer,omitempty"`
	GramPanchayat       primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village             primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	Version             string             `json:"version"  bson:"version,omitempty"`
	Description         string             `json:"description"  bson:"description,omitempty"`
	IrrigationType      string             `json:"irrigationType"  bson:"irrigationType,omitempty"`
	CultivationPractice string             `json:"cultivationPractice"  bson:"cultivationPractice,omitempty"`
	LandType            string             `json:"landType"  bson:"landType,omitempty"`
	LandPosition        string             `json:"landPosition"  bson:"landPosition,omitempty"`
	OwnerShip           string             `json:"ownerShip"  bson:"ownerShip,omitempty"`
	Unit                string             `json:"unit"  bson:"unit,omitempty"`
	CultivatedArea      float64            `json:"cultivatedArea"  bson:"cultivatedArea,omitempty"`
	VacantArea          float64            `json:"vacantArea"  bson:"vacantArea,omitempty"`
	AreaInAcre          float64            `json:"areaInAcre"  bson:"areaInAcre,omitempty"`
	State               primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block               primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District            primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	SoilType            primitive.ObjectID `json:"soilType"  bson:"soilType,omitempty"`
	ParcelNumber        string             `json:"parcelNumber"  bson:"parcelNumber,omitempty"`
	KhasraNumber        string             `json:"khasraNumber"  bson:"khasraNumber,omitempty"`
}

//RefFarmerLand : "FarmerLand with refrence data such as language..."
type RefFarmerLand struct {
	FarmerLand `bson:",inline"`
	Ref        struct {
		State         State         `json:"state"  bson:"state,omitempty"`
		Block         Block         `json:"block"  bson:"block,omitempty"`
		District      District      `json:"district"  bson:"district,omitempty"`
		SoilType      SoilType      `json:"soilType"  bson:"soilType,omitempty"`
		GramPanchayat GramPanchayat `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
		Village       Village       `json:"village"  bson:"village,omitempty"`
		LandCrop      LandCrop      `json:"landCrop"  bson:"landCrop,omitempty"`
		Crop          Commodity     `json:"crop"  bson:"crop,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type Land struct {
	CultivatedArea float64 `json:"cultivatedArea"  bson:"cultivatedArea,omitempty"`
	VacantArea     float64 `json:"vacantArea"  bson:"vacantArea,omitempty"`
	TotalLand      float64 `json:"totalLand"  bson:"totalLand,omitempty"`
}

//FarmerLandFilter : "Used for constructing filter query"
type FarmerLandFilter struct {
	ActiveStatus        []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	GramPanchayat       []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village             []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	State               []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block               []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District            []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	SoilType            []primitive.ObjectID `json:"soilType"  bson:"soilType,omitempty"`
	Farmer              []primitive.ObjectID `json:"farmer"  bson:"farmer,omitempty"`
	Status              []string             `json:"status" bson:"status,omitempty"`
	SortBy              string               `json:"sortBy"`
	SortOrder           int                  `json:"sortOrder"`
	IrrigationType      []string             `json:"irrigationType"  bson:"irrigationType,omitempty"`
	CultivationPractice []string             `json:"cultivationPractice"  bson:"cultivationPractice,omitempty"`
	LandType            []string             `json:"landType"  bson:"landType,omitempty"`
	LandPosition        []string             `json:"landPosition"  bson:"landPosition,omitempty"`
	OwnerShip           []string             `json:"ownerShip"  bson:"ownerShip,omitempty"`

	Regex struct {
		Type string `json:"type"  bson:"type"`

		ParcelNumber string `json:"parcelNumber"  bson:"parcelNumber"`
	} `json:"regex" bson:"regex"`
}
