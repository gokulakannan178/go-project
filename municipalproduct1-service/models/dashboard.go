package models

// DashboardDemandAndCollection
type DashboardDemandAndCollection struct {
	TotalDemandArrear      float64 `json:"totalDemandArrear" bson:"totalDemandArrear,omitempty"`
	TotalDemandCurrent     float64 `json:"totalDemandCurrent" bson:"totalDemandCurrent,omitempty"`
	TotalDemandTax         float64 `json:"totalDemandTax" bson:"totalDemandTax,omitempty"`
	TotalCollectionArrear  float64 `json:"totalCollectionArrear" bson:"totalCollectionArrear,omitempty"`
	TotalCollectionCurrent float64 `json:"totalCollectionCurrent" bson:"totalCollectionCurrent,omitempty"`
	TotalCollectionTax     float64 `json:"totalCollectionTax" bson:"totalCollectionTax,omitempty"`
}

type DashboardDemandAndCollectionFilter struct {
	PropertyFilter
}
