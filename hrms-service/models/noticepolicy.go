package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NoticePolicy struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Desc           string             `json:"description,omitempty" bson:"description,omitempty"`
	EmployeeID     string             `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	NoticeDays     int                `json:"noticeDays,omitempty" bson:"noticeDays,omitempty"`
	OrganisationId string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        *Created           `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefNoticePolicy struct {
	NoticePolicy `bson:",inline"`
	Ref          struct {
		OrganisationId Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterNoticePolicy struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
