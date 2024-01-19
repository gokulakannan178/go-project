package models

type MobiletowerDashboardDayWise struct {
	UniqueID           string                          `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                          `json:"status" bson:"status,omitempty"`
	Demand             DashBoardMobileTowerDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardMobileTowerCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardMobileTowerPending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardMobileTowerOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type MobileTowerDashboardDayWiseFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefMobileTowerDayWiseDashboard struct {
	MobiletowerDashboardDayWise `bson:",inline"`
}
