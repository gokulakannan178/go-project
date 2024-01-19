package models

type PropertyDemandStoredCalc struct {
	UniqueID              string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Property              string `json:"property" bson:"property,omitempty"`
	OverallPropertyDemand struct {
		Arrear     float64 `json:"arrear,omitempty" bson:"arrear,omitempty"`
		Current    float64 `json:"current,omitempty" bson:"current,omitempty"`
		BoreCharge float64 `json:"boreCharge" bson:"boreCharge"`
		FormFee    float64 `json:"formFee" bson:"formFee"`
	} `json:"overallPropertyDemand" bson:"overallPropertyDemand,omitempty"`
	Collection struct {
		Arrear         float64 `json:"arrear,omitempty" bson:"arrear,omitempty"`
		ArrearPenalty  float64 `json:"arrearPenalty,omitempty" bson:"arrearPenalty,omitempty"`
		Current        float64 `json:"current,omitempty" bson:"current,omitempty"`
		CurrentPenalty float64 `json:"currentPenalty,omitempty" bson:"currentPenalty,omitempty"`
		BoreCharge     float64 `json:"boreCharge" bson:"boreCharge"`
		FormFee        float64 `json:"formFee" bson:"formFee"`
	} `json:"collection" bson:"collection,omitempty"`
	Pending struct {
		Arrear         float64 `json:"arrear,omitempty" bson:"arrear,omitempty"`
		ArrearPenalty  float64 `json:"arrearPenalty,omitempty" bson:"arrearPenalty,omitempty"`
		Current        float64 `json:"current,omitempty" bson:"current,omitempty"`
		CurrentPenalty float64 `json:"currentPenalty,omitempty" bson:"currentPenalty,omitempty"`
		BoreCharge     float64 `json:"boreCharge" bson:"boreCharge"`
		FormFee        float64 `json:"formFee" bson:"formFee"`
	} `json:"pending" bson:"pending,omitempty"`
	Status  string  `json:"status"  bson:"status,omitempty"`
	Created Created `json:"created"  bson:"created,omitempty"`
}
type PropertyDemandFyStoredCalc struct {
	UniqueID              string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Property              string `json:"property" bson:"property,omitempty"`
	OverallPropertyDemand struct {
		Fy            string  `json:"fy" bson:"fy"`
		FloorTax      float64 `json:"floorTax" bson:"floorTax"`
		VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
		OtherDemand   float64 `json:"otherDemand" bson:"otherDemand"`
		TotalTax      float64 `json:"totalTax" bson:"totalTax"`
	} `json:"overallPropertyDemand" bson:"overallPropertyDemand,omitempty"`
	Collection struct {
		Fy            string  `json:"fy" bson:"fy"`
		FloorTax      float64 `json:"floorTax" bson:"floorTax"`
		VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
		OtherDemand   float64 `json:"otherDemand" bson:"otherDemand"`
		Penalty       float64 `json:"penalty" bson:"penalty"`
		TotalTax      float64 `json:"totalTax" bson:"totalTax"`
		TotalDemand   float64 `json:"totalDemand" bson:"totalDemand"`
	} `json:"collection" bson:"collection,omitempty"`
	Pending struct {
		Fy            string  `json:"fy" bson:"fy"`
		FloorTax      float64 `json:"floorTax" bson:"floorTax"`
		VacantLandTax float64 `json:"vacantLandTax" bson:"vacantLandTax"`
		OtherDemand   float64 `json:"otherDemand" bson:"otherDemand"`
		Penalty       float64 `json:"penalty" bson:"penalty"`
		Legacy        float64 `json:"legacy" bson:"legacy"`
		TotalTax      float64 `json:"totalTax" bson:"totalTax"`
		TotalDemand   float64 `json:"totalDemand" bson:"totalDemand"`
	} `json:"pending" bson:"pending,omitempty"`
	Status  string  `json:"status"  bson:"status,omitempty"`
	Created Created `json:"created"  bson:"created,omitempty"`
}
type StoredCalculationDemandfy struct {
	UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
	FyId       string `json:"fyId" bson:"fyId,omitempty"`
	IsCurrent  bool   `json:"isCurrent" bson:"isCurrent,omitempty"`
	Actual     struct {
		FloorTax             float64 `json:"floorTax" bson:"floorTax"`
		VacantLandTax        float64 `json:"vacantLandTax" bson:"vacantLandTax"`
		OtherDemand          float64 `json:"otherDemand" bson:"otherDemand"`
		CompositeTax         float64 `json:"compositeTax" bson:"compositeTax"`
		Ecess                float64 `json:"ecess" bson:"ecess,omitempty"`
		PenaltyChargeableTax float64 `json:"penaltyChargeableTax" bson:"penaltyChargeableTax"`
		PenaltyNonChargeable float64 `json:"penaltyNonChargeable" bson:"penaltyNonChargeable"`
		Total                float64 `json:"total" bson:"total,omitempty"`
	} `json:"actual" bson:"actual,omitempty"`
	Paid struct {
		AmountWithPenalty    float64 `json:"amountWithPenalty" bson:"amountWithPenalty"`
		AmountWithOutPenalty float64 `json:"amountWithOutPenalty" bson:"amountWithOutPenalty"`
		PenaltyChargeableTax float64 `json:"penaltyChargeableTax" bson:"penaltyChargeableTax"`
		PenaltyNonChargeable float64 `json:"penaltyNonChargeable" bson:"penaltyNonChargeable"`
		Penalty              float64 `json:"penalty" bson:"penalty"`
		PanelCharge          float64 `json:"panelCharge" bson:"panelCharge"`
		Legacy               float64 `json:"legacy" bson:"legacy"`
		Rebate               float64 `json:"rebate" bson:"rebate"`
	} `json:"paid" bson:"paid,omitempty"`
	Pending struct {
		Rebate               float64 `json:"rebate" bson:"rebate"`
		PenaltyChargeableTax float64 `json:"penaltyChargeableTax" bson:"penaltyChargeableTax"`
		PenaltyNonChargeable float64 `json:"penaltyNonChargeable" bson:"penaltyNonChargeable"`
		Penalty              float64 `json:"penalty" bson:"penalty"`
		TotalWithPenalty     float64 `json:"totalWithPenalty" bson:"totalWithPenalty"`
		TotalWithOutPenalty  float64 `json:"totalWithOutPenalty" bson:"totalWithOutPenalty"`
	} `json:"pending" bson:"pending,omitempty"`
	Status  string  `json:"status"  bson:"status,omitempty"`
	Created Created `json:"created"  bson:"created,omitempty"`
}
type StoredCalculationDemand struct {
	UniqueID   string `json:"uniqueId" bson:"uniqueId,omitempty"`
	PropertyID string `json:"propertyId" bson:"propertyId,omitempty"`
	Actual     struct {
		Arrear     float64 `json:"arrear" bson:"arrear"`
		Current    float64 `json:"current" bson:"current"`
		BoreCharge float64 `json:"boreCharge" bson:"boreCharge"`
		FormFee    float64 `json:"formFee" bson:"formFee"`
		Total      float64 `json:"total" bson:"total"`
	} `json:"actual" bson:"actual,omitempty"`
	Paid struct {
		BoreCharge float64 `json:"boreCharge" bson:"boreCharge"`
		FormFee    float64 `json:"formFee" bson:"formFee"`
		Arrear     struct {
			Amount       float64 `json:"amount" bson:"amount"`
			FYTax        float64 `json:"fyTax" bson:"fyTax"`
			VLTax        float64 `json:"vlTax" bson:"vlTax"`
			CompositeTax float64 `json:"compositeTax" bson:"compositeTax"`
			Ecess        float64 `json:"ecess" bson:"ecess"`
			Penalty      float64 `json:"penalty" bson:"penalty"`
			Rebate       float64 `json:"rebate" bson:"rebate"`
			PanelCharge  float64 `json:"panelCharge" bson:"panelCharge"`
			PaidTax      float64 `json:"paidTax" bson:"paidTax"`
			OtherDemand  float64 `json:"otherDemand" bson:"otherDemand"`
		} `json:"arrear" bson:"arrear,omitempty"`
		Current struct {
			Amount       float64 `json:"amount" bson:"amount"`
			FYTax        float64 `json:"fyTax" bson:"fyTax"`
			VLTax        float64 `json:"vlTax" bson:"vlTax"`
			CompositeTax float64 `json:"compositeTax" bson:"compositeTax"`
			Ecess        float64 `json:"ecess" bson:"ecess"`
			Penalty      float64 `json:"penalty" bson:"penalty"`
			Rebate       float64 `json:"rebate" bson:"rebate"`
			PanelCharge  float64 `json:"panelCharge" bson:"panelCharge"`
			PaidTax      float64 `json:"paidTax" bson:"paidTax"`
			OtherDemand  float64 `json:"otherDemand" bson:"otherDemand"`
		} `json:"current" bson:"current,omitempty"`
		Total struct {
			Total               float64 `json:"total" bson:"total"`
			TotalWithOutPenalty float64 `json:"totalWithTotalWithOutPenalty" bson:"totalWithTotalWithOutPenalty"`
			Penalty             float64 `json:"penalty" bson:"penalty"`
			Rebate              float64 `json:"rebate" bson:"rebate"`
			PanelCharge         float64 `json:"panelCharge" bson:"panelCharge"`
			PaidTax             float64 `json:"paidTax" bson:"paidTax"`
			OtherDemand         float64 `json:"otherDemand" bson:"otherDemand"`
		} `json:"total" bson:"total,omitempty"`
	}
	Status  string  `json:"status"  bson:"status,omitempty"`
	Created Created `json:"created"  bson:"created,omitempty"`
}
