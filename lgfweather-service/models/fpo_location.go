package models

type FPOUpdateLocation struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Location Location `json:"location" bson:"location,omitempty"`
	By       string   `json:"by" bson:"by,omitempty"`
	ByType   string   `json:"byType" bson:"byType,omitempty"`
}
