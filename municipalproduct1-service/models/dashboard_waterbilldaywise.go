package models

type WaterBillDashboardDayWise struct {
	UniqueID           string                        `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                        `json:"status" bson:"status,omitempty"`
	Demand             DashBoardWaterbillDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardWaterbillCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardWaterbillPending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardWaterbillOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type WaterBillDashboardDayWiseFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefWaterBillDashboardDayWise struct {
	WaterBillDashboardDayWise `bson:",inline"`
}
