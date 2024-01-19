package models

type DashboardContentCountFilter struct {
	ContentFilter `bson:",inline"`
}
type DashboardContentCountReport struct {
	UnReviewed int `json:"unReviewed" bson:"unReviewed,omitempty"`
	Reviewed   int `json:"reviewed" bson:"reviewed,omitempty"`
	Rejected   int `json:"rejected" bson:"rejected,omitempty"`
	Deleted    int `json:"deleted" bson:"deleted,omitempty"`
}
type DayWiseContentDemandChartReport struct {
	//ID   bool `json:"id" bson:"_id,omitempty"`
	Days []struct {
		ID   int `json:"id" bson:"_id,omitempty"`
		Data struct {
			Approved   float64 `json:"approved" bson:"A,omitempty"`
			UnReviewed float64 `json:"unReviewed" bson:"U,omitempty"`
			Rejected   float64 `json:"rejected" bson:"R,omitempty"`
			Deleted    float64 `json:"deleted" bson:"Deleted,omitempty"`
		} `json:"data" bson:"data"`
	} `json:"days,omitempty" bson:"days"`
}
