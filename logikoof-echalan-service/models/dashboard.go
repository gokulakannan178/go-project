package models

import (
	"time"
)

//PaymentWidget : ""
type PaymentWidget struct {
	Payments struct {
		Recovered int64 `json:"recovered" bson:"Completed,omitempty"`
		Pending   int64 `json:"pending" bson:"Pending,omitempty"`
	} `json:"payments" bson:"payments,omitempty"`
}

//PaymentWidgetFilter : ""
type PaymentWidgetFilter struct {
	VehicleType []string `json:"vehicleType" bson:"vehicleType,omitempty"`
	OffenceType []string `json:"offenceType" bson:"offenceType,omitempty"`
	DateRange   *struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"dateRange" bson:"dateRange,omitempty"`
}

//TodaysOffenceWidget : ""
type TodaysOffenceWidget struct {
	Types struct {
		Car  int64 `json:"car" bson:"CAR,omitempty"`
		Bike int64 `json:"bike" bson:"BIKE,omitempty"`
		HV   int64 `json:"hv" bson:"HEAVYVEHICLE,omitempty"`
	} `json:"types" bson:"types,omitempty"`
}

//TodaysOffenceWidgetFilter : ""
type TodaysOffenceWidgetFilter struct {
	Date *time.Time `json:"date" bson:"date,omitempty"`
}

//TopOffencesWidget : ""
type TopOffencesWidget struct {
	OffenceType `bson:",inline"`
	// Name     string `json:"name" bson:"name,omitempty"`
	// UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Challans struct {
		Total   int64   `json:"total" bson:"total,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
	} `json:"challans" bson:"challans,omitempty"`
}

//TopOffencesWidgetFilter : ""
type TopOffencesWidgetFilter struct {
	Status      []string `json:"status" bson:"status,omitempty"`
	VehicleType []string `json:"vehicleType" bson:"vehicleType,omitempty"`
	OffenceID   []string `json:"offenceId" bson:"offenceId,omitempty"`
	SortBy      string   `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder   int      `json:"sortOrder" bson:"sortOrder,omitempty"`
}
