package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//JobTimeline : ""
type JobTimeline struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID  string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	State     string             `json:"state,omitempty" bson:"state,omitempty"`
	StartDate *time.Time         `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   *time.Time         `json:"endDate,omitempty" bson:"endDate,omitempty"`
	Assigned  struct {
		UserID   string    `json:"userId,omitempty" bson:"userId,omitempty"`
		UserType string    `json:"userType,omitempty" bson:"userType,omitempty"`
		Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
	} `json:"assigned,omitempty" bson:"assigned,omitempty"`
	Removed struct {
		UserID   string    `json:"userId,omitempty" bson:"userId,omitempty"`
		UserType string    `json:"userType,omitempty" bson:"userType,omitempty"`
		Date     time.Time `json:"date,omitempty" bson:"date,omitempty"`
	} `json:"removed,omitempty" bson:"removed,omitempty"`
	LineManager    string   `json:"lineManager,omitempty" bson:"lineManager,omitempty"`
	JobTitle       string   `json:"jobTitle,omitempty" bson:"jobTitle,omitempty"`
	PositionType   string   `json:"positionType,omitempty" bson:"positionType,omitempty"`
	EmployeeId     string   `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	OrganisationId string   `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	DepartmentId   string   `json:"departmentId,omitempty" bson:"departmentId,omitempty"`
	BranchId       string   `json:"branchId,omitempty" bson:"branchId,omitempty"`
	DesignationId  string   `json:"designationId,omitempty" bson:"designationId,omitempty"`
	Office         string   `json:"office,omitempty" bson:"office,omitempty"`
	Remark         string   `json:"remark,omitempty" bson:"remark,omitempty"`
	Status         string   `json:"status,omitempty" bson:"status,omitempty"`
	Created        *Created `json:"createdOn" bson:"createdOn,omitempty"`
	Updated        Updated  `json:"updated" form:"id," bson:"updated,omitempty"`
}

//RefJobTimeline :""
type RefJobTimeline struct {
	JobTimeline `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//JobTimelineFilter : ""
type JobTimelineFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	UniqueID  []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Regex     struct {
		LineManager string `json:"lineManager" bson:"lineManager,omitempty"`
	} `json:"regex" bson:"regex"`
}
