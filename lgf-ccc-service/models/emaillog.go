package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//EmailLog : ""
type EmailLog struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	SentFor  string             `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	IsJob    bool               `json:"isJob"  bson:"isJob,omitempty"`
	Message  string             `json:"message,omitempty"  bson:"message,omitempty"`
	Status   string             `json:"status,omitempty"  bson:"status,omitempty"`
	SentDate *time.Time         `json:"sentDate,omitempty"  bson:"sentDate,omitempty"`
	To       ToEmailLog         `json:"to,omitempty"  bson:"to,omitempty"`
	Created  *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type EmailLogFilter struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	IsJob  []bool   `json:"isJob,omitempty"  bson:"isJob,omitempty"`
	Email  []string `json:"email,omitempty" bson:"email,omitempty"`
	Regex  struct {
		SentFor string `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type RefEmailLog struct {
	EmailLog `bson:",inline"`
	Ref      struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type ToEmailLog struct {
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Name  string `json:"name"  bson:"name,omitempty"`
}
