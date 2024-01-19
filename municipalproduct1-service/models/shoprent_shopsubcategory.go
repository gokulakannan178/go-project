package models

//ShopRentShopSubCategory : ""
type ShopRentShopSubCategory struct {
	UniqueID       string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name           string    `json:"name" bson:"name,omitempty"`
	Desc           string    `json:"desc" bson:"desc,omitempty"`
	Status         string    `json:"status" bson:"status,omitempty"`
	Created        Created   `json:"created,omitempty"  bson:"created,omitempty"`
	Updated        []Updated `json:"updated,omitempty"  bson:"updated,omitempty"`
	ShopCategoryID string    `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
}

//ShopRentShopSubCategoryFilter : ""
type ShopRentShopSubCategoryFilter struct {
	ShopCategoryID []string `json:"shopCategoryId,omitempty"  bson:"shopCategoryId,omitempty"`
	UniqueID       []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name           []string `json:"name,omitempty" bson:"name,omitempty"`
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
}

// RefShopRentShopSubCategory : ""
type RefShopRentShopSubCategory struct {
	ShopRentShopSubCategory `bson:",inline"`
	Ref                     struct {
		ShopRentShopCategory ShopRentShopCategory `json:"shopRentShopCategory,omitempty" bson:"shopRentShopCategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
