package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DemoUser : ""
type DemoUser struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Mobile         string             `json:"mobile" bson:"mobile,omitempty"`
	Email          string             `json:"email" bson:"email,omitempty"`
	DOB            *time.Time         `json:"dob" bson:"dob,omitempty"`
	OrganisationID string             `json:"organisationId" bson:"organisationId,omitempty"`
	Password       string             `json:"-" bson:"password,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Gender         string             `json:"gender" bson:"gender,omitempty"`
	Salary         bool               `json:"salary" bson:"salary,omitempty"`
	Hobbies        []string           `json:"hobbies" bson:"hobbies,omitempty"`
}

//RefDemoUser :""
type RefDemoUser struct {
	DemoUser `bson:",inline"`
	Ref      struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//DemoUserFilter : ""
type DemoUserFilter struct {
	Status            []string `json:"status"`
	UniqueID          []string `json:"uniqueId"`
	OmitID            []string `json:"omitId"`
	OrganisationID    []string `json:"organisationId" bson:"organisationId,omitempty"`
	Manager           []string `json:"manager" bson:"manager,omitempty"`
	Type              []string `json:"type" bson:"type"`
	GetRecentLocation bool     `json:"getRecentLocation"`
	Regex             struct {
		Name         string `json:"name" bson:"name"`
		Contact      string `json:"contact" bson:"contact"`
		DemoUserName string `json:"demouserName" bson:"demouserName"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
