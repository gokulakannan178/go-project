package models

import "time"

//Report : ""
type MonthWiseDumphistoryCount struct {
	ID       int64 `json:"id" bson:"_id,omitempty"`
	Date     int64 `json:"date" bson:"date,omitempty"`
	Quantity int64 `json:"quantity" bson:"quantity,omitempty"`
}

type CircleWiseHouseVisitedv2 struct {
	Date     *time.Time `json:"date" bson:"date,omitempty"`
	Code     string     `json:"code" bson:"code,omitempty"`
	Quantity int64      `json:"quantity" bson:"quantity,omitempty"`
}
