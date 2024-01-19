package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type SolidWasteUserChargeDemand struct {
	RefSolidWasteUserCharge `bson:",inline"`
	FY                      []RefSolidWasteUserChargeDemandFYLog `json:"fy" bson:"fy,omitempty"`
	Demand                  SolidWasteUserChargeTotalDemand      `json:"demand" bson:"demand,omitempty"`
}
type RefSolidWasteUserChargeDemandFYLog struct {
	SolidWasteUserChargeDemandFYLog `bson:",inline"`
}
type SolidWasteUserChargeDemandFYLog struct {
	SolidWasteUserChargeDemandLog `bson:",inline"`
	Details                       struct {
		Tax            float64 `json:"tax" bson:"tax,omitempty"`
		Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
		Other          float64 `json:"other" bson:"other,omitempty"`
		TotalTaxAmount float64 `json:"totalTaxAmount,omitempty" bson:"totalTaxAmount,omitempty"`
		Months         int     `json:"months" bson:"months,omitempty"`
	} `json:"details,omitempty" bson:"details,omitempty"`
}

type SolidWasteUserChargeDemandLog struct {
	FinancialYear `bson:",inline"`
	Months        []struct {
		Month   `bson:",inline"`
		Rate    SolidWasteUserChargeRate `json:"rate" bson:"rate,omitempty"`
		SOM     *time.Time               `json:"som" bson:"som,omitempty"`
		EOM     *time.Time               `json:"eom" bson:"eom,omitempty"`
		Details struct {
			Tax            float64 `json:"tax" bson:"tax,omitempty"`
			Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
			Other          float64 `json:"other" bson:"other,omitempty"`
			TotalTaxAmount float64 `json:"totalTaxAmount,omitempty" bson:"totalTaxAmount,omitempty"`
			Months         int     `json:"months" bson:"months,omitempty"`
		} `json:"details,omitempty" bson:"details,omitempty"`
	} `json:"months" bson:"months,omitempty"`
}

type SolidWasteUserChargeDemandCalcQueryFilter struct {
	SolidWasteChargeID string                  `json:"solidWasteChargeId" bson:"solidWasteChargeId,omitempty"`
	OmitFy             []SingleMonthIdentifier `json:"omitFy" bson:"omitFy,omitempty"`
	AddFy              []SingleMonthIdentifier `json:"addFy" bson:"addFy,omitempty"`
	OmitPayedMonths    bool                    `json:"omitPayedMonths" bson:"omitPayedMonths,omitempty"`
}

type SolidWasteUserChargeTotalDemand struct {
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

func (demand *SolidWasteUserChargeDemand) CalcQuery(filter *SolidWasteUserChargeDemandCalcQueryFilter) ([]bson.M, error) {
	var mainPipeline []bson.M
	//select the solidwaste
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": "1"}})
	//generate datefrom and dateTo month
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dateFromMonth": bson.M{"$month": "$dateFrom"}, "dateToMonth": bson.M{"$month": "$dateTo"}}})

	//lookups for datefrom and dateto months
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"from": "months", "as": "dateFromMonth", "let": bson.M{"month": "$dateFromMonth"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$$month", "$month"}},
			}}}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"from": "months", "as": "dateToMonth", "let": bson.M{"month": "$dateToMonth"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$$month", "$month"}},
			}}}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"dateFromMonth": bson.M{"$arrayElemAt": []interface{}{"$dateFromMonth", 0}},
		//    "dateFromFy":{"$arrayElemAt":["$dateFromFy",0]},
		"dateToMonth": bson.M{"$arrayElemAt": []interface{}{"$dateToMonth", 0}},
		//    "dateToFy":{"$arrayElemAt":["$dateToFy",0]}
	}})

	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from": "financialyears",
			"as":   "fy",
			"let": bson.M{
				//        "dateFromFy":"$dateFromFy","dateToFy":"$dateToFy",
				"dateFromMonth": "$dateFromMonth", "dateToMonth": "$dateToMonth",
				"dateFrom": "$dateFrom", "dateTo": "$dateTo", "swuc": "$$ROOT"},
			"pipeline": []bson.M{
				func() bson.M {
					if len(filter.AddFy) > 0 {
						var fyIds []string
						for _, v := range filter.AddFy {
							fyIds = append(fyIds, v.FYID)
						}
						return bson.M{"$match": bson.M{"uniqueId": bson.M{"$in": fyIds}}}

					}
					return bson.M{"$match": bson.M{}}
				}(),
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$lte": []string{"$$dateFrom", "$to"}},
					bson.M{"$lte": []string{"$from", "$$dateTo"}},
				}}}},
				bson.M{"$lookup": bson.M{
					"from": "months",
					"as":   "months",
					"let": bson.M{"dateFromMonth": "$$dateFromMonth", "dateToMonth": "$$dateToMonth",
						"fy": "$$ROOT", "dateFrom": "$$dateFrom", "dateTo": "$$dateTo", "swuc": "$$swuc"},
					"pipeline": []bson.M{
						// func()bson.M{

						// }(),
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							bson.M{"$gte": []interface{}{"$fyOrder",
								bson.M{
									"$cond": bson.M{"if": bson.M{"$and": []bson.M{
										bson.M{"$gte": []string{"$$dateFrom", "$$fy.from"}},
										bson.M{"$lte": []string{"$$dateFrom", "$$fy.to"}},
									}}, "then": "$$dateFromMonth.fyOrder", "else": 1},
								},
							}},
							bson.M{"$lte": []interface{}{"$fyOrder", bson.M{
								"$cond": bson.M{"if": bson.M{"$and": []bson.M{
									bson.M{"$gte": []string{"$$dateTo", "$$fy.from"}},
									bson.M{"$lte": []string{"$$dateTo", "$$fy.to"}},
								}}, "then": "$$dateToMonth.fyOrder", "else": 12},
							},
							}},
						}}}},
						bson.M{"$sort": bson.M{"fyOrder": 1}},
						bson.M{"$addFields": bson.M{
							"som": bson.M{"$dateFromParts": bson.M{
								"day":   1,
								"month": "$month",
								"year": bson.M{"$cond": bson.M{
									"if": bson.M{"$in": []interface{}{"$month", []int{1, 2, 3}}}, "then": bson.M{"$year": "$$fy.to"}, "else": bson.M{"$year": "$$fy.from"},
								}},
							}},
							"eom": bson.M{"$dateFromParts": bson.M{
								"day":   0,
								"month": bson.M{"$sum": []interface{}{"$month", 1}},
								"year": bson.M{"$cond": bson.M{
									"if": bson.M{"$in": []interface{}{"$month", []int{1, 2, 3}}}, "then": bson.M{"$year": "$$fy.to"}, "else": bson.M{"$year": "$$fy.from"},
								}},
							}},
						}},
						bson.M{"$lookup": bson.M{
							"from": "solidwasteuserchargerate", "as": "rate",
							"let": bson.M{"som": "$som", "eom": "$eom", "swuc": "$$swuc"},
							"pipeline": []bson.M{
								bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
									bson.M{"$eq": []string{"$status", "Active"}},
									bson.M{"$eq": []string{"$categoryId", "$$swuc.categoryId"}},
									bson.M{"$eq": []string{"$subCategoryId", "$$swuc.subCategoryId"}},
									//                           {"$gte":["$doe","$$som"]},
									bson.M{"$lt": []string{"$doe", "$$eom"}},
								}}}},
								bson.M{"$sort": bson.M{"doe": -1}},
							},
						}},
						bson.M{"$addFields": bson.M{
							"rate": bson.M{"$arrayElemAt": []interface{}{"$rate", 0}},
						}},
					},
				}},
			},
		}})
	return mainPipeline, nil
}

func (demand *SolidWasteUserChargeDemand) CalcDemand() error {
	if demand != nil {
		demand.Demand = SolidWasteUserChargeTotalDemand{}
		for k, v := range demand.FY {
			for k2, v2 := range v.Months {
				demand.FY[k].Months[k2].Details.Tax = v2.Rate.Rate
				demand.FY[k].Months[k2].Details.Penalty = 0
				demand.FY[k].Months[k2].Details.Other = 0
				demand.FY[k].Months[k2].Details.TotalTaxAmount = demand.FY[k].Months[k2].Details.Tax + demand.FY[k].Months[k2].Details.Penalty + demand.FY[k].Months[k2].Details.Other
				if demand.FY[k].FinancialYear.IsCurrent {
					demand.Demand.Current.Tax = demand.Demand.Current.Tax + demand.FY[k].Months[k2].Details.Tax
					demand.Demand.Current.Penalty = demand.Demand.Current.Penalty + demand.FY[k].Months[k2].Details.Penalty
					demand.Demand.Current.Other = demand.Demand.Current.Other + demand.FY[k].Months[k2].Details.Other
					demand.Demand.Current.Total = demand.Demand.Current.Total + demand.FY[k].Months[k2].Details.TotalTaxAmount
				} else {
					demand.Demand.Arrear.Tax = demand.Demand.Arrear.Tax + demand.FY[k].Months[k2].Details.Tax
					demand.Demand.Arrear.Penalty = demand.Demand.Arrear.Penalty + demand.FY[k].Months[k2].Details.Penalty
					demand.Demand.Arrear.Other = demand.Demand.Arrear.Other + demand.FY[k].Months[k2].Details.Other
					demand.Demand.Arrear.Total = demand.Demand.Arrear.Total + demand.FY[k].Months[k2].Details.TotalTaxAmount
				}
				demand.FY[k].Details.Tax = demand.FY[k].Details.Tax + demand.FY[k].Months[k2].Details.Tax
				demand.FY[k].Details.Penalty = demand.FY[k].Details.Penalty + demand.FY[k].Months[k2].Details.Penalty
				demand.FY[k].Details.Other = demand.FY[k].Details.Other + demand.FY[k].Months[k2].Details.Other
			}
			demand.FY[k].Details.TotalTaxAmount = demand.FY[k].Details.Tax + demand.FY[k].Details.Other - demand.FY[k].Details.Penalty
			demand.FY[k].Details.Months = len(demand.FY[k].Months)
		}
		demand.Demand.Total.Tax = demand.Demand.Total.Tax + demand.Demand.Arrear.Tax + demand.Demand.Current.Tax
		demand.Demand.Total.Penalty = demand.Demand.Total.Penalty + demand.Demand.Arrear.Penalty + demand.Demand.Current.Penalty
		demand.Demand.Total.Other = demand.Demand.Total.Other + demand.Demand.Arrear.Other + demand.Demand.Current.Other
		demand.Demand.Total.Total = demand.Demand.Total.Total + demand.Demand.Arrear.Total + demand.Demand.Current.Total
	}
	return nil
}
