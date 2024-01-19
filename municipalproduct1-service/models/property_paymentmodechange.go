package models

import "time"

// PropertyPaymentModeChange : ""
type PropertyPaymentModeChangeRequest struct {
	TnxID         string                  `json:"tnxId" bson:"tnxId,omitempty"`
	UniqueID      string                  `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	PropertyID    string                  `json:"propertyId" bson:"propertyId,omitempty"`
	OwnerName     string                  `json:"ownerName" bson:"ownerName,omitempty"`
	Mobile        string                  `json:"mobile" bson:"mobile,omitempty"`
	ReciptNo      string                  `json:"reciptNo" bson:"reciptNo,omitempty"`
	ReciptDate    *time.Time              `json:"reciptDate" bson:"reciptDate,omitempty"`
	Previous      *PropertyPaymentDetails `json:"previous" bson:"previous,omitempty"`
	New           *PropertyPaymentDetails `json:"new" bson:"new,omitempty"`
	Status        string                  `json:"status" bson:"status,omitempty"`
	Created       *CreatedV2              `json:"created" bson:"created,omitempty"`
	PaymentMode   string                  `json:"paymentMode" bson:"paymentMode,omitempty"`
	Requested     Action                  `json:"requested,omitempty" bson:"requested,omitempty"`
	Approved      Action                  `json:"approved,omitempty" bson:"approved,omitempty"`
	Rejected      Action                  `json:"rejected,omitempty" bson:"rejected,omitempty"`
	NewPropertyID string                  `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string                  `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

// PropertyPaymentModeChangeFilter : ""
type PropertyPaymentModeChangeRequestFilter struct {
	PropertyID []string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy     string   `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder  int      `json:"sortOrder" bson:"sortOrder,omitempty"`
}

// RefPropertyPaymentModeChange : ""
type RefPropertyPaymentModeChangeRequest struct {
	PropertyPaymentModeChangeRequest `bson:",inline"`
	Ref                              struct {
		RequestedBy     User     `json:"requestedBy,omitempty" bson:"requestedBy,omitempty"`
		RequestedByType UserType `json:"requestedByType,omitempty" bson:"requestedByType,omitempty"`
		ApprovedBy      User     `json:"approvedBy,omitempty" bson:"approvedBy,omitempty"`
		ApprovedByType  UserType `json:"approvedByType,omitempty" bson:"approvedByType,omitempty"`
		RejectedBy      User     `json:"rejectedBy,omitempty" bson:"rejectedBy,omitempty"`
		RejectedByType  User     `json:"rejectedByType,omitempty" bson:"rejectedByType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptPropertyPaymentModeChange : ""
type AcceptPropertyPaymentModeChangeRequest struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectPropertyPaymentModeChangeRequest : ""
type RejectPropertyPaymentModeChangeRequest struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}
