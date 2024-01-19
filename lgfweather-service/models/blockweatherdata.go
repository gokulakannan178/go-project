package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BlockWeatherData : ""
type BlockWeatherDataV2 struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Block        primitive.ObjectID `json:"Block" bson:"Block,omitempty"`
	Sunrise      float64            `json:"sunrise,omitempty"  bson:"sunrise,omitempty"`
	Sunset       float64            `json:"sunset,omitempty"  bson:"sunset,omitempty"`
	Moonrise     float64            `json:"moonrise,omitempty"  bson:"moonrise,omitempty"`
	Moonset      float64            `json:"moonset,omitempty"  bson:"moonset,omitempty"`
	Temp         Temp               `json:"temp,omitempty"  bson:"temp,omitempty"`
	Pressure     float64            `json:"pressure,omitempty"  bson:"pressure,omitempty"`
	Humidity     float64            `json:"humidity,omitempty"  bson:"humidity,omitempty"`
	Wind_speed   float64            `json:"wind_speed,omitempty"  bson:"wind_speed,omitempty"`
	Wind_deg     float64            `json:"wind_deg,omitempty"  bson:"wind_deg,omitempty"`
	Wind_gust    float64            `json:"wind_gust,omitempty"  bson:"wind_gust,omitempty"`
	Weather      []WeatherType      `json:"weather,omitempty"  bson:"weather,omitempty"`
	Clouds       float64            `json:"clouds,omitempty"  bson:"clouds,omitempty"`
	Pop          float64            `json:"pop,omitempty"  bson:"pop,omitempty"`
	Rain         float64            `json:"rain,omitempty"  bson:"rain,omitempty"`
	Uvi          float64            `json:"uvi,omitempty"  bson:"uvi,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	CreatedDate  *time.Time         `json:"createdDate,omitempty"  bson:"createdDate,omitempty"`
	Date         *time.Time         `json:"date,omitempty"  bson:"date,omitempty"`
}
type BlockWeatherData struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	State        primitive.ObjectID `json:"state" bson:"state,omitempty"`
	Block        primitive.ObjectID `json:"Block" bson:"Block,omitempty"`
	District     primitive.ObjectID `json:"district" bson:"district,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	CreatedDate  *time.Time         `json:"createdDate,omitempty"  bson:"createdDate,omitempty"`
	Location     Location           `json:"location" bson:"location,omitempty"`
	Alto         float64            `json:"alto,omitempty"  bson:"alto,omitempty"`
	Pcod         string             `json:"pcod,omitempty"  bson:"pcod,omitempty"`
	Mspl         float64            `json:"mspl,omitempty"  bson:"mspl,omitempty"`
	Icid         float64            `json:"icid,omitempty"  bson:"icid,omitempty"`
	Date         time.Time          `json:"date,omitempty"  bson:"date,omitempty"`
	WeatherData  DailyWeatherDataV2 `json:"weatherData,omitempty"  bson:"weatherData,omitempty"`
}

type BlockWeatherDataFilter struct {
	Status       []string             `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Block        []primitive.ObjectID `json:"Block" bson:"Block,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	DateRange    *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefBlockWeatherData struct {
	BlockWeatherData `bson:",inline"`
	Ref              struct {
		Block Block `json:"Block" bson:"Block,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
