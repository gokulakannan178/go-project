package models

type MonthlyTarget struct {
	UniqueID string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string    `json:"status" bson:"status,omitempty"`
	UserName string    `json:"userName" bson:"userName,omitempty`
	Amount   float64   `json:"amount" bson:"amount,omitempty`
	Created  CreatedV2 `json:"created" bson:"created,omitempty`
	Updated  Updated   `json:"updated" bson:"updated,omitempty`
}

type MonthlyTargetFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefMonthlyTarget struct {
	MonthlyTarget `bson:",inline"`
	RefUser       `bson:",inline"`
}
