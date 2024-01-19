package models

import (
	"time"
)

//LetterGenerate : ""
type LetterGenerate struct {
	UniqueID    string               `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name        string               `json:"name" bson:"name,omitempty"`
	Desc        string               `json:"desc" bson:"desc,omitempty"`
	Status      string               `json:"status" bson:"status,omitempty"`
	Subject     string               `json:"subject" bson:"subject,omitempty"`
	Date        *time.Time           `json:"date" bson:"date,omitempty"`
	NO          string               `json:"no" bson:"no,omitempty"`
	Content     string               `json:"content" bson:"content,omitempty"`
	ContentHtml string               `json:"contentHtml" bson:"contentHtml,omitempty"`
	URL         string               `json:"url" bson:"url,omitempty"`
	Created     *CreatedV2           `json:"created,omitempty"  bson:"created,omitempty"`
	Updated     []Updated            `json:"update,omitempty"  bson:"update,omitempty"`
	Submitted   LetterGenerateAction `json:"submitted,omitempty"  bson:"submitted,omitempty"`
	Activated   LetterGenerateAction `json:"activated,omitempty"  bson:"activated,omitempty"`
	Blocked     LetterGenerateAction `json:"blocked,omitempty"  bson:"blocked,omitempty"`
}

//LetterGenerateAction
type LetterGenerateAction struct {
	UniqueID string     `json:"uniqueId" bson:"-"`
	Action   string     `json:"action" form:"action" bson:"action,omitempty"`
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	By       string     `json:"by,omitempty" form:"by" bson:"by,omitempty"`
	ByType   string     `json:"bytype,omitempty" form:"bytype" bson:"bytype,omitempty"`
	Remarks  string     `json:"remarks" bson:"remarks,omitempty"`
}

//RefLetterGenerate : ""
type RefLetterGenerate struct {
	LetterGenerate `bson:",inline"`
	Ref            struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//LetterGenerateFilter : ""
type LetterGenerateFilter struct {
	Status    []string  `json:"status,omitempty" bson:"status,omitempty"`
	DateRange DateRange `json:"dateRange,omitempty"  bson:"dateRange,omitempty"`
}
