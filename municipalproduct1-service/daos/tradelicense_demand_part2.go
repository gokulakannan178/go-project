package daos

import (
	"context"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

// TradeLicenseDemandPart2 : ""
func (d *Daos) TradeLicenseDemandPart2(ctx *models.Context, filter *models.TradeLicenseDemandPart2Filter) (*models.RefTradeLicenseDemandPart2, error) {
	var mainPipeline []bson.M
	// Lookup

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": filter.TradeLicenseID}})
	// if len(filter.FyID) > 0 {
	// 	query = append(query, bson.M{"$eq": []interface{}{"$uniqueId", filter.FyID}})
	// }

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"as":   "fys",
		"from": constants.COLLECTIONFINANCIALYEAR,
		"let":  bson.M{"sd": "$dateFrom", "varTlbtId": "$tlbtId", "varTlctId": "$tlctId", "tlId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": func() bson.M {
				if len(filter.FyID) > 0 {
					return bson.M{"uniqueId": bson.M{"$in": filter.FyID}}
				}
				return bson.M{}
			}(),
			},
			bson.M{"$sort": bson.M{"to": -1}},
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": bson.M{"$lte": []interface{}{"$$sd", "$to"}}}}},
			bson.M{"$lookup": bson.M{
				"as":   "tlRebate",
				"from": constants.COLLECTIONTRADELICENSEREBATE,
				"let":  bson.M{"varTlTo": "$to"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$gte": []interface{}{"$$varTlTo", "$doe"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				},
			},
			},
			bson.M{"$addFields": bson.M{"tlRebate": bson.M{"$arrayElemAt": []interface{}{"$tlRebate", 0}}}},
			bson.M{"$lookup": bson.M{
				"as":   "rateMaster",
				"from": constants.COLLECTIONTRADELICENSERATEMASTER,
				"let":  bson.M{"varTo": "$to"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []interface{}{"$$varTlbtId", "$tlbtId"}},
						bson.M{"$eq": []interface{}{"$$varTlctId", "$tlctId"}},
						bson.M{"$gte": []interface{}{"$$varTo", "$doe"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
				},
			},
			},
			bson.M{"$addFields": bson.M{"rateMaster": bson.M{"$arrayElemAt": []interface{}{"$rateMaster", 0}}}},
			// initiating the value of fyTax for a financial year
			bson.M{"$addFields": bson.M{"fyTax": "$rateMaster.rate"}},
			//
			bson.M{"$lookup": bson.M{
				"as":   "alreadyPaied",
				"from": constants.COLLECTIONTRADELICENSEPAYMENTSFYPART2,
				"let":  bson.M{"fyId": "$uniqueId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []interface{}{"$$fyId", "$fy.uniqueId"}},
						bson.M{"$eq": []interface{}{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
						bson.M{"$eq": []interface{}{"$$tlId", "$tradeLicenseId"}},
					}}}},
					bson.M{"$sort": bson.M{"doe": -1}},
					bson.M{"$group": bson.M{"_id": nil,
						"penalty":        bson.M{"$sum": "$fy.details.penalty"},
						"totalTaxAmount": bson.M{"$sum": "$fy.details.totalTaxAmount"},
						"tax":            bson.M{"$sum": "$fy.details.tax"},
						"other":          bson.M{"$sum": "$fy.details.other"},
						"penaltyValue":   bson.M{"$sum": "$fy.details.penaltyValue"},
						"penaltyType":    bson.M{"$sum": "$fy.details.penaltyType"},
						"rebate":         bson.M{"$sum": "$fy.details.rebate"},
						"rebateValue":    bson.M{"$sum": "$fy.details.rebateValue"},
						"rebateType":     bson.M{"$sum": "$fy.details.rebateType"},
					}},
				}}},
			bson.M{"$addFields": bson.M{"alreadyPaied": bson.M{"$arrayElemAt": []interface{}{"$alreadyPaied", 0}}}},
		},
	},
	})
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSECATEGORYTYPE, "tlctId", "uniqueId", "ref.tradeLicenseCategoryType", "ref.tradeLicenseCategoryType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.previous.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.previous.ward", "ref.address.ward")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Trade License Payment =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var demands []models.RefTradeLicenseDemandPart2
	var demand *models.RefTradeLicenseDemandPart2
	if err = cursor.All(context.TODO(), &demands); err != nil {
		return nil, err
	}
	if len(demands) > 0 {
		demand = &demands[0]
	}
	return demand, nil

}
