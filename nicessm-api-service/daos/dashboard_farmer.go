package daos

import (
	"context"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) DashboardFarmerCount(ctx *models.Context, farmerfilter *models.DashboardFarmerCountFilter) ([]models.DashboardFarmerCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"active": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{constants.FARMERSTATUSACTIVE}}}},
			bson.M{"$count": "active"},
		},
		"inActive": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{constants.FARMERSTATUSDISABLED}}}},
			bson.M{"$count": "inActive"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"active":   bson.M{"$arrayElemAt": []interface{}{"$active", 0}},
			"inActive": bson.M{"$arrayElemAt": []interface{}{"$inActive", 0}},
		}},
		bson.M{"$addFields": bson.M{
			"active":   "$active.active",
			"inActive": "$inActive.inActive",
		}})

	d.Shared.BsonToJSONPrintTag("DashboardFarmerCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbucr []models.DashboardFarmerCountReport
	if err := cursor.All(ctx.CTX, &dbucr); err != nil {
		return nil, err
	}
	return dbucr, nil

}
func (d *Daos) DayWiseFarmerDemandChart(ctx *models.Context, farmerfilter *models.DashboardFarmerCountFilter) (farmer *models.DayWiseFarmerDemandChartReport, err error) {

	mainPipeline := []bson.M{}
	mainPipeline, err = d.FilterFarmerQuery(ctx, &farmerfilter.FarmerFilter)
	if err != nil {
		return
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"day": bson.M{"$dayOfMonth": "$createdDate"}, "status": "$status"},
		"count": bson.M{"$sum": 1}}})
	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "mobileTowerCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "data": bson.M{"$push": bson.M{"k": "$_id.status", "v": "$count"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data": bson.M{"$arrayToObject": "$data"}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{farmerfilter.CreatedDate.From.Day(), farmerfilter.CreatedDate.To.Day() + 1, 1}},
				"as":    "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days._id", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"_id": "$$rangeDay",
									"data": bson.M{
										"Active":   0.0,
										"Disabled": 0.0,
									}},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})
	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	//	var farmer *models.DayWiseFarmerDemandChartReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return farmer, err
	}
	var data []models.DayWiseFarmerDemandChartReport

	if err = cursor.All(context.TODO(), &data); err != nil {
		return farmer, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return farmer, nil

}
