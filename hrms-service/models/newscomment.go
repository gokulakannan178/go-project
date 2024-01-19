package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsComment struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	NewsId         string             `json:"newsId,omitempty" bson:"newsId,omitempty"`
	Comment        string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Date           *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefNewsComment struct {
	NewsComment `bson:",inline"`
	Ref         struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterNewsComment struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Comment string `json:"comment,omitempty" bson:"comment,omitempty"`
	} `json:"regex" bson:"regex"`
}
