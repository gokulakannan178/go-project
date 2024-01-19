package models


//LeaseRentShopCategory : ""
type LeaseRentShopCategory struct {
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Desc     string             `json:"desc" bson:"desc,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated  []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}


//LeaseRentShopCategoryFilter : ""
type LeaseRentShopCategoryFilter struct {
	UniqueID   []string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Name       []string `json:"name,omitempty" bson:"name,omitempty"`
	Status     []string `json:"status,omitempty" bson:"status,omitempty"`
}
