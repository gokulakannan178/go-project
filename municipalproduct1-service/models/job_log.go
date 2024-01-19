package models

import "time"

// JobLog : ""
type JobLog struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	JobID     string     `json:"jobId" bson:"jobId,omitempty"`
	Title     string     `json:"title" bson:"title,omitempty"`
	Desc      string     `json:"desc" bson:"desc,omitempty"`
	ErrorMsg  string     `json:"errorMsg" bson:"errorMsg,omitempty"`
	StartTime *time.Time `json:"startTime" bson:"startTime,omitempty"`
	EndTime   *time.Time `json:"endTime" bson:"endTime,omitempty"`
	Created   CreatedV2  `json:"created" bson:"created,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
}

// JobLogFilter : ""
type JobLogFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

type RefJobLog struct {
	JobLog `bson:",inline"`
}
