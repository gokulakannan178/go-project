package models

//PropertyConfiguration : ""
type PropertyConfiguration struct {
	UniqueID                 string  `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	TaxableVacantLandConfig  float64 `json:"taxableVacantLandConfig,omitempty" bson:"taxableVacantLandConfig,omitempty"`
	VacantLandRatePercentage float64 `json:"vacantLandRatePercentage,omitempty" bson:"vacantLandRatePercentage,omitempty"`
	ServiceCharge            float64 `json:"serviceCharge,omitempty" bson:"serviceCharge,omitempty"`
}
