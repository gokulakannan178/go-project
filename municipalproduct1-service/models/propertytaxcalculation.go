package models

import "time"

//PropertyTaxCalculation : "property tax calculation response"
type PropertyTaxCalculation struct {
	UniqueID           string     `json:"uniqueId"`
	AreaOfPlot         int        `json:"areaOfPlot"`
	BuildUpArea        int        `json:"buildUpArea"`
	RoadTypeID         string     `json:"roadTypeId"`
	MunicipalTypeID    string     `json:"municipalTypeId"`
	Doa                *time.Time `json:"doa"`
	Status             string     `json:"status"`
	PropertyConfig     `json:"propertyConfig"`
	PercentAreaBuildup float64 `json:"percentAreaBuildup"`
	TaxableVacantLand  float64 `json:"taxableVacantLand"`
	Fys                []Fys   `json:"fys"`
}

//PropertyConfig : "property configuration"
type PropertyConfig struct {
	UniqueID                 string  `json:"uniqueId"`
	TaxableVacantLandConfig  float64 `json:"taxableVacantLandConfig"`
	VacantLandRatePercentage int     `json:"vacantLandRatePercentage"`
}

//Fys : "financial years"
type Fys struct {
	UniqueID      string     `json:"uniqueId"`
	Name          string     `json:"name"`
	From          *time.Time `json:"from"`
	To            *time.Time `json:"to"`
	Status        string     `json:"status"`
	Created       Created    `json:"created"`
	Vlr           Vlr        `json:"vlr"`
	VacantLandTax float64    `json:"vacantLandTax"`
}

//Vlr : "vacant land rate"
type Vlr struct {
	UniqueID           string     `json:"uniqueId"`
	Name               string     `json:"name"`
	Desc               string     `json:"desc"`
	Status             string     `json:"status"`
	MunicipalityTypeID string     `json:"municipalityTypeId"`
	RoadTypeID         string     `json:"roadTypeId"`
	Rate               float64    `json:"rate"`
	RateType           string     `json:"rateType"`
	Doe                *time.Time `json:"doe"`
	Created            Created    `json:"created"`
}
