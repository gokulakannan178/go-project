package models

type VendorDashBoard struct {
	VendorID string                  `json:"vendorId" bson:"vendorId,omitempty"`
	Order    VendorDashBoardOrder    `json:"order" bson:"order,omitempty"`
	Product  VendorDashBoardProduct  `json:"product" bson:"product,omitempty"`
	LowStock VendorDashBoardLowStock `json:"lowStock" bson:"lowStock,omitempty"`
}
type VendorDashBoardFilter struct {
	VendorID string `json:"vendorId" bson:"vendorId,omitempty"`
}
type VendorDashBoardOrder struct {
	NoOfOrders     float64 `json:"noOfOrders" bson:"noOfOrders,omitempty"`
	PendingAmounts float64 `json:"pendingAmounts" bson:"pendingAmounts,omitempty"`
	SoldItems      float64 `json:"soldItems" bson:"soldItems,omitempty"`
	TotalSales     float64 `json:"totalSales" bson:"totalSales,omitempty"`
}
type VendorDashBoardProduct struct {
	Products float64 `json:"products" bson:"products,omitempty"`
}

type VendorDashBoardLowStock struct {
	LowStock float64 `json:"lowStocks" bson:"lowStocks,omitempty"`
}
