package models

import "time"

type Vechile struct {
	UniqueID       string       `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	MinUser        MinUser      `json:"minUser" bson:"minUser,omitempty"`
	Driver         Driver       `json:"driver" bson:"driver,omitempty"`
	DumpSite       DumpSiteName `json:"dumpSite" bson:"dumpSite,omitempty"`
	VechileName    string       `json:"vechileName" bson:"vechileName,omitempty"`
	VechileType    string       `json:"vechileType" bson:"vechileType,omitempty"`
	VechileImage   string       `json:"vechileImage" bson:"vechileImage,omitempty"`
	VechileTypeID  string       `json:"vechileTypeId" bson:"vechileTypeId,omitempty"`
	Location       Location     `json:"location" bson:"location,omitempty"`
	LicenseNo      string       `json:"licenseNo" bson:"licenseNo,omitempty"`
	Chassis        string       `json:"chassis" bson:"chassis,omitempty"`
	InsuranceValid *time.Time   `json:"insuranceValid" bson:"insuranceValid,omitempty"`
	PermitNo       string       `json:"permitNo" bson:"permitNo,omitempty"`
	RcNO           string       `json:"rcNO" bson:"rcNO,omitempty"`
	RcNOTill       *time.Time   `json:"rcNOTill" bson:"rcNOTill,omitempty"`
	PUCNo          string       `json:"pUCNo" bson:"pUCNo,omitempty"`
	PUCValidDate   *time.Time   `json:"pUCValidDate,omitempty" bson:"pUCValidDate,omitempty"`
	Date           *time.Time   `json:"date" bson:"date,omitempty"`
	EndDate        *time.Time   `json:"endDate" bson:"endDate,omitempty"`
	CreatedBy      CreatedBy    `json:"createdBy" bson:"createdBy,omitempty"`
	Created        Created      `json:"created" bson:"created,omitempty"`
	Status         string       `json:"status" bson:"status,omitempty"`
	Remark         string       `json:"remark" bson:"remark,omitempty"`
	EngineNumber   string       `json:"engineNumber" bson:"engineNumber,omitempty"`
	IsDriverAssign string       `json:"isDriverAssign" bson:"isDriverAssign,omitempty"`
}

type Driver struct {
	Name string `json:"name" bson:"name,omitempty"`
	ID   string `json:"id" form:"id," bson:"id,omitempty"`
}

type RefVechile struct {
	Vechile `bson:",inline"`
	Ref     struct {
		VechileTypeID VehicleType `json:"vechileTypeId,omitempty" bson:"vechileTypeId,omitempty"`
		Vehiclelog    VehicleLog  `json:"vehiclelog" bson:"vehiclelog,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterVechile struct {
	Status      []string `json:"status" bson:"status,omitempty"`
	Name        []string `json:"name" bson:"name,omitempty"`
	UniqueID    []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	VechileType []string `json:"vechileType" bson:"vechileType,omitempty"`
	Model       []string `json:"model,omitempty" bson:"model,omitempty"`
	LicenseNo   []string `json:"licenseNo,omitempty" bson:"licenseNo,omitempty"`
	ManagerID   []string `json:"managerId" bson:"managerId,omitempty"`
	DumbsiteID  []string `json:"dumbsiteId" bson:"dumbsiteId,omitempty"`
	IsExpDate   string   `json:"isExpdate" bson:"isExpdate,omitempty"`

	Regex struct {
		Name      string `json:"name" bson:"name"`
		LicenseNo string `json:"licenseNo" bson:"licenseNo,omitempty"`
		PUCNo     string `json:"pUCNo" bson:"pUCNo,omitempty"`
		RcNO      string `json:"rcNO" bson:"rcNO,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type VechileAssign struct {
	VehicleID     string `json:"vehicleId" bson:"vehicleId,omitempty"`
	VechileTypeID string `json:"vechileTypeId" bson:"vechileType,omitempty"`
	DriverID      string `json:"driverId,omitempty" bson:"driverId,omitempty"`
}

type VehicleLocationUpdate struct {
	VehicleID string   `json:"vehicleId" bson:"vehicleId,omitempty"`
	Location  Location `json:"location" bson:"location,omitempty"`
}

type AreaAssign struct {
	PMID string  `json:"pmId" bson:"pmId,omitempty"`
	Area Address `json:"area" bson:"area,omitempty"`
}
