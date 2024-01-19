package models

import (
	"time"
)

//VehicleChallan : ""
type VehicleChallan struct {
	Vehicle       Vehicle        `json:"vehicle" bson:"vehicle,omitempty"`
	UniqueID      string         `json:"uniqueId" bson:"uniqueId,omitempty"`
	OffenceDate   *time.Time     `json:"offenceDate" bson:"offenceDate,omitempty"`
	OffenceAt     string         `json:"offenceAt" bson:"offenceAt,omitempty"`
	OffenceDetail string         `json:"offenceDetail" bson:"offenceDetail,omitempty"`
	OffenceType   string         `json:"offenceType" bson:"offenceType,omitempty"`
	Images        []string       `json:"images" bson:"images,omitempty"`
	Videos        []string       `json:"videos" bson:"videos,omitempty"`
	Pelalty       float64        `json:"pelalty" bson:"pelalty,omitempty"`
	Status        string         `json:"status" bson:"status,omitempty"`
	Payment       VehiclePayment `json:"payment" bson:"payment,omitempty"`
	Created       Created        `json:"createdOn" bson:"createdOn,omitempty"`
	Updated       Updated        `json:"updated"  bson:"updated,omitempty"`
	UpdateLog     []Updated      `json:"updatedLog" bson:"updatedLog,omitempty"`
}

//RefVehicleChallan :""
type RefVehicleChallan struct {
	VehicleChallan `bson:",inline"`
	Ref            struct {
		OffenceType *OffenceType `json:"offenceType"  bson:"offenceType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//VehicleChallanFilter : ""
type VehicleChallanFilter struct {
	Status         []string `json:"status"`
	OffenceType    []string `json:"offenceType"`
	VehicleType    []string `json:"vehicleType"`
	IsOffenceVideo bool     `json:"isOffenceVideo"`
	Mobile         []string `json:"mobile"`
	TN04AS9101     []string `json:"TN04AS9101"`
	RegNo          []string `json:"regNo"`
	OmitID         []string `json:"omitId"`
	SortBy         string   `json:"sortBy"`
	SortOrder      int      `json:"sortOrder"`
}
