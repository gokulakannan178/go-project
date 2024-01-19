package models

// LeaseRentDemandFYLog : ""
type LeaseRentDemandFYLog struct {
	FinancialYear `bson:",inline"`
	LeaseRentID string `json:"leaseRentId" bson:"leaseRentId,omitempty"`
	Status        string `json:"status" bson:"status,omitempty"`
	Details       struct {
		Tax            float64 `json:"tax" bson:"tax,omitempty"`
		Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
		Rebate         float64 `json:"rebate" bson:"rebate,omitempty"`
		Other          float64 `json:"other" bson:"other,omitempty"`
		Total          float64 `json:"total,omitempty" bson:"total,omitempty"`
	} `json:"details,omitempty" bson:"details,omitempty"`
	Ref struct {
		Rate LeaseRentRateMaster `json:"rate,omitempty" bson:"rate,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//RefLeaseRentDemandFYLog :""
type RefLeaseRentDemandFYLog struct{
	LeaseRentDemandFYLog `bson:",inline"`

}

// func (mtd *LeaseRentDemandFYLog) CalcDemandQuery() ([]bson.M, error) {

// }