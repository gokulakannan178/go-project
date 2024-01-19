package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PropertyDeleterequest : "Used to show properties requested for delete"
type PropertyDeleteRequest struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	PropertyID string             `json:"propertyID" bson:"propertyID,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Requester  Updated            `json:"requester" bson:"requester,omitempty"`
	Action     Updated            `json:"action" bson:"action,omitempty"`
	Created    Created            `json:"created,omitempty" bson:"created,omitempty"`
}

//PropertyDeleteRequestFilter : ""
type PropertyDeleteRequestFilter struct {
	PropertyID []string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	SearchText struct {
		PropertyID string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
		UniqueID   string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	} `json:"searchText,omitempty" bson:"searchText,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty"`
}

type RefPropertyDeleteRequest struct {
	PropertyDeleteRequest `bson:",inline"`
	Ref                   struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptPropertyDeleteRequestUpdate : ""
type AcceptPropertyDeleteRequestUpdate struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectPropertyDeleteRequestUpdate : ""
type RejectPropertyDeleteRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}
