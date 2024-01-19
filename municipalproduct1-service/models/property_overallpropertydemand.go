package models

type OverallPropertyDemand struct {
	PropertyID string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	Actual     struct {
		Arrear struct {
			VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
			Tax           float64 `json:"tax" bson:"tax"`
			TotalTax      float64 `json:"totalTax" bson:"totalTax"`
		} `json:"arrear,omitempty" bson:"arrear,omitempty"`
		Current struct {
			VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
			Tax           float64 `json:"tax" bson:"tax"`
			TotalTax      float64 `json:"totalTax" bson:"totalTax"`
		} `json:"current,omitempty" bson:"current,omitempty"`
		Total struct {
			VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
			Tax           float64 `json:"tax" bson:"tax"`
			TotalTax      float64 `json:"totalTax" bson:"totalTax"`
		} `json:"total,omitempty" bson:"total,omitempty"`
	} `json:"actual" bson:"actual"`
	Arrear struct {
		VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
		Rebate        float64 `json:"rebate" bson:"rebate"`
		Penalty       float64 `json:"penanty" bson:"penanty"`
		Tax           float64 `json:"tax" bson:"tax"`
		CompositeTax  float64 `json:"compositeTax" bson:"compositeTax"`
		Ecess         float64 `json:"ecess" bson:"ecess"`
		PanelCh       float64 `json:"panelCh" bson:"panelCh"`

		TotalTax float64 `json:"totalTax" bson:"totalTax"`
	} `json:"arrear,omitempty" bson:"arrear,omitempty"`
	Current struct {
		VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
		Rebate        float64 `json:"rebate" bson:"rebate"`
		Penalty       float64 `json:"penanty" bson:"penanty"`
		CompositeTax  float64 `json:"compositeTax" bson:"compositeTax"`
		Ecess         float64 `json:"ecess" bson:"ecess"`
		PanelCh       float64 `json:"panelCh" bson:"panelCh"`

		Tax      float64 `json:"tax" bson:"tax"`
		TotalTax float64 `json:"totalTax" bson:"totalTax"`
	} `json:"current,omitempty" bson:"current,omitempty"`
	Total struct {
		VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
		Rebate        float64 `json:"rebate" bson:"rebate"`
		Penalty       float64 `json:"penanty" bson:"penanty"`
		CompositeTax  float64 `json:"compositeTax" bson:"compositeTax"`
		Ecess         float64 `json:"ecess" bson:"ecess"`
		PanelCh       float64 `json:"panelCh" bson:"panelCh"`

		Tax      float64 `json:"tax" bson:"tax"`
		Other    float64 `json:"other" bson:"other"`
		TotalTax float64 `json:"totalTax" bson:"totalTax"`
	} `json:"total,omitempty" bson:"total,omitempty"`
	Other struct {
		BoreCharge  float64 `json:"boreCharge" bson:"boreCharge"`
		FormFee     float64 `json:"formFee" bson:"formFee"`
		OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
	} `json:"other,omitempty" bson:"other,omitempty"`
	NewPropertyID string `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

// RefOverallPropertyDemand : ""
type RefOverallPropertyDemand struct {
	OverallPropertyDemand `bson:",inline"`
	Ref                   struct {
	} `json:"ref" bson:"ref"`
}

// MobileTowerTaxFilter : ""
type OverallPropertyDemandFilter struct {
	PropertyID []string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	TotalTax   *struct {
		From float64 `json:"from"`
		To   float64 `json:"to"`
	} `json:"totalTax" bson:"totalTax"`
}
