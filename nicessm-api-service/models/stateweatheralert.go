package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StateWeatherAlert struct {
	ID               primitive.ObjectID      `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus     bool                    `json:"activeStatus" bson:"activeStatus,omitempty"`
	UniqueID         string                  `json:"uniqueId" bson:"uniqueId,omitempty"`
	State            State                   `json:"state" form:"state," bson:"state,omitempty"`
	ParameterId      WeatherParameter        `json:"parameter" form:"parameter," bson:"parameter,omitempty"`
	SeverityType     WeatherAlertType        `json:"severityType,omitempty"  bson:"severityType,omitempty"`
	WeatherData      StateWeatherData        `json:"weatherData,omitempty"  bson:"weatherData,omitempty"`
	WeatherDataAlert StateWeatherAlertMaster `json:"weatherDataAlert,omitempty"  bson:"weatherDataAlert,omitempty"`
	Month            MonthSeason             `json:"month,omitempty"  bson:"month,omitempty"`
	Date             *time.Time              `json:"date,omitempty"  bson:"date,omitempty"`
	Status           string                  `json:"status,omitempty"  bson:"status,omitempty"`
	Tittle           string                  `json:"tittle,omitempty"  bson:"tittle,omitempty"`
	ValueMax         float64                 `json:"valueMax,omitempty"  bson:"valueMax,omitempty"`
	Value            float64                 `json:"value,omitempty"  bson:"value,omitempty"`
	Created          *Created                `json:"created,omitempty"  bson:"created,omitempty"`
}

type StateWeatherAlertFilter struct {
	Status           []string             `json:"status,omitempty" bson:"status,omitempty"`
	State            []primitive.ObjectID `json:"state" form:"state," bson:"state,omitempty"`
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

type RefStateWeatherAlert struct {
	StateWeatherAlert `bson:",inline"`
	Ref               struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type UpdateStateWeatherAlert struct {
	ID         primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	IsSms      string             `json:"isSms"  bson:"isSms,omitempty"`
	IsNicessm  string             `json:"isNicessm"  bson:"isNicessm,omitempty"`
	IsWhatsApp string             `json:"isWhatsApp"  bson:"isWhatsApp,omitempty"`
	IsTelegram string             `json:"isTelegram"  bson:"isTelegram,omitempty"`
	///	IsAutomatic  string             `json:"isAutomatic"  bson:"isAutomatic,omitempty"`
	WeatherAlert string `json:"weatherAlert"  bson:"weatherAlert,omitempty"`
	IsUpdateMode string `json:"isUpdateMode"  bson:"isUpdateMode,omitempty"`
}
type SendStateWeatherAlert struct {
	WeatherAlert StateWeatherAlert `json:"weatherAlert"  bson:"weatherAlert,omitempty"`
}
type WeatherDisseminationChennal struct {
	Users        []DissiminateUser   `json:"users" bson:"users,omitempty"`
	Farmers      []DissiminateFarmer `json:"farmers" bson:"farmers,omitempty"`
	WeatherAlert StateWeatherAlert   `json:"weatherAlert" form:"weatherAlert," bson:"weatherAlert,omitempty"`
}
type DissseminationUserFarmer struct {
	Farmers     []DissiminateFarmer `json:"farmers"  bson:"farmers,omitempty"`
	NoOfFarmers int                 `json:"noOfFarmers"  bson:"noOfFarmers,omitempty"`
	Users       []DissiminateUser   `json:"users"  bson:"users,omitempty"`
	NoOfUsers   int                 `json:"noOfUsers"  bson:"noOfUsers,omitempty"`
}
