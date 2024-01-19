package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CitizenGrevience struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	MinUser        *MinUser           `json:"minUser,omitempty" bson:"minUser,omitempty"`
	Citizen        *Citizen           `json:"citizen,omitempty" bson:"citizen,omitempty"`
	GCUser         *GCUser            `json:"gcUser,omitempty" bson:"gcUser,omitempty"`
	NfcID          string             `json:"nfcId,omitempty" bson:"nfcId,omitempty"`
	Mobile         string             `json:"mobile" bson:"mobile,omitempty"`
	WardNo         string             `json:"wardNo,omitempty" bson:"wardNo,omitempty"`
	HoldingNumber  string             `json:"holdingNumber" bson:"holdingNumber,omitempty"`
	Location       Location           `json:"location,omitempty" bson:"location,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	OrganisationID string             `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	Date           *time.Time         `json:"date" bson:"date,omitempty"`
	EndDate        *time.Time         `json:"endDate" bson:"endDate,omitempty"`
	Created        *CreatedV2         `json:"created" bson:"created,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
}

type RefCitizenGrevience struct {
	CitizenGrevience `bson:",inline"`
	Ref              struct {
		OrganisationID Organisation `json:"organisationId,omitempty" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterCitizenGrevience struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	Name      string   `json:"name" bson:"name,omitempty"`
	UniqueID  []string `json:"uniqueID" bson:"uniqueID,omitempty"`
	GCID      []string `json:"gcId" bson:"gcId,omitempty"`
	ManagerID []string `json:"managerId" bson:"managerId,omitempty"`

	Regex struct {
		Name   string `json:"name" bson:"name"`
		Email  string `json:"email" bson:"email"`
		Mobile string `json:"mobile" bson:"mobile"`
	} `json:"regex" bson:"regex"`
	DateRange struct {
		From *time.Time `json:"from,omitempty"  bson:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"  bson:"to,omitempty"`
	} `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
