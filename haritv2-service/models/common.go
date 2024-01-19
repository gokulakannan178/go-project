package models

import "time"

//Created : "Used To store created On and created by details"
type Created struct {
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	By       string     `json:"by,omitempty" form:"by" bson:"by,omitempty"`
	Scenario string     `json:"scenario" bson:"scenario,omitempty"`
}

//Updated : ""
type Updated struct {
	On       *time.Time `json:"on" bson:"updatedOn,omitempty"`
	By       string     `json:"by" bson:"by,omitempty"`
	Scenario string     `json:"scenario" bson:"scenario,omitempty"`
	ByType   string     `json:"byType,omitempty" form:"byType" bson:"byType,omitempty"`
	Remarks  string     `json:"remarks" bson:"remarks,omitempty"`
}
type CreatedV2 struct {
	On      *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	By      string     `json:"by,omitempty" form:"by" bson:"by,omitempty"`
	ByType  string     `json:"bytype,omitempty" form:"bytype" bson:"bytype,omitempty"`
	Remarks string     `json:"remarks" bson:"remarks,omitempty"`
}

//DateRange : ""
type DateRange struct {
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

//Action
type Action struct {
	On      *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	By      string     `json:"by,omitempty" form:"by" bson:"by,omitempty"`
	ByType  string     `json:"bytype,omitempty" form:"bytype" bson:"bytype,omitempty"`
	Remarks string     `json:"remarks" bson:"remarks,omitempty"`
}
type Contact struct {
	Salutation   string `json:"salutation" bson:"salutation,omitempty"`
	SalutationID string `json:"salutationID" bson:"salutationID,omitempty"`
	Firstname    string `json:"firstname" bson:"firstname,omitempty"`
	LastName     string `json:"lastName" bson:"lastName,omitempty"`
	Email        string `json:"email" bson:"email,omitempty"`
	Ph           string `json:"ph" bson:"ph,omitempty"`
	Mobile       string `json:"mobile" bson:"mobile,omitempty"`
	Skype        string `json:"skype" bson:"skype,omitempty"`
	Designation  string `json:"designation" bson:"designation,omitempty"`
	Dept         string `json:"dept" bson:"dept,omitempty"`
}
