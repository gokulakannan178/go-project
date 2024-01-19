package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ContentViewLog : ""
type ContentViewLog struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	UniqueId     string             `json:"uniqueId,omitempty"  bson:"uniqueId,omitempty"`
	ContentId    primitive.ObjectID `json:"contentId,omitempty"  bson:"contentId,omitempty"`
	Count        int64              `json:"count,omitempty"  bson:"count,omitempty"`
	Date         *time.Time         `json:"date" form:"dateReviewed" bson:"date,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type ContentViewLogFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefContentViewLog struct {
	ContentViewLog `bson:",inline"`
	Ref            struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type FilterDaywiseViewChart struct {
	ContentId   primitive.ObjectID `json:"contentId,omitempty"  bson:"contentId,omitempty"`
	CreatedFrom struct {
		StartDate *time.Time `json:"startDate"`
		EndDate   *time.Time `json:"endDate"`
	} `json:"createdFrom"`
}
type DayWiseContentViewChartReport struct {
	Days []struct {
		Date  *time.Time `json:"date" form:"dateReviewed" bson:"date,omitempty"`
		Count int64      `json:"count"  bson:"count,omitempty"`
	} `json:"days,omitempty" bson:"days"`
}
