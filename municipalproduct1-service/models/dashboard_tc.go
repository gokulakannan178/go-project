package models

type TcDashboardFilter struct {
	CollectorID string `json:"collectorId,omitempty" bson:"collectorId,omitempty"`
}

type Tcdashboard struct {
	InhandBalance             float64                       `json:"inhandbalance" bson:"inhandbalance,omitempty"`
	TodaysTarget              float64                       `json:"todaysTarget" bson:"todaysTarget,omitempty"`
	TodaysCollection          float64                       `json:"todaysCollection" bson:"todaysCollection,omitempty"`
	TodaysTargetCollection    float64                       `json:"todaysTargetCollection" bson:"todaysTargetCollection,omitempty"`
	YesterdayTarget           float64                       `json:"yesterdayTarget" bson:"yesterdayTarget,omitempty"`
	YesterdayTargetCollection float64                       `json:"yesterdayTargetCollection" bson:"yesterdayTargetCollection,omitempty"`
	YesterdayCollection       float64                       `json:"yesterdayCollection" bson:"yesterdayCollection,omitempty"`
	Overall                   *DashboardDemandAndCollection `json:"overall" bson:"ovarall,omitempty"`
	MyAccess                  struct {
		TotalDemand        float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection    float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		TotalSurvey        float64 `json:"totalSurvey" bson:"totalSurvey,omitempty"`
		CollectionOfHouses float64 `json:"collectionOfHouses" bson:"collectionOfHouses,omitempty"`
	} `json:"myaccess" bson:"myaccess,omitempty"`
}
