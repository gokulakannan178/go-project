package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DistrictWeatherAlertDissimination struct {
	ID           primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool                 `json:"activeStatus" bson:"activeStatus,omitempty"`
	UniqueID     string               `json:"uniqueId" bson:"uniqueId,omitempty"`
	WeatherAlert DistrictWeatherAlert `json:"weatherAlert" form:"weatherAlert," bson:"weatherAlert,omitempty"`
	Farmers      []DissiminateFarmer  `json:"farmers"  bson:"farmers,omitempty"`
	NoOfFarmers  int                  `json:"noOfFarmers"  bson:"noOfFarmers,omitempty"`
	Users        []DissiminateUser    `json:"users"  bson:"users,omitempty"`
	NoOfUsers    int                  `json:"noOfUsers"  bson:"noOfUsers,omitempty"`
	Date         *time.Time           `json:"date,omitempty"  bson:"date,omitempty"`
	Status       string               `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created             `json:"created,omitempty"  bson:"created,omitempty"`
}

type DistrictWeatherAlertDissiminationFilter struct {
	Status                 []string             `json:"status,omitempty" bson:"status,omitempty"`
	District               []primitive.ObjectID `json:"district" form:"district," bson:"district,omitempty"`
	ParameterId            []primitive.ObjectID `json:"parameter" form:"parameter," bson:"parameter,omitempty"`
	SeverityType           []primitive.ObjectID `json:"severityType,omitempty"  bson:"severityType,omitempty"`
	WeatherData            []primitive.ObjectID `json:"weatherData,omitempty"  bson:"weatherData,omitempty"`
	WeatherDataAlert       []primitive.ObjectID `json:"weatherDataAlert,omitempty"  bson:"weatherDataAlert,omitempty"`
	Month                  []primitive.ObjectID `json:"month,omitempty"  bson:"month,omitempty"`
	DateDisseminationRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateDisseminationRange"`
	ActiveStatus []bool `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string `json:"sortBy"`
	SortOrder    int    `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefDistrictWeatherAlertDissimination struct {
	DistrictWeatherAlertDissimination `bson:",inline"`
	Ref                               struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
