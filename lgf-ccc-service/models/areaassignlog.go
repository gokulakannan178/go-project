package models

import (
	"time"
)

type AreaAssignLog struct {
	UniqueID    string     `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Area        Address    `json:"area"  bson:"area,omitempty"`
	User        User       `json:"user"  bson:"user,omitempty"`
	Description string     `json:"description" bson:"description,omitempty"`
	AssignDate  *time.Time `json:"assignDate" bson:"assignDate,omitempty"`
	StartDate   *time.Time `json:"startDate" bson:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate" bson:"endDate,omitempty"`
	Status      string     `json:"status" bson:"status,omitempty"`
	Created     Created    `json:"created"  bson:"created,omitempty"`
	Updated     Updated    `json:"updated"  bson:"updated,omitempty"`
}
type AreaAssignLogFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId"  bson:"uniqueId,omitempty"`
	UserId   []string `json:"userId"  bson:"userId,omitempty"`
	Regex    struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type RefAreaAssignLog struct {
	AreaAssignLog `bson:",inline"`
	Ref           struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
