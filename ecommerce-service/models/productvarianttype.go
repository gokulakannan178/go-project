package models

type ProductVariantType struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status   string `json:"status" bson:"status,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	Desc     string `json:"desc" bson:"desc,omitempty"`
}

type ProductVariantTypeFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefProductVariantType struct {
	ProductVariantType `bson:",inline"`
}
