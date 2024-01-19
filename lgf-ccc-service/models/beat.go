package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Beat struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	BeatMasterID       string             `json:"beatMasterId,omitempty" bson:"beatMasterId,omitempty"`
	AssignBeatdetails  BeatMaster         `json:"assignBeatdetails,omitempty" bson:"assignBeatdetails,omitempty"`
	AbsorveBeatDetails string             `json:"absorveBeatDetails,omitempty" bson:"absorveBeatDetails,omitempty"`
	Name               string             `json:"name" bson:"name,omitempty"`
	Date               *time.Time         `json:"date" bson:"date,omitempty"`
	StartTime          *time.Time         `json:"startTime" bson:"startTime,omitempty"`
	EndTime            *time.Time         `json:"endTime" bson:"endTime,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            *Created           `json:"created"  bson:"created,omitempty"`
}

type RefBeat struct {
	Beat `bson:",inline"`
	Ref  struct {
		Beatmaster      BeatMaster        `json:"beatmaster" bson:"beatmaster,omitempty"`
		Route           RouteMaster       `json:"route" bson:"route,omitempty"`
		VehicleLocation []VehicleLocation `json:"vehicleLocation" bson:"vehicleLocation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterBeat struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId []string `json:"employeeID" bson:"employeeID,omitempty"`
	VehicleId  []string `json:"vehicleId" bson:"vehicleId,omitempty"`

	Regex struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`

	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
