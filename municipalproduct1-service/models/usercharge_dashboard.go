package models

type UserChargeSAFDashboard struct {
	Pending struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"pending,omitempty" bson:"pending,omitempty"`
	Active struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"active,omitempty" bson:"active,omitempty"`
	Rejected struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"rejected,omitempty" bson:"rejected,omitempty"`
	Deleted struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"deleted,omitempty" bson:"deleted,omitempty"`
}

type GetUserChargeSAFDashboardFilter struct {
}
