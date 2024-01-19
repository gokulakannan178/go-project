package models

type ProductVariant struct {
	UniqueID    string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name        string   `json:"name" bson:"name,omitempty"`
	Images      []string `json:"images" bson:"images,omitempty"`
	Desc        string   `json:"desc" bson:"desc,omitempty"`
	DispalyName string   `json:"dispalyName" bson:"dispalyName,omitempty"`
	ProductID   string   `json:"productId" bson:"productId,omitempty"`
	Status      string   `json:"status" bson:"status,omitempty"`
}

type ProductVariantFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	ProductID []string `json:"productId" bson:"productId,omitempty"`
}
type ProductVariantInventoryFilter struct {
	Status    []string `json:"status" bson:"status,omitempty"`
	VendorID  []string `json:"vendorId" bson:"vendorId,omitempty"`
	ProductID []string `json:"productId" bson:"productId,omitempty"`
}

type RefProductVariant struct {
	ProductVariant `bson:",inline"`
	Ref            struct {
		Product     Product     `json:"product" bson:"product,omitempty"`
		Category    Category    `json:"category" bson:"category,omitempty"`
		Inventory   Inventory   `json:"inventory" bson:"inventory,omitempty"`
		SubCategory SubCategory `json:"subCategory" bson:"subCategory,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}
type RegProductVariant struct {
	ProductVariant `bson:",inline"`
	Mesh           []ProductVariantMesh `json:"mesh" bson:"mesh,omitempty"`
}
