package daos

import (
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) GetUserChargeDemand(ctx *models.Context, ucmcf *models.UserChargeMonthlyCalcQueryFilter) (*models.UserChargeDemand, error) {
	mainPipeline := []bson.M{}

	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": ucmcf.UserChargeID}})
	if len(ucmcf.AddFy) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"neededFys": ucmcf.AddFy}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "financialyears",
		"as":   "ref.fy",
		"let": bson.M{
			"doa":        "$userCharge.doa",
			"neededFys":  "$neededFys",
			"categoryId": "$userCharge.categoryId",
		},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Active"}},
				bson.M{"$lt": []interface{}{"$$doa", "$to"}},
			}}}},
			bson.M{"$sort": bson.M{"order": 1}},
			bson.M{
				"$addFields": bson.M{
					"neededMonthsIndex": bson.M{
						"$indexOfArray": []string{
							"$$neededFys.fyId",
							"$uniqueId",
						},
					},
				},
			},
			bson.M{
				"$addFields": bson.M{
					"neededMonths": bson.M{
						"$arrayElemAt": []string{
							"$$neededFys",
							"$neededMonthsIndex",
						},
					},
				},
			},
			func() bson.M {
				if len(ucmcf.AddFy) > 0 {
					return bson.M{
						"$match": bson.M{
							"neededMonthsIndex": bson.M{
								"$gte": 0,
							},
						},
					}
				}
				return bson.M{
					"$match": bson.M{},
				}
			}(),

			bson.M{"$lookup": bson.M{
				"from": "months",
				"as":   "fymonth",
				"let": bson.M{
					"fromYear":     bson.M{"$year": "$from"},
					"toYear":       bson.M{"$year": "$to"},
					"fyId":         "$uniqueId",
					"neededMonths": "$neededMonths",
					"categoryId":   "$$categoryId",
				},
				"pipeline": []bson.M{
					func() bson.M {
						if len(ucmcf.AddFy) > 0 {
							return bson.M{
								"$match": bson.M{
									"$expr": bson.M{
										"$in": []string{
											"$month",
											"$$neededMonths.months",
										},
									},
								},
							}
						}
						return bson.M{
							"$match": bson.M{},
						}
					}(),

					bson.M{"$sort": bson.M{"fyOrder": 1}},
					bson.M{"$addFields": bson.M{
						"from": bson.M{"$dateFromParts": bson.M{
							"year":  bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$year", "from"}}, "then": "$$fromYear", "else": "$$toYear"}},
							"month": "$month", "day": 1, "hour": 0, "minute": 0, "second": 0, "millisecond": 0}},
						"to": bson.M{"$dateFromParts": bson.M{
							"year":  bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$year", "from"}}, "then": "$$fromYear", "else": "$$toYear"}},
							"month": bson.M{"$sum": []interface{}{"$month", 1}}, "day": 0, "hour": 23, "minute": 59, "second": 59, "millisecond": 0}},
					}},
					bson.M{"$lookup": bson.M{
						"from": "userchargeratemaster",
						"as":   "rate",
						"let": bson.M{
							"monthstart": "$from",
							"monthend":   "$to",
							"categoryId": "$$categoryId",
						},
						"pipeline": []bson.M{
							bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

								bson.M{"$lte": []interface{}{"$doe", "$$monthend"}},
								bson.M{"$eq": []interface{}{"$categoryId", "$$categoryId"}},
								//
							}}}},
							bson.M{"$sort": bson.M{"doe": -1}},
						},
					}},
					bson.M{"$addFields": bson.M{"rate": bson.M{"$arrayElemAt": []interface{}{"$rate", 0}}}},

					bson.M{"$lookup": bson.M{
						"from": "userchargepaymentfy",
						"as":   "alreadypaid",
						"let":  bson.M{"fyId": "$$fyId", "month": "$month"},
						"pipeline": []bson.M{
							bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

								bson.M{"$eq": []string{"$status", "Completed"}},
								bson.M{"$eq": []string{"$fy.uniqueId", "$$fyId"}},
								bson.M{"$eq": []string{"$month.month", "$$month"}},
								bson.M{"$eq": []string{"$userchargeId", ucmcf.UserChargeID}},
								//
							}}}},
							bson.M{"$group": bson.M{"_id": nil,
								"total": bson.M{"$sum": "$month.paidTax"},
							}},
						},
					}},
					bson.M{"$addFields": bson.M{"alreadypaid": bson.M{"$arrayElemAt": []interface{}{"$alreadypaid.total", 0}}}},
				},
			}},
			//
		},
	}},
	)

	d.Shared.BsonToJSONPrintTag("property payment query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var pps []models.UserChargeDemand
	var pp *models.UserChargeDemand
	if err = cursor.All(ctx.CTX, &pps); err != nil {
		return nil, err
	}
	if len(pps) > 0 {
		pp = &pps[0]
	}
	fmt.Println("daos - pp.Demand.TotalTax", pp.UCDemand.TotalTax)
	return pp, nil
}
