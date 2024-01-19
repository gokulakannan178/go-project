package models

import "time"

//Preregistration : ""
type Preregistration struct {
	Applicant `bson:",inline"`
	UniqueID  string                    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Submitted Created                   `json:"submitted" bson:"submitted,omitempty"`
	Reapplied Created                   `json:"reapplied" bson:"-"`
	Log       []PreregistrationTimeline `json:"log" bson:"log,omitempty"`
}

//RefPreregistration :""
type RefPreregistration struct {
	Preregistration `bson:",inline"`
	Ref             struct {
		Address       RefAddress     `json:"address,omitempty" bson:"address,omitempty"`
		ApplicantType *ApplicantType `json:"applicantType,omitempty" bson:"applicantType,omitempty"`
		DaysSince     int64          `json:"dayssince" bson:"dayssince"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//PreregistrationFilter : ""
type PreregistrationFilter struct {
	Status            []string       `json:"status"`
	UniqueID          []string       `json:"uniqueId"`
	Address           *AddressSearch `json:"address"`
	ApplicantType     []string       `json:"applicantType"`
	SortBy            string         `json:"sortBy"`
	SortOrder         int            `json:"sortOrder"`
	IsGetExpiredDraft bool           `json:"isGetExpiredDraft"`
}

//PreregistrationTimeline : ""
type PreregistrationTimeline struct {
	On      *time.Time `json:"on" bson:"on,omitempty"`
	By      string     `json:"by" bson:"by,omitempty"`
	ByType  string     `json:"byType" bson:"byType,omitempty"`
	ByName  string     `json:"byName" bson:"byName,omitempty"`
	Remarks string     `json:"remarks" bson:"remarks,omitempty"`
	Type    string     `json:"type" bson:"type,omitempty"`
}

//PreregistrationStatusChange : ""
type PreregistrationStatusChange struct {
	ApplicantID string     `json:"applicantId,omitempty" bson:"applicantId,omitempty"`
	Status      string     `json:"status,omitempty" bson:"status,omitempty"`
	Remarks     string     `json:"remarks,omitempty" bson:"remarks,omitempty"`
	On          *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	By          string     `json:"by,omitempty" bson:"by,omitempty"`
	ByType      string     `json:"byType,omitempty" bson:"byType,omitempty"`
	ByName      string     `json:"byName,omitempty" bson:"byName,omitempty"`
}

//PreregistrationPayment : ""
type PreregistrationPayment struct {
	ApplicantID string     `json:"applicantId,omitempty" bson:"applicantId,omitempty"`
	Remarks     string     `json:"remarks,omitempty" bson:"remarks,omitempty"`
	On          *time.Time `json:"on,omitempty" bson:"on,omitempty"`
	By          string     `json:"by,omitempty" bson:"by,omitempty"`
	ByType      string     `json:"byType,omitempty" bson:"byType,omitempty"`
	ByName      string     `json:"byName,omitempty" bson:"byName,omitempty"`
}
