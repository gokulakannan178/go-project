package models

import "time"

type TradeLicenseDashboard struct {
	UniqueID           string                           `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                           `json:"status" bson:"status,omitempty"`
	Demand             DashBoardTradeLicenseDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardTradeLicenseCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardTradeLicensePending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardTradeLicenseOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type DashBoardTradeLicenseDemand struct {
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

type DashBoardTradeLicenseCollections struct {
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

type DashBoardTradeLicenseOutstanding struct {
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

type DashBoardTradeLicensePending struct {
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

type TradeLicenseDashboardFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefTradeLicenseDashboard struct {
	TradeLicenseDashboard `bson:",inline"`
}

type DashboardTradeLicenseDemandAndCollectionFilter struct {
	TradeLicenseFilter
	TodayRange struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"todayRange,omitempty" bson:"todayRange,omitempty"`
	YesterdayRange struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"yesterdayRange,omitempty" bson:"yesterdayRange,omitempty"`
}
type DashboardDayWiseTradeLicenseCollectionChartFilter struct {
	Status    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

type DashboardDayWiseTradeLicenseCollectionChart struct {
	Records []struct {
		ID                int     `json:"day" bson:"_id,omitempty"`
		TradeLicenseCount int64   `json:"tradeLicenseCount" bson:"tradeLicenseCount,omitempty"`
		Amount            float64 `json:"amount" bson:"amount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
}

type DayWiseTradeLicenseDemandChartFilter struct {
	TradeLicenseFilter
	Status    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

type DayWiseTradeLicenseDemandChart struct {
	Records []struct {
		ID                int     `json:"day" bson:"_id,omitempty"`
		TradeLicenseCount int64   `json:"tradeLicenseCount" bson:"tradeLicenseCount,omitempty"`
		Amount            float64 `json:"amount" bson:"amount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
	Total float64 `json:"total" bson:"total,omitempty"`
}

type DashboardTradeLicenseDemandAndCollection struct {
	TotalDemandArrear      float64                                   `json:"totalDemandArrear" bson:"totalDemandArrear,omitempty"`
	TotalDemandCurrent     float64                                   `json:"totalDemandCurrent" bson:"totalDemandCurrent,omitempty"`
	TotalDemandTax         float64                                   `json:"totalDemandTax" bson:"totalDemandTax,omitempty"`
	TotalCollectionArrear  float64                                   `json:"totalCollectionArrear" bson:"totalCollectionArrear,omitempty"`
	TotalCollectionCurrent float64                                   `json:"totalCollectionCurrent" bson:"totalCollectionCurrent,omitempty"`
	TotalCollectionTax     float64                                   `json:"totalCollectionTax" bson:"totalCollectionTax,omitempty"`
	SAFCount               DashBoardStatusWiseTradeLicenseCollection `json:"safCount" bson:"safCount,omitempty"`
}

type DashBoardStatusWiseTradeLicenseCollection struct {
	Active    float64 `json:"active" bson:"active,omitempty"`
	Pending   float64 `json:"pending" bson:"pending,omitempty"`
	Expired   float64 `json:"expired" bson:"expired,omitempty"`
	Disabled  float64 `json:"disabled" bson:"disabled,omitempty"`
	Rejected  float64 `json:"rejected" bson:"rejected,omitempty"`
	Today     float64 `json:"today" bson:"-"`
	Yesterday float64 `json:"yesterday" bson:"-"`
}

// WardDayWiseTradeLicenseCollectionReportFilter : ""
type WardDayWiseTradeLicenseCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWiseTradeLicenseCollectionReport : ""
type WardDayWiseTradeLicenseCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWiseTradeLicenseCollectionReport) Inc(a int) int {
	return a + 1
}

type WardMonthWiseTradeLicenseCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWiseTradeLicenseCollectionReport : ""
type WardMonthWiseTradeLicenseCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWiseTradeLicenseCollectionReport) Inc(a int) int {
	return a + 1
}

// WardDayWiseTradeLicenseDemandReportFilter : ""
type WardDayWiseTradeLicenseDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

//  WardDayWiseTradeLicenseDemandReport : ""
type WardDayWiseTradeLicenseDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TradeLicenses int64   `json:"tradeLicenses" bson:"tradeLicenses,omitempty"`
		TotalDemand   float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWiseTradeLicenseDemandReport) Inc(a int) int {
	return a + 1
}

// WardMonthWiseTradeLicenseDemandReportFilter : ""
type WardMonthWiseTradeLicenseDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

//  WardMonthWiseTradeLicenseDemandReport : ""
type WardMonthWiseTradeLicenseDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TradeLicenses int64   `json:"tradeLicenses" bson:"tradeLicenses,omitempty"`
		TotalDemand   float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWiseTradeLicenseDemandReport) Inc(a int) int {
	return a + 1
}

// TeamDayWiseTradeLicenseCollectionReportFilter : ""
type TeamDayWiseTradeLicenseCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamMonthWiseTradeLicenseCollectionReport : ""
type TeamDayWiseTradeLicenseCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoTradeLicenses float64 `json:"totalNoTradeLicenses" bson:"totalNoTradeLicenses,omitempty"`
		TotalNoPayments      float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections     float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamDayWiseTradeLicenseCollectionReport) Inc(a int) int {
	return a + 1
}

// TeamMonthWiseTradeLicenseCollectionReportFilter : ""
type TeamMonthWiseTradeLicenseCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamMonthWiseTradeLicenseCollectionReport : ""
type TeamMonthWiseTradeLicenseCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoTradeLicenses float64 `json:"totalNoTradeLicenses" bson:"totalNoTradeLicenses,omitempty"`
		TotalNoPayments      float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections     float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamMonthWiseTradeLicenseCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWiseTradeLicenseCollectionReportFilter : ""
type WardYearWiseTradeLicenseCollectionReportFilter struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward     []string `json:"ward" bson:"ward,omitempty"`
	Zone     []string `json:"zone" bson:"zone,omitempty"`
	// Date     *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// WardDayWiseTradeLicenseCollectionReport : ""
type WardYearWiseTradeLicenseCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWiseTradeLicenseCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWiseTradeLicenseDemandReportFilter : ""
type WardYearWiseTradeLicenseDemandReportFilter struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward     []string `json:"ward" bson:"ward,omitempty"`
	Zone     []string `json:"zone" bson:"zone,omitempty"`
	// Date     *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// WardYearWiseTradeLicenseDemandReport : ""

type WardYearWiseTradeLicenseDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TradeLicenses int64   `json:"tradeLicenses" bson:"tradeLicenses,omitempty"`
		TotalDemand   float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWiseTradeLicenseDemandReport) Inc(a int) int {
	return a + 1
}

// TeamYearWiseTradeLicenseCollectionReportFilter : ""
type TeamYearWiseTradeLicenseCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

//  TeamYearWiseTradeLicenseCollectionReport : ""
type TeamYearWiseTradeLicenseCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoTradeLicenses float64 `json:"totalNoTradeLicenses" bson:"totalNoTradeLicenses,omitempty"`
		TotalNoPayments      float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections     float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamYearWiseTradeLicenseCollectionReport) Inc(a int) int {
	return a + 1
}
