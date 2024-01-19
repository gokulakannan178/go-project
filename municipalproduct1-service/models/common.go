package models

import "time"

//Created : "Used To store created On and created by details"
type Created struct {
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	By       string     `json:"by,omitempty" form:"by" bson:"by,omitempty"`
	Scenario string     `json:"scenario" bson:"scenario,omitempty"`
}

// // IsZero : ""
// func (c Created) IsZero() bool {
// 	return false
// }

//Updated : ""
type Updated struct {
	On       *time.Time `json:"on" bson:"on,omitempty"`
	By       string     `json:"by" bson:"by,omitempty"`
	Scenario string     `json:"scenario" bson:"scenario,omitempty"`
	ByType   string     `json:"byType" form:"byType" bson:"byType,omitempty"`
	Remarks  string     `json:"remarks" bson:"remarks,omitempty"`
}
type CreatedV2 struct {
	On      *time.Time `json:"on" bson:"on,omitempty"`
	By      string     `json:"by" bson:"by,omitempty"`
	ByType  string     `json:"bytype"  bson:"bytype,omitempty"`
	Remarks string     `json:"remarks" bson:"remarks,omitempty"`
	Name    string     `json:"name" bson:"name,omitempty"`
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
type Countstruct struct {
	Count int64 `json:"count" form:"count" bson:"count,omitempty"`
}
