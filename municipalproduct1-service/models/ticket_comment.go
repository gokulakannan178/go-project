package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//TicketComment : ""
type TicketComment struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID    string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	TicketID    string             `json:"ticketId" bson:"ticketId,omitempty"`
	PostedBy    string             `json:"postedBy" bson:"postedBy,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
	Created     *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
	Updated     []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	Description string             `json:"description,omitempty"  bson:"description,omitempty"`
	Document    string             `json:"document,omitempty"  bson:"document,omitempty"`
	ParentID    string             `json:"parentID,omitempty"  bson:"parentID,omitempty"`
}

//TicketCommentFilter : ""
type TicketCommentFilter struct {
	UniqueID         string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	TicketID         string     `json:"ticketId" bson:"ticketId,omitempty"`
	Status           []string   `json:"status" bson:"status,omitempty"`
	CreatedDateRange *DateRange `json:"createdDateRange,omitempty"  bson:"createdDateRange,omitempty"`
}

// RefTicketComment : ""
type RefTicketComment struct {
	TicketComment `bson:",inline"`
	Ref           struct {
	} `json:"ref" bson:"ref,omitempty"`
}
