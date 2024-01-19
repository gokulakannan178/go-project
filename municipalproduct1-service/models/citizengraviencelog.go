package models

import (
	"time"
)

//CitizenGraviansLog : ""
type CitizenGraviansLog struct {
	UniqueID          string     `json:"uniqueId"  bson:"uniqueId,omitempty"`
	CitizenGraviansID string     `json:"citizenGraviansId"  bson:"citizenGraviansId,omitempty"`
	Status            string     `json:"status"  bson:"status,omitempty"`
	Type              string     `json:"type"  bson:"type,omitempty"`
	On                *time.Time `json:"on"  bson:"on,omitempty"`
	By                string     `json:"by"  bson:"by,omitempty"`
	ByType            string     `json:"byType"  bson:"byType,omitempty"`
	ByID              string     `json:"byId"  bson:"byId,omitempty"`
	Remarks           string     `json:"remarks"  bson:"remarks,omitempty"`
	Desc              string     `json:"desc"  bson:"desc,omitempty"`
	PreviousStatus    string     `json:"previousStatus"  bson:"previousStatus,omitempty"`
	NewStatus         string     `json:"newStatus"  bson:"newStatus,omitempty"`
}

//CitizenGraviansLogFilter : ""
type CitizenGraviansLogFilter struct {
	UniqueID         []string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name             []string   `json:"name" bson:"name,omitempty"`
	Description      []string   `json:"description" bson:"description,omitempty"`
	Status           []string   `json:"status" bson:"status,omitempty"`
	AssignedTo       []string   `json:"assignedTo,omitempty"  bson:"assignedTo,omitempty"`
	AssignedBy       []string   `json:"assignedBy,omitempty"  bson:"assignedBy,omitempty"`
	FromDateRange    *DateRange `json:"fromDateRange,omitempty"  bson:"fromDateRange,omitempty"`
	ToDateRange      *DateRange `json:"toDateRange,omitempty"  bson:"toDateRange,omitempty"`
	CreatedDateRange *DateRange `json:"createdDateRange,omitempty"  bson:"createdDateRange,omitempty"`
}

// RefCitizenGraviansLog : ""
type RefCitizenGraviansLog struct {
	CitizenGraviansLog `bson:",inline"`
	Ref                struct {
	} `json:"ref" bson:"ref,omitempty"`
}
