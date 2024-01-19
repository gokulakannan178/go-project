package models

import "time"

type BulkPrint struct {
	UniqueID            string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Title               string     `json:"title" bson:"title,omitempty"`
	Amount              float64    `json:"amount" bson:"amount,omitempty"`
	Created             *CreatedV2 `json:"created" bson:"created,omitempty"`
	Status              string     `json:"status" bson:"status,omitempty"`
	WaterConnectionType string     `json:"waterConnectionType" bson:"waterConnectionType,omitempty"`
}

//BulkPrintReceiptsRequest
type BulkPrintReceiptsRequest struct {
	TnxIds []string `json:"tnxIds" bson:"tnxIds,omitempty"`
}

// Boring Charges Filter : ""
type BulkPrintFilter struct {
	Status    []string   `json:"status" bson:"status,omitempty"`
	UserType  []string   `json:"userType" bson:"userType,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder int        `json:"sortOrder" bson:"sortOrder,omitempty"`
}

// RefBulkPrint : ""
type RefBulkPrint struct {
	BulkPrint `bson:",inline"`
	Ref       struct {
	} `json:"ref" bson:"ref,omitempty"`
}
type BulkPrintDetail struct {
	UniqueID string   `json:"name" bson:"name,omitempty"`
	UserName string   `json:"userName" bson:"userName,omitempty"`
	TnxId    []string `json:"tnxId" bson:"tnxId,omitempty"`
	Amount   float64  `json:"amount" bson:"amount,omitempty"`
	Payments int64    `json:"payments" bson:"payments,omitempty"`
}
