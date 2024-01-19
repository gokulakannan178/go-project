package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserLoginLog : ""
type UserLoginLog struct {
	ID         primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UserType   string             `json:"userType,omitempty"  bson:"userType,omitempty"`
	UserId     primitive.ObjectID `json:"userId,omitempty"  bson:"userId,omitempty"`
	FarmerId   primitive.ObjectID `json:"farmerId,omitempty"  bson:"farmerId,omitempty"`
	LoginTime  *time.Time         `json:"loginTime" form:"loginTime" bson:"loginTime,omitempty"`
	LogOutTime *time.Time         `json:"logOutTime" form:"logOutTime" bson:"logOutTime,omitempty"`
	Status     string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created    *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type UserLoginLogFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefUserLoginLog struct {
	UserLoginLog `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
