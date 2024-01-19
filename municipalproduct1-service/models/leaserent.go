package models

import "time"

//LeaseRent : ""
type LeaseRent struct {
	UniqueID           string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ShopCategoryID     string			  `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID  string			  `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Sqft               float64            `json:"sqft" bson:"sqft,omitempty"`
	DateFrom           *time.Time         `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo             *time.Time         `json:"dateTo" bson:"dateTo,omitempty"`
	Status             string             `json:"status" bson:"status,omitempty"`
	Created            Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated            []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}


//LeaseRentFilter : ""
type LeaseRentFilter struct {
	UniqueID          []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ShopCategoryID    []string `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID []string `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Status            []string `json:"status,omitempty" bson:"status,omitempty"`
}


//RefLeaseRent : ""
type RefLeaseRent struct {
	LeaseRent `bson:",inline"`
	Ref     struct {
		LeaseRentShopCategory    *LeaseRentShopCategory    `json:"leaseRentShopCategory,omitempty" bson:"leaseRentShopCategory,omitempty"`
		LeaseRentShopSubCategory *LeaseRentShopSubCategory `json:"leaseRentShopSubCategory,omitempty" bson:"leaseRentShopSubCategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
