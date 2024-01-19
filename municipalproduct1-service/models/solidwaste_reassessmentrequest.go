package models

// ReassessmentRequest : ""
type SolidWasteReassessmentRequest struct {
	SolidWasteID string                  `json:"solidWasteId,omitempty" bson:"solidWasteId,omitempty"`
	UniqueID     string                  `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous     RefSolidWasteUserCharge `json:"previous,omitempty" bson:"previous,omitempty"`
	New          RefSolidWasteUserCharge `json:"new,omitempty" bson:"new,omitempty"`
	UserName     string                  `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType     string                  `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester    Updated                 `json:"requester" bson:"requester,omitempty"`
	Action       Updated                 `json:"action" bson:"action,omitempty"`
	Proof        []string                `json:"proof,omitempty" bson:"proof,omitempty"`
	Status       string                  `json:"status,omitempty" bson:"status,omitempty"`
}

type SolidWasteReassessmentRequestFilter struct {
	SolidWasteID []string `json:"solidWasteId,omitempty" bson:"solidWasteId,omitempty"`
	UniqueID     []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
}

type RefSolidWasteReassessmentRequest struct {
	SolidWasteReassessmentRequest `bson:",inline"`
	Ref                           struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptReassessmentRequestUpdate : ""
type AcceptSolidWasteReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectReassessmentRequestUpdate : ""
type RejectSolidWasteReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}
