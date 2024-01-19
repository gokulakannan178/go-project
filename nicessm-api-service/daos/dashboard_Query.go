package daos

import (
	"context"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) DashboardQueryCount(ctx *models.Context, queryfilter *models.DashboardQueryCountFilter) ([]models.DashboardQueryCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"unresolvedQueries": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{constants.QUERYSTATUSCREATED}}}},
			bson.M{"$count": "unresolvedQueries"},
		},
		"assinged": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{constants.QUERYSTATUSASSIGN}}}},
			bson.M{"$count": "assinged"},
		},
		"resolvedQueries": []bson.M{
			bson.M{"$match": bson.M{"status": bson.M{"$in": []string{constants.QUERYSTATUSRESOLVED}}}},
			bson.M{"$count": "resolvedQueries"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"unresolvedQueries": bson.M{"$arrayElemAt": []interface{}{"$unresolvedQueries", 0}},
			"assinged":          bson.M{"$arrayElemAt": []interface{}{"$assinged", 0}},
			"resolvedQueries":   bson.M{"$arrayElemAt": []interface{}{"$resolvedQueries", 0}},
		}},
		bson.M{"$addFields": bson.M{
			"unresolvedQueries": "$unresolvedQueries.unresolvedQueries",
			"assinged":          "$assinged.assinged",
			"resolvedQueries":   "$resolvedQueries.resolvedQueries",
		}})

	d.Shared.BsonToJSONPrintTag("DashboardQueryCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONQUERY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbucr []models.DashboardQueryCountReport
	if err := cursor.All(ctx.CTX, &dbucr); err != nil {
		return nil, err
	}
	return dbucr, nil

}
func (d *Daos) DayWiseQueryDemandChart(ctx *models.Context, queryfilter *models.DashboardQueryCountFilter) (query *models.DayWiseQueryDemandChartReport, err error) {

	mainPipeline := []bson.M{}
	mainPipeline, err = d.QueryFilter(ctx, &queryfilter.QueryFilter)
	if err != nil {
		return
	}
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"day": bson.M{"$dayOfMonth": "$date"}, "status": "$status"},
		"count": bson.M{"$sum": 1}}})
	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "mobileTowerCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "data": bson.M{"$push": bson.M{"k": "$_id.status", "v": "$count"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data": bson.M{"$arrayToObject": "$data"}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{queryfilter.CreatedFrom.StartDate.Day(), queryfilter.CreatedFrom.EndDate.Day() + 1, 1}},
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
	//	var query *models.DayWiseQueryDemandChartReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONQUERY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return query, err
	}
	var data []models.DayWiseQueryDemandChartReport

	if err = cursor.All(context.TODO(), &data); err != nil {
		return query, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return query, nil

}
