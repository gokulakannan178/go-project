package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID    string             `json:"uniqueId"  bson:"uniqueId,omitempty"`
	ProjectName string             `json:"projectName"  bson:"projectName,omitempty"`
	StartDate   *time.Time         `json:"startDate" bson:"startDate,omitempty"`
	EndDate     *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	Rate        float64            `json:"rate"  bson:"rate,omitempty"`
	Created     Created            `json:"created"  bson:"created,omitempty"`
	Updated     Updated            `json:"updated"  bson:"updated,omitempty"`
	Priority    string             `json:"priority"  bson:"priority,omitempty"`
	LeaderID    string             `json:"leaderId"  bson:"leaderId,omitempty"`
	TeamMember  []string           `json:"teamMember"  bson:"-"`
	Status      string             `json:"status" bson:"status,omitempty"`
	ClientID    string             `json:"clientId"  bson:"clientId,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Files       []string           `json:"files"  bson:"files,omitempty"`
}
type ProjectFilter struct {
	Status    string `json:"status" bson:"status,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

//ProjectMember:""
type ProjectMember struct {
	Status    string  `json:"status" bson:"status,omitempty"`
	UniqueID  string  `json:"uniqueId"  bson:"uniqueId,omitempty"`
	ProjectID string  `json:"projectId" bson:"projectId,omitempty"`
	UserName  string  `json:"userName" bson:"userName,omitempty"`
	Created   Created `json:"created"  bson:"created,omitempty"`
}
