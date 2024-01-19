package models

type GST struct {
	UniqueID   string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Label      string     `json:"label" bson:"label,omitempty"`
	Desc       string     `json:"desc" bson:"desc,omitempty"`
	Status     string     `json:"status" bson:"status,omitempty"`
	Percentage float64    `json:"percentage" bson:"percentage"`
	IsDefault  bool       `json:"isDefault" bson:"isDefault,omitempty"`
	Updated    Updated    `json:"updated" bson:"updated,omitempty"`
	Created    *CreatedV2 `json:"created" bson:"created,omitempty"`
}

type GSTFilter struct {
	Status   []string `json:"status" bson:"status,omitempty"`
	UniqueID []string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

type RefGST struct {
	GST `bson:",inline"`
}
