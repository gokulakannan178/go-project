package daos

import (
	"context"
	"haritv2-service/constants"
	"haritv2-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) ULBMasterReportV2(ctx *models.Context, filter *models.ULBMasterReportV2Filter) ([]models.RefULBMasterReportV2, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		// query = append(query, bson.M{"uniqueId": bson.M{"$in": "ULB00136"}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline,
		bson.M{
			"$lookup": bson.M{
				"from": "months",
				"as":   "months",
				"let":  bson.M{"companyId": "$uniqueId", "months": filter.Months, "year": filter.Year},
				"pipeline": []bson.M{
					{"$sort": bson.M{"month": 1}},
					{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

						{"$in": []string{"$name", "$$months"}},
					},
					}}},
					{"$addFields": bson.M{
						"startDate": bson.M{"$dateFromParts": bson.M{
							"year": "$$year", "month": "$month", "day": 1,
							"hour": 0, "minute": 0, "second": 0,
						}},
						"endDate": bson.M{"$dateFromParts": bson.M{
							"year": "$$year", "month": bson.M{"$add": []interface{}{"$month", 1}}, "day": 1,
							"hour": 0, "minute": 0, "second": 0,
						}},
					}},
					{"$lookup": bson.M{
						"from": "batch",
						"as":   "compostGenerated",
						"let": bson.M{
							"productId": "PRODUCT1",
							"pkgType":   "PKGTYPE00001",
							"companyId": "$$companyId",
							"startDate": "$startDate",
							"endDate":   "$endDate",
						},
						"pipeline": []bson.M{
							{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
								{"$eq": []interface{}{"$productId", "$$productId"}},
								{"$eq": []interface{}{"$pkgType", "$$pkgType"}},
								{"$eq": []interface{}{"$companyId", "$$companyId"}},
								{"$gte": []interface{}{"$created.on", "$$startDate"}},
								{"$lt": []interface{}{"$created.on", "$$endDate"}},
							}}}},
							{"$group": bson.M{"_id": bson.M{"month": "$created.on"}, "quantity": bson.M{"$sum": "$quantity"}}},
						},
					}},
					{"$lookup": bson.M{
						"from": "sale",
						"as":   "sale",
						"let": bson.M{
							"productId": "PRODUCT1",
							"pkgType":   "PKGTYPE00001",
							"companyId": "$$companyId",
							"startDate": "$startDate",
							"endDate":   "$endDate",
						},
						"pipeline": []bson.M{
							{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
								{"$eq": []interface{}{"$company.id", "$$companyId"}},
								{"$in": []interface{}{"$status", []interface{}{"Active"}}},
								{"$gte": []interface{}{"$createdOn.on", "$$startDate"}},
								{"$lt": []interface{}{"$createdOn.on", "$$endDate"}},
							}}}},
							{"$unwind": "$items"},
							{"$group": bson.M{"_id": bson.M{"customerType": "$customer.type", "customer": "$customer.contact"},
								"quantity": bson.M{"$sum": "$items.quantity"},
								"amount":   bson.M{"$sum": "$totalAmount"},
							}},
							{"$group": bson.M{
								"_id":           "$_id.customerType",
								"customerCount": bson.M{"$sum": 1},
								"quantity":      bson.M{"$sum": "$quantity"},
								"amount":        bson.M{"$sum": "$amount"},
							}},
							{"$addFields": bson.M{
								"k": "$_id",
								"v": bson.M{
									"customerCount": "$customerCount",
									"quantity":      "$quantity",
									"amount":        "$amount",
								},
							}},
							{"$project": bson.M{
								"k": 1, "v": 1, "_id": 0,
							}},
						},
					}},
					{"$addFields": bson.M{
						"compostGenerated": bson.M{"$arrayElemAt": []interface{}{"$compostGenerated", 0}},
						"sale":             bson.M{"$arrayToObject": "$sale"},
					}},
				},
			}},
	)

	d.Shared.BsonToJSONPrintTag("ulbmasterreportv2 query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONULB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ulbs []models.RefULBMasterReportV2
	if err = cursor.All(context.TODO(), &ulbs); err != nil {
		return nil, err
	}
	return ulbs, nil
}
