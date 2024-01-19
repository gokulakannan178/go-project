package models

import "time"

//DashboardTotalCollectionChart : ""
type DashboardTotalCollectionChart struct {
	Month    `bson:",inline"`
	Payments struct {
		Collection float64 `json:"collection" bson:"collection,omitempty"`
	} `json:"payments,omitempty" bson:"payments,omitempty"`
}

//DashboardTotalCollectionChartFilter : ""
type DashboardTotalCollectionChartFilter struct {
	Fy string `json:"fy,omitempty" bson:"fy,omitempty"`
}

type DashboardTotalCollectionOverviewFilter struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
}

type DashboardTotalCollectionOverview struct {
	ArrearTotalTax  float64 `json:"arrearTotalTax,omitempty" bson:"arrearTotalTax,omitempty"`
	CurrentTotalTax float64 `json:"currentTotalTax,omitempty" bson:"currentTotalTax,omitempty"`
}

type DashboardDayWiseCollectionChartFilter struct {
	Status    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

type DashboardDayWiseCollectionChart struct {
	Records []struct {
		ID struct {
			Day    int    `json:"day,omitempty" bson:"day,omitempty"`
			DayStr string `json:"dayStr,omitempty" bson:"dayStr,omitempty"`
		} `json:"id,omitempty" bson:"_id,omitempty"`
		TotalTax          float64 `json:"totalTax" bson:"totalTax,omitempty"`
		PropertyCount     int64   `json:"propertyCount" bson:"propertyCount,omitempty"`
		ArrearCollection  float64 `json:"arrearCollection" bson:"arrearCollection,omitempty"`
		CurrentCollection float64 `json:"currentCollection" bson:"currentCollection,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
		TotalCollection   float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		Date              string  `json:"date" bson:"date,omitempty"`
		ArrearPenalty     float64 `json:"arrearPenalty" bson:"arrearPenalty,omitempty"`
		CurrentPenalty    float64 `json:"currentPenalty" bson:"currentPenalty,omitempty"`
		RebateAmount      float64 `json:"rebateAmount" bson:"rebateAmount,omitempty"`
		AdvanceAmount     float64 `json:"advanceAmount" bson:"advanceAmount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
}

type WardWiseCollectionReportFilter struct {
	Zone      []string   `json:"zone,omitempty" bson:"zone,omitempty"`
	Ward      []string   `json:"ward,omitempty" bson:"ward,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

type WardWiseCollectionReport struct {
	Ward             `bson:",inline"`
	PropertyPayments struct {
		Properties int64   `json:"properties" bson:"properties,omitempty"`
		Payments   float64 `json:"payments" bson:"payments,omitempty"`
	} `json:"propertypayments,omitempty" bson:"propertypayments,omitempty"`

	Properties struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
}

type TCCollectionSummaryFilter struct {
	DateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"dateRange"`
	UserType   []string `json:"userType" bson:"userType,omitempty"`
	UserName   []string `json:"userName" bson:"userName,omitempty"`
	UserStatus []string `json:"userStatus" bson:"userStatus,omitempty"`
}
type TCCollectionSummaryReport struct {
	User             `bson:",inline"`
	Manager          User     `json:"manager" bson:"manager,omitempty"`
	UserType         UserType `json:"userType" bson:"userType,omitempty"`
	TotalAmount      float64  `json:"totalAmount" bson:"totalAmount,omitempty"`
	TotalConsumer    int64    `json:"totalConsumer" bson:"totalConsumer,omitempty"`
	PropertyPayments struct {
		PropertyCount     int64      `json:"propertyCount" bson:"propertyCount,omitempty"`
		RecentTransaction *time.Time `json:"recentTransaction" bson:"recentTransaction,omitempty"`
		Payments          float64    `json:"payments" bson:"payments,omitempty"`
	} `json:"propertypayments" bson:"propertypayments,omitempty"`
}

type TCCollectionSummaryReportV2 struct {
	TCCollectionSummaryReport []TCCollectionSummaryReport `json:"tCCollectionSummaryReport" bson:"tCCollectionSummaryReport,omitempty"`
	TotalAmount               float64                     `json:"totalAmount" bson:"totalAmount,omitempty"`
	TotalConsumer             int64                       `json:"totalConsumer" bson:"totalConsumer,omitempty"`
}

// PropertyMonthWiseCollectionReportFilter:""
type PropertyMonthWiseCollectionReportFilter struct {
	Status []string `json:"status,omitempty" bson:"status,omitempty"`
	FYID   string   `json:"fyId" bson:"fyId,omitempty"`
}

// PropertyMonthWiseCollectionReport : ""
type PropertyMonthWiseCollectionReport struct {
	PropertyPayment `bson:",inline"`

	// FinancialYear `bson:",inline"`
	Records []struct {
		ID struct {
			Month int `json:"month,omitempty" bson:"month,omitempty"`
		} `json:"id,omitempty" bson:"_id,omitempty"`
		PropertyCount         int64   `json:"propertyCount" bson:"propertyCount,omitempty"`
		CurrentAmount         float64 `json:"currentAmount" bson:"currentAmount,omitempty"`
		ArrearAmount          float64 `json:"arrearAmount" bson:"arrearAmount,omitempty"`
		TotalDetailsAmount    float64 `json:"totalDetailsAmount" bson:"totalDetailsAmount,omitempty"`
		TotalCollectionAmount float64 `json:"totalCollectionAmount" bson:"totalCollectionAmount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
}

type PropertyWiseCollectionReportFilter struct {
	Status    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
}

type PropertyWiseCollectionReport struct {
	Property `bson:",inline"`
	Ref      struct {
		Address       RefAddress         `json:"address" bson:"address,omitempty"`
		PropertyOwner []RefPropertyOwner `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		Payments      struct {
			FyFrom          string  `json:"fyFrom" bson:"fyFrom,omitempty"`
			FyTo            string  `json:"fyTo" bson:"fyTo,omitempty"`
			TotalTax        float64 `json:"totalTax" bson:"totalTax,omitempty"`
			TotalPenalty    float64 `json:"totalPenalty" bson:"totalPenalty,omitempty"`
			TotalCollection float64 `json:"totalCollection" bson:"totalCollection,omitempty"`
		} `json:"payments,omitempty" bson:"payments,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
