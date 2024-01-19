package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MobileTowerDeleterequest : "Used to show tradelicense requested for delete"
type MobileTowerDeleteRequest struct {
	ID            primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	MobileTowerID string             `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	Requester     Updated            `json:"requester" bson:"requester,omitempty"`
	Action        Updated            `json:"action" bson:"action,omitempty"`
	Created       Created            `json:"created,omitempty" bson:"created,omitempty"`
}

//MobileTowerDeleteRequestFilter : ""
type MobileTowerDeleteRequestFilter struct {
	MobileTowerID []string `json:"mobileTowerId,omitempty" bson:"mobileTowerId,omitempty"`
	UniqueID      []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status        []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy        string   `json:"sortBy,omitempty"`
	SortOrder     int      `json:"sortOrder,omitempty"`
}

type RefMobileTowerDeleteRequest struct {
	MobileTowerDeleteRequest `bson:",inline"`
	Ref                      struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptMobileTowerDeleteRequestUpdate : ""
type AcceptMobileTowerDeleteRequestUpdate struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectMobileTowerDeleteRequestUpdate : ""
type RejectMobileTowerDeleteRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}
