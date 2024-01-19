package models

type TodayTips struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Title    string     `json:"title" bson:"title,omitempty"`
	Status   string     `json:"status" bson:"status,omitempty"`
	Created  *CreatedV2 `json:"created" bson:"created,omitempty"`
}

type TodayTipsFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

type RefTodayTips struct {
	TodayTips `bson:",inline"`
}
