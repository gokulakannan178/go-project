package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DistrictWeatherAlert struct {
	ID               primitive.ObjectID         `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus     bool                       `json:"activeStatus" bson:"activeStatus,omitempty"`
	UniqueID         string                     `json:"uniqueId" bson:"uniqueId,omitempty"`
	District         District                   `json:"district" form:"district," bson:"district,omitempty"`
	ParameterId      WeatherParameter           `json:"parameter" form:"parameter," bson:"parameter,omitempty"`
	SeverityType     WeatherAlertType           `json:"severityType,omitempty"  bson:"severityType,omitempty"`
	WeatherData      DistrictWeatherData        `json:"weatherData,omitempty"  bson:"weatherData,omitempty"`
	WeatherDataAlert DistrictWeatherAlertMaster `json:"weatherDataAlert,omitempty"  bson:"weatherDataAlert,omitempty"`
	Month            MonthSeason                `json:"month,omitempty"  bson:"month,omitempty"`
	Date             *time.Time                 `json:"date,omitempty"  bson:"date,omitempty"`
	Status           string                     `json:"status,omitempty"  bson:"status,omitempty"`
	Tittle           string                     `json:"tittle,omitempty"  bson:"tittle,omitempty"`
	ValueMax         float64                    `json:"valueMax,omitempty"  bson:"valueMax,omitempty"`
	Value            float64                    `json:"value,omitempty"  bson:"value,omitempty"`
	Created          *Created                   `json:"created,omitempty"  bson:"created,omitempty"`
}

type DistrictWeatherAlertFilter struct {
	Status           []string             `json:"status,omitempty" bson:"status,omitempty"`
	District         []primitive.ObjectID `json:"district" form:"district," bson:"district,omitempty"`
	ParameterId      []primitive.ObjectID `json:"parameter" form:"parameter," bson:"parameter,omitempty"`
	SeverityType     []primitive.ObjectID `json:"severityType,omitempty"  bson:"severityType,omitempty"`
	WeatherData      []primitive.ObjectID `json:"weatherData,omitempty"  bson:"weatherData,omitempty"`
	WeatherDataAlert []primitive.ObjectID `json:"weatherDataAlert,omitempty"  bson:"weatherDataAlert,omitempty"`
	Month            []primitive.ObjectID `json:"month,omitempty"  bson:"month,omitempty"`
	ActiveStatus     []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy           string               `json:"sortBy"`
	SortOrder        int                  `json:"sortOrder"`
	SearchBox        struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefDistrictWeatherAlert struct {
	DistrictWeatherAlert `bson:",inline"`
	Ref                  struct {
		State State `json:"state" form:"state," bson:"state,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type UpdateDistrictWeatherAlert struct {
	ID         primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	IsSms      string             `json:"isSms"  bson:"isSms,omitempty"`
	IsNicessm  string             `json:"isNicessm"  bson:"isNicessm,omitempty"`
	IsWhatsApp string             `json:"isWhatsApp"  bson:"isWhatsApp,omitempty"`
	IsTelegram string             `json:"isTelegram"  bson:"isTelegram,omitempty"`
	///	IsAutomatic  string             `json:"isAutomatic"  bson:"isAutomatic,omitempty"`
	WeatherAlert string `json:"weatherAlert"  bson:"weatherAlert,omitempty"`
	IsUpdateMode string `json:"isUpdateMode"  bson:"isUpdateMode,omitempty"`
}
type SendDistrictWeatherAlert struct {
	WeatherAlert DistrictWeatherAlert `json:"weatherAlert"  bson:"weatherAlert,omitempty"`
}
