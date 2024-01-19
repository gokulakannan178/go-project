package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PropertyFixedArv : ""
type PropertyFixedArv struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID string             `json:"propertyId" bson:"propertyId,omitempty"`
	FyID       string             `json:"fyId" bson:"fyId,omitempty"`
	FyIDs      []string           `json:"fyIds" bson:"fyIds,omitempty"`
	FyFrom     *time.Time         `json:"fyFrom" bson:"fyFrom,omitempty"`
	FyTo       *time.Time         `json:"fyTo" bson:"fyTo,omitempty"`
	ARV        float64            `json:"arv" bson:"arv,omitempty"`
	Tax        float64            `json:"tax" bson:"tax,omitempty"`
	Total      float64            `json:"total" bson:"total,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    CreatedV2          `json:"created" bson:"created,omitempty"`
	Requester  Updated            `json:"requester" bson:"requester,omitempty"`
	Approved   Updated            `json:"approved" bson:"approved,omitempty"`
	Rejected   Updated            `json:"rejected" bson:"rejected,omitempty"`
}

//PropertyFixedArvFilter : ""
type PropertyFixedArvFilter struct {
	UniqueIDs   []string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID  []string   `json:"propertyId" bson:"propertyId,omitempty"`
	Status      []string   `json:"status" bson:"status,omitempty"`
	FyID        []string   `json:"fyId" bson:"fyId,omitempty"`
	ARV         []string   `json:"arv" bson:"arv,omitempty"`
	CreatedBy   []string   `json:"createdBy" bson:"createdBy,omitempty"`
	CreatedDate *DateRange `json:"createdDate"`

	Regex struct {
		UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
		PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
		Name       string `json:"name" bson:"name,omitempty"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// RefPropertyFixedArv : ""
type RefPropertyFixedArv struct {
	PropertyFixedArv `bson:",inline"`
	Ref              struct {
		FinancialYear   FinancialYear `json:"financialYear" bson:"financialYear,omitempty"`
		CreatedBy       User          `json:"createdBy" bson:"createdBy,omitempty"`
		CreatedByType   UserType      `json:"createdByType" bson:"createdByType,omitempty"`
		RequestedBy     User          `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType      `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User          `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType      `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		RejectedBy      User          `json:"rejectedBy,omitempty" bson:"rejectedBy,omitempty"`
		RejectedByType  User          `json:"rejectedByType,omitempty" bson:"rejectedByType,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

type PropertyFixedDemand struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID string             `json:"propertyId" bson:"propertyId,omitempty"`
	FyID       string             `json:"fyId" bson:"fyId,omitempty"`
	FyIDs      []string           `json:"fyIds" bson:"fyIds,omitempty"`
	FyFrom     *time.Time         `json:"fyFrom" bson:"fyFrom,omitempty"`
	FyTo       *time.Time         `json:"fyTo" bson:"fyTo,omitempty"`
	ARV        float64            `json:"arv" bson:"arv,omitempty"`
	Tax        float64            `json:"tax" bson:"tax,omitempty"`
	Total      float64            `json:"total" bson:"total,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    CreatedV2          `json:"created" bson:"created,omitempty"`
}

//  AcceptReassessmentRequestUpdate : ""
type AcceptPropertyFixedArv struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectReassessmentRequestUpdate : ""
type RejectPropertyFixedArv struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

type AcceptMultiplePropertyFixedArv struct {
	UniqueID []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string   `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string   `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string   `json:"userType,omitempty" bson:"userType,omitempty"`
}
