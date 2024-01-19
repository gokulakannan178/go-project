package models

import (
	"time"
)

// PenaltyLogs : ""
type PenaltyLogs struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on" bson:"on,omitempty"`
	By       string     `json:"by" bson:"by,omitempty"`
	Message  string     `json:"message" bson:"message,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
}

// RefPropertyMobileTower : ""
type RefPropertyPenalty struct {
	PenaltyLogs `bson:",inline"`
}

// PropertyMobileTowerFilter : ""
type PropertyPenaltyFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

// func(pl *PenaltyLog) SetMobileTowerPenaltyLog {

// }

// func(pl *PenaltyLog) SetPropertyPenaltyLog {

// }
