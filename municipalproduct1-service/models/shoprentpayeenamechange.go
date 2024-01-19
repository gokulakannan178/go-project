package models

import "time"

type ShoprentPayeeNameChange struct {
	UniqueID          string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	TransactionId     string     `json:"txnId" bson:"txnId,omitempty"`
	ShoprentId        string     `json:"shoprentId" bson:"shoprentId,omitempty"`
	ReceiptNo         string     `json:"receiptNo" bson:"receiptNo,omitempty"`
	PreviousPayeeName string     `json:"previousPayeeName" bson:"previousPayeeName,omitempty"`
	ChangeData        ChangeName `json:"changedData" bson:"changedData,omitempty"`
	ApprovedBy        CreatedBy  `json:"approvedBy" bson:"approvedBy,omitempty"`
	RejectedBy        CreatedBy  `json:"rejectedBy" bson:"rejectedBy,omitempty"`
	CreatedBy         CreatedBy  `json:"createdBy" bson:"createdBy,omitempty"`
	Status            string     `json:"status" bson:"status,omitempty"`
	CreatedOn         CreatedV2  `json:"createdOn" bson:"createdOn,omitempty"`
}

type ShoprentPayeeNameChangeFilter struct {
	Status     []string   `json:"status" bson:"status,omitempty"`
	ApprovedBy []string   `json:"approvedBy" bson:"approvedBy,omitempty"`
	RejectedBy []string   `json:"rejectedBy" bson:"rejectedBy,omitempty"`
	CreatedBy  []string   `json:"createdBy" bson:"createdBy,omitempty"`
	PropertyId []string   `json:"propertyId" bson:"propertyId,omitempty"`
	ReceiptNo  []string   `json:"receiptNo" bson:"receiptNo,omitempty"`
	DateRange  *DateRange `json:"dateRange"`
	CreatedOn  string     `json:"createdOn" bson:"createdOn,omitempty"`
	SortBy     string     `json:"sortBy" bson:"sortBy,omitempty"`
	SortOrder  string     `json:"sortOrder" bson:"sortOrder,omitempty"`
}
type RefShoprentPayeeNameChange struct {
	ShoprentPayeeNameChange `bson:",inline"`
	Ref                     struct {
	} `json:"ref" bson:"ref,omitempty"`
}

type ApproveShoprentPayeeNameChange struct {
	UniqueID string     `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	UserName string     `json:"userName,omitempty" bson:"userName,omitempty"`
	UserType string     `json:"userType,omitempty" bson:"userType,omitempty"`
}
