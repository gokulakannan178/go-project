package models

import (
	"fmt"
	"municipalproduct1-service/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type MobileTowerDemand struct {
	RefPropertyMobileTower `bson:",inline"`
	Property               RefProperty                 `json:"property" bson:"property,omitempty"`
	FY                     []RefMobileTowerDemandFYLog `json:"fy" bson:"fy,omitempty"`
	ProductConfiguration   *RefProductConfiguration    `json:"-" bson:"productConfiguration,omitempty"`
	IsRegPaid              int                         `json:"isRegPaid" bson:"isRegPaid,omitempty"`
	UnPaid                 float64                     `json:"unPaid" bson:"unPaid,omitempty"`
}

type MobileTowerCalcQueryFilter struct {
	MobileTowerID  string   `json:"mobileTowerId" bson:"mobileTowerId,omitempty"`
	OmitFy         []string `json:"omitFy" bson:"omitFy,omitempty"`
	AddFy          []string `json:"addFy" bson:"addFy,omitempty"`
	OmitPayedYears bool     `json:"omitPayedYears" bson:"omitPayedYears4,omitempty"`
}

func (mtd *MobileTowerDemand) CalcQuery(filter *MobileTowerCalcQueryFilter) ([]bson.M, error) {
	var mainPipeline []bson.M
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": mtd.RefPropertyMobileTower.UniqueID}})

	//Finding Property
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "properties",
		"as":   "property",
		"let":  bson.M{"propertyId": "$propertyId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$uniqueId", "$$propertyId"}},
			}}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONMOBILETOWERPAYMENTS,
		"as":   "isRegPaid",
		"let":  bson.M{"mobiletowerId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$mobileTowerId", "$$mobiletowerId"}},
				bson.M{"$eq": []string{"$scenario", constants.MOBILETOWERPAYMENTREGISTRATIONPAYMENT}},
				bson.M{"$eq": []string{"$status", constants.MOBILETOWERPAYMENRSTATUSCOMPLETED}},
			}}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"isRegPaid": bson.M{"$size": "$isRegPaid"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"property": bson.M{"$arrayElemAt": []interface{}{"$property", 0}}}})

	//Geting COmpleted financial year ids
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONMOBILETOWERPAYMENTSFY,
		"as":   "completedFys",
		"let":  bson.M{"mobileTowerId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$mobileTowerId", "$$mobileTowerId"}},
				bson.M{"$eq": []string{"$status", constants.MOBILETOWERPAYMENRSTATUSCOMPLETED}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "ids": bson.M{"$push": "$fy.uniqueId"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"completedFys": bson.M{"$arrayElemAt": []interface{}{"$completedFys", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"completedFys.ids": bson.M{"$cond": bson.M{"if": bson.M{"$not": []interface{}{"$completedFys.ids"}}, "then": []interface{}{}, "else": "$completedFys.ids"}}}})
	// financial; year calculation
	fyQuery := []bson.M{}
	fyQuery = append(fyQuery, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
		bson.M{"$gte": []string{"$status", "Active"}},
		bson.M{"$gte": []string{"$to", "$$mtDateFrom"}},
		bson.M{"$eq": []interface{}{true, bson.M{"$cond": bson.M{"if": bson.M{"$not": []string{"$$mtDateTo"}}, "then": true, "else": bson.M{
			"$cond": bson.M{"if": bson.M{"$gte": []string{"$$mtDateTo", "$from"}}, "then": true, "else": false},
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
	/*****************/
	if filter.OmitPayedYears {
		fmt.Println("omiting paid financial years")

		fyQuery = append(fyQuery, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$cond": bson.M{"if": bson.M{"$in": []interface{}{"$uniqueId", "$$completedFys"}}, "then": false, "else": true}},
		}}}})
	}

	/****************/
	fyQuery = append(fyQuery, bson.M{"$sort": bson.M{"order": 1}})
	fyQuery = append(fyQuery, bson.M{
		"$lookup": bson.M{
			"from": "mobiletowertaxs",
			"as":   "ref.mobileTowerTax",
			"let":  bson.M{"fyFrom": "$from", "fyTo": "$to"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$gte": []string{"$status", "Active"}},
					bson.M{"$lte": []string{"$doe", "$$fyTo"}},
				}}}},
				bson.M{"$sort": bson.M{"doe": -1}},
			},
		},
	})
	fyQuery = append(fyQuery, bson.M{"$addFields": bson.M{"ref.mobileTowerTax": bson.M{"$arrayElemAt": []interface{}{"$ref.mobileTowerTax", 0}}}})
	//Finding Demand For Financial Years
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from":     "financialyears",
			"as":       "fy",
			"let":      bson.M{"mtDateFrom": "$dateFrom", "mtDateTo": "$dateTo", "completedFys": "$completedFys.ids"},
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
func (mtd *MobileTowerDemand) CalcDemand() error {
	mtd.PropertyMobileTower.Demand = MobileTowerTotalDemand{}
	ct := time.Now()
	if mtd != nil {
		for k, _ := range mtd.FY {
			mtd.FY[k].Details.Tax = mtd.FY[k].Ref.MobileTowerTax.Value
			if !mtd.FY[k].FinancialYear.IsCurrent {
				//Current Year
				mtd.FY[k].Details.Penalty = ((mtd.FY[k].Details.Tax * mtd.FY[k].FinancialYear.MobileTowerPenalty) / 100)
			} else {
				//Arrear Years
				d := *mtd.FY[k].FinancialYear.LastDate
				if ct.After(d) {
					// Find total no of months and multiply 1.5 % of tax  * months
				}
			}
			mtd.FY[k].Details.TotalTaxAmount = mtd.FY[k].Details.TotalTaxAmount + mtd.FY[k].Details.Tax + mtd.FY[k].Details.Penalty + mtd.FY[k].Details.Other
			if mtd.FY[k].FinancialYear.IsCurrent {
				mtd.Demand.Current.Tax = mtd.Demand.Current.Tax + mtd.FY[k].Details.Tax
				mtd.Demand.Current.Penalty = mtd.Demand.Current.Penalty + mtd.FY[k].Details.Penalty
				mtd.Demand.Current.Other = mtd.Demand.Current.Other + mtd.FY[k].Details.Other
				mtd.Demand.Current.Total = mtd.Demand.Current.Total + mtd.FY[k].Details.TotalTaxAmount
			} else {
				mtd.Demand.Arrear.Tax = mtd.Demand.Arrear.Tax + mtd.FY[k].Details.Tax
				mtd.Demand.Arrear.Penalty = mtd.Demand.Arrear.Penalty + mtd.FY[k].Details.Penalty
				mtd.Demand.Arrear.Other = mtd.Demand.Arrear.Other + mtd.FY[k].Details.Other
				mtd.Demand.Arrear.Total = mtd.Demand.Arrear.Total + mtd.FY[k].Details.TotalTaxAmount
			}
		}
		mtd.Demand.Total.Tax = mtd.Demand.Total.Tax + mtd.Demand.Arrear.Tax + mtd.Demand.Current.Tax
		mtd.Demand.Total.Penalty = mtd.Demand.Total.Penalty + mtd.Demand.Arrear.Penalty + mtd.Demand.Current.Penalty
		mtd.Demand.Total.Other = mtd.Demand.Total.Other + mtd.Demand.Arrear.Other + mtd.Demand.Current.Other
		mtd.Demand.Total.Total = mtd.Demand.Total.Total + mtd.Demand.Arrear.Total + mtd.Demand.Current.Total
	}
	return nil
}
func (mtd *MobileTowerDemand) CalcCollection(mtps []RefMobileTowerPayments) error {
	for _, v := range mtps {
		fmt.Println("inside MTPS")
		mtd.Collections.Current.Tax = mtd.Collections.Current.Tax + v.Demand.Current.Tax
		mtd.Collections.Current.Penalty = mtd.Collections.Current.Penalty + v.Demand.Current.Penalty
		mtd.Collections.Current.Total = mtd.Collections.Current.Total + v.Demand.Current.Total
		mtd.Collections.Current.Other = mtd.Collections.Current.Other + v.Demand.Current.Other

		mtd.Collections.Arrear.Tax = mtd.Collections.Arrear.Tax + v.Demand.Arrear.Tax
		mtd.Collections.Arrear.Penalty = mtd.Collections.Arrear.Penalty + v.Demand.Arrear.Penalty
		mtd.Collections.Arrear.Total = mtd.Collections.Arrear.Total + v.Demand.Arrear.Total
		mtd.Collections.Arrear.Other = mtd.Collections.Arrear.Other + v.Demand.Arrear.Other

		mtd.Collections.Total.Tax = mtd.Collections.Total.Tax + v.Demand.Total.Tax
		mtd.Collections.Total.Penalty = mtd.Collections.Total.Penalty + v.Demand.Total.Penalty
		mtd.Collections.Total.Total = mtd.Collections.Total.Total + v.Demand.Total.Total
		mtd.Collections.Total.Other = mtd.Collections.Total.Other + v.Demand.Total.Other
	}
	fmt.Println("inside MTPS", mtd.Collections)

	return nil
}
func (mtd *MobileTowerDemand) CalcCollectionQuery() ([]bson.M, error) {
	var mainPipeline []bson.M
	//Fetch The Mobile Tower Payment
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"mobileTowerId": mtd.PropertyMobileTower.UniqueID,
			"status":        constants.MOBILETOWERPAYMENRSTATUSCOMPLETED,
		}})

	return mainPipeline, nil
}
func (mtd *MobileTowerDemand) CalcPendingCollection(mtps []RefMobileTowerPayments) error {
	for _, v := range mtps {
		mtd.PendingCollections.Current.Tax = mtd.PendingCollections.Current.Tax + v.Demand.Current.Tax
		mtd.PendingCollections.Current.Penalty = mtd.PendingCollections.Current.Penalty + v.Demand.Current.Penalty
		mtd.PendingCollections.Current.Total = mtd.PendingCollections.Current.Total + v.Demand.Current.Total
		mtd.PendingCollections.Current.Other = mtd.PendingCollections.Current.Other + v.Demand.Current.Other
	}
	return nil
}
func (mtd *MobileTowerDemand) CalcPendingCollectionQuery() ([]bson.M, error) {
	var mainPipeline []bson.M
	//Fetch The Mobile Tower Payment
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{
			"mobileTowerId": mtd.PropertyMobileTower.UniqueID,
			"status":        constants.MOBILETOWERPAYMENRSTATUSPENDING,
		}})

	return mainPipeline, nil
}
func (mtd *MobileTowerDemand) CalcOutStanding() error {
	mtd.OutStanding.Current.Tax = mtd.Demand.Current.Tax - mtd.Collections.Current.Tax
	mtd.OutStanding.Current.Penalty = mtd.Demand.Current.Penalty - mtd.Collections.Current.Penalty
	mtd.OutStanding.Current.Total = mtd.Demand.Current.Total - mtd.Collections.Current.Total
	mtd.OutStanding.Current.Other = mtd.Demand.Current.Other - mtd.Collections.Current.Other

	mtd.OutStanding.Arrear.Tax = mtd.Demand.Arrear.Tax - mtd.Collections.Arrear.Tax
	mtd.OutStanding.Arrear.Penalty = mtd.Demand.Arrear.Penalty - mtd.Collections.Arrear.Penalty
	mtd.OutStanding.Arrear.Total = mtd.Demand.Arrear.Total - mtd.Collections.Arrear.Total
	mtd.OutStanding.Arrear.Other = mtd.Demand.Arrear.Other - mtd.Collections.Arrear.Other

	mtd.OutStanding.Total.Tax = mtd.Demand.Total.Tax - mtd.Collections.Total.Tax
	mtd.OutStanding.Total.Penalty = mtd.Demand.Total.Penalty - mtd.Collections.Total.Penalty
	mtd.OutStanding.Total.Total = mtd.Demand.Total.Total - mtd.Collections.Total.Total
	mtd.OutStanding.Total.Other = mtd.Demand.Total.Other - mtd.Collections.Total.Other
	return nil
}
