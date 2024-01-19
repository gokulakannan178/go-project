package models

import "time"

type UserLocationTracker struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName  string     `json:"userName" bson:"userName,omitempty"`
	UserType  string     `json:"userType" bson:"userType,omitempty"`
	Location  Location   `json:"location" bson:"location,omitempty"`
	TimeStamp *time.Time `json:"timeStamp" bson:"timeStamp,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	Created   *CreatedV2 `json:"created" bson:"created,omitempty"`
}
type UserLocationTrackerCoordinates struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	StartDate *time.Time `json:"startDate" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate" bson:"endDate,omitempty"`
}
type UserLocationTrackerFilter struct {
	UniqueID  []string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status    []string  `json:"status" bson:"status,omitempty"`
	UserName  []string  `json:"userName" bson:"userName,omitempty"`
	UserType  []string  `json:"userType" bson:"userType,omitempty"`
	Date      DateRange `json:"date" bson:"date,omitempty"`
	SortBy    string    `json:"sortBy"`
	SortOrder int       `json:"sortOrder"`
}
type RefUserLocationTracker struct {
	UserLocationTracker `bson:",inline"`
	Ref                 struct {
		User     User     `json:"user" bson:"user,omitempty"`
		UserType UserType `json:"userType" bson:"userType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
