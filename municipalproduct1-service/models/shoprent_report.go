package models

import "time"

type WardWiseShoprentReportFilter struct {
	ZoneCode  []string   `json:"zoneCode,omitempty" bson:"zoneCode,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   *time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

type WardWiseShoprentReport struct {
	Ward             `bson:",inline"`
	ShoprentPayments struct {
		Properties int64   `json:"properties" bson:"properties,omitempty"`
		Payments   float64 `json:"payments" bson:"payments,omitempty"`
	} `json:"propertypayments,omitempty" bson:"propertypayments,omitempty"`

	Properties struct {
		Count int64 `json:"count" bson:"count,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
}

// WardDayWiseShopRentCollectionReportFilter : ""
type WardDayWiseShopRentCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWiseShopRentCollectionReport : ""
type WardDayWiseShopRentCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWiseShopRentCollectionReport) Inc(a int) int {
	return a + 1
}

type WardMonthWiseShopRentCollectionReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

// WardDayWiseShopRentCollectionReport : ""
type WardMonthWiseShopRentCollectionReport struct {
	Ward   `bson:",inline"`
	Report struct {
		TotalNoProperties float64 `json:"totalNoProperties" bson:"totalNoProperties,omitempty"`
		TotalNoPayments   float64 `json:"totalNoPayments" bson:"totalNoPayments,omitempty"`
		TotalCollections  float64 `json:"totalCollections" bson:"totalCollections,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWiseShopRentCollectionReport) Inc(a int) int {
	return a + 1
}

// WardDayWiseShopRentDemandReportFilter : ""
type WardDayWiseShopRentDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

//  WardDayWiseShopRentDemandReport : ""
type WardDayWiseShopRentDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		ShopRents   int64   `json:"shoprents" bson:"shoprents,omitempty"`
		TotalDemand float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardDayWiseShopRentDemandReport) Inc(a int) int {
	return a + 1
}

// WardMonthWiseShopRentDemandReportFilter : ""
type WardMonthWiseShopRentDemandReportFilter struct {
	Ward      []string   `json:"ward" bson:"ward,omitempty"`
	Zone      []string   `json:"zone" bson:"zone,omitempty"`
	Date      *time.Time `json:"date" bson:"date,omitempty"`
	SortBy    string     `json:"sortBy"`
	SortOrder int        `json:"sortOrder"`
}

//  WardMonthWiseShopRentDemandReport : ""
type WardMonthWiseShopRentDemandReport struct {
	Ward   `bson:",inline"`
	Report struct {
		ShopRents   int64   `json:"shoprents" bson:"shoprents,omitempty"`
		TotalDemand float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	} `json:"report" bson:"report,omitempty"`
}

func (res *WardMonthWiseShopRentDemandReport) Inc(a int) int {
	return a + 1
}
