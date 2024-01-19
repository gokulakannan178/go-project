package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeatherAlertNotInRange struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	State        State              `json:"state" form:"state," bson:"state,omitempty"`
	ParameterId  WeatherParameter   `json:"parameter" form:"parameter," bson:"parameter,omitempty"`
	WeatherData  StateWeatherData   `json:"weatherData,omitempty"  bson:"weatherData,omitempty"`
	Month        MonthSeason        `json:"month,omitempty"  bson:"month,omitempty"`
	Date         *time.Time         `json:"date,omitempty"  bson:"date,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Tittle       string             `json:"tittle,omitempty"  bson:"tittle,omitempty"`
	Value        float64            `json:"value,omitempty"  bson:"value,omitempty"`
	ValueMax     float64            `json:"valueMax,omitempty"  bson:"valueMax,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type WeatherAlertNotInRangeFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefWeatherAlertNotInRange struct {
	WeatherAlertNotInRange `bson:",inline"`
	Ref                    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
