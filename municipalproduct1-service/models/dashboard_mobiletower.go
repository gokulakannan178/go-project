package models

import "time"

type MobileTowerDashboard struct {
	UniqueID           string                          `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                          `json:"status" bson:"status,omitempty"`
	Demand             DashBoardMobileTowerDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardMobileTowerCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardMobileTowerPending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardMobileTowerOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type DashBoardMobileTowerDemand struct {
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

type DashBoardMobileTowerCollections struct {
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

type DashBoardMobileTowerOutstanding struct {
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

type DashBoardMobileTowerPending struct {
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

type MobileTowerDashboardFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefMobileTowerDashboard struct {
	MobileTowerDashboard `bson:",inline"`
}

type DashboardMobileTowerDemandAndCollectionFilter struct {
	PropertyMobileTowerFilter
	TodayRange struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"todayRange,omitempty" bson:"todayRange,omitempty"`
	YesterdayRange struct {
		From *time.Time `json:"from" bson:"from,omitempty"`
		To   *time.Time `json:"to" bson:"to,omitempty"`
	} `json:"yesterdayRange,omitempty" bson:"yesterdayRange,omitempty"`
}

type DashboardDayWiseMobileTowerCollectionChartFilter struct {
	Status    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

// DashboardDemandAndCollection
type DashboardMobileTowerDemandAndCollection struct {
	TotalDemandArrear      float64                                  `json:"totalDemandArrear" bson:"totalDemandArrear,omitempty"`
	TotalDemandCurrent     float64                                  `json:"totalDemandCurrent" bson:"totalDemandCurrent,omitempty"`
	TotalDemandTax         float64                                  `json:"totalDemandTax" bson:"totalDemandTax,omitempty"`
	TotalCollectionArrear  float64                                  `json:"totalCollectionArrear" bson:"totalCollectionArrear,omitempty"`
	TotalCollectionCurrent float64                                  `json:"totalCollectionCurrent" bson:"totalCollectionCurrent,omitempty"`
	TotalCollectionTax     float64                                  `json:"totalCollectionTax" bson:"totalCollectionTax,omitempty"`
	SAFCount               DashBoardStatusWiseMobileTowerCollection `json:"safCount" bson:"safCount,omitempty"`
}

// StatusWiseMobileTowerCollection : ""
type DashBoardStatusWiseMobileTowerCollection struct {
	Active    float64 `json:"active" bson:"active,omitempty"`
	Pending   float64 `json:"pending" bson:"pending,omitempty"`
	Rejected  float64 `json:"rejected" bson:"rejected,omitempty"`
	Disabled  float64 `json:"disabled" bson:"disabled,omitempty"`
	Today     float64 `json:"today" bson:"-"`
	Yesterday float64 `json:"yesterday" bson:"-"`
}

type DashboardDayWiseMobileTowerCollectionChart struct {
	Records []struct {
		ID               int     `json:"day" bson:"_id,omitempty"`
		MobileTowerCount int64   `json:"mobileTowerCount" bson:"mobileTowerCount,omitempty"`
		Amount           float64 `json:"amount" bson:"amount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
}

// DayWiseMobileTowerDemandChartFilter : ""
type DayWiseMobileTowerDemandChartFilter struct {
	PropertyMobileTowerFilter `bson:",inline"`
	Status                    []string   `json:"status,omitempty" bson:"status,omitempty"`
	StartDate                 *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
}

// DashboardDayWiseMobileTowerDemandChart : ""
type DayWiseMobileTowerDemandChart struct {
	Records []struct {
		ID               int     `json:"day" bson:"_id,omitempty"`
		MobileTowerCount int64   `json:"mobileTowerCount" bson:"mobileTowerCount,omitempty"`
		Amount           float64 `json:"amount" bson:"amount,omitempty"`
	} `json:"records,omitempty" bson:"records,omitempty"`
	Total float64 `json:"total" bson:"total,omitempty"`
}

// WardDayWiseMobileTowerDemandReportFilter : ""
type WardDayWiseMobileTowerDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

//  WardDayWiseMobileTowerDemandReport : ""
type WardDayWiseMobileTowerDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		MobileTowers int64   `json:"mobiletowers" bson:"mobiletowers,omitempty"`
		TotalDemand  float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWiseMobileTowerDemandReport) Inc(a int) int {
	return a + 1
}

// WardDayWiseMobileTowerCollectionReportFilter : ""
type WardDayWiseMobileTowerCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWiseMobileTowerCollectionReport : ""
type WardDayWiseMobileTowerCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWiseMobileTowerCollectionReport) Inc(a int) int {
	return a + 1
}

// WardDayWiseMobileTowerCollectionReportFilter : ""
type WardMonthWiseMobileTowerCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWiseMobileTowerCollectionReport : ""
type WardMonthWiseMobileTowerCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWiseMobileTowerCollectionReport) Inc(a int) int {
	return a + 1
}

// WardMonthWiseMobileTowerDemandReportFilter : ""
type WardMonthWiseMobileTowerDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

//  WardMonthWiseMobileTowerDemandReport : ""
type WardMonthWiseMobileTowerDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		MobileTowers int64   `json:"mobileTowers" bson:"mobileTowers,omitempty"`
		TotalDemand  float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWiseMobileTowerDemandReport) Inc(a int) int {
	return a + 1
}

// TeamDayWiseMobileTowerCollectionReportFilter : ""
type TeamDayWiseMobileTowerCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamDayWiseMobileTowerCollectionReport : ""
type TeamDayWiseMobileTowerCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoMobileTowers float64 `json:"totalNoMobileTowers" bson:"totalNoMobileTowers,omitempty"`
		TotalNoPayments     float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections    float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamDayWiseMobileTowerCollectionReport) Inc(a int) int {
	return a + 1
}

// TeamMonthWiseMobileTowerCollectionReportFilter : ""
type TeamMonthWiseMobileTowerCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamMonthWiseMobileTowerCollectionReport : ""
type TeamMonthWiseMobileTowerCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoMobileTowers float64 `json:"totalNoMobileTowers" bson:"totalNoMobileTowers,omitempty"`
		TotalNoPayments     float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections    float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamMonthWiseMobileTowerCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWiseMobileTowerCollectionReportFilter : ""
type WardYearWiseMobileTowerCollectionReportFilter struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward     []string `json:"ward" bson:"ward,omitempty"`
	Zone     []string `json:"zone" bson:"zone,omitempty"`
	// Date     *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// WardDayWiseMobileTowerCollectionReport : ""
type WardYearWiseMobileTowerCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWiseMobileTowerCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWiseMobileTowerDemandReportFilter : ""
type WardYearWiseMobileTowerDemandReportFilter struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward     []string `json:"ward" bson:"ward,omitempty"`
	Zone     []string `json:"zone" bson:"zone,omitempty"`
	// Date     *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// WardYearWiseMobileTowerDemandReport : ""

type WardYearWiseMobileTowerDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		MobileTowers int64   `json:"mobileTowers" bson:"mobileTowers,omitempty"`
		TotalDemand  float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWiseMobileTowerDemandReport) Inc(a int) int {
	return a + 1
}

// TeamYearWiseMobileTowerCollectionReportFilter : ""
type TeamYearWiseMobileTowerCollectionReportFilter struct {
	UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserFilter `bson:",inline"`
}

//  TeamYearWiseMobileTowerCollectionReport : ""
type TeamYearWiseMobileTowerCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoMobileTowers float64 `json:"totalNoMobileTowers" bson:"totalNoMobileTowers,omitempty"`
		TotalNoPayments     float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections    float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamYearWiseMobileTowerCollectionReport) Inc(a int) int {
	return a + 1
}
