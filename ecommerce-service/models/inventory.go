package models

type Inventory struct {
	UniqueID         string  `json:"uniqueId" bson:"uniqueId,omitempty"`
	VendorID         string  `json:"vendorId" bson:"vendorId,omitempty"`
	ProductVarientID string  `json:"productVarientId" bson:"productVarientId,omitempty"`
	Status           string  `json:"status" bson:"status,omitempty"`
	Quantity         float64 `json:"quantity" bson:"quantity,omitempty"`
	Price            struct {
		Selling float64 `json:"selling" bson:"selling,omitempty"`
		Buying  float64 `json:"buying" bson:"buying,omitempty"`
	} `json:"price" bson:"price,omitempty"`
	LowStock float64 `json:"lowStock" bson:"lowStock,omitempty"`
}

type InventoryFilter struct {
	Status        []string `json:"status" bson:"status,omitempty"`
	VendorID      []string `json:"vendorId" bson:"vendorId,omitempty"`
	CategoryID    []string `json:"categoryId" bson:"categoryId,omitempty"`
	SubCategoryID []string `json:"subCategoryId" bson:"subCategoryId,omitempty"`
	PVCombination []string `json:"pVCombination" bson:"pVCombination,omitempty"`
	ProductID     []string `json:"productId" bson:"productId,omitempty"`
	QuantityRange *struct {
		From float64 `json:"from" bson:"from,omitempty"`
		To   float64 `json:"to" bson:"to,omitempty"`
	} `json:"quantityRange" bson:"quantityRange,omitempty"`
}

type RefInventory struct {
	Inventory `bson:",inline"`
	Ref       struct {
		Product *RefProduct `json:"product,omitempty" bson:"product,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
	Mesh []RefInventoryMesh `json:"mesh" bson:"mesh,omitempty"`
}

type InventoryMesh struct {
	InventoryID   string `json:"inventoryId" bson:"inventoryId,omitempty"`
	MeshID        string `json:"meshId" bson:"meshId,omitempty"`
	VarientTypeID string `json:"varientTypeId" bson:"varientTypeId,omitempty"`
	Name          string `json:"name" bson:"name,omitempty"`
	ProductID     string `json:"productId" bson:"productId,omitempty"`
	VendorID      string `json:"vendorId" bson:"vendorId,omitempty"`
}

type InventoryData struct {
	Inventory     Inventory          `json:"inventory" bson:"inventory,omitempty"`
	InventoryMesh []RefInventoryMesh `json:"inventoryMesh" bson:"inventoryMesh,omitempty"`
}

type InventoryMeshCreate struct {
	Product        Product             `json:"product" bson:"product,omitempty"`
	ProductVarient map[string][]string `json:"productVarient" bson:"productVarient,omitempty"`
	VendorID       string              `json:"vendorId" bson:"vendorId,omitempty"`
}

type RefInventoryMesh struct {
	InventoryMesh `bson:",inline"`
}
