package models

import "time"

type DriverDetails struct {
	//ID          string   `json:"id" form:"id," bson:"id,omitempty"`
	UniqueID      string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	UserName      string     `json:"userName,omitempty" bson:"userName,omitempty"`
	Name          string     `json:"name" bson:"name,omitempty"`
	Mobile        string     `json:"mobile" bson:"mobile,omitempty"`
	Email         string     `json:"email" bson:"email,omitempty"`
	Gender        string     `json:"gender" bson:"gender,omitempty"`
	DOB           *time.Time `json:"dob" bson:"dob,omitempty"`
	DateofJoining *time.Time `json:"dateOfJoining" bson:"dateOfJoining,omitempty"`
	LicenseNo     string     `json:"licenseNo,omitempty" bson:"licenseNo,omitempty"`
	LicenseType   string     `json:"licenseType" bson:"licenseType,omitempty"`
	LicenseExpiry *time.Time `json:"licenseExpiry" bson:"licenseExpiry,omitempty"`
	AadharNo      string     `json:"AadharNo" bson:"AadharNo,omitempty"`
	Language      string     `json:"language" bson:"language,omitempty"`
	ProfileImg    string     `json:"profileImg" bson:"profileImg,omitempty"`
	Description   string     `json:"description,omitempty" bson:"description,omitempty"`
	MinUser       *MinUser   `json:"minUser" bson:"minUser,omitempty"`
	//Vechile        Vechile    `json:"vechile,omitempty" bson:"vechile,omitempty"`
	OrganisationID string     `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time `json:"date" bson:"date,omitempty"`
	EndDate        *time.Time `json:"endDate" bson:"endDate,omitempty"`
	Created        CreatedV2  `json:"created" bson:"created,omitempty"`
	Status         string     `json:"status,omitempty" bson:"status,omitempty"`
	Remark         string     `json:"remark,omitempty" bson:"remark,omitempty"`
}

type RefDriverDetails struct {
	DriverDetails `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDriverDetails struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	Name      string   `json:"name" bson:"name,omitempty"`
	UniqueID  []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	IsExpDate string   `json:"isExpdate" bson:"isExpdate,omitempty"`
	Regex     struct {
		Name        string `json:"name" bson:"name"`
		LicenseType string `json:"licenseType" bson:"licenseType,omitempty"`
		Mobile      string `json:"mobile" bson:"mobile"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
