package models

import "time"

// CronLog : ""
type CronLog struct {
	UniqueID        string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name            string     `json:"name" bson:"name,omitempty"`
	DateStr         string     `json:"dateStr" bson:"dateStr,omitempty"`
	Date            *time.Time `json:"date" bson:"date,omitempty"`
	StartTime       *time.Time `json:"startTime" bson:"startTime,omitempty"`
	EndTime         *time.Time `json:"endTime" bson:"endTime,omitempty"`
	Status          string     `json:"status" bson:"status,omitempty"`
	IsCurrentScript bool       `json:"isCurrentScript" bson:"isCurrentScript,omitempty"`
	ErrorMessage    string     `json:"errorMessage" bson:"errorMessage,omitempty"`
}

// CronLogFilter : ""
type CronLogFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	Name      []string `json:"name" bson:"name,omitempty"`
	UniqueID  []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	SearchBox struct {
		Name     string `json:"name" bson:"name,omitempty"`
		DateStr  string `json:"dateStr" bson:"dateStr,omitempty"`
		UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	} `json:"searchBox,omitempty" bson:"searchBox,omitempty"`
	StartTime *DateRange `json:"startTime" bson:"startTime,omitempty"`
	EndTime   *DateRange `json:"endTime" bson:"endTime,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

type RefCronLog struct {
	CronLog `bson:",inline"`
}
