package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Order : ""
type Order struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       string             `json:"status"  bson:"status,omitempty"`
	From         OrderFrom          `json:"orderFrom"  bson:"orderFrom,omitempty"`
	To           OrderTo            `json:"orderTo"  bson:"orderTo,omitempty"`
	Items        []Items            `json:"items"  bson:"items,omitempty"`
	Total        float64            `json:"total"  bson:"total,omitempty"`
	Subtotal     float64            `json:"subtotal"  bson:"subtotal,omitempty"`
	Discount     float64            `json:"discount"  bson:"discount,omitempty"`
	Tax          float64            `json:"tax"  bson:"tax,omitempty"`
	Created      *Created           `json:"created"  bson:"created,omitempty"`
	Date         time.Time          `json:"date"  bson:"date,omitempty"`
	Payment      struct {
		PendingAmount float64 `json:"pendingAmount"  bson:"pendingAmount,omitempty"`
		Status        string  `json:"status"  bson:"status,omitempty"`
	} `json:"payment"  bson:"payment,omitempty"`
}

type OrderFilter struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus   []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy         string   `json:"sortBy"`
	SortOrder      int      `json:"sortOrder"`
	DateOredrRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateOrderRange"`
	SearchBox struct {
		FromName   string `json:"fromName" bson:"fromName"`
		FromMobile string `json:"fromMobile"  bson:"fromMobile,omitempty"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefOrder struct {
	Order `bson:",inline"`
	Ref   struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type OrderFrom struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Mobile        string             `json:"mobile"  bson:"mobile,omitempty"`
	Email         string             `json:"email"  bson:"email,omitempty"`
	Type          string             `json:"type"  bson:"type,omitempty"`
	Address       Address            `json:"address" bson:"address,omitempty"`
	GramPanchayat primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	State         primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block         primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District      primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	PinCode       float64            `json:"pinCode"  bson:"pinCode,omitempty"`
}

type OrderTo struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Mobile        string             `json:"mobile"  bson:"mobile,omitempty"`
	Email         string             `json:"email"  bson:"email,omitempty"`
	Type          string             `json:"type"  bson:"type,omitempty"`
	Address       Address            `json:"address" bson:"address,omitempty"`
	GramPanchayat primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	State         primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block         primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District      primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	PinCode       float64            `json:"pinCode"  bson:"pinCode,omitempty"`
}
type Items struct {
	Name         string             `json:"name" bson:"name,omitempty"`
	ProductID    primitive.ObjectID `json:"productId"  bson:"productId,omitempty"`
	BuyingPrice  float64            `json:"buyingPrice"  bson:"buyingPrice,omitempty"`
	SellingPrice float64            `json:"sellingPrice"  bson:"sellingPrice,omitempty"`
	Quantity     float64            `json:"quantity"  bson:"quantity,omitempty"`
	Amount       float64            `json:"amount"  bson:"amount,omitempty"`
}
