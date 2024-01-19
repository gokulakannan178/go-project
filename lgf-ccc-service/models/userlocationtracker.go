package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLocationTracker struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	OrganisationId string             `json:"organisationId" bson:"organisationId,omitempty"`
	UserName       string             `json:"userName" bson:"userName,omitempty"`
	UserType       string             `json:"userType" bson:"userType,omitempty"`
	Location       Location           `json:"location" bson:"location,omitempty"`
	TimeStamp      *time.Time         `json:"timeStamp" bson:"timeStamp,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Created        *Created           `json:"created" bson:"created,omitempty"`
}

type RefUserLocationTracker struct {
	UserLocationTracker `bson:",inline"`
	Ref                 struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterUserLocationTracker struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	UniqueID       []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName       []string `json:"userName" bson:"userName,omitempty"`
	UserType       []string `json:"userType" bson:"userType,omitempty"`

	Regex struct {
		UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// type Location struct {
// 	Type        string    `json:"type" bson:"type,omitempty"`
// 	Coordinates []float64 `json:"coordinates" bson:"coordinates,omitempty"`
// }
