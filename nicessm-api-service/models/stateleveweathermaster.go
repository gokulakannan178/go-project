package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type StateWeatherAlertMaster struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	State        primitive.ObjectID `json:"state" form:"state," bson:"state,omitempty"`
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

type StateWeatherAlertMasterFilter struct {
	Status       []string             `json:"status,omitempty" bson:"status,omitempty"`
	State        []primitive.ObjectID `json:"state" form:"state," bson:"state,omitempty"`
	District     []primitive.ObjectID `json:"district" form:"district," bson:"district,omitempty"`
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
type StateWeatherAlertMasterFilterv2 struct {
	ParameterId primitive.ObjectID `json:"parameterid" form:"parameterid," bson:"parameterid,omitempty"`
	Month       primitive.ObjectID `json:"month" form:"month," bson:"month,omitempty"`
	Season      primitive.ObjectID `json:"season" form:"season," bson:"season,omitempty"`
}

type RefStateWeatherAlertMaster struct {
	StateWeatherAlertMaster `bson:",inline"`
	Ref                     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type GetWeatherAlertMaster struct {
	Block  string   `json:"block" form:"block," bson:"block,omitempty"`
	Values []Values `json:"values" form:"values," bson:"values,omitempty"`
}
type Values struct {
	SeverityType string  `json:"severityType,omitempty"  bson:"severityType,omitempty"`
	Max          float64 `json:"max,omitempty"  bson:"max,omitempty"`
	Min          float64 `json:"min,omitempty"  bson:"min,omitempty"`
}
type GetStateLeveWeatherDataAlert struct {
	State       `bson:",inline"`
	ServentType struct {
		Medium   ServentType `json:"medium,omitempty"  bson:"medium,omitempty"`
		Minimum  ServentType `json:"minimum,omitempty"  bson:"minimum,omitempty"`
		Disaster ServentType `json:"disaster,omitempty"  bson:"disaster,omitempty"`
	} `json:"severitytype,omitempty"  bson:"severitytype,omitempty"`
}
type ServentType struct {
	ID           primitive.ObjectID      `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool                    `json:"activeStatus" bson:"activeStatus,omitempty"`
	Name         string                  `json:"name,omitempty"  bson:"name,omitempty"`
	Unit         string                  `json:"unit,omitempty"  bson:"unit,omitempty"`
	Status       string                  `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created                `json:"created,omitempty"  bson:"created,omitempty"`
	Weatherdata  StateWeatherAlertMaster `json:"weatherdata,omitempty"  bson:"weatherdata,omitempty"`
}
type Weatherdata struct {
	Max float64 `json:"max,omitempty"  bson:"max,omitempty"`
	Min float64 `json:"min,omitempty"  bson:"min,omitempty"`
}
