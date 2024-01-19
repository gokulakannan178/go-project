package models

import "time"

//LeaseRentRateMaster : ""
type LeaseRentRateMaster struct {
	UniqueID           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ShopCategoryID     string			  `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID  string			  `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Rate               float64            `json:"rate" bson:"rate,omitempty"`
	DOE               *time.Time          `json:"doe" bson:"doe,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated            []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}


//LeaseRentRateMasterFilter : ""
type LeaseRentRateMasterFilter struct {
	UniqueID          []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ShopCategoryID    []string `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID []string `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Status            []string `json:"status,omitempty" bson:"status,omitempty"`
}


//RefLeaseRentRateMaster : ""
type RefLeaseRentRateMaster struct {
	LeaseRentRateMaster `bson:",inline"`
	Ref     struct {
		LeaseRentShopCategory    *LeaseRentShopCategory    `json:"leaseRentShopCategory,omitempty" bson:"leaseRentShopCategory,omitempty"`
		LeaseRentShopSubCategory *LeaseRentShopSubCategory `json:"leaseRentShopSubCategory,omitempty" bson:"leaseRentShopSubCategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
