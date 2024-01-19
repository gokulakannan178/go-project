package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//StateWeatherData : ""
type StateWeatherDataV2 struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	State        primitive.ObjectID `json:"state" bson:"state,omitempty"`
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
type StateWeatherData struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	State        primitive.ObjectID `json:"state" bson:"state,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
	CreatedDate  *time.Time         `json:"createdDate,omitempty"  bson:"createdDate,omitempty"`
	Date         time.Time          `json:"date,omitempty"  bson:"date,omitempty"`
	WeatherData  DailyWeatherData   `json:"weatherData,omitempty"  bson:"weatherData,omitempty"`
}

type StateWeatherDataFilter struct {
	Status       []string             `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	State        []primitive.ObjectID `json:"state" bson:"state,omitempty"`
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

type RefStateWeatherData struct {
	StateWeatherData `bson:",inline"`
	Ref              struct {
		State State `json:"state" bson:"state,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type Temp struct {
	Day   float64 `json:"day,omitempty"  bson:"day,omitempty"`
	Min   float64 `json:"min,omitempty"  bson:"min,omitempty"`
	Max   float64 `json:"max,omitempty"  bson:"max,omitempty"`
	Night float64 `json:"night,omitempty"  bson:"night,omitempty"`
	Eve   float64 `json:"eve,omitempty"  bson:"eve,omitempty"`
	Morn  float64 `json:"morn,omitempty"  bson:"morn,omitempty"`
}
type WeatherType struct {
	Id          int    `json:"id,omitempty"  bson:"id,omitempty"`
	Main        string `json:"main,omitempty"  bson:"main,omitempty"`
	Description string `json:"description,omitempty"  bson:"description,omitempty"`
	Icon        string `json:"icon,omitempty"  bson:"icon,omitempty"`
}
type WeatherDataMaster struct {
	Lat            float64              `json:"lat,omitempty"  bson:"lat,omitempty"`
	Lon            float64              `json:"lon,omitempty"  bson:"lon,omitempty"`
	Timezone       string               `json:"timezone,omitempty"  bson:"timezone,omitempty"`
	Timezoneoffset float64              `json:"timezone_offset,omitempty"  bson:"timezone_offset,omitempty"`
	Current        CurrentWeatherData   `json:"current,omitempty"  bson:"current,omitempty"`
	Minutely       []MinutelyWeatheData `json:"minutely,omitempty"  bson:"minutely,omitempty"`
	Hourly         []hourlyWeatherData  `json:"hourly,omitempty"  bson:"hourly,omitempty"`
	Daily          []DailyWeatherData   `json:"daily,omitempty"  bson:"daily,omitempty"`
}
type CurrentWeatherData struct {
	Dt         float64       `json:"dt,omitempty"  bson:"dt,omitempty"`
	Sunrise    float64       `json:"sunrise,omitempty"  bson:"sunrise,omitempty"`
	Sunset     float64       `json:"sunset,omitempty"  bson:"sunset,omitempty"`
	Temp       float64       `json:"temp,omitempty"  bson:"temp,omitempty"`
	Feelslike  float64       `json:"feels_like,omitempty"  bson:"feels_like,omitempty"`
	Pressure   float64       `json:"pressure,omitempty"  bson:"pressure,omitempty"`
	Humidity   float64       `json:"humidity,omitempty"  bson:"humidity,omitempty"`
	Dewpoint   float64       `json:"dew_point,omitempty"  bson:"dew_point,omitempty"`
	Uvi        float64       `json:"uvi,omitempty"  bson:"uvi,omitempty"`
	Clouds     float64       `json:"clouds,omitempty"  bson:"clouds,omitempty"`
	Visibility float64       `json:"visibility,omitempty"  bson:"visibility,omitempty"`
	Windspeed  float64       `json:"wind_speed,omitempty"  bson:"wind_speed,omitempty"`
	Winddeg    float64       `json:"wind_deg,omitempty"  bson:"wind_deg,omitempty"`
	Windgust   float64       `json:"wind_gust,omitempty"  bson:"wind_gust,omitempty"`
	Weather    []WeatherType `json:"weather,omitempty"  bson:"weather,omitempty"`
}
type MinutelyWeatheData struct {
	Dt            float64 `json:"dt,omitempty"  bson:"dt,omitempty"`
	Precipitation float64 `json:"precipitation,omitempty"  bson:"precipitation,omitempty"`
}
type hourlyWeatherData struct {
	Dt         float64       `json:"dt,omitempty"  bson:"dt,omitempty"`
	Temp       float64       `json:"temp,omitempty"  bson:"temp,omitempty"`
	Feelslike  float64       `json:"feels_like,omitempty"  bson:"feels_like,omitempty"`
	Pressure   float64       `json:"pressure,omitempty"  bson:"pressure,omitempty"`
	Humidity   float64       `json:"humidity,omitempty"  bson:"humidity,omitempty"`
	Dewpoint   float64       `json:"dew_point,omitempty"  bson:"dew_point,omitempty"`
	Uvi        float64       `json:"uvi,omitempty"  bson:"uvi,omitempty"`
	Clouds     float64       `json:"clouds,omitempty"  bson:"clouds,omitempty"`
	Visibility float64       `json:"visibility,omitempty"  bson:"visibility,omitempty"`
	Windspeed  float64       `json:"wind_speed,omitempty"  bson:"wind_speed,omitempty"`
	Winddeg    float64       `json:"wind_deg,omitempty"  bson:"wind_deg,omitempty"`
	Windgust   float64       `json:"wind_gust,omitempty"  bson:"wind_gust,omitempty"`
	Weather    []WeatherType `json:"weather,omitempty"  bson:"weather,omitempty"`
	Pop        float64       `json:"pop,omitempty"  bson:"pop,omitempty"`
}
type DailyWeatherData struct {
	Dt         float64       `json:"dt,omitempty"  bson:"dt,omitempty"`
	Sunrise    int           `json:"sunrise,omitempty"  bson:"sunrise,omitempty"`
	Sunset     float64       `json:"sunset,omitempty"  bson:"sunset,omitempty"`
	Moonrise   float64       `json:"moonrise,omitempty"  bson:"moonrise,omitempty"`
	Moonset    float64       `json:"moonset,omitempty"  bson:"moonset,omitempty"`
	Moon_phase float64       `json:"moon_phase,omitempty"  bson:"moon_phase,omitempty"`
	Temp       Temp          `json:"temp,omitempty"  bson:"temp,omitempty"`
	Feelslike  FeelslikeData `json:"feels_like,omitempty"  bson:"feels_like,omitempty"`
	Pressure   float64       `json:"pressure,omitempty"  bson:"pressure,omitempty"`
	Humidity   float64       `json:"humidity,omitempty"  bson:"humidity,omitempty"`
	Dewpoint   float64       `json:"dew_point,omitempty"  bson:"dew_point,omitempty"`
	Windspeed  float64       `json:"wind_speed,omitempty"  bson:"wind_speed,omitempty"`
	Winddeg    float64       `json:"wind_deg,omitempty"  bson:"wind_deg,omitempty"`
	Windgust   float64       `json:"wind_gust,omitempty"  bson:"wind_gust,omitempty"`
	Weather    []WeatherType `json:"weather,omitempty"  bson:"weather,omitempty"`
	Clouds     float64       `json:"clouds,omitempty"  bson:"clouds,omitempty"`
	Pop        float64       `json:"pop,omitempty"  bson:"pop,omitempty"`
	Rain       float64       `json:"rain,omitempty"  bson:"rain,omitempty"`
	Uvi        float64       `json:"uvi,omitempty"  bson:"uvi,omitempty"`
}

//for use struct
type DailyWeatherDataV2 struct {
	Dt          float64       `json:"dt,omitempty"  bson:"dt,omitempty"`
	Sunrise     int           `json:"sunrise,omitempty"  bson:"sunrise,omitempty"`
	Sunset      float64       `json:"sunset,omitempty"  bson:"sunset,omitempty"`
	Moonrise    float64       `json:"moonrise,omitempty"  bson:"moonrise,omitempty"`
	Moonset     float64       `json:"moonset,omitempty"  bson:"moonset,omitempty"`
	Moon_phase  float64       `json:"moon_phase,omitempty"  bson:"moon_phase,omitempty"`
	Temp        Temp          `json:"temp,omitempty"  bson:"temp,omitempty"`
	Feelslike   FeelslikeData `json:"feels_like,omitempty"  bson:"feels_like,omitempty"`
	Pressure    float64       `json:"pressure,omitempty"  bson:"pressure,omitempty"`
	HumidityMax float64       `json:"humidityMax,omitempty"  bson:"humidityMax,omitempty"`
	HumidityMin float64       `json:"humidityMin,omitempty"  bson:"humidityMin,omitempty"`
	Dewpoint    float64       `json:"dew_point,omitempty"  bson:"dew_point,omitempty"`
	Windspeed   float64       `json:"wind_speed,omitempty"  bson:"wind_speed,omitempty"`
	Winddeg     float64       `json:"wind_deg,omitempty"  bson:"wind_deg,omitempty"`
	Windgust    float64       `json:"wind_gust,omitempty"  bson:"wind_gust,omitempty"`
	Weather     []WeatherType `json:"weather,omitempty"  bson:"weather,omitempty"`
	Clouds      float64       `json:"clouds,omitempty"  bson:"clouds,omitempty"`
	Pop         float64       `json:"pop,omitempty"  bson:"pop,omitempty"`
	Rain        float64       `json:"rain,omitempty"  bson:"rain,omitempty"`
	Uvi         float64       `json:"uvi,omitempty"  bson:"uvi,omitempty"`
}
type FeelslikeData struct {
	Day   float64 `json:"day,omitempty"  bson:"day,omitempty"`
	Night float64 `json:"night,omitempty"  bson:"night,omitempty"`
	Eve   float64 `json:"eve,omitempty"  bson:"eve,omitempty"`
	Morn  float64 `json:"morn,omitempty"  bson:"morn,omitempty"`
}
