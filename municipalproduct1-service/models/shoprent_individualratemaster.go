package models

import "time"

//ShopRentindividualRateMaster : ""
type ShopRentindividualRateMaster struct {
	UniqueID     string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	ShopRentID   string     `json:"shopRentId" bson:"shopRentId,omitempty"`
	Rate         float64    `json:"rate" bson:"rate,omitempty"`
	ShopRentArea float64    `json:"shopRentArea" bson:"shopRentArea,omitempty"`
	DOE          *time.Time `json:"doe" bson:"doe,omitempty"`
	Status       string     `json:"status" bson:"status,omitempty"`
	Created      *Created   `json:"created,omitempty"  bson:"created,omitempty"`
	Updated      []Updated  `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//ShopRentindividualRateMasterFilter : ""
type ShopRentindividualRateMasterFilter struct {
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ShopRentID []string `json:"shopRentId" bson:"shopRentId,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
}

//RefShopRentindividualRateMaster : ""
type RefShopRentindividualRateMaster struct {
	ShopRentindividualRateMaster `bson:",inline"`
	Ref                          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
