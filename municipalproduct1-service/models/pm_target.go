package models

type PmTarget struct {
	PmId         string     `json:"pmId" bson:"pmId,omitempty"`
	TargetAmount float64    `json:"targetAmount" bson:"targetAmount,omitempty"`
	Status       string     `json:"status" bson:"status,omitempty"`
	UniqueID     string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Created      *CreatedV2 `json:"created" bson:"created,omitempty"`
	Updated      Updated    `json:"updated" bson:"updated,omitempty"`
}

type PmTargetFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefPmTarget struct {
	PmTarget `bson:",inline"`
}
