package models

type DashboardQueryCountFilter struct {
	QueryFilter `bson:",inline"`
}
type DashboardQueryCountReport struct {
	UnresolvedQueries int `json:"unresolvedQueries" bson:"unresolvedQueries,omitempty"`
	Assinged          int `json:"assinged" bson:"assinged,omitempty"`
	ResolvedQueries   int `json:"resolvedQueries" bson:"resolvedQueries,omitempty"`
}
type DayWiseQueryDemandChartReport struct {
	//ID   bool `json:"id" bson:"_id,omitempty"`
	Days []struct {
		ID   int `json:"id" bson:"_id,omitempty"`
		Data struct {
			UnresolvedQueries float64 `json:"unresolvedQueries" bson:"U,omitempty"`
			AssingedQueries   float64 `json:"assingedQueries" bson:"O,omitempty"`
			ResolvedQueries   float64 `json:"resolvedQueries" bson:"R,omitempty"`
		} `json:"data" bson:"data"`
	} `json:"days,omitempty" bson:"days"`
}
