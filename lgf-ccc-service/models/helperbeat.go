package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HelperBeat struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	UniqueID     string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	VehicleId    string             `json:"vehicleId,omitempty" bson:"vehicleId,omitempty"`
	BeatID       string             `json:"beatId,omitempty" bson:"beatId,omitempty"`
	EmployeeUser []string           `json:"employeeUser" bson:"_"`
	Manager      *MinUser           `json:"manager,omitempty" bson:"manager,omitempty"`
	Employee     MinUser            `json:"employee,omitempty" bson:"employee,omitempty"`
	Desc         string             `json:"description,omitempty" bson:"description,omitempty"`
	StartDate    *time.Time         `json:"startDate,omitempty" bson:"startDate,omitempty"`
	AssignDate   *time.Time         `json:"assignDate,omitempty" bson:"assignDate,omitempty"`
	Logo         string             `json:"logo,omitempty" bson:"logo,omitempty"`
	IsChecked    string             `json:"isChecked" bson:"isChecked,omitempty"`
	Created      *Created           `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefHelperBeat struct {
	HelperBeat `bson:",inline"`
	Ref        struct {
		//OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
		User       User       `json:"user,omitempty" bson:"user,omitempty"`
		BeatMaster BeatMaster `json:"beatmaster,omitempty" bson:"beatmaster,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterHelperBeat struct {
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	Managerid  []string `json:"managerId,omitempty" bson:"managerId,omitempty"`
	EmployeeId []string `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	TicketId   []string `json:"ticketId,omitempty" bson:"ticketId,omitempty"`
	ProjectId  []string `json:"projectId,omitempty" bson:"projectId,omitempty"`
	BeatID     []string `json:"beatId,omitempty" bson:"beatId,omitempty"`
	IsStatus   []string `json:"isStatus,omitempty" bson:"isStatus,omitempty"`

	Regex struct {
		Name     string `json:"name,omitempty" bson:"name,omitempty"`
		TicketId string `json:"ticketId,omitempty" bson:"ticketId,omitempty"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
