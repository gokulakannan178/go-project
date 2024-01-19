package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type News struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID         string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Title            string             `json:"title,omitempty" bson:"title,omitempty"`
	Message          string             `json:"message,omitempty" bson:"message,omitempty"`
	Date             *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	NewsLikeCount    int64              `json:"newslikeCount,omitempty" bson:"newslikeCount,omitempty"`
	NewsCommentCount int64              `json:"newscommentCount,omitempty" bson:"newscommentCount,omitempty"`
	OrganisationId   string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	SendTo           struct {
		DepartmentId []string `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
		Employee     []string `json:"employee,omitempty" bson:"employee,omitempty"`
	} `json:"sendTo,omitempty" bson:"sendTo,omitempty"`
	Created   *Created `json:"createdOn" bson:"createdOn,omitempty"`
	CreatedBy string   `json:"createdBy" bson:"createdBy,omitempty"`
	Updated   Updated  `json:"updated" form:"id," bson:"updated,omitempty"`
	Status    string   `json:"status,omitempty" bson:"status,omitempty"`
}

type RefNews struct {
	News `bson:",inline"`
	Ref  struct {
		OrganisationId Organisation  `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
		NewsLike       []NewsLike    `json:"newslike,omitempty" bson:"newslike,omitempty"`
		CreatedBy      Employee      `json:"createdBy" bson:"createdBy,omitempty"`
		NewsComment    []NewsComment `json:"newscomment,omitempty" bson:"newscomment,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterNews struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	DepartmentId   []string `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	Employee       []string `json:"employee,omitempty" bson:"employee,omitempty"`
	SortBy         string   `json:"sortBy"`
	SortOrder      int      `json:"sortOrder"`
	Regex          struct {
		Title string `json:"title,omitempty" bson:"title,omitempty"`
	} `json:"regex" bson:"regex"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
type PublishedNews struct {
	UniqueID       string   `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	OrganisationId []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	DepartmentId   []string `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
}
