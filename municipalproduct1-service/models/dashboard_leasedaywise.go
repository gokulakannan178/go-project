package models

type LeaseDashboardDayWise struct {
	UniqueID           string                    `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                    `json:"status" bson:"status,omitempty"`
	Demand             DashBoardLeaseDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardLeaseCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardLeasePending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardLeaseOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type LeaseDashboardDayWiseFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefLeaseDashboardDayWise struct {
	LeaseDashboardDayWise `bson:",inline"`
}
