package models

type PmDashboardFilter struct {
	ProjectManagerID string `json:"projectManagerId,omitempty" bson:"projectManagerId,omitempty"`
}

type PmDashboard struct {
	TodaysCollectedAmount float64 `json:"todaysCollectedAmount,omitempty" bson:"todaysCollectedAmount,omitempty"`
	CurrentMonthTarget    struct {
		AcheivedAmount float64 `json:"acheivedAmount,omitempty" bson:"acheivedAmount,omitempty"`
		PendingAmount  float64 `json:"pendingAmount,omitempty" bson:"pendingAmount,omitempty"`
		TargetAmount   float64 `json:"targetAmount,omitempty" bson:"targetAmount,omitempty"`
	} `json:"currentMonthTarget,omitempty" bson:"currentMonthTarget,omitempty"`
	Overall  *DashboardDemandAndCollection `json:"overall" bson:"overall,omitempty"`
	MyAccess struct {
		TotalDemand        float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection    float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		TotalSurvey        float64 `json:"totalSurvey" bson:"totalSurvey,omitempty"`
		CollectionOfHouses float64 `json:"collectionOfHouses" bson:"collectionOfHouses,omitempty"`
	} `json:"myaccess" bson:"myaccess,omitempty"`
}
