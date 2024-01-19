package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAttendanceAction struct {
	UserName string   `json:"userName" bson:"userName,omitempty"`
	Image    string   `json:"image" bson:"image,omitempty"`
	Location Location `json:"location" bson:"location,omitempty"`
}

type UserAttendance struct {
	Id       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName string             `json:"userName" bson:"userName,omitempty"`
	Date     *time.Time         `json:"date" bson:"date,omitempty"`
	PunchIn  struct {
		Time     *time.Time `json:"time" bson:"time,omitempty"`
		Location Location   `json:"location" bson:"location,omitempty"`
		Image    string     `json:"image" bson:"image,omitempty"`
	} `json:"punchIn" bson:"punchIn,omitempty"`
	PunchOut struct {
		Time     *time.Time `json:"time" bson:"time,omitempty"`
		Location Location   `json:"location" bson:"location,omitempty"`
		Image    string     `json:"image" bson:"image,omitempty"`
	} `json:"punchOut" bson:"punchOut,omitempty"`
	Status    string `json:"status" bson:"status,omitempty"`
	PunchFrom string `json:"punchFrom" bson:"punchFrom,omitempty"`
}
