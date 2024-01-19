package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Asset : ""
type NotificationLog struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueID,omitempty"  bson:"uniqueID,omitempty"`
	SentFor  string             `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	IsJob    bool               `json:"isJob"  bson:"isJob,omitempty"`
	Message  string             `json:"message,omitempty"  bson:"message,omitempty"`
	Tittle   string             `json:"tittle,omitempty"  bson:"tittle,omitempty"`
	Body     string             `json:"body,omitempty"  bson:"body,omitempty"`
	Image    string             `json:"image,omitempty"  bson:"image,omitempty"`
	Topic    string             `json:"topic,omitempty"  bson:"topic,omitempty"`
	Data     map[string]string  `json:"data,omitempty"  bson:"data,omitempty"`
	Status   string             `json:"status,omitempty"  bson:"status,omitempty"`
	SentDate *time.Time         `json:"sentDate,omitempty"  bson:"sentDate,omitempty"`
	To       ToNotificationLog  `json:"to,omitempty"  bson:"to,omitempty"`
	Created  *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type NotificationLogFilter struct {
	Status               []string             `json:"status,omitempty" bson:"status,omitempty"`
	IsJob                []bool               `json:"isJob"  bson:"isJob,omitempty"`
	No                   []string             `json:"no,omitempty" bson:"no,omitempty"`
	Name                 []string             `json:"name"  bson:"name,omitempty"`
	UserName             []primitive.ObjectID `json:"userName" bson:"userName,omitempty"`
	UserType             []string             `json:"userType" bson:"userType,omitempty"`
	AppRegistrationToken []string             `json:"appRegistrationToken" bson:"appRegistrationToken,omitempty"`
	SortBy               string               `json:"sortBy"`
	SortOrder            int                  `json:"sortOrder"`
	Regex                struct {
		SentFor string `json:"sentFor,omitempty"  bson:"sentFor,omitempty"`
	} `json:"regex" bson:"regex"`
}

type RefNotificationLog struct {
	NotificationLog `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type ToNotificationLog struct {
	No                   string `json:"no,omitempty" bson:"no,omitempty"`
	Name                 string `json:"name"  bson:"name,omitempty"`
	UserName             string `json:"userName" bson:"userName,omitempty"`
	AppRegistrationToken string `json:"appRegistrationToken" bson:"appRegistrationToken,omitempty"`
	UserType             string `json:"userType" bson:"userType,omitempty"`
	UserId               string `json:"userId" bson:"userId,omitempty"`
	Gender               string `json:"gender" bson:"gender,omitempty"`
	State                string `json:"state"  bson:"state,omitempty"`
	District             string `json:"district"  bson:"district,omitempty"`
	Block                string `json:"block"  bson:"block,omitempty"`
	GramPanchayat        string `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village              string `json:"village"  bson:"village,omitempty"`
	StateCode            string `json:"stateCode"  bson:"stateCode,omitempty"`
	DistricCode          string `json:"districtCode"  bson:"districtCode,omitempty"`
	BlockCode            string `json:"blockCode"  bson:"blockCode,omitempty"`
	GramPanchayatCode    string `json:"gramPanchayatCode"  bson:"gramPanchayatCode,omitempty"`
	VillageCode          string `json:"villageCode"  bson:"villageCode,omitempty"`
}
