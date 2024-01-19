package models

type TradeLicenseDashboardDayWise struct {
	UniqueID           string                           `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                           `json:"status" bson:"status,omitempty"`
	Demand             DashBoardTradeLicenseDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardTradeLicenseCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardTradeLicensePending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardTradeLicenseOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type TradeLicenseDashboardDayWiseFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefTradeLicenseDashboardDayWise struct {
	TradeLicenseDashboardDayWise `bson:",inline"`
}
