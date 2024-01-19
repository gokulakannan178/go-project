package models

import "time"

//ShopRentRateMaster : ""
type ShopRentRateMaster struct {
	UniqueID          string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	ShopRentID        string     `json:"shopRentId" bson:"shopRentId,omitempty"`
	ShopCategoryID    string     `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID string     `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Rate              float64    `json:"rate" bson:"rate,omitempty"`
	DOE               *time.Time `json:"doe" bson:"doe,omitempty"`
	Status            string     `json:"status" bson:"status,omitempty"`
	Created           *Created   `json:"created,omitempty"  bson:"created,omitempty"`
	Updated           []Updated  `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//ShopRentRateMasterFilter : ""
type ShopRentRateMasterFilter struct {
	UniqueID          []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ShopCategoryID    []string `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID []string `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	ShopRentID        []string `json:"shopRentId" bson:"shopRentId,omitempty"`
	Status            []string `json:"status,omitempty" bson:"status,omitempty"`
}

//RefShopRentRateMaster : ""
type RefShopRentRateMaster struct {
	ShopRentRateMaster `bson:",inline"`
	Ref                struct {
		ShopRentShopCategory    ShopRentShopCategory    `json:"shopRentShopCategory,omitempty" bson:"shopRentShopCategory,omitempty"`
		ShopRentShopSubCategory ShopRentShopSubCategory `json:"shopRentShopSubCategory,omitempty" bson:"shopRentShopSubCategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
