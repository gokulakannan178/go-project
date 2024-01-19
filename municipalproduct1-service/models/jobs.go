package models

// Job : ""
type Job struct {
	UniqueID   string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Title      string    `json:"title" bson:"title,omitempty"`
	WarningMsg string    `json:"warningMsg" bson:"warningMsg,omitempty"`
	Created    CreatedV2 `json:"created" bson:"created,omitempty"`
	Status     string    `json:"status" bson:"status,omitempty"`
}

// JobFilter : ""
type JobFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

type RefJob struct {
	Job `bson:",inline"`
}
