package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoadBlockWeatherReport struct {
	ID                  primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name                string             `json:"name" bson:"name,omitempty"`
	Alto                float64            `json:"alto" bson:"alto"`
	Cumrain             float64            `json:"cumrain" bson:"cumrain,omitempty"`
	Block               primitive.ObjectID `json:"block,omitempty"  bson:"block,omitempty"`
	CreatedAt           *time.Time         `json:"createdAt" bson:"createdAt,omitempty"`
	MaxRelativeHumidity float64            `json:"maxRelativeHumidity,omitempty"  bson:"maxRelativeHumidity,omitempty"`
	Icld                string             `json:"icld,omitempty"  bson:"icld,omitempty"`
	MaxTemperature      float64            `json:"maxTemperature,omitempty"  bson:"maxTemperature,omitempty"`
	Latitude            float64            `json:"latitude,omitempty"  bson:"latitude,omitempty"`
	MinRelativeHumidity float64            `json:"minRelativeHumidity,omitempty"  bson:"minRelativeHumidity,omitempty"`
	Longitude           float64            `json:"longitude,omitempty"  bson:"longitude,omitempty"`
	MinTemperature      float64            `json:"minTemperature,omitempty"  bson:"minTemperature,omitempty"`
	Month               int                `json:"month,omitempty"  bson:"month,omitempty"`
	Rainfall            float64            `json:"rainfall,omitempty"  bson:"rainfall,omitempty"`
	Mslp                float64            `json:"mslp,omitempty"  bson:"mslp,omitempty"`
	Version             int                `json:"version,omitempty"  bson:"version,omitempty"`
	Date                *time.Time         `json:"date,omitempty"  bson:"date,omitempty"`
	Day                 int                `json:"day,omitempty"  bson:"day,omitempty"`
	District            primitive.ObjectID `json:"district,omitempty"  bson:"district,omitempty"`
	Hour                int                `json:"hour,omitempty"  bson:"hour,omitempty"`
	Pcod                string             `json:"pcod,omitempty"  bson:"pcod,omitempty"`
	State               primitive.ObjectID `json:"state,omitempty"  bson:"state,omitempty"`
	Type                string             `json:"type,omitempty"  bson:"type,omitempty"`
	WindDirection       float64            `json:"windDirection,omitempty"  bson:"windDirection,omitempty"`
	WindSpeed           float64            `json:"windSpeed,omitempty"  bson:"windSpeed,omitempty"`
	Year                int                `json:"year,omitempty"  bson:"year,omitempty"`
	Status              string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created             *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}
