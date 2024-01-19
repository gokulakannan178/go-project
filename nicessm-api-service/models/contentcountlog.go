package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ContentCountLog : ""
type ContentCountLog struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	UserType     string             `json:"userType,omitempty"  bson:"userType,omitempty"`
	ContentId    primitive.ObjectID `json:"contentId,omitempty"  bson:"contentId,omitempty"`
	UserId       primitive.ObjectID `json:"userId,omitempty"  bson:"userId,omitempty"`
	FarmerId     primitive.ObjectID `json:"farmerId,omitempty"  bson:"farmerId,omitempty"`
	Count        int64              `json:"count,omitempty"  bson:"count,omitempty"`
	Date         *time.Time         `json:"date" form:"dateReviewed" bson:"date,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type ContentCountLogFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefContentCountLog struct {
	ContentCountLog `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
