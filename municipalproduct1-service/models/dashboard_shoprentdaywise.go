package models

type ShopRentDashboardDayWise struct {
	UniqueID           string                       `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                       `json:"status" bson:"status,omitempty"`
	Demand             DashBoardShopRentDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardShopRentCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardShopRentPending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardShopRentOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type ShopRentDashboardDayWiseFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefShopRentDashboardDayWise struct {
	ShopRentDashboardDayWise `bson:",inline"`
}
