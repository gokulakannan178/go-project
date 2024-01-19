package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DistrictWeatherAlertMaster struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	District     primitive.ObjectID `json:"district" form:"district," bson:"district,omitempty"`
	ParameterId  primitive.ObjectID `json:"parameterid" form:"parameterid," bson:"parameterid,omitempty"`
	Month        primitive.ObjectID `json:"month" form:"month," bson:"month,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	SeverityType primitive.ObjectID `json:"severityType,omitempty"  bson:"severityType,omitempty"`
	Max          float64            `json:"max"  bson:"max,omitempty"`
	Min          float64            `json:"min"  bson:"min,omitempty"`
	IsSms        string             `json:"isSms"  bson:"isSms,omitempty"`
	IsNicessm    string             `json:"isNicessm"  bson:"isNicessm,omitempty"`
	IsWhatsApp   string             `json:"isWhatsApp"  bson:"isWhatsApp,omitempty"`
	IsTelegram   string             `json:"isTelegram"  bson:"isTelegram,omitempty"`
	IsAutomatic  string             `json:"isAutomatic"  bson:"isAutomatic,omitempty"`
	WeatherAlert string             `json:"weatherAlert"  bson:"weatherAlert,omitempty"`
	Unit         string             `json:"unit,omitempty"  bson:"unit,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type DistrictWeatherAlertMasterFilter struct {
	Status       []string             `json:"status,omitempty" bson:"status,omitempty"`
	District     []primitive.ObjectID `json:"district" form:"district," bson:"district,omitempty"`
	State        []primitive.ObjectID `json:"state" form:"state," bson:"state,omitempty"`
	Block        []primitive.ObjectID `json:"block" form:"block," bson:"block,omitempty"`
	ParameterId  []primitive.ObjectID `json:"parameterid" form:"parameterid," bson:"parameterid,omitempty"`
	Month        []primitive.ObjectID `json:"month" form:"month," bson:"month,omitempty"`
	Season       []primitive.ObjectID `json:"season" form:"season," bson:"season,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	// SearchBox    struct {
	// 	Max string `json:"max" bson:"max"`
	// } `json:"searchbox" bson:"searchbox"`
}
type DistrictWeatherAlertMasterFilterv2 struct {
	ParameterId primitive.ObjectID `json:"parameterid" form:"parameterid," bson:"parameterid,omitempty"`
	Month       primitive.ObjectID `json:"month" form:"month," bson:"month,omitempty"`
	Season      primitive.ObjectID `json:"season" form:"season," bson:"season,omitempty"`
}

type RefDistrictWeatherAlertMaster struct {
	DistrictWeatherAlertMaster `bson:",inline"`
	Ref                        struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type GetDistrictLeveWeatherDataAlert struct {
	District    `bson:",inline"`
	ServentType struct {
		Medium   ServentType `json:"medium,omitempty"  bson:"medium,omitempty"`
		Minimum  ServentType `json:"minimum,omitempty"  bson:"minimum,omitempty"`
		Disaster ServentType `json:"disaster,omitempty"  bson:"disaster,omitempty"`
	} `json:"severitytype,omitempty"  bson:"severitytype,omitempty"`
}
