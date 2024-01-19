package models

import (
	"fmt"
	"municipalproduct1-service/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ShopRentMonthlyDemand struct {
	RefShopRent          `bson:",inline"`
	DOAMonth             Month                           `json:"doamonth" bson:"doamonth,omitempty"`
	DOAFY                FinancialYear                   `json:"doafy" bson:"doafy,omitempty"`
	FY                   []RefShopRentMonthlyDemandFYLog `json:"fy" bson:"fy,omitempty"`
	ProductConfiguration *RefProductConfiguration        `json:"-" bson:"productConfiguration,omitempty"`
}

type ShopRentMonthlyCalcQueryFilter struct {
	ShopRentID      string                    `json:"shopRentId" bson:"shopRentId,omitempty"`
	OmitFy          []SingleMonthIdentifierV2 `json:"omitFy" bson:"omitFy,omitempty"`
	AddFy           []SingleMonthIdentifierV2 `json:"addFy" bson:"addFy,omitempty"`
	OmitPayedMonths bool                      `json:"omitPayedMonths" bson:"omitPayedMonths,omitempty"`
}

func (demand *ShopRentMonthlyDemand) CalcQuery(filter *ShopRentMonthlyCalcQueryFilter) ([]bson.M, error) {
	var mainPipeline []bson.M
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": demand.RefShopRent.UniqueID}})
	//get datefrom month number
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"doamonth": bson.M{"$month": "$dateFrom"}},
	})
	//get datefrom month
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from":         constants.COLLECTIONMONTH,
			"as":           "doamonth",
			"localField":   "doamonth",
			"foreignField": "month",
		}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONFINANCIALYEAR,
			"as":   "doaFy",
			"let":  bson.M{"doa": "$dateFrom"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$gte": []string{"$$doa", "$from"}},
					{"$lte": []string{"$$doa", "$to"}},
					{"$eq": []string{"$status", constants.FINANCIALYEARSTATUSACTIVE}},
				}}}},
			},
		},
	})
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$addFields": bson.M{"doaFy": bson.M{"$arrayElemAt": []interface{}{"$doaFy", 0}}},
	// })
	//get oth element
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{
			"doaFy":    bson.M{"$arrayElemAt": []interface{}{"$doaFy", 0}},
			"doamonth": bson.M{"$arrayElemAt": []interface{}{"$doamonth", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"addMonths": filter.AddFy}})

	if filter.OmitPayedMonths {
		mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
			"from": constants.COLLECTIONSHOPRENTPAYMENTSFY,
			"as":   "payedYears",
			"let":  bson.M{"shopRentId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$shopRentId", "$$shopRentId"}},
					bson.M{"$eq": []string{"$status", "Completed"}},
				}}}},
				bson.M{"$unwind": "$fy.months"},
				bson.M{"$group": bson.M{
					"_id":    "$fy.uniqueId",
					"months": bson.M{"$push": "$fy.months.name"},
				}},
			},
		}})
	}
	fyQuery := []bson.M{}

	//Calculating Financial Years
	if filter != nil {
		if len(filter.AddFy) > 0 {
			fyQuery = append(fyQuery, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$in": []interface{}{"$uniqueId", "$$addMonths.fyId"}},
			}}}})
		}
	}

	fyQuery = append(fyQuery, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
		bson.M{"$gte": []string{"$status", "Active"}},
		bson.M{"$gte": []string{"$to", "$$calcDateFrom"}},
		bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$not": []string{"$$calcDateTo"}}, "then": true, "else": bson.M{
			"$cond": bson.M{"if": bson.M{"$gte": []string{"$$calcDateTo", "$from"}}, "then": true, "else": false},
		}}}}},
	}}}})
	fyQuery = append(fyQuery, bson.M{"$sort": bson.M{"order": 1}})
	avoidunwantedmonths := func() bson.M {
		return bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{
				"if":   bson.M{"$ne": []string{"$$fyId", "$$doaFy.uniqueId"}},
				"then": true,
				"else": bson.M{"$cond": bson.M{
					"if": bson.M{
						"$lt": []string{"$fyOrder", "$$doamonth.fyOrder"},
					},
					"then": false,
					"else": true,
				}}},
			}},
			},
		}}}}
	}

	getV2TaxP1 := func() bson.M {
		return bson.M{
			"$lookup": bson.M{
				"from": constants.COLLECTIONSHOPRENTINDIVIDUALRATEMASTER,
				"as":   "ref.shopRentTaxV2",
				"let":  bson.M{"endMonth": "$endMonth"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$status", "Active"}},
						bson.M{"$lte": []string{"$doe", "$$endMonth"}},

						bson.M{"$eq": []string{"$shopRentId", filter.ShopRentID}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				},
			},
		}
	}
	getV2TaxP2 := func() bson.M {
		return bson.M{"$addFields": bson.M{"ref.shopRentTaxV2": bson.M{"$arrayElemAt": []interface{}{"$ref.shopRentTaxV2", 0}}}}
	}

	if filter.OmitPayedMonths {

		fyQuery = append(fyQuery, bson.M{"$addFields": bson.M{"correspondingPaiedFy": bson.M{"$filter": bson.M{
			"input": "$$payedYears",
			"as":    "payedYears2",
			"cond":  bson.M{"$eq": []string{"$$payedYears2._id", "$uniqueId"}},
		}}}})

		fyQuery = append(fyQuery, bson.M{"$addFields": bson.M{"correspondingPaiedFy": bson.M{"$arrayElemAt": []interface{}{"$correspondingPaiedFy", 0}}}})

		fyQuery = append(fyQuery, bson.M{"$addFields": bson.M{"correspondingPaiedFy.months": bson.M{
			"$cond": bson.M{"if": bson.M{
				"$isArray": "$correspondingPaiedFy.months"},
				"then": "$correspondingPaiedFy.months",
				"else": []interface{}{}}}}})

		fyQuery = append(fyQuery, bson.M{"$lookup": bson.M{
			"from": constants.COLLECTIONMONTH,
			"as":   "months",
			"let":  bson.M{"correspondingPaiedFy": "$correspondingPaiedFy", "fyId": "$uniqueId", "doamonth": "$$doamonth", "doaFy": "$$doaFy", "fyFrom": "$from", "fyTo": "$to"},
			"pipeline": []bson.M{
				avoidunwantedmonths(),
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$not": bson.M{"$in": []string{"$name", "$$correspondingPaiedFy.months"}}},
				}}}},
				{"$sort": bson.M{"fyOrder": 1}},
				bson.M{"$addFields": bson.M{"doamonth": "$$doamonth", "doaFy": "$$doaFy", "fyFrom": "$$fyFrom"}},
				bson.M{"$addFields": bson.M{"startMonth": bson.M{

					"$dateFromParts": bson.M{"year": bson.M{"$year": "$$fyFrom"}, "month": "$month", "day": 1, "hour": 12},
				}}},
				bson.M{"$addFields": bson.M{"endMonth": bson.M{

					"$dateFromParts": bson.M{"year": bson.M{"$year": "$$fyTo"}, "month": bson.M{"$sum": []interface{}{"$month", 1}}, "day": 1, "hour": 12},
				}}},
				// bson.M{"$addFields": bson.M{"endMonth": bson.M{"$dateSubtract": bson.M{
				// 	"startDate": "$endMonth",
				// 	"unit":      "day",
				// 	"amount":    1,
				// }}}},
				bson.M{"$addFields": bson.M{"endMonth": bson.M{"$subtract": []interface{}{"$endMonth", 86400000}}}},
				getV2TaxP1(),
				getV2TaxP2(),
			},
		}})
		fyQuery = append(fyQuery, bson.M{"$addFields": bson.M{"monthsSize": bson.M{"$size": "$months"}}})
		fyQuery = append(fyQuery, bson.M{"$match": bson.M{"monthsSize": bson.M{"$gt": 0}}})

	} else {
		if len(filter.AddFy) > 0 {
			fyQuery = append(fyQuery, bson.M{"$addFields": bson.M{"correspondingFy": bson.M{"$filter": bson.M{
				"input": "$$addMonths",
				"as":    "month",
				"cond":  bson.M{"$eq": []string{"$$month.fyId", "$uniqueId"}},
			}}}})
			fyQuery = append(fyQuery, bson.M{"$addFields": bson.M{"correspondingFy": bson.M{"$arrayElemAt": []interface{}{"$correspondingFy", 0}}}})
			fyQuery = append(fyQuery, bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONMONTH,
				"as":   "months",
				"let": bson.M{"correspondingFy": "$correspondingFy",
					"monthNo": "$monthnumber", "fyId": "$uniqueId", "doamonth": "$$doamonth",
					"doaFy": "$$doaFy", "fyFrom": "$from", "fyTo": "$to",
				},
				"pipeline": []bson.M{
					avoidunwantedmonths(),
					{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						{"$in": []string{"$name", "$$correspondingFy.months"}},
					}}}},
					{"$sort": bson.M{"fyOrder": 1}},
					bson.M{"$addFields": bson.M{"doamonth": "$$doamonth", "doaFy": "$$doaFy", "fyFrom": "$$fyFrom"}},
					bson.M{"$addFields": bson.M{"startMonth": bson.M{

						"$dateFromParts": bson.M{"year": bson.M{"$year": "$$fyFrom"}, "month": "$month", "day": 1, "hour": 12},
					}}},
					bson.M{"$addFields": bson.M{"endMonth": bson.M{

						"$dateFromParts": bson.M{"year": bson.M{"$year": "$$fyTo"}, "month": bson.M{"$sum": []interface{}{"$month", 1}}, "day": 1, "hour": 12},
					}}},
					bson.M{"$addFields": bson.M{"endMonth": bson.M{"$subtract": []interface{}{"$endMonth", 86400000}}}},

					// bson.M{"$addFields": bson.M{"endMonth": bson.M{"$dateSubtract": bson.M{
					// 	"startDate": "$endMonth",
					// 	"unit":      "day",
					// 	"amount":    1,
					// }}}},
					getV2TaxP1(),
					getV2TaxP2(),
				},
			}})
		} else {
			fyQuery = append(fyQuery, bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONMONTH,
				"as":   "months",
				"let": bson.M{
					"monthNo": "$monthnumber", "fyId": "$uniqueId", "doamonth": "$$doamonth",
					"doaFy": "$$doaFy", "fyFrom": "$from", "fyTo": "$$to",
				},
				"pipeline": []bson.M{
					avoidunwantedmonths(),
					{"$sort": bson.M{"fyOrder": 1}},
					bson.M{"$addFields": bson.M{"doamonth": "$$doamonth", "doaFy": "$$doaFy", "fyFrom": "$$fyFrom"}},
					bson.M{"$addFields": bson.M{"startMonth": bson.M{

						"$dateFromParts": bson.M{"year": bson.M{"$year": "$$fyFrom"}, "month": "$month", "day": 1, "hour": 12},
					}}},
					bson.M{"$addFields": bson.M{"endMonth": bson.M{

						"$dateFromParts": bson.M{"year": bson.M{"$year": "$$fyTo"}, "month": bson.M{"$sum": []interface{}{"$month", 1}}, "day": 1, "hour": 12},
					}}},
					bson.M{"$addFields": bson.M{"endMonth": bson.M{"$subtract": []interface{}{"$endMonth", 86400000}}}},
					// bson.M{"$addFields": bson.M{"endMonth": bson.M{"$dateSubtract": bson.M{
					// 	"startDate": "$endMonth",
					// 	"unit":      "day",
					// 	"amount":    1,
					// }}}},
					getV2TaxP1(),
					getV2TaxP2(),
				},
			}})
		}
	}

	// if filter != nil {
	// 	if len(filter.OmitFy) > 0 {
	// 		fyQuery = append(fyQuery, bson.M{"$match": bson.M{"uniqueId": bson.M{"$nin": filter.OmitFy}}})
	// 	}

	// 	if len(filter.AddFy) > 0 {
	// 		fyQuery = append(fyQuery, bson.M{"$match": bson.M{"uniqueId": bson.M{"$in": filter.AddFy}}})
	// 	}
	// }
	//Find Tax for current year
	fyQuery = append(fyQuery, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONSHOPRENTRATEMASTER,
			"as":   "ref.shopRentTax",
			"let": bson.M{
				// "fyFrom": "$from",
				"fyTo": "$to", "shopCategoryId": "$$shopCategoryId",
				"shopSubCategoryId": "$$shopSubCategoryId", "id": "$uniqueId",
			},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$lte": []string{"$doe", "$$fyTo"}},
					bson.M{"$eq": []string{"$shopCategoryId", "$$shopCategoryId"}},
					bson.M{"$eq": []string{"$shopSubCategoryId", "$$shopSubCategoryId"}},
				}}}},
				bson.M{"$sort": bson.M{"doe": -1}},
			},
		},
	},
		bson.M{"$addFields": bson.M{"ref.shopRentTax": bson.M{"$arrayElemAt": []interface{}{"$ref.shopRentTax", 0}}}})

	//Finding Demand For Financial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "financialyears",
			"as":   "fy",
			"let": bson.M{"payedYears": "$payedYears", "addMonths": "$addMonths",
				"shopCategoryId": "$shopCategoryId", "shopSubCategoryId": "$shopSubCategoryId",
				"calcDateFrom": "$dateFrom", "calcDateTo": "$dateTo",
				"monthNo": "$monthnumber", "doaFy": "$doaFy", "doamonth": "$doamonth",
				"fyId": "$uniqueId",
			},
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

func (demand *ShopRentMonthlyDemand) CalcDemand() error {
	// ct := time.Now()
	if demand != nil {
		demand.ShopRent.Demand = ShopRentTotalDemand{}
		for k, v := range demand.FY {
			for k2, v2 := range v.Months {
				demand.FY[k].Months[k2].Details.Tax = demand.RefShopRent.RentAmount
				if demand.IndividualRateMaster == "Yes" {
					demand.FY[k].Months[k2].Details.Tax = v2.Ref.ShopRentTaxV2.Rate
				}
				demand.FY[k].Months[k2].Details.Penalty = 0
				demand.FY[k].Months[k2].Details.Other = 0

				// Penalty calculation
				// TODO
				demand.FY[k].Months[k2].Details.TotalTaxAmount = demand.FY[k].Months[k2].Details.TotalTaxAmount + demand.FY[k].Months[k2].Details.Tax + demand.FY[k].Months[k2].Details.Penalty + demand.FY[k].Months[k2].Details.Other
				if demand.FY[k].FinancialYear.IsCurrent {
					demand.Demand.Current.Tax = demand.Demand.Current.Tax + demand.FY[k].Months[k2].Details.Tax
					demand.Demand.Current.Penalty = demand.Demand.Current.Penalty + demand.FY[k].Months[k2].Details.Penalty
					demand.Demand.Current.Other = demand.Demand.Current.Other + demand.FY[k].Months[k2].Details.Other
					demand.Demand.Current.Total = demand.Demand.Current.Total + demand.FY[k].Months[k2].Details.TotalTaxAmount
				} else {
					demand.Demand.Arrear.Tax = demand.Demand.Arrear.Tax + demand.FY[k].Months[k2].Details.Tax
					fmt.Println("arrear penalty", demand.Demand.Arrear.Penalty, "-", demand.FY[k].Months[k2].Details.Penalty)
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
func (demand *ShopRentMonthlyDemand) CalcCollectionQuery() ([]bson.M, error) {
	var mainPipeline []bson.M
	//Fetch The Shop Rent Payments
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"shopRentId": demand.ShopRent.UniqueID,
			"status":     constants.SHOPRENTPAYMENTSTATUSCOMPLETED,
		}})

	return mainPipeline, nil
}

func (demand *ShopRentMonthlyDemand) CalcCollection(mtps []RefShopRentPayments) error {
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
func (demand *ShopRentMonthlyDemand) CalcPendingCollectionQuery() ([]bson.M, error) {
	var mainPipeline []bson.M
	//Fetch The Mobile Tower Payment
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"shopRentId": demand.ShopRent.UniqueID,
			"status":     constants.SHOPRENTPAYMENTSTATUSPENDING,
		}})

	return mainPipeline, nil
}

func (demand *ShopRentMonthlyDemand) CalcPendingCollection(mtps []RefShopRentPayments) error {
	for _, v := range mtps {
		demand.PendingCollections.Current.Tax = demand.PendingCollections.Current.Tax + v.Demand.Current.Tax
		demand.PendingCollections.Current.Penalty = demand.PendingCollections.Current.Penalty + v.Demand.Current.Penalty
		demand.PendingCollections.Current.Total = demand.PendingCollections.Current.Total + v.Demand.Current.Total
		demand.PendingCollections.Current.Other = demand.PendingCollections.Current.Other + v.Demand.Current.Other
	}
	return nil
}
func (demand *ShopRentMonthlyDemand) CalcOutStanding() error {
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

type RefShopRentMonthlyDemandFYLog struct {
	ShopRentMonthlyDemandFYLog `bson:",inline"`
}

type ShopRentMonthlyDemandFYLog struct {
	ShopRentMonthlyDemandLog `bson:",inline"`
	PropertyID               string `json:"propertyId" bson:"propertyId,omitempty"`
	ShopRentID               string `json:"shopRentId" bson:"shopRentId,omitempty"`
	//Status                   string `json:"status" bson:"status,omitempty"`
	Details struct {
		Tax            float64 `json:"tax" bson:"tax,omitempty"`
		Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
		Other          float64 `json:"other" bson:"other,omitempty"`
		TotalTaxAmount float64 `json:"totalTaxAmount,omitempty" bson:"totalTaxAmount,omitempty"`
		Months         int     `json:"months" bson:"months,omitempty"`
	} `json:"details,omitempty" bson:"details,omitempty"`
	Ref struct {
		ShopRentTax ShopRentRateMaster `json:"shopRentTax,omitempty" bson:"shopRentTax,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type ShopRentMonthlyDemandLog struct {
	FinancialYear `bson:",inline"`
	Months        []struct {
		Month      `bson:",inline"`
		DOAFY      bson.M     `json:"doaFy" bson:"doaFy,omitempty"`
		DOAMonth   bson.M     `json:"doamonth" bson:"doamonth,omitempty"`
		FYFrom     *time.Time `json:"fyFrom" bson:"fyFrom,omitempty"`
		StartMonth *time.Time `json:"startMonth" bson:"startMonth,omitempty"`
		EndMonth   *time.Time `json:"endMonth" bson:"endMonth,omitempty"`
		Details    struct {
			Tax            float64 `json:"tax" bson:"tax,omitempty"`
			Penalty        float64 `json:"penalty" bson:"penalty,omitempty"`
			Other          float64 `json:"other" bson:"other,omitempty"`
			TotalTaxAmount float64 `json:"totalTaxAmount,omitempty" bson:"totalTaxAmount,omitempty"`
		} `json:"details,omitempty" bson:"details,omitempty"`
		Ref struct {
			ShopRentTaxV2 RefShopRentindividualRateMaster `json:"shopRentTaxV2,omitempty" bson:"shopRentTaxV2,omitempty"`
		} `json:"ref,omitempty" bson:"ref,omitempty"`
	} `json:"months" bson:"months,omitempty"`
}
