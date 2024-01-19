package models

import (
	"fmt"
	"municipalproduct1-service/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type TradeLicenseDemand struct {
	RefTradeLicense      `bson:",inline"`
	FY                   []RefTradeLicenseDemandFYLog `json:"fy" bson:"fy,omitempty"`
	ProductConfiguration *RefProductConfiguration     `json:"-" bson:"productConfiguration,omitempty"`
}
type TradeLicenseCalcQueryFilter struct {
	TradeLicenseID string   `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	OmitFy         []string `json:"omitFy" bson:"omitFy,omitempty"`
	AddFy          []string `json:"addFy" bson:"addFy,omitempty"`
	OmitPayedYears bool     `json:"omitPayedYears" bson:"omitPayedYears4,omitempty"`
}
type RefTradeLicenseDemandFYLog struct {
	TradeLicenseDemandFYLog `bson:",inline"`
}

type TradeLicenseDemandFYLog struct {
	FinancialYear  `bson:",inline"`
	PropertyID     string `json:"propertyId" bson:"propertyId,omitempty"`
	TradeLicenseID string `json:"tradeLicenseId" bson:"tradeLicenseId,omitempty"`
	Status         string `json:"status" bson:"status,omitempty"`
	Details        struct {
		Tax            float64 `json:"tax" bson:"tax,omitempty"`
		Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
		Other          float64 `json:"other" bson:"other,omitempty"`
		TotalTaxAmount float64 `json:"totalTaxAmount,omitempty" bson:"totalTaxAmount,omitempty"`
		PenaltyValue   float64 `json:"penaltyValue" bson:"penaltyValue,omitempty"`
		PenaltyType    float64 `json:"penaltyType" bson:"penaltyType,omitempty"`
		Rebate         float64 `json:"rebate" bson:"rebate,omitempty"`
		AfterRebate    float64 `json:"afterRebate" bson:"afterRebate,omitempty"`
		RebateValue    float64 `json:"rebateValue" bson:"rebateValue,omitempty"`
		RebateType     float64 `json:"rebateType" bson:"rebateType,omitempty"`
	} `json:"details,omitempty" bson:"details,omitempty"`
	Ref struct {
		TradeLicenseTax TradeLicenseRateMaster `json:"tradeLicenseTax,omitempty" bson:"tradeLicenseTax,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (demand *TradeLicenseDemand) CalcQuery(filter *TradeLicenseCalcQueryFilter) ([]bson.M, error) {
	var mainPipeline []bson.M
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": demand.RefTradeLicense.UniqueID}})

	//Calculating Financial Years
	fyQuery := []bson.M{}
	fyQuery = append(fyQuery, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
		bson.M{"$gte": []string{"$status", "Active"}},
		bson.M{"$gte": []string{"$to", "$$calcDateFrom"}},
		bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$not": []string{"$$calcDateTo"}}, "then": true, "else": bson.M{
			"$cond": bson.M{"if": bson.M{"$gte": []string{"$$calcDateTo", "$from"}}, "then": true, "else": false},
		}}}}},
	}}}})
	if filter != nil {
		if len(filter.OmitFy) > 0 {
			fyQuery = append(fyQuery, bson.M{"$match": bson.M{"uniqueId": bson.M{"$nin": filter.OmitFy}}})
		}

		if len(filter.AddFy) > 0 {
			fyQuery = append(fyQuery, bson.M{"$match": bson.M{"uniqueId": bson.M{"$in": filter.AddFy}}})
		}
	}
	fyQuery = append(fyQuery, bson.M{"$sort": bson.M{"order": 1}})
	//Find Tax for current year
	fyQuery = append(fyQuery, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONTRADELICENSERATEMASTER,
			"as":   "ref.tradeLicenseTax",
			"let":  bson.M{"fyFrom": "$from", "fyTo": "$to", "tlbtId": "$$tlbtId", "tlctId": "$$tlctId", "id": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$gte": []string{"$status", "Active"}},
					bson.M{"$lte": []string{"$doe", "$$fyTo"}},
					bson.M{"$eq": []string{"$tlbtId", "$$tlbtId"}},
					bson.M{"$eq": []string{"$tlctId", "$$tlctId"}},
				}}}},
				bson.M{"$sort": bson.M{"doe": -1}},
			},
		},
	},
		bson.M{"$addFields": bson.M{"ref.tradeLicenseTax": bson.M{"$arrayElemAt": []interface{}{"$ref.tradeLicenseTax", 0}}}})
	//Finding Demand For Financial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from":     "financialyears",
			"as":       "fy",
			"let":      bson.M{"tlbtId": "$tlbtId", "tlctId": "$tlctId", "calcDateFrom": "$dateFrom", "calcDateTo": "$dateTo"},
			"pipeline": fyQuery,
		},
	})

	//Finding Product Conf
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from":     constants.COLLECTIONPRODUCTCONFIGURATION,
		"as":       "productConfiguration",
		"let":      bson.M{},
		"pipeline": []bson.M{},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"productConfiguration": bson.M{"$arrayElemAt": []interface{}{"$productConfiguration", 0}}}})

	return mainPipeline, nil
}

func (demand *TradeLicenseDemand) CalcDemand() error {
	demand.TradeLicense.Demand = TradeLicenseTotalDemand{}
	ct := time.Now()
	if demand != nil {
		for k, _ := range demand.FY {
			demand.FY[k].Details.Tax = demand.FY[k].Ref.TradeLicenseTax.Rate
			//demand.FY[k].Details.Tax = demand.RefTradeLicense.LicenseAmount
			fmt.Println("Licence amt", demand.RefTradeLicense.LicenseAmount)
			if !demand.FY[k].FinancialYear.IsCurrent {
				//Current Year
				demand.FY[k].Details.Penalty = ((demand.FY[k].Details.Tax * demand.FY[k].FinancialYear.MobileTowerPenalty) / 100)
			} else {
				//Arrear Years
				d := *demand.FY[k].FinancialYear.LastDate
				if ct.After(d) {
					// Find total no of months and multiply 1.5 % of tax  * months
				}
			}
			demand.FY[k].Details.TotalTaxAmount = demand.FY[k].Details.TotalTaxAmount + demand.FY[k].Details.Tax + demand.FY[k].Details.Penalty + demand.FY[k].Details.Other
			if demand.FY[k].FinancialYear.IsCurrent {
				demand.Demand.Current.Tax = demand.Demand.Current.Tax + demand.FY[k].Details.Tax
				demand.Demand.Current.Penalty = demand.Demand.Current.Penalty + demand.FY[k].Details.Penalty
				demand.Demand.Current.Other = demand.Demand.Current.Other + demand.FY[k].Details.Other
				demand.Demand.Current.Total = demand.Demand.Current.Total + demand.FY[k].Details.TotalTaxAmount
			} else {
				demand.Demand.Arrear.Tax = demand.Demand.Arrear.Tax + demand.FY[k].Details.Tax
				demand.Demand.Arrear.Penalty = demand.Demand.Arrear.Penalty + demand.FY[k].Details.Penalty
				demand.Demand.Arrear.Other = demand.Demand.Arrear.Other + demand.FY[k].Details.Other
				demand.Demand.Arrear.Total = demand.Demand.Arrear.Total + demand.FY[k].Details.TotalTaxAmount
			}
		}
		demand.Demand.Total.Tax = demand.Demand.Total.Tax + demand.Demand.Arrear.Tax + demand.Demand.Current.Tax
		demand.Demand.Total.Penalty = demand.Demand.Total.Penalty + demand.Demand.Arrear.Penalty + demand.Demand.Current.Penalty
		demand.Demand.Total.Other = demand.Demand.Total.Other + demand.Demand.Arrear.Other + demand.Demand.Current.Other
		demand.Demand.Total.Total = demand.Demand.Total.Total + demand.Demand.Arrear.Total + demand.Demand.Current.Total
	}
	return nil
}

func (demand *TradeLicenseDemand) CalcCollectionQuery() ([]bson.M, error) {
	var mainPipeline []bson.M
	//Fetch The Shop Rent Payments
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tradeLicenseId": demand.TradeLicense.UniqueID,
			"status":         constants.SHOPRENTPAYMENTSTATUSCOMPLETED,
		}})

	return mainPipeline, nil
}

func (demand *TradeLicenseDemand) CalcCollection(mtps []RefTradeLicensePayments) error {
	for _, v := range mtps {
		fmt.Println("inside MTPS")
		demand.Collections.Current.Tax = demand.Collections.Current.Tax + v.Demand.Current.Tax
		demand.Collections.Current.Penalty = demand.Collections.Current.Penalty + v.Demand.Current.Penalty
		demand.Collections.Current.Total = demand.Collections.Current.Total + v.Demand.Current.Total
		demand.Collections.Current.Other = demand.Collections.Current.Other + v.Demand.Current.Other

		demand.Collections.Arrear.Tax = demand.Collections.Arrear.Tax + v.Demand.Arrear.Tax
		demand.Collections.Arrear.Penalty = demand.Collections.Arrear.Penalty + v.Demand.Arrear.Penalty
		demand.Collections.Arrear.Total = demand.Collections.Arrear.Total + v.Demand.Arrear.Total
		demand.Collections.Arrear.Other = demand.Collections.Arrear.Other + v.Demand.Arrear.Other

		demand.Collections.Total.Tax = demand.Collections.Total.Tax + v.Demand.Total.Tax
		demand.Collections.Total.Penalty = demand.Collections.Total.Penalty + v.Demand.Total.Penalty
		demand.Collections.Total.Total = demand.Collections.Total.Total + v.Demand.Total.Total
		demand.Collections.Total.Other = demand.Collections.Total.Other + v.Demand.Total.Other
	}
	fmt.Println("inside MTPS", demand.Collections)

	return nil
}
func (demand *TradeLicenseDemand) CalcPendingCollectionQuery() ([]bson.M, error) {
	var mainPipeline []bson.M
	//Fetch The Mobile Tower Payment
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"tradeLicenseId": demand.TradeLicense.UniqueID,
			"status":         constants.SHOPRENTPAYMENTSTATUSPENDING,
		}})

	return mainPipeline, nil
}

func (demand *TradeLicenseDemand) CalcPendingCollection(mtps []RefTradeLicensePayments) error {
	for _, v := range mtps {
		demand.PendingCollections.Current.Tax = demand.PendingCollections.Current.Tax + v.Demand.Current.Tax
		demand.PendingCollections.Current.Penalty = demand.PendingCollections.Current.Penalty + v.Demand.Current.Penalty
		demand.PendingCollections.Current.Total = demand.PendingCollections.Current.Total + v.Demand.Current.Total
		demand.PendingCollections.Current.Other = demand.PendingCollections.Current.Other + v.Demand.Current.Other
	}
	return nil
}
func (demand *TradeLicenseDemand) CalcOutStanding() error {
	// demand.OutStanding.Current.Tax = demand.Demand.Current.Tax - demand.Collections.Current.Tax
	// demand.OutStanding.Current.Penalty = demand.Demand.Current.Penalty - demand.Collections.Current.Penalty
	// demand.OutStanding.Current.Total = demand.Demand.Current.Total - demand.Collections.Current.Total
	// demand.OutStanding.Current.Other = demand.Demand.Current.Other - demand.Collections.Current.Other

	// demand.OutStanding.Arrear.Tax = demand.Demand.Arrear.Tax - demand.Collections.Arrear.Tax
	// demand.OutStanding.Arrear.Penalty = demand.Demand.Arrear.Penalty - demand.Collections.Arrear.Penalty
	// demand.OutStanding.Arrear.Total = demand.Demand.Arrear.Total - demand.Collections.Arrear.Total
	// demand.OutStanding.Arrear.Other = demand.Demand.Arrear.Other - demand.Collections.Arrear.Other

	// demand.OutStanding.Total.Tax = demand.Demand.Total.Tax - demand.Collections.Total.Tax
	// demand.OutStanding.Total.Penalty = demand.Demand.Total.Penalty - demand.Collections.Total.Penalty
	// demand.OutStanding.Total.Total = demand.Demand.Total.Total - demand.Collections.Total.Total
	// demand.OutStanding.Total.Other = demand.Demand.Total.Other - demand.Collections.Total.Other

	demand.OutStanding.Current.Tax = demand.Demand.Current.Tax
	demand.OutStanding.Current.Penalty = demand.Demand.Current.Penalty
	demand.OutStanding.Current.Total = demand.Demand.Current.Total
	demand.OutStanding.Current.Other = demand.Demand.Current.Other

	demand.OutStanding.Arrear.Tax = demand.Demand.Arrear.Tax
	demand.OutStanding.Arrear.Penalty = demand.Demand.Arrear.Penalty
	demand.OutStanding.Arrear.Total = demand.Demand.Arrear.Total
	demand.OutStanding.Arrear.Other = demand.Demand.Arrear.Other

	demand.OutStanding.Total.Tax = demand.Demand.Total.Tax
	demand.OutStanding.Total.Penalty = demand.Demand.Total.Penalty
	demand.OutStanding.Total.Total = demand.Demand.Total.Total
	demand.OutStanding.Total.Other = demand.Demand.Total.Other
	return nil
}
