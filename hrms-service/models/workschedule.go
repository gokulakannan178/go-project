package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//WorkSchedule : ""
type WorkSchedule struct {
	ID                 primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	OrganisationID     string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Name               string             `json:"name,omitempty" bson:"name,omitempty"`
	Schedule           string             `json:"schedule,omitempty" bson:"schedule,omitempty"`
	ScheduleType       string             `json:"scheduleType,omitempty" bson:"scheduleType,omitempty"`
	Monday             bool               `json:"monday" bson:"monday,omitempty"`
	Tuesday            bool               `json:"tuesday" bson:"tuesday,omitempty"`
	Wednesday          bool               `json:"wednesday" bson:"wednesday,omitempty"`
	Thursday           bool               `json:"thursday" bson:"thursday,omitempty"`
	Friday             bool               `json:"friday" bson:"friday,omitempty"`
	Saturday           bool               `json:"saturday" bson:"saturday,omitempty"`
	Sunday             bool               `json:"sunday" bson:"sunday,omitempty"`
	WorkingHoursinDay  float64            `json:"workingHoursinDay,omitempty" bson:"workingHoursinDay,omitempty"`
	WorkingDaysinWeek  float64            `json:"workingDaysinWeek,omitempty" bson:"workingDaysinWeek,omitempty"`
	WorkingHoursinWeek float64            `json:"workingHoursinWeek,omitempty" bson:"workingHoursinWeek,omitempty"`
	LossOfPay          float64            `json:"lossOfPay,omitempty" bson:"lossOfPay,omitempty"`
	PartialPay         float64            `json:"partialPay,omitempty" bson:"partialPay,omitempty"`
	FullPay            float64            `json:"fullPay,omitempty" bson:"fullPay,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            Created            `json:"createdOn" bson:"createdOn,omitempty"`
}

//RefWorkSchedule :""
type RefWorkSchedule struct {
	WorkSchedule `bson:",inline"`
	Ref          struct {
		OrganisationID Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//WorkScheduleFilter : ""
type WorkScheduleFilter struct {
	Status         []string `json:"status" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	UniqueID       []string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
