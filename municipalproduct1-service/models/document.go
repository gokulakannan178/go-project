package models

type Document struct {
	Name       string    `json:"name,omitempty" bson:"name,omitempty"`
	Status     string    `json:"status,omitempty" bson:"status,omitempty"`
	UniqueID   string    `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Created    CreatedV2 `json:"created,omitempty" bson:"created,omitempty"`
	UpdatedLog []Updated `json:"updatedLog,omitempty" bson:"updatedLog,omitempty"`
}
