package models

import "time"

type PropertyDashboardDayWise struct {
	UniqueID           string                       `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                       `json:"status" bson:"status,omitempty"`
	Demand             DashBoardPropertyDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardPropertyCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardPropertyPending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardPropertyOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
	Date               *time.Time                   `json:"date" bson:"date,omitempty"`
}

// PropertyDashboardFilter : ""
type PropertyDashboardDayWiseFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefPropertyDashboardDayWise struct {
	PropertyDashboardDayWise `bson:",inline"`
}
