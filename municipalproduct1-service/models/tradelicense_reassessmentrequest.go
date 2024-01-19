package models

// TradeLicenseReassessmentRequest : ""
type TradeLicenseReassessmentRequest struct {
	TradeLicenseID string          `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
	UniqueID       string          `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous       RefTradeLicense `json:"previous,omitempty" bson:"previous,omitempty"`
	New            RefTradeLicense `json:"new,omitempty" bson:"new,omitempty"`
	UserName       string          `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType       string          `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester      Updated         `json:"requester" bson:"requester,omitempty"`
	Action         Updated         `json:"action" bson:"action,omitempty"`
	Proof          []string        `json:"proof,omitempty" bson:"proof,omitempty"`
	Status         string          `json:"status,omitempty" bson:"status,omitempty"`
}

// TradeLicenseReassessmentRequestFilter : ""
type TradeLicenseReassessmentRequestFilter struct {
	TradeLicenseID []string `json:"tradeLicenseId,omitempty" bson:"tradeLicenseId,omitempty"`
	UniqueID       []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
}

// RefTradeLicenseReassessmentRequest : ""
type RefTradeLicenseReassessmentRequest struct {
	TradeLicenseReassessmentRequest `bson:",inline"`
	Ref                             struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptTradeLicenseReassessmentRequestUpdate : ""
type AcceptTradeLicenseReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectTradeLicenseReassessmentRequestUpdate : ""
type RejectTradeLicenseReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}
