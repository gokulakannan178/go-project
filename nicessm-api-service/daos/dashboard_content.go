package daos

import (
	"context"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) DashboardContentSmsCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) ([]models.DashboardContentCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.ContentFilter(ctx, &contentfilter.ContentFilter)
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"unReviewed": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPESMS}}, "status": bson.M{"$in": []string{"U"}}}},
			bson.M{"$count": "unReviewed"},
		},
		"reviewed": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPESMS}}, "status": bson.M{"$in": []string{"A", "E"}}}},

			bson.M{"$count": "reviewed"},
		},
		"rejected": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPESMS}}, "status": bson.M{"$in": []string{"R"}}}},

			bson.M{"$count": "rejected"},
		},

		"deleted": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPESMS}}, "status": bson.M{"$in": []string{"Deleted"}}}},

			bson.M{"$count": "deleted"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"unReviewed": bson.M{"$arrayElemAt": []interface{}{"$unReviewed", 0}},
			"reviewed":   bson.M{"$arrayElemAt": []interface{}{"$reviewed", 0}},
			"rejected":   bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}},
			"deleted":    bson.M{"$arrayElemAt": []interface{}{"$deleted", 0}},
		}},
		bson.M{"$addFields": bson.M{"unReviewed": "$unReviewed.unReviewed", "reviewed": "$reviewed.reviewed", "rejected": "$rejected.rejected",
			"deleted": "$deleted.deleted"}})

	d.Shared.BsonToJSONPrintTag("DashboardContentCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbccr []models.DashboardContentCountReport
	if err := cursor.All(ctx.CTX, &dbccr); err != nil {
		return nil, err
	}
	return dbccr, nil

}
func (d *Daos) DashboardContentVoiceCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) ([]models.DashboardContentCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.ContentFilter(ctx, &contentfilter.ContentFilter)

	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"unReviewed": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVOICE}}, "status": bson.M{"$in": []string{"U"}}}},
			bson.M{"$count": "unReviewed"},
		},
		"reviewed": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVOICE}}, "status": bson.M{"$in": []string{"A", "E"}}}},

			bson.M{"$count": "reviewed"},
		},
		"rejected": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVOICE}}, "status": bson.M{"$in": []string{"R"}}}},

			bson.M{"$count": "rejected"},
		},

		"deleted": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVOICE}}, "status": bson.M{"$in": []string{"Deleted"}}}},

			bson.M{"$count": "deleted"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"unReviewed": bson.M{"$arrayElemAt": []interface{}{"$unReviewed", 0}},
			"reviewed":   bson.M{"$arrayElemAt": []interface{}{"$reviewed", 0}},
			"rejected":   bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}},
			"deleted":    bson.M{"$arrayElemAt": []interface{}{"$deleted", 0}},
		}},
		bson.M{"$addFields": bson.M{"unReviewed": "$unReviewed.unReviewed", "reviewed": "$reviewed.reviewed", "rejected": "$rejected.rejected",
			"deleted": "$deleted.deleted"}})

	d.Shared.BsonToJSONPrintTag("DashboardContentCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbccr []models.DashboardContentCountReport
	if err := cursor.All(ctx.CTX, &dbccr); err != nil {
		return nil, err
	}
	return dbccr, nil

}
func (d *Daos) DashboardContentVideoCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) ([]models.DashboardContentCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.ContentFilter(ctx, &contentfilter.ContentFilter)

	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"unReviewed": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVIDEO}}, "status": bson.M{"$in": []string{"U"}}}},
			bson.M{"$count": "unReviewed"},
		},
		"reviewed": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVIDEO}}, "status": bson.M{"$in": []string{"A", "E"}}}},

			bson.M{"$count": "reviewed"},
		},
		"rejected": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVIDEO}}, "status": bson.M{"$in": []string{"R"}}}},

			bson.M{"$count": "rejected"},
		},

		"deleted": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEVIDEO}}, "status": bson.M{"$in": []string{"Deleted"}}}},

			bson.M{"$count": "deleted"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"unReviewed": bson.M{"$arrayElemAt": []interface{}{"$unReviewed", 0}},
			"reviewed":   bson.M{"$arrayElemAt": []interface{}{"$reviewed", 0}},
			"rejected":   bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}},
			"deleted":    bson.M{"$arrayElemAt": []interface{}{"$deleted", 0}},
		}},
		bson.M{"$addFields": bson.M{"unReviewed": "$unReviewed.unReviewed", "reviewed": "$reviewed.reviewed", "rejected": "$rejected.rejected",
			"deleted": "$deleted.deleted"}})

	d.Shared.BsonToJSONPrintTag("DashboardContentCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbccr []models.DashboardContentCountReport
	if err := cursor.All(ctx.CTX, &dbccr); err != nil {
		return nil, err
	}
	return dbccr, nil

}
func (d *Daos) DashboardContentPosterCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) ([]models.DashboardContentCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.ContentFilter(ctx, &contentfilter.ContentFilter)

	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"unReviewed": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEPOSTER}}, "status": bson.M{"$in": []string{"U"}}}},
			bson.M{"$count": "unReviewed"},
		},
		"reviewed": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEPOSTER}}, "status": bson.M{"$in": []string{"A", "E"}}}},

			bson.M{"$count": "reviewed"},
		},
		"rejected": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEPOSTER}}, "status": bson.M{"$in": []string{"R"}}}},

			bson.M{"$count": "rejected"},
		},

		"deleted": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEPOSTER}}, "status": bson.M{"$in": []string{"Deleted"}}}},

			bson.M{"$count": "deleted"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"unReviewed": bson.M{"$arrayElemAt": []interface{}{"$unReviewed", 0}},
			"reviewed":   bson.M{"$arrayElemAt": []interface{}{"$reviewed", 0}},
			"rejected":   bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}},
			"deleted":    bson.M{"$arrayElemAt": []interface{}{"$deleted", 0}},
		}},
		bson.M{"$addFields": bson.M{"unReviewed": "$unReviewed.unReviewed", "reviewed": "$reviewed.reviewed", "rejected": "$rejected.rejected",
			"deleted": "$deleted.deleted"}})

	d.Shared.BsonToJSONPrintTag("DashboardContentCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbccr []models.DashboardContentCountReport
	if err := cursor.All(ctx.CTX, &dbccr); err != nil {
		return nil, err
	}
	return dbccr, nil

}
func (d *Daos) DashboardContentDocmentCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) ([]models.DashboardContentCountReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.ContentFilter(ctx, &contentfilter.ContentFilter)

	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{
		"unReviewed": []bson.M{
			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEDOCMENT}}, "status": bson.M{"$in": []string{"U"}}}},
			bson.M{"$count": "unReviewed"},
		},
		"reviewed": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEDOCMENT}}, "status": bson.M{"$in": []string{"A", "E"}}}},

			bson.M{"$count": "reviewed"},
		},
		"rejected": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEDOCMENT}}, "status": bson.M{"$in": []string{"R"}}}},

			bson.M{"$count": "rejected"},
		},

		"deleted": []bson.M{

			bson.M{"$match": bson.M{"type": bson.M{"$in": []string{constants.CONTENTTYPEDOCMENT}}, "status": bson.M{"$in": []string{"Deleted"}}}},

			bson.M{"$count": "deleted"},
		},
	}},
		bson.M{"$addFields": bson.M{
			"unReviewed": bson.M{"$arrayElemAt": []interface{}{"$unReviewed", 0}},
			"reviewed":   bson.M{"$arrayElemAt": []interface{}{"$reviewed", 0}},
			"rejected":   bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}},
			"deleted":    bson.M{"$arrayElemAt": []interface{}{"$deleted", 0}},
		}},
		bson.M{"$addFields": bson.M{"unReviewed": "$unReviewed.unReviewed", "reviewed": "$reviewed.reviewed", "rejected": "$rejected.rejected",
			"deleted": "$deleted.deleted"}})

	d.Shared.BsonToJSONPrintTag("DashboardContentCount Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dbccr []models.DashboardContentCountReport
	if err := cursor.All(ctx.CTX, &dbccr); err != nil {
		return nil, err
	}
	return dbccr, nil

}
func (d *Daos) DayWiseContentDemandChart(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) (*models.DayWiseContentDemandChartReport, error) {

	mainPipeline := []bson.M{}
	mainPipeline = d.ContentFilter(ctx, &contentfilter.ContentFilter)
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"day": bson.M{"$dayOfMonth": "$dateCreated"}, "status": "$status"},
		"count": bson.M{"$sum": 1}}})
	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "mobileTowerCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
	//mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "data": bson.M{"$push": bson.M{"k": "$_id.status", "v": "$count"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"data": bson.M{"$arrayToObject": "$data"}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{contentfilter.CreatedFrom.StartDate.Day(), contentfilter.CreatedFrom.EndDate.Day() + 1, 1}},
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
										"R":       0.0,
										"Deleted": 0.0,
										"U":       0.0,
										"A":       0.0}},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})
	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var emptyData *models.DayWiseContentDemandChartReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DayWiseContentDemandChartReport

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return emptyData, nil

}
