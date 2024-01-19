package models

import "time"

type PropertyDashBoard struct {
	UniqueID           string                       `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status             string                       `json:"status" bson:"status,omitempty"`
	Demand             DashBoardPropertyDemand      `json:"demand" bson:"demand,omitempty"`
	Collections        DashBoardPropertyCollections `json:"collection" bson:"collection,omitempty"`
	PendingCollections DashBoardPropertyPending     `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding        DashBoardPropertyOutstanding `json:"outstanding" bson:"outstanding,omitempty"`
}

type DashBoardPropertyDemand struct {
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

type DashBoardPropertyCollections struct {
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

type DashBoardPropertyOutstanding struct {
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

type DashBoardPropertyPending struct {
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

// PropertyDashboardFilter : ""
type PropertyDashBoardFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefPropertyDashBoard struct {
	PropertyDashBoard `bson:",inline"`
}

// WardDayWisePropertyCollectionReportFilter : ""
type WardDayWisePropertyCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWisePropertyCollectionReport : ""
type WardDayWisePropertyCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWisePropertyCollectionReport) Inc(a int) int {
	return a + 1
}

// WardDayWisePropertyCollectionReportFilter : ""
type WardMonthWisePropertyCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWisePropertyCollectionReport : ""
type WardMonthWisePropertyCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWisePropertyCollectionReport) Inc(a int) int {
	return a + 1
}

// WardDayWisePropertyDemandReportFilter : ""
type WardDayWisePropertyDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWisePropertyDemandReport : ""
type WardDayWisePropertyDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWisePropertyDemandReport) Inc(a int) int {
	return a + 1
}

// WardMonthWisePropertyDemandReportFilter : ""
type WardMonthWisePropertyDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardMonthWisePropertyDemandReport : ""
type WardMonthWisePropertyDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWisePropertyDemandReport) Inc(a int) int {
	return a + 1
}

// TeamDayWisePropertyCollectionReportFilter : ""
type TeamDayWisePropertyCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamDayWisePropertyCollectionReport : ""
type TeamDayWisePropertyCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamDayWisePropertyCollectionReport) Inc(a int) int {
	return a + 1
}

// TeamMonthWisePropertyCollectionReportFilter : ""
type TeamMonthWisePropertyCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	Date       *time.Time `json:"date" bson:"date,omitempty"`
}

//  TeamMonthWisePropertyCollectionReport : ""
type TeamMonthWisePropertyCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamMonthWisePropertyCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWisePropertyCollectionReportFilter : ""
type WardYearWisePropertyCollectionReportFilter struct {
	UniqueID string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward     []string `json:"ward" bson:"ward,omitempty"`
	Zone     []string `json:"zone" bson:"zone,omitempty"`
	// Date     *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

// WardDayWisePropertyCollectionReport : ""
type WardYearWisePropertyCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWisePropertyCollectionReport) Inc(a int) int {
	return a + 1
}

// WardYearWisePropertyDemandReportFilter : ""
type WardYearWisePropertyDemandReportFilter struct {
	UniqueID  string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Ward      []string `json:"ward" bson:"ward,omitempty"`
	Zone      []string `json:"zone" bson:"zone,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`

	// Date     *time.Time `json:"date" bson:"date,omitempty"`
}

// WardYearWisePropertyDemandReport : ""
type WardYearWisePropertyDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalDemand       float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardYearWisePropertyDemandReport) Inc(a int) int {
	return a + 1
}

// TeamYearWisePropertyCollectionReportFilter : ""
type TeamYearWisePropertyCollectionReportFilter struct {
	UserFilter `bson:",inline"`
	UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
}

//  TeamYearWisePropertyCollectionReport : ""
type TeamYearWisePropertyCollectionReport struct {
	User   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *TeamYearWisePropertyCollectionReport) Inc(a int) int {
	return a + 1
}

type OverallDashBoard struct {
	Property struct {
		Overall struct {
			Demand struct {
				Current float64 `json:"current" bson:"current,omitempty"`
				Arrear  float64 `json:"arrear" bson:"arrear,omitempty"`
				Total   float64 `json:"total" bson:"total,omitempty"`
			} `json:"demand" bson:"demand,omitempty"`
			Collection struct {
				Current       float64 `json:"current" bson:"current,omitempty"`
				Arrear        float64 `json:"arrear" bson:"arrear,omitempty"`
				Total         float64 `json:"total" bson:"total,omitempty"`
				PropertyCount float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			} `json:"collection" bson:"collection,omitempty"`
			Survey int64 `json:"survey" bson:"survey,omitempty"`
		} `json:"overall" bson:"overall,omitempty"`
		Today struct {
			Demand struct {
				Current float64 `json:"current" bson:"current,omitempty"`
				Arrear  float64 `json:"arrear" bson:"arrear,omitempty"`
				Total   float64 `json:"total" bson:"total,omitempty"`
			} `json:"demand" bson:"demand,omitempty"`
			Collection struct {
				Current       float64 `json:"current" bson:"current,omitempty"`
				Arrear        float64 `json:"arrear" bson:"arrear,omitempty"`
				Total         float64 `json:"total" bson:"total,omitempty"`
				PropertyCount float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			} `json:"collection" bson:"collection,omitempty"`
			Survey int64 `json:"survey" bson:"survey,omitempty"`
		} `json:"today" bson:"today,omitempty"`
		Month struct {
			Demand struct {
				Current float64 `json:"current" bson:"current,omitempty"`
				Arrear  float64 `json:"arrear" bson:"arrear,omitempty"`
				Total   float64 `json:"total" bson:"total,omitempty"`
			} `json:"demand" bson:"demand,omitempty"`
			Collection struct {
				Current       float64 `json:"current" bson:"current,omitempty"`
				Arrear        float64 `json:"arrear" bson:"arrear,omitempty"`
				Total         float64 `json:"total" bson:"total,omitempty"`
				PropertyCount float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			} `json:"collection" bson:"collection,omitempty"`
			Survey int64 `json:"survey" bson:"survey,omitempty"`
		} `json:"month" bson:"month,omitempty"`
		Year struct {
			Demand struct {
				Current float64 `json:"current" bson:"current,omitempty"`
				Arrear  float64 `json:"arrear" bson:"arrear,omitempty"`
				Total   float64 `json:"total" bson:"total,omitempty"`
			} `json:"demand" bson:"demand,omitempty"`
			Collection struct {
				Current       float64 `json:"current" bson:"current,omitempty"`
				Arrear        float64 `json:"arrear" bson:"arrear,omitempty"`
				Total         float64 `json:"total" bson:"total,omitempty"`
				PropertyCount float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			} `json:"collection" bson:"collection,omitempty"`
			Survey int64 `json:"survey" bson:"survey,omitempty"`
		} `json:"year" bson:"year,omitempty"`
		Week struct {
			Demand struct {
				Current float64 `json:"current" bson:"current,omitempty"`
				Arrear  float64 `json:"arrear" bson:"arrear,omitempty"`
				Total   float64 `json:"total" bson:"total,omitempty"`
			} `json:"demand" bson:"demand,omitempty"`
			Collection struct {
				Current       float64 `json:"current" bson:"current,omitempty"`
				Arrear        float64 `json:"arrear" bson:"arrear,omitempty"`
				Total         float64 `json:"total" bson:"total,omitempty"`
				PropertyCount float64 `json:"propertyCount" bson:"propertyCount,omitempty"`
			} `json:"collection" bson:"collection,omitempty"`
			Survey int64 `json:"survey" bson:"survey,omitempty"`
		} `json:"week" bson:"week,omitempty"`
	} `json:"property" bson:"property,omitempty"`

	IsDefault bool `json:"isDefault" bson:"isDefault,omitempty"`
}
