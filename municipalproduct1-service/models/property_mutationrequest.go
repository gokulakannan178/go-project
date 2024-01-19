package models

import "time"

// PropertyMutationRequest : ""
type PropertyMutationRequest struct {
	PropertyID          string       `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	UniqueID            string       `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Property            *RefProperty `json:"property,omitempty" bson:"property,omitempty"`
	UserName            string       `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType            string       `json:"userType,omitempty" bson:"userType,omitempty"`
	RequestedDate       *time.Time   `json:"requestedDate,omitempty" bson:"requestedDate,omitempty"`
	PropertyMutatedDate *time.Time   `json:"propertyMutatedDate,omitempty" bson:"propertyMutatedDate,omitempty"`
	Requester           Updated      `json:"requester" bson:"requester,omitempty"`
	Action              Updated      `json:"action" bson:"action,omitempty"`
	Status              string       `json:"status,omitempty" bson:"status,omitempty"`
}

type PropertyMutationRequestFilter struct {
	PropertyID []string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy     string   `json:"sortBy,omitempty"`
	SortOrder  int      `json:"sortOrder,omitempty"`
}

type RefPropertyMutationRequest struct {
	PropertyMutationRequest `bson:",inline"`
	Ref                     struct {
		RequestedUser     User     `json:"requestedUser,omitempty" bson:"requestedUser,omitempty"`
		RequestedUserType UserType `json:"requestedUserType,omitempty" bson:"requestedUserType,omitempty"`
		ActionUser        User     `json:"actionUser,omitempty" bson:"actionUser,omitempty"`
		ActionUserType    UserType `json:"actionUserType,omitempty" bson:"actionUserType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//  AcceptPropertyMutationRequestUpdate : ""
type AcceptPropertyMutationRequestUpdate struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	Remark   string     `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}

// RejectPropertyMutationRequestUpdate : ""
type RejectPropertyMutationRequestUpdate struct {
	UniqueID string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Remark   string `json:"remark,omitempty" bson:"remark,omitempty"`
	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string `json:"userType,omitempty" bson:"userType,omitempty"`
}

// MutatedProperty : ""
type MutatedProperty struct {
	ChildID                  string     `json:"childId,omitempty" bson:"childId,omitempty"`
	ParentID                 string     `json:"parentId,omitempty" bson:"parentId,omitempty"`
	UniqueID                 string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Status                   string     `json:"status,omitempty" bson:"status,omitempty"`
	Property                 *Property  `json:"property,omitempty" bson:"property,omitempty"`
	RemainingAreaOfPlot      float64    `json:"remainingAreaOfPlot" bson:"remainingAreaOfPlot,omitempty"`
	RemainingBuiltUpArea     float64    `json:"remainingBuiltUpArea" bson:"remainingBuiltUpArea,omitempty"`
	PercentBuiltUpAreaFilled float64    `json:"percentBuiltUpAreaFilled" bson:"percentBuiltUpAreaFilled,omitempty"`
	PercentAreaOfPlotFilled  float64    `json:"percentAreaOfPlotFilled" bson:"percentAreaOfPlotFilled,omitempty"`
	TotalAreaOfPlot          float64    `json:"totalAreaOfPlot" bson:"totalAreaOfPlot,omitempty"`
	TotalBuiltUpArea         float64    `json:"totalBuiltUpArea" bson:"totalBuiltUpArea,omitempty"`
	BuiltUpArea              float64    `json:"builtUpArea" bson:"builtUpArea,omitempty"`
	Created                  *CreatedV2 `json:"created,omitempty" bson:"created,omitempty"`
}

// MutatedPropertyFilter : ""
type MutatedPropertyFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	UniqueID  []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ParentID  []string `json:"parentId,omitempty" bson:"parentId,omitempty"`
	SortBy    string   `json:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty"`
}

// RefMutatedProperties : ""
type RefMutatedProperty struct {
	MutatedProperty `bson:",inline"`
	Ref             struct {
		Property RefProperty `json:"property,omitempty" bson:"property,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type RefRemainingOfMutatedProperty struct {
	RemainingAreaOfPlot      float64 `json:"remainingAreaOfPlot" bson:"remainingAreaOfPlot,omitempty"`
	RemainingBuiltUpArea     float64 `json:"remainingBuiltUpArea" bson:"remainingBuiltUpArea,omitempty"`
	TotalAreaOfPlot          float64 `json:"totalAreaOfPlot" bson:"totalAreaOfPlot,omitempty"`
	TotalBuiltUpArea         float64 `json:"totalBuiltUpArea" bson:"totalBuiltUpArea,omitempty"`
	TotalBuiltUpAreaOfChild  float64 `json:"totalBuiltUpAreaOfChild" bson:"totalBuiltUpAreaOfChild,omitempty"`
	TotalAreaOfPlotOfChild   float64 `json:"totalAreaOfPlotOfChild" bson:"totalAreaOfPlotOfChild,omitempty"`
	PercentBuiltUpAreaFilled float64 `json:"percentBuiltUpAreaFilled" bson:"percentBuiltUpAreaFilled,omitempty"`
	PercentAreaOfPlotFilled  float64 `json:"percentAreaOfPlotFilled" bson:"percentAreaOfPlotFilled,omitempty"`
}
