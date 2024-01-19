package models

// Boring Charges : ""
type BoringCharges struct {
	UniqueID            string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Title               string     `json:"title" bson:"title,omitempty"`
	Amount              float64    `json:"amount" bson:"amount,omitempty"`
	Created             *CreatedV2 `json:"created" bson:"created,omitempty"`
	Status              string     `json:"status" bson:"status,omitempty"`
	WaterConnectionType string     `json:"waterConnectionType" bson:"waterConnectionType,omitempty"`
}

// Boring Charges Filter : ""
type BoringChargesFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder" bson:"sortOrder,omitempty"`
}

// RefBoringCharges : ""
type RefBoringCharges struct {
	BoringCharges `bson:",inline"`
	Ref           struct {
	} `json:"ref" bson:"ref,omitempty"`
}
