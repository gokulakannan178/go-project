package models

type UserChargeDashboardDayWise struct {
	UniqueID           string                         `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                         `json:"status" bson:"status,omitempty"`
	Demand             DashBoardUserChargeDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardUserChargeCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardUserChargePending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardUserChargeOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type UserChargeDashboardDayWiseFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefUserChargeDashboardDayWise struct {
	UserChargeDashboardDayWise `bson:",inline"`
}
