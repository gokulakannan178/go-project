package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Asset : ""
type TelegramLog struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	SentFor  string             `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	IsJob    bool               `json:"isJob"  bson:"isJob,omitempty"`
	Message  string             `json:"message,omitempty"  bson:"message,omitempty"`
	Status   string             `json:"status,omitempty"  bson:"status,omitempty"`
	SentDate *time.Time         `json:"sentDate,omitempty"  bson:"sentDate,omitempty"`
	To       []ToTelegramLog    `json:"to,omitempty"  bson:"to,omitempty"`
	Created  *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type TelegramLogFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	IsJob     []bool   `json:"isJob"  bson:"isJob,omitempty"`
	No        []string `json:"no,omitempty" bson:"no,omitempty"`
	Name      []string `json:"name"  bson:"name,omitempty"`
	UserName  []string `json:"userName" bson:"userName,omitempty"`
	UserType  []string `json:"userType" bson:"userType,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Regex     struct {
		SentFor string `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	} `json:"regex" bson:"regex"`
}

type RefTelegramLog struct {
	TelegramLog `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type ToTelegramLog struct {
	No       string `json:"no,omitempty" bson:"no,omitempty"`
	Name     string `json:"name"  bson:"name,omitempty"`
	UserName string `json:"userName" bson:"userName,omitempty"`
	UserType string `json:"userType" bson:"userType,omitempty"`
}
