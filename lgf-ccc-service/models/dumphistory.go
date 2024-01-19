package models

import "time"

type DumpHistory struct {
	UniqueID       string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string     `json:"name" bson:"name,omitempty"`
	Description    string     `json:"description,omitempty" bson:"description,omitempty"`
	MinUser        *MinUser   `json:"minUser" bson:"minUser,omitempty"`
	DumbUser       *GCUser    `json:"dumbUser" bson:"dumbUser,omitempty"`
	Driver         Driver     `json:"driver"  bson:"driver,omitempty"`
	VehicleID      MinUser    `json:"vehicleId"  bson:"vehicleId,omitempty"`
	OrganisationID string     `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time `json:"date" bson:"date,omitempty"`
	EndDate        *time.Time `json:"endDate" bson:"endDate,omitempty"`
	Time           *time.Time `json:"time" bson:"time,omitempty"`
	Quantity       float64    `json:"quantity" bson:"quantity,omitempty"`
	Created        Created    `json:"created" bson:"created,omitempty"`
	Status         string     `json:"status,omitempty" bson:"status,omitempty"`
	Remark         string     `json:"remark,omitempty" bson:"remark,omitempty"`
}

type RefDumpHistory struct {
	DumpHistory `bson:",inline"`
	Ref         struct {
		DumbSite DumpSite `json:"dumbsite"  bson:"dumbsite,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDumpHistory struct {
	Status    []string   `json:"status" bson:"status,omitempty"`
	Name      []string   `json:"name" bson:"name,omitempty"`
	UniqueID  []string   `json:"uniqueID" bson:"uniqueID,omitempty"`
	ManagerId []string   `json:"managerId" bson:"managerId,omitempty"`
	DumbUser  []string   `json:"dumbUser" bson:"dumbUser,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`

	Regex struct {
		Name        string `json:"name" bson:"name"`
		DriverName  string `json:"driverName"  bson:"driverName,omitempty"`
		VehicleName string `json:"vehicleName"  bson:"vehicleName,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
type GetQuantity struct {
	Quantity int64 `json:"quantity" bson:"quantity,omitempty"`
}
type DayWiseDumpHistory struct {
	Date *time.Time `json:"date,omitempty"  bson:"date,omitempty"`
}
