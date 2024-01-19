package models

//ShopRentShopCategory : ""
type ShopRentShopCategory struct {
	UniqueID string    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string    `json:"name" bson:"name,omitempty"`
	Desc     string    `json:"desc" bson:"desc,omitempty"`
	Status   string    `json:"status" bson:"status,omitempty"`
	Created  *Created  `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//ShopRentShopCategoryFilter : ""
type ShopRentShopCategoryFilter struct {
	UniqueID []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name     []string `json:"name,omitempty" bson:"name,omitempty"`
	Status   []string `json:"status,omitempty" bson:"status,omitempty"`
}

type RefShopRentShopCategory struct {
	ShopRentShopCategory `bson:",inline"`
	Ref                  struct {
		ShopRentShopSubCategory ShopRentShopSubCategory `json:"shopRentShopSubCategory" bson:"shopRentShopSubCategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
