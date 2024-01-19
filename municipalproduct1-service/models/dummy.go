package models

//AssetTestModelForInterview : ""
type AssetTestModelForInterview struct {
	InhandBalance float64 `json:"inhandbalance" bson:"inhandbalance,omitempty"`
	Todaystarget  float64 `json:"todaystarget" bson:"todaystarget,omitempty"`
	UserTypeID    string  `json:"userType" bson:"userType,omitempty"`
	Name          string  `json:"name" bson:"name,omitempty"`
	Achieved      bool    `json:"achieved" bson:"achieved,omitempty"`
	//	Status              string               `json:"status" bson:"status,omitempty"`
	YesterdayCollection float64              `json:"yesterdayCollection,omitempty"  bson:"yesterdayCollection,omitempty"`
	TodayCollection     float64              `json:"todayCollection,omitempty"  bson:"todayCollection,omitempty"`
	Array               []map[string]float64 `json:"array,omitempty"  bson:"array,omitempty"`
	ProjectWise         struct {
		TotalDemand     float64 `json:"totalDemand,omitempty"  bson:"totalDemand,omitempty"`
		TotalCollection float64 `json:"totalCollection,omitempty"  bson:"totalCollection,omitempty"`
	} `json:"projectWise,omitempty"  bson:"projectWise,omitempty"`

	Arrear struct {
		ArrearDemand     float64 `json:"demand,omitempty"  bson:"demand,omitempty"`
		ArrearCollection float64 `json:"collection,omitempty"  bson:"collection,omitempty"`
	} `json:"arrear,omitempty"  bson:"arrear,omitempty"`
}
