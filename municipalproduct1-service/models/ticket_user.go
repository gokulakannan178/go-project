package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//TicketUser : ""
type TicketUser struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	TicketID string             `json:"ticketId" bson:"ticketId,omitempty"`
	UserID   string             `json:"userId" bson:"userId,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//TicketUserFilter : ""
type TicketUserFilter struct {
	UniqueID         []string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	TicketID         []string   `json:"ticketId" bson:"ticketId,omitempty"`
	UserID           []string   `json:"userId" bson:"userId,omitempty"`
	Status           []string   `json:"status" bson:"status,omitempty"`
	CreatedDateRange *DateRange `json:"createdDateRange,omitempty"  bson:"createdDateRange,omitempty"`
}

// RefTicketUser : ""
type RefTicketUser struct {
	TicketUser `bson:",inline"`
	Ref        struct {
	} `json:"ref" bson:"ref,omitempty"`
}
