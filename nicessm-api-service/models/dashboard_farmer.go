package models

type DashboardFarmerCountFilter struct {
	FarmerFilter `bson:",inline"`
}
type DashboardFarmerCountReport struct {
	Active   int `json:"active" bson:"active,omitempty"`
	InActive int `json:"inActive" bson:"inActive,omitempty"`
}
type DayWiseFarmerDemandChartReport struct {
	//ID   bool `json:"id" bson:"_id,omitempty"`
	Days []struct {
		ID   int `json:"id" bson:"_id,omitempty"`
		Data struct {
			Active   float64 `json:"active" bson:"Active,omitempty"`
			Disabled float64 `json:"disabled" bson:"Disabled,omitempty"`
		} `json:"data" bson:"data"`
	} `json:"days,omitempty" bson:"days"`
}
