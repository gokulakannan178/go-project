package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Name       string             `json:"name" bson:"name,omitempty"`
	AssignedBy string             `json:"assignedBy" bson:"assignedBy,omitempty"`
	AssignedTo string             `json:"assignedTo" bson:"assignedTo,omitempty"`
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StartDate  *time.Time         `json:"startDate" bson:"startDate,omitempty"`
	EndDate    *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	TeamMember []string           `json:"teamMember"  bson:"_"`
	TeamLeader string             `json:"teamLeader"  bson:"teamLeader",omitempty"`
	UniqueID   string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created" bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	ProjectID  string             `json:"projectId,omitempty"  bson:"projectId,omitempty"`
}

type RefTask struct {
	Task `bson:",inline"`
	Ref  struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterTask struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	ProjectId []string `json:"projectId,omitempty" bson:"projectId,omitempty"`
}

type TaskMember struct {
	Status   string  `json:"status" bson:"status,omitempty"`
	UniqueID string  `json:"uniqueId"  bson:"uniqueId,omitempty"`
	TaskID   string  `json:"taskId" bson:"taskId,omitempty"`
	UserName string  `json:"userName" bson:"userName,omitempty"`
	Created  Created `json:"created"  bson:"created,omitempty"`
}
