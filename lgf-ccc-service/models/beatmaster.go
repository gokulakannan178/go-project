package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BeatMaster struct {
	ID         primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	RouteID    string             `json:"routeId,omitempty" bson:"routeId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Date       *time.Time         `json:"date" bson:"date,omitempty"`
	StartTime  *time.Time         `json:"startTime" bson:"startTime,omitempty"`
	EndTime    *time.Time         `json:"EndTime" bson:"EndTime,omitempty"`
	Driver     MinUser            `json:"driver" bson:"driver,omitempty"`
	Vehicle    MinUser            `json:"vehicle" bson:"vehicle,omitempty"`
	ContactNo  string             `json:"contactNo" bson:"contactNo,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	EmployeeId string             `json:"employeeID" bson:"employeeID,omitempty"`
	Created    *Created           `json:"created"  bson:"created,omitempty"`
	Area       Address            `json:"area" bson:"area,omitempty"`
}

type RefBeatMaster struct {
	BeatMaster `bson:",inline"`
	Ref        struct {
		Route RouteMaster `json:"route" bson:"route,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterBeatMaster struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId []string `json:"employeeID" bson:"employeeID,omitempty"`
	RouteID    []string `json:"routeId,omitempty" bson:"routeId,omitempty"`
	WardCode   []string `json:"wardCode" bson:"wardCode,omitempty"`
	ZoneCode   []string `json:"zoneCode" bson:"zoneCode,omitempty"`
	VehicleId  []string `json:"vehicleId" bson:"vehicleId,omitempty"`
	DriverId   []string `json:"driverId" bson:"driverId,omitempty"`
	DateRange  struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	Regex struct {
		Name      string `json:"name" bson:"name"`
		Type      string `json:"type" bson:"type"`
		ContactNo string `json:"ContactNo" bson:"ContactNo"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type VehicleAssignBeat struct {
	UniqueID  string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	VehicleID string `json:"vehicleId,omitempty" bson:"vehicleId,omitempty"`
}
