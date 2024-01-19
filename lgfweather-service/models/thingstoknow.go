package models

type ThingsToKnow struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string     `json:"name" bson:"name,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Created  *CreatedV2 `json:"byType" bson:"byType,omitempty"`
	FileURL  []string   `json:"fileUrl" bson:"v,omitempty"`
}

type RefThingsToKnow struct {
	ThingsToKnow `bson:",inline"`
}

type FilterThingsToKnow struct {
	Status []string `json:"status" bson:"status,omitempty"`
}
