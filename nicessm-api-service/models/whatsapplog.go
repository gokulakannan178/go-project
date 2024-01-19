package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Asset : ""
type WhatsappLog struct {
	ID                   primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	Content              primitive.ObjectID   `json:"content,omitempty" bson:"content,omitempty"`
	Query                primitive.ObjectID   `json:"query,omitempty"  bson:"query,omitempty"`
	ContentRecordId      string               `json:"contentRecordId,omitempty" bson:"contentRecordId,omitempty"`
	QueryRecordId        string               `json:"queryRecordId,omitempty" bson:"queryRecordId,omitempty"`
	WeatherAlert         StateWeatherAlert    `json:"weatherAlert,omitempty" bson:"weatherAlert,omitempty"`
	DistrictWeatherAlert DistrictWeatherAlert `json:"districtWeatherAlert,omitempty" bson:"districtWeatherAlert,omitempty"`
	SentFor              string               `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	IsJob                bool                 `json:"isJob"  bson:"isJob,omitempty"`
	Message              string               `json:"message,omitempty"  bson:"message,omitempty"`
	Status               string               `json:"status,omitempty"  bson:"status,omitempty"`
	SentDate             *time.Time           `json:"sentDate,omitempty"  bson:"sentDate,omitempty"`
	To                   ToWhatsappLog        `json:"to,omitempty"  bson:"to,omitempty"`
	Created              *Created             `json:"created,omitempty"  bson:"created,omitempty"`
}

type WhatsappLogFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	IsJob     []bool   `json:"isJob"  bson:"isJob,omitempty"`
	No        []string `json:"no,omitempty" bson:"no,omitempty"`
	Name      []string `json:"name"  bson:"name,omitempty"`
	UserName  []string `json:"userName" bson:"userName,omitempty"`
	UserType  []string `json:"userType" bson:"userType,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Regex     struct {
		SentFor string `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	} `json:"regex" bson:"regex"`
}

type RefWhatsappLog struct {
	WhatsappLog `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type ToWhatsappLog struct {
	No                string             `json:"no,omitempty" bson:"no,omitempty"`
	Name              string             `json:"name"  bson:"name,omitempty"`
	UserName          string             `json:"userName" bson:"userName,omitempty"`
	UserType          string             `json:"userType" bson:"userType,omitempty"`
	UserId            primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Gender            string             `json:"gender" bson:"gender,omitempty"`
	State             string             `json:"state"  bson:"state,omitempty"`
	District          string             `json:"district"  bson:"district,omitempty"`
	Block             string             `json:"block"  bson:"block,omitempty"`
	GramPanchayat     string             `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village           string             `json:"village"  bson:"village,omitempty"`
	StateCode         primitive.ObjectID `json:"stateCode"  bson:"stateCode,omitempty"`
	DistricCode       primitive.ObjectID `json:"districtCode"  bson:"districtCode,omitempty"`
	BlockCode         primitive.ObjectID `json:"blockCode"  bson:"blockCode,omitempty"`
	GramPanchayatCode primitive.ObjectID `json:"gramPanchayatCode"  bson:"gramPanchayatCode,omitempty"`
	VillageCode       primitive.ObjectID `json:"villageCode"  bson:"villageCode,omitempty"`
}
