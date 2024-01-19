package models

import "time"

type WasteCollected struct {
	UniqueID       string        `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string        `json:"name" bson:"name,omitempty"`
	Description    string        `json:"description,omitempty" bson:"description,omitempty"`
	Citizen        *Citizen      `json:"citizen,omitempty" bson:"citizen,omitempty"`
	DumbSite       *DumpSiteName `json:"dumbSite,omitempty" bson:"dumbSite,omitempty"`
	MinUser        *MinUser      `json:"minUser,omitempty" bson:"minUser,omitempty"`
	GCUser         *GCUser       `json:"gcUser,omitempty" bson:"gcUser,omitempty"`
	VisitedHome    *VisitedHome  `json:"visitedHome,omitempty" bson:"visitedHome,omitempty"`
	OrganisationID string        `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time    `json:"date" bson:"date,omitempty"`
	EndDate        *time.Time    `json:"endDate" bson:"endDate,omitempty"`
	Created        Created       `json:"created" bson:"created,omitempty"`
	Status         string        `json:"status,omitempty" bson:"status,omitempty"`
	Remark         string        `json:"remark,omitempty" bson:"remark,omitempty"`
}

type DumpSiteName struct {
	Name     string   `json:"name" bson:"name,omitempty"`
	ID       string   `json:"id" form:"id," bson:"id,omitempty"`
	Location Location `json:"location" bson:"location,omitempty"`
}

type VisitedHome struct {
	Name     string   `json:"name" bson:"name,omitempty"`
	ID       string   `json:"id" form:"id," bson:"id,omitempty"`
	Location Location `json:"location" bson:"location,omitempty"`
}

type RefWasteCollected struct {
	WasteCollected `bson:",inline"`
	Ref            struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterWasteCollected struct {
	Status     []string `json:"status" bson:"status,omitempty"`
	Name       string   `json:"name" bson:"name,omitempty"`
	UniqueID   []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	ManagerID  []string `json:"managerId" bson:"managerId,omitempty"`
	DumbsiteID []string `json:"dumbsiteId" bson:"dumbsiteId,omitempty"`
	GCID       []string `json:"gcId" bson:"gcId,omitempty"`

	Regex struct {
		Name        string       `json:"name" bson:"name"`
		ManagerName MinUser      `json:"managerName" bson:"managerName"`
		GCName      GCUser       `json:"gcName" bson:"gcName"`
		CitizenName Citizen      `json:"citizenName" bson:"citizenName"`
		DumbSite    DumpSiteName `json:"dumbSite,omitempty" bson:"dumbSite,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
