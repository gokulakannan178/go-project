package models

import "time"

type Cart struct {
	ULBID     string     `json:"ulbId" bson:"ulbId,omitempty"`
	FPOID     string     `json:"fpoId" bson:"fpoId,omitempty"`
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	Created   Created    `json:"created,omitempty"  bson:"created,omitempty"`
	DateTime  *time.Time `json:"updated,omitempty"  bson:"updated,omitempty"`
	Quantity  float64    `json:"quantity,omitempty"  bson:"quantity,omitempty"`
	ProductID string     `json:"productId" bson:"productId,omitempty"`
}

//RefCart : ""
type RefCart struct {
	Cart `bson:",inline"`
	Ref  struct {
		Ulb             RefULB  `json:"ulb" bson:"ulb,omitempty"`
		CalculatedPrice float64 `json:"calculatedPrice" bson:"calculatedPrice,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//CartFilter : ""
type CartFilter struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	ULBID  []string `json:"ulbId,omitempty" bson:"ulbId,omitempty"`
	FPOID  []string `json:"fpoId,omitempty" bson:"fpoId,omitempty"`
}
