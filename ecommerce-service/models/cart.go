package models

type Cart struct {
	CreateOrder `bson:",inline"`
	UniqueID    string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status      string `json:"status"  bson:"status,omitempty"`
}
type RefCart struct {
	Cart `bson:",inline"`
	// Ref     struct {
	// } `json:"ref" bson:"ref,omitempty"`
}
type CartFilter struct {
	Status []string `json:"status"  bson:"status,omitempty"`
}
type UpdateCart struct {
	CustomerId   string `json:"customerId" bson:"customerId,omitempty"`
	CustomerType string `json:"customerType" bson:"customerType,omitempty"`
	InventoryID  string `json:"inventoryid" bson:"inventoryid,omitempty"`
	//	CartUniqueID string  `json:"cartUniqueId" bson:"cartUniqueId,omitempty"`
	VendorID   string  `json:"vendorId" bson:"vendorId,omitempty"`
	VendorType string  `json:"vendorType" bson:"vendorType,omitempty"`
	Quantity   float64 `json:"quantity" bson:"quantity,omitempty"`
}
