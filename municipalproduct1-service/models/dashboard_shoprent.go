package models

import "time"

type ShopRentDashboard struct {
	UniqueID           string                       `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                       `json:"status" bson:"status,omitempty"`
	Demand             DashBoardShopRentDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardShopRentCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardShopRentPending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardShopRentOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type DashBoardShopRentDemand struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type DashBoardShopRentCollections struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type DashBoardShopRentOutstanding struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type DashBoardShopRentPending struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type ShopRentDashboardFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefShopRentDashboard struct {
	ShopRentDashboard `bson:",inline"`
}

type DashboardDayWiseShoprentCollectionChartFilter struct {
	Status    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

type DashboardDayWiseShoprentCollectionChart struct {
	Records []struct {
		ID            int     `json:"day" bson:"_id,omitempty"`
		ShopRentCount int64   `json:"shopRentCount" bson:"shopRentCount,omitempty"`
		Amount        float64 `json:"amount" bson:"amount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
}

type DayWiseShoprentDemandChartFilter struct {
	ShopRentFilter `bson:",inline"`
	Status         []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

type DayWiseShoprentDemandChart struct {
	Records []struct {
		ID            int     `json:"day" bson:"_id,omitempty"`
		ShopRentCount int64   `json:"shopRentCount" bson:"shopRentCount,omitempty"`
		Amount        float64 `json:"amount" bson:"amount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
	Total float64 `json:"total" bson:"total,omitempty"`
}

type DashboardShopRentDemandAndCollectionFilter struct {
	ShopRentFilter
	TodayRange struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"todayRange,omitempty" bson:"todayRange,omitempty"`
	YesterdayRange struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"yesterdayRange,omitempty" bson:"yesterdayRange,omitempty"`
}

// DashboardDemandAndCollection
type DashboardShopRentDemandAndCollection struct {
	TotalDemandArrear      float64                               `json:"totalDemandArrear" bson:"totalDemandArrear,omitempty"`
	TotalDemandCurrent     float64                               `json:"totalDemandCurrent" bson:"totalDemandCurrent,omitempty"`
	TotalDemandTax         float64                               `json:"totalDemandTax" bson:"totalDemandTax,omitempty"`
	TotalCollectionArrear  float64                               `json:"totalCollectionArrear" bson:"totalCollectionArrear,omitempty"`
	TotalCollectionCurrent float64                               `json:"totalCollectionCurrent" bson:"totalCollectionCurrent,omitempty"`
	TotalCollectionTax     float64                               `json:"totalCollectionTax" bson:"totalCollectionTax,omitempty"`
	SAFCount               DashBoardStatusWiseShopRentCollection `json:"safCount" bson:"safCount,omitempty"`
}

// StatusWiseShopRentCollection : ""
type DashBoardStatusWiseShopRentCollection struct {
	Active    float64 `json:"active" bson:"active,omitempty"`
	Pending   float64 `json:"pending" bson:"pending,omitempty"`
	Disabled  float64 `json:"disabled" bson:"disabled,omitempty"`
	Today     float64 `json:"today" bson:"today,omitempty"`
	Yesterday float64 `json:"yesterday" bson:"yesterday,omitempty"`
}

// TeamDayWiseShopRentCollectionReportFilter : ""
type TeamDayWiseShopRentCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamMonthWiseShopRentCollectionReport : ""
type TeamDayWiseShopRentCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoShopRents float64 `json:"totalNoShopRents" bson:"totalNoShopRents,omitempty"`
		TotalNoPayments  float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamDayWiseShopRentCollectionReport) Inc(a int) int {
	return a + 1
}

// TeamMonthWiseShopRentCollectionReportFilter : ""
type TeamMonthWiseShopRentCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamMonthWiseShopRentCollectionReport : ""
type TeamMonthWiseShopRentCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoShopRents float64 `json:"totalNoShopRents" bson:"totalNoShopRents,omitempty"`
		TotalNoPayments  float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamMonthWiseShopRentCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWiseShopRentCollectionReportFilter : ""
type WardYearWiseShopRentCollectionReportFilter struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward     []string `json:"ward" bson:"ward,omitempty"`
	Zone     []string `json:"zone" bson:"zone,omitempty"`
	// Date     *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// WardDayWiseShopRentCollectionReport : ""
type WardYearWiseShopRentCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWiseShopRentCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWiseShopRentDemandReportFilter : ""
type WardYearWiseShopRentDemandReportFilter struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward     []string `json:"ward" bson:"ward,omitempty"`
	Zone     []string `json:"zone" bson:"zone,omitempty"`
	// Date     *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// WardYearWiseShopRentDemandReport : ""

type WardYearWiseShopRentDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		ShopRents   int64   `json:"shopRents" bson:"shopRents,omitempty"`
		TotalDemand float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWiseShopRentDemandReport) Inc(a int) int {
	return a + 1
}

// TeamYearWiseShopRentCollectionReportFilter : ""
type TeamYearWiseShopRentCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

//  TeamYearWiseShopRentCollectionReport : ""
type TeamYearWiseShopRentCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoShopRents float64 `json:"totalNoShopRents" bson:"totalNoShopRents,omitempty"`
		TotalNoPayments  float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamYearWiseShopRentCollectionReport) Inc(a int) int {
	return a + 1
}
