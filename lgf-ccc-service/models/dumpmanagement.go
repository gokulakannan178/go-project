package models

import "time"

type DumpSite struct {
	UniqueID       string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string     `json:"name" bson:"name,omitempty"`
	Mobile         string     `json:"mobile" bson:"mobile,omitempty"`
	Email          string     `json:"email" bson:"email,omitempty"`
	OwnerName      string     `json:"ownerName" bson:"ownerName,omitempty"`
	License        string     `json:"license" bson:"license,omitempty"`
	LicenseUpload  string     `json:"licenseUpload" bson:"licenseUpload,omitempty"`
	Description    string     `json:"description,omitempty" bson:"description,omitempty"`
	Location       Location   `json:"location,omitempty" bson:"location,omitempty"`
	MinUser        *MinUser   `json:"minUser,omitempty" bson:"minUser,omitempty"`
	OrganisationID string     `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time `json:"date" bson:"date,omitempty"`
	EndDate        *time.Time `json:"endDate" bson:"endDate,omitempty"`
	Created        Created    `json:"created" bson:"created,omitempty"`
	Status         string     `json:"status,omitempty" bson:"status,omitempty"`
	Remark         string     `json:"remark,omitempty" bson:"remark,omitempty"`
	Address        Address    `json:"address" bson:"address,omitempty"`
}

type RefDumpSite struct {
	DumpSite `bson:",inline"`
	Ref      struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
		Village  *District `json:"village,omitempty" bson:"village,omitempty"`
		Zone     *Zone     `json:"zone,omitempty" bson:"zone,omitempty"`
		Ward     *Ward     `json:"ward,omitempty" bson:"ward,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterDumpSite struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	Name     []string `json:"name" bson:"name,omitempty"`
	UniqueID []string `json:"uniqueID" bson:"uniqueID,omitempty"`

	Regex struct {
		Name        string `json:"name" bson:"name"`
		OwnerName   string `json:"ownerName" bson:"ownerName,omitempty"`
		ManagerName string `json:"managerName"  bson:"managerName,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
