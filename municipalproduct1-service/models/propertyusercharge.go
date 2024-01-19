package models

import (
	"time"
)

//PropertyUserCharge : ""
type PropertyUserCharge struct {
	UniqueID     string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID   string     `json:"propertyId" bson:""`
	CategoryID   string     `json:"categoryId" bson:"categoryId,omitempty"`
	DOA          *time.Time `json:"doa" bson:"doa,omitempty"`
	IsUserCharge string     `json:"isUserCharge" bson:"isUserCharge,omitempty"`
	Status       string     `json:"status" bson:"status,omitempty"`
	Createdby    CreatedV2  `json:"createdBy" bson:"createdBy,omitempty"`
}

//RefPropertyUserCharge :""
type RefPropertyUserCharge struct {
	PropertyUserCharge `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PropertyUserChargeFilter : ""
type PropertyUserChargeFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
}

type UserChargeAction struct {
	By         string     `json:"by" bson:"by,omitempty"`
	ByType     string     `json:"byType" bson:"byType,omitempty"`
	ActionDate *time.Time `json:"actionDate" bson:"actionDate,omitempty"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
	UniqueId   string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Remarks    string     `json:"remarks" bson:"remarks,omitempty"`
}
