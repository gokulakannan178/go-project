package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TradeLicenseDeleterequest : "Used to show tradelicense requested for delete"
type TradeLicenseDeleteRequest struct {
	ID             primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID       string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	TradeLicenseID string             `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
	Requester      Updated            `json:"requester" bson:"requester,omitempty"`
	Action         Updated            `json:"action" bson:"action,omitempty"`
	Created        Created            `json:"created,omitempty" bson:"created,omitempty"`
}

//TradeLicenseDeleteRequestFilter : ""
type TradeLicenseDeleteRequestFilter struct {
	TradeLicenseID []string `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
	UniqueID       []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	SearchText     struct {
		TradeLicenseID string `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
		UniqueID       string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	} `json:"searchText,omitempty" bson:"searchText,omitempty"`
	SortBy    string `json:"sortBy,omitempty"`
	SortOrder int    `json:"sortOrder,omitempty"`
}

type RefTradeLicenseDeleteRequest struct {
	TradeLicenseDeleteRequest `bson:",inline"`
	Ref                       struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptTradeLicenseDeleteRequestUpdate : ""
type AcceptTradeLicenseDeleteRequestUpdate struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectTradeLicenseDeleteRequestUpdate : ""
type RejectTradeLicenseDeleteRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}
