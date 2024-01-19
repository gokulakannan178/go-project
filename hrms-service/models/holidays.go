package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Holidays struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name" bson:"name,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	Date           *time.Time         `json:"date,omitempty" bson:"date,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Created        Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefHolidays struct {
	Holidays `bson:",inline"`
	Ref      struct {
		OrganisationID Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type HolidaysList struct {
	UniqueID            string            `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name                string            `json:"name" bson:"name,omitempty"`
	Day                 string            `json:"day" bson:"day,omitempty"`
	Description         string            `json:"description,omitempty" bson:"description,omitempty"`
	Date                *time.Time        `json:"date,omitempty" bson:"date,omitempty"`
	Holidays            []Holidays        `json:"holidays" bson:"holidays,omitempty"`
	BrithdateEmployee   []WeekCalEmployee `json:"brithdateEmployee" bson:"brithdateEmployee,omitempty"`
	AnniversaryEmployee []WeekCalEmployee `json:"anniversaryEmployee" bson:"anniversaryEmployee,omitempty"`
}
type FilterHolidays struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationID []string `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           struct {
		StartDate *time.Time `json:"startDate,omitempty"  bson:"startDate,omitempty"`
		EndDate   *time.Time `json:"EndDate,omitempty"  bson:"EndDate,omitempty"`
	} `json:"date" bson:"date"`
	Regex struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
	SortBy     string            `json:"sortBy"`
	SortOrder  int               `json:"sortOrder"`
	DataAccess DataAccessRequest `json:"dataAccess" bson:"dataAccess,omitempty"`
}
