package models

import "time"

type TaskMessage struct {
	From        string    `json:"from" bson:"from,omitempty"` // write filter
	Desc        string    `json:"desc" bson:"desc,omitempty"`
	Message     string    `json:"message" bson:"message,omitempty"`
	MessageType string    `json:"messageType" bson:"messageType,omitempty"`
	UniqueID    string    `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status      string    `json:"status" bson:"status,omitempty"`
	File        []string  `json:"fileType" bson:"fileType,omitempty"`
	Mention     []string  `json:"mention" bson:"mention,omitempty"`
	On          time.Time `json:"on" bson:"on,omitempty"`
	TaskID      string    `json:"taskId" bson:"taskId,omitempty"`
	ProjectID   string    `json:"projectId" bson:"projectId,omitempty"`
}

type RefTaskMessage struct {
	TaskMessage `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterTaskMessage struct {
	Status    []string  `json:"status" bson:"status,omitempty"`
	TaskID    []string  `json:"taskId" bson:"taskId,omitempty"`
	ProjectID []string  `json:"projectId" bson:"projectId,omitempty"`
	On        time.Time `json:"on" bson:"on,omitempty"`
}
