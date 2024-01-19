package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SmsLog : ""
type SmsLog struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	SentFor  string             `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	IsJob    bool               `json:"isJob"  bson:"isJob,omitempty"`
	Message  string             `json:"message,omitempty"  bson:"message,omitempty"`
	Status   string             `json:"status,omitempty"  bson:"status,omitempty"`
	SentDate *time.Time         `json:"sentDate,omitempty"  bson:"sentDate,omitempty"`
	To       To                 `json:"to,omitempty"  bson:"to,omitempty"`
	Created  *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type RefSmsLog struct {
	SmsLog `bson:",inline"`
	Ref    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type SmsLogFilter struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	IsJob  []bool   `json:"isJob,omitempty"  bson:"isJob,omitempty"`
	No     []string `json:"no,omitempty" bson:"no,omitempty"`
	Regex  struct {
		SentFor string `json:"sentFor"  bson:"sentFor,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type To struct {
	No       string `json:"no,omitempty" bson:"no,omitempty"`
	Name     string `json:"name"  bson:"name,omitempty"`
	UserName string `json:"userName" bson:"userName,omitempty"`
	UserType string `json:"userType" bson:"userType,omitempty"`
	Gender   string `json:"gender" bson:"gender,omitempty"`
}
