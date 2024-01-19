package models

type RegisterProduct struct {
	Product       Product                  `json:"product" bson:"product,omitempty"`
	Vendor        Vendor                   `json:"vendor" bson:"vendor,omitempty"`
	InventoryData map[string]InventoryData `json:"inventoryData" bson:"inventoryData,omitempty"`
	Varients      []ProductVariant         `json:"varients" bson:"varients,omitempty"`
}
