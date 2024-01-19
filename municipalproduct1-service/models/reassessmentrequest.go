package models

// ReassessmentRequest : ""
type ReassessmentRequest struct {
	PropertyID    string      `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	UniqueID      string      `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous      RefProperty `json:"previous,omitempty" bson:"previous,omitempty"`
	New           RefProperty `json:"new,omitempty" bson:"new,omitempty"`
	Created       *Created    `json:"created,omitempty" bson:"created,omitempty"`
	UserName      string      `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType      string      `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester     Updated     `json:"requester" bson:"requester,omitempty"`
	Action        Updated     `json:"action" bson:"action,omitempty"`
	Proof         []string    `json:"proof,omitempty" bson:"proof,omitempty"`
	Status        string      `json:"status,omitempty" bson:"status,omitempty"`
	NewPropertyID string      `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string      `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

type ReassessmentRequestFilter struct {
	PropertyID []string       `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	UniqueID   []string       `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Address    *AddressSearch `json:"address"`
	IsLocation bool           `json:"isLocation"`
	Regex      struct {
		Mobile        string `json:"mobile" bson:"mobile,omitempty"`
		PropertyNo    string `json:"propertyNo" bson:"propertyNo"`
		ApplicationNo string `json:"applicationNo" bson:"applicationNo"`
		OwnerName     string `json:"ownerName" bson:"ownerName"`
	} `json:"regex" bson:"regex"`
	AppliedRange *PropertyAppliedRange `json:"appliedRange" bson:"appliedRange"`
	Status       []string              `json:"status,omitempty" bson:"status,omitempty"`
	SortBy       string                `json:"sortBy,omitempty"`
	SortOrder    int                   `json:"sortOrder,omitempty"`
}

type RefReassessmentRequest struct {
	ReassessmentRequest `bson:",inline"`
	Ref                 struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptReassessmentRequestUpdate : ""
type AcceptReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectReassessmentRequestUpdate : ""
type RejectReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}
