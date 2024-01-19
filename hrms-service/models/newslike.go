package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsLike struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	EmployeeId     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	NewsId         string             `json:"newsId,omitempty" bson:"newsId,omitempty"`
	Date           *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefNewsLike struct {
	NewsLike `bson:",inline"`
	Ref      struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterNewsLike struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	EmployeeId     []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
	} `json:"regex" bson:"regex"`
}
