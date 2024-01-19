package models

//TradeLicenseDemandPart2Filter : ""
type TradeLicenseDemandPart2Filter struct {
	TradeLicenseID string   `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	By             string   `json:"by" bson:"by,omitempty"`
	ByType         string   `json:"byType" bson:"byType,omitempty"`
	FyID           []string `json:"fyId" bson:"fyId,omitempty"`
}

//RefTradeLicenseDemandPart2 : ""
type RefTradeLicenseDemandPart2 struct {
	RefTradeLicense `bson:",inline"`
	TotalDemand     float64                `json:"totalDemand" bson:"totalDemand,omitempty"`
	FYs             []TradeLicenseFYDemand `json:"fys" bson:"fys,omitempty"`
}

func (resTlDemand *RefTradeLicenseDemandPart2) CalcDemand() {
	var totalDemandYearWise, totalDemand float64
	var fys2 []TradeLicenseFYDemand
	if resTlDemand != nil {
		if len(resTlDemand.FYs) > 0 {
			for i, v := range resTlDemand.FYs {
				resTlDemand.FYs[i].AlreadyPaidFYTax = v.AlreadyPaied.TotalTaxAmount + v.AlreadyPaied.Rebate - v.AlreadyPaied.Penalty
				resTlDemand.FYs[i].RemainingFYTax = v.FYTax - resTlDemand.FYs[i].AlreadyPaidFYTax
				a := v.TlRebate.Rate / 100
				resTlDemand.FYs[i].FYRebateValue = v.TlRebate.Rate
				resTlDemand.FYs[i].RebateType = v.TlRebate.Type
				resTlDemand.FYs[i].FYRebate = resTlDemand.FYs[i].RemainingFYTax * a
				resTlDemand.FYs[i].PenaltyValue = v.TLPenaltyRate
				resTlDemand.FYs[i].FYAfterRebate = resTlDemand.FYs[i].RemainingFYTax - resTlDemand.FYs[i].FYRebate
				resTlDemand.FYs[i].FYTLPenalty = v.TLPenaltyRate
				if resTlDemand.FYs[i].RemainingFYTax != 0 {
					resTlDemand.FYs[i].FYTotal = resTlDemand.FYs[i].RemainingFYTax - resTlDemand.FYs[i].FYRebate + v.TLPenaltyRate
					totalDemandYearWise = resTlDemand.FYs[i].FYAfterRebate + v.TLPenaltyRate
				} else {
					resTlDemand.FYs[i].FYTotal = resTlDemand.FYs[i].RemainingFYTax - resTlDemand.FYs[i].FYRebate
					totalDemandYearWise = resTlDemand.FYs[i].FYAfterRebate
				}
				totalDemand = totalDemand + totalDemandYearWise

				// if resTlDemand.FYs[i].RemainingFYTax == 0 {
				// 	resTlDemand.FYs = append(resTlDemand.FYs[:i], resTlDemand.FYs[i+1:]...)
				// 	fmt.Println("resTlDemand.FYs ==========>", resTlDemand.FYs)
				// }
				if resTlDemand.FYs[i].RemainingFYTax != 0 {
					fys2 = append(fys2, resTlDemand.FYs[i])
				}

			}
			resTlDemand.FYs = fys2
		}

	}
	if resTlDemand != nil {
		resTlDemand.TotalDemand = totalDemand
	}
}

//TradeLicenseFYDemand
type TradeLicenseFYDemand struct {
	FinancialYear    `bson:",inline"`
	PenaltyValue     float64                        `json:"penaltyValue" bson:"penaltyValue,omitempty"`
	PenaltyType      float64                        `json:"penaltyType" bson:"penaltyType,omitempty"`
	RebateValue      float64                        `json:"rebateValue" bson:"rebateValue,omitempty"` // % or amount to be discounted depending on the rebate type
	RebateType       string                         `json:"rebateType" bson:"rebateType,omitempty"`   // type of rebate (% or amount)
	FYTotal          float64                        `json:"fyTotal" bson:"fyTotal,omitempty"`
	FYTLPenalty      float64                        `json:"fyTlPenalty" bson:"fyTlPenalty,omitempty"`
	FYRebate         float64                        `json:"fyRebate" bson:"fyRebate,omitempty"` // FYRebate Contains the discounted amount
	FYRebateValue    float64                        `json:"fyRebateValue" bson:"fyRebateValue,omitempty"`
	FYAfterRebate    float64                        `json:"fyAfterRebate" bson:"fyAfterRebate,omitempty"` // Amount to be paied after rebate discounted
	FYTax            float64                        `json:"fyTax" bson:"fyTax,omitempty"`
	AlreadyPaidFYTax float64                        `json:"alreadyPaidFYTax" bson:"alreadyPaidFYTax,omitempty"`
	RemainingFYTax   float64                        `json:"remainingFYTax" bson:"remainingFYTax,omitempty"`
	FYOther          float64                        `json:"fyOther" bson:"fyOther,omitempty"`
	TlRebate         RefTradeLicenseRebate          `json:"tlRebate" bson:"tlRebate,omitempty"`
	RateMaster       RefTradeLicenseRateMaster      `json:"rateMaster" bson:"rateMaster,omitempty"`
	AlreadyPaied     RefTradeLicensePaymentsFyPart2 `json:"alreadyPaied" bson:"alreadyPaied,omitempty"`
}

//RefTradeLicenseDemandCal : ""
type RefTradeLicenseDemandCal struct {
	TotalDemand   float64 `json:"totalDemand" bson:"totalDemand,omitempty"`
	FYTotal       float64 `json:"fyTotal" bson:"fyTotal,omitempty"`
	FYPenalty     float64 `json:"fyPenalty" bson:"fyPenalty,omitempty"`
	FYRebateValue float64 `json:"fyRebateValue" bson:"fyRebateValue,omitempty"`
}

//RefTradeLicensePaymentsFyPart2 : ""
type RefTradeLicensePaymentsFyPart2 struct {
	Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
	TotalTaxAmount float64 `json:"totalTaxAmount" bson:"totalTaxAmount,omitempty"`
	Tax            float64 `json:"tax" bson:"tax,omitempty"`
	Other          float64 `json:"other" bson:"other,omitempty"`
	PenaltyValue   float64 `json:"penaltyValue" bson:"penaltyValue,omitempty"`
	PenaltyType    float64 `json:"penaltyType" bson:"penaltyType,omitempty"`
	RebateValue    float64 `json:"rebateValue" bson:"rebateValue,omitempty"`
	RebateType     float64 `json:"rebateType" bson:"rebateType,omitempty"`
	Rebate         float64 `json:"rebate" bson:"rebate,omitempty"`
}
