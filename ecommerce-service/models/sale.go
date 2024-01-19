package models

import "time"

type Sale struct {
	From            SaleFrom           `json:"from" bson:"from,omitempty"`
	To              SaleTo             `json:"to" bson:"to,omitempty"`
	UniqueID        string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status          string             `json:"status" bson:"status,omitempty"`
	Created         Created            `json:"created,omitempty"  bson:"created,omitempty"`
	BillingAddress  SaleBillingAddress `json:"billingaddress" bson:"billingaddress,omitempty"`
	ShippingAddress SaleBillingAddress `json:"shippingAddress" bson:"shippingAddress,omitempty"`
	Item            []SaleItem         `json:"Items" bson:"Items,omitempty"`
	SubTotal        float64            `json:"subTotal" bson:"subTotal,omitempty"`
	Discount        Discount           `json:"discount" bson:"discount,omitempty"`
	Total           float64            `json:"total" bson:"total,omitempty"`
	Payment         SalePayment        `json:"payment" bson:"payment,omitempty"`
	Date            *time.Time         `json:"date" bson:"date,omitempty"`
}
type SaleFrom struct {
	ID      string  `json:"id" bson:"id,omitempty"`
	Type    string  `json:"type" bson:"type,omitempty"`
	Name    string  `json:"name" bson:"name,omitempty"`
	Logo    string  `json:"logo" bson:"logo,omitempty"`
	Moblie  string  `json:"moblie" bson:"moblie,omitempty"`
	Email   string  `json:"email" bson:"email,omitempty"`
	Address Address `json:"address" bson:"address,omitempty"`
}
type SaleTo struct {
	ID      string  `json:"id" bson:"id,omitempty"`
	Type    string  `json:"type" bson:"type,omitempty"`
	Name    string  `json:"name" bson:"name,omitempty"`
	Moblie  string  `json:"moblie" bson:"moblie,omitempty"`
	Logo    string  `json:"logo" bson:"logo,omitempty"`
	Email   string  `json:"email" bson:"email,omitempty"`
	Address Address `json:"address" bson:"address,omitempty"`
}
type SaleBillingAddress struct {
	Address Address `json:"address" bson:"address,omitempty"`
}
type SaleShippingAddress struct {
	Address Address `json:"address" bson:"address,omitempty"`
}
type SaleItem struct {
	Item     RefInventory `json:"item" bson:"item,omitempty"`
	Quantity float64      `json:"quantity" bson:"quantity,omitempty"`
	Total    float64      `json:"total" bson:"total,omitempty"`
}
type Discount struct {
	Type  string  `json:"type" bson:"type,omitempty"`
	Rate  float64 `json:"rate" bson:"rate,omitempty"`
	Total float64 `json:"total" bson:"total,omitempty"`
}
type SalePayment struct {
	Status        string  `json:"status" bson:"status,omitempty"`
	Amount        float64 `json:"amount" bson:"amount,omitempty"`
	PendingAmount float64 `json:"pendingAmount" bson:"pendingAmount,omitempty"`
}
