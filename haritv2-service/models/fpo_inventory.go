package models

import "time"

type FPOInventory struct {
	UniqueID              string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	CompanyID             string     `json:"companyId" bson:"companyId,omitempty"`
	Quantity              float64    `json:"quantity" bson:"quantity,omitempty"`
	BuyingPrice           float64    `json:"buyingPrice" bson:"buyingPrice,omitempty"`
	Sellingprice          float64    `json:"sellingPrice" bson:"sellingPrice,omitempty"`
	Status                string     `json:"status" bson:"status,omitempty"`
	Created               *CreatedV2 `json:"created" bson:"created,omitempty"`
	ProductID             string     `json:"productId" bson:"productId,omitempty"`
	IsSellingPriceUpdated bool       `json:"isSellingPriceUpdated" bson:"isSellingPriceUpdated,omitempty"`
}

type FPOInventoryLog struct {
	Old      FPOInventory `json:"old" bson:"old,omitempty"`
	New      FPOInventory `json:"new" bson:"new,omitempty"`
	Msg      string       `json:"msg" bson:"msg,omitempty"`
	Scenario string       `json:"scenario" bson:"scenario,omitempty"`
	Date     *time.Time   `json:"date" bson:"date,omitempty"`
	By       string       `json:"by" bson:"by,omitempty"`
	ByType   string       `json:"byType" bson:"byType,omitempty"`
}

type RefFPOINVENTORY struct {
	FPOInventory `bson:",inline"`
}
