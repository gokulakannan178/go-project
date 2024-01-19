package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Ticket : ""
type Ticket struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID    string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	DateFrom    *time.Time         `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo      *time.Time         `json:"dateTo" bson:"dateTo,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Created     *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
	Updated     []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	Documents   []string           `json:"documents,omitempty"  bson:"documents,omitempty"`
	AssignedBy  string             `json:"assignedBy,omitempty"  bson:"assignedBy,omitempty"`
}

//TicketFilter : ""
type TicketFilter struct {
	UniqueID         []string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name             []string   `json:"name" bson:"name,omitempty"`
	Description      []string   `json:"description" bson:"description,omitempty"`
	Status           []string   `json:"status" bson:"status,omitempty"`
	AssignedTo       []string   `json:"assignedTo,omitempty"  bson:"assignedTo,omitempty"`
	AssignedBy       []string   `json:"assignedBy,omitempty"  bson:"assignedBy,omitempty"`
	FromDateRange    *DateRange `json:"fromDateRange,omitempty"  bson:"fromDateRange,omitempty"`
	ToDateRange      *DateRange `json:"toDateRange,omitempty"  bson:"toDateRange,omitempty"`
	CreatedDateRange *DateRange `json:"createdDateRange,omitempty"  bson:"createdDateRange,omitempty"`
}

// RefTicket : ""
type RefTicket struct {
	Ticket `bson:",inline"`
	Ref    struct {
	} `json:"ref" bson:"ref,omitempty"`
}
