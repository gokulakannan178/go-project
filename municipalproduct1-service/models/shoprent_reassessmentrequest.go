package models

// ReassessmentRequest : ""
type ShoprentReassessmentRequest struct {
	ShoprentID string      `json:"shoprentId,omitempty" bson:"shoprentId,omitempty"`
	UniqueID   string      `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Previous   RefShopRent `json:"previous,omitempty" bson:"previous,omitempty"`
	New        RefShopRent `json:"new,omitempty" bson:"new,omitempty"`
	UserName   string      `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType   string      `json:"userType,omitempty" bson:"userType,omitempty"`
	Requester  Updated     `json:"requester" bson:"requester,omitempty"`
	Action     Updated     `json:"action" bson:"action,omitempty"`
	Proof      []string    `json:"proof,omitempty" bson:"proof,omitempty"`
	Status     string      `json:"status,omitempty" bson:"status,omitempty"`
}

type ShoprentReassessmentRequestFilter struct {
	ShoprentID []string `json:"shoprentId,omitempty" bson:"shoprentId,omitempty"`
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy     string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder  int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

type RefShoprentReassessmentRequest struct {
	ShoprentReassessmentRequest `bson:",inline"`
	Ref                         struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptReassessmentRequestUpdate : ""
type AcceptShoprentReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectReassessmentRequestUpdate : ""
type RejectShoprentReassessmentRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}
