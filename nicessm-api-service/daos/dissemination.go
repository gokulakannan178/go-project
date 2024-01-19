package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveDissemination :""
func (d *Daos) SaveDissemination(ctx *models.Context, dissemination *models.Dissemination) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).InsertOne(ctx.CTX, dissemination)
	if err != nil {
		return err
	}
	dissemination.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDissemination : ""
func (d *Daos) GetSingleDissemination(ctx *models.Context, UniqueID string) (*models.RefDissemination, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "createdBy", "_id", "ref.createdBy", "ref.createdBy")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var disseminations []models.RefDissemination
	var dissemination *models.RefDissemination
	if err = cursor.All(ctx.CTX, &disseminations); err != nil {
		return nil, err
	}
	if len(disseminations) > 0 {
		dissemination = &disseminations[0]
	}
	return dissemination, nil
}

//UpdateDissemination : ""
func (d *Daos) UpdateDissemination(ctx *models.Context, dissemination *models.Dissemination) error {

	selector := bson.M{"_id": dissemination.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dissemination}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDissemination : ""
func (d *Daos) FilterDissemination(ctx *models.Context, disseminationfilter *models.DisseminationFilter, pagination *models.Pagination) ([]models.RefDissemination, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)
	query := []bson.M{}
	if disseminationfilter != nil {

		if len(disseminationfilter.IsSent) > 0 {
			query = append(query, bson.M{"isSent": bson.M{"$in": disseminationfilter.IsSent}})
		}
		if len(disseminationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": disseminationfilter.Status}})
		}
		if len(disseminationfilter.Content) > 0 {
			query = append(query, bson.M{"content": bson.M{"$in": disseminationfilter.Content}})
		}
		if len(disseminationfilter.CreatedBy) > 0 {
			query = append(query, bson.M{"createdBy": bson.M{"$in": disseminationfilter.CreatedBy}})
		}
		if len(disseminationfilter.Mode) > 0 {
			query = append(query, bson.M{"mode": bson.M{"$in": disseminationfilter.Mode}})
		}
		if len(disseminationfilter.State) > 0 {
			query = append(query, bson.M{"ref.content.indexingData.STATE": bson.M{"$in": disseminationfilter.State}})
		}
		if len(disseminationfilter.District) > 0 {
			query = append(query, bson.M{"ref.content.indexingData.DISTRICT": bson.M{"$in": disseminationfilter.District}})
		}
		if len(disseminationfilter.Block) > 0 {
			query = append(query, bson.M{"ref.content.indexingData.BLOCK": bson.M{"$in": disseminationfilter.Block}})
		}
		if len(disseminationfilter.Organisation) > 0 {
			query = append(query, bson.M{"ref.content.organisation": bson.M{"$in": disseminationfilter.Organisation}})
		}
		if len(disseminationfilter.Project) > 0 {
			query = append(query, bson.M{"ref.content.project": bson.M{"$in": disseminationfilter.Project}})
		}
		//daterange
		if disseminationfilter.DateDisseminationRange != nil {
			//var sd,ed time.Time
			if disseminationfilter.DateDisseminationRange.From != nil {
				sd := time.Date(disseminationfilter.DateDisseminationRange.From.Year(), disseminationfilter.DateDisseminationRange.From.Month(), disseminationfilter.DateDisseminationRange.From.Day(), 0, 0, 0, 0, disseminationfilter.DateDisseminationRange.From.Location())
				ed := time.Date(disseminationfilter.DateDisseminationRange.From.Year(), disseminationfilter.DateDisseminationRange.From.Month(), disseminationfilter.DateDisseminationRange.From.Day(), 23, 59, 59, 0, disseminationfilter.DateDisseminationRange.From.Location())
				if disseminationfilter.DateDisseminationRange.To != nil {
					ed = time.Date(disseminationfilter.DateDisseminationRange.To.Year(), disseminationfilter.DateDisseminationRange.To.Month(), disseminationfilter.DateDisseminationRange.To.Day(), 23, 59, 59, 0, disseminationfilter.DateDisseminationRange.To.Location())
				}
				query = append(query, bson.M{"dateOfDissemination": bson.M{"$gte": sd, "$lte": ed}})

			}
		}

		//Regex
		if disseminationfilter.Regex.Message != "" {
			query = append(query, bson.M{"message": primitive.Regex{Pattern: disseminationfilter.Regex.Message, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if disseminationfilter != nil {
		if disseminationfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{disseminationfilter.SortBy: disseminationfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("Commodity Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "createdBy", "_id", "ref.createdBy", "ref.createdBy")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Dissemination query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var disseminations []models.RefDissemination
	if err = cursor.All(context.TODO(), &disseminations); err != nil {
		return nil, err
	}
	return disseminations, nil
}

//EnableDissemination :""
func (d *Daos) EnableDissemination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isSent": true, "status": constants.DISSEMINATIONSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDissemination :""
func (d *Daos) DisableDissemination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isSent": false, "status": constants.DISSEMINATIONSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDissemination :""
func (d *Daos) DeleteDissemination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isSent": false, "status": constants.DISSEMINATIONSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) SendLaterDissemination(ctx *models.Context) ([]models.Dissemination, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	//	if disseminationfilter.DateDisseminationRange.From != nil {
	t := time.Now()
	sd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())

	query = append(query, bson.M{"dateOfDissemination": bson.M{"$gte": sd, "$lte": ed}})

	//}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Dissemination Send later query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var disseminations []models.Dissemination
	if err = cursor.All(context.TODO(), &disseminations); err != nil {
		return nil, err
	}
	return disseminations, nil
}
func (d *Daos) DisseminationReport(ctx *models.Context, disseminationfilter *models.DisseminationReportFilter) ([]models.RefDisseminationReport, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{
			"item":                1,
			"content":             1,
			"dateOfDissemination": 1,
			"numberOFfarmers": bson.M{"$cond": bson.M{
				"if": bson.M{"$isArray": "$farmers"}, "then": bson.M{"$size": "$farmers"}, "else": 0}},
		},
	})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)
	query := []bson.M{}
	if disseminationfilter != nil {

		if len(disseminationfilter.State) > 0 {
			query = append(query, bson.M{"ref.content.indexingData.STATE": bson.M{"$in": disseminationfilter.State}})
		}

		//daterange
		if disseminationfilter.DateDisseminationRange != nil {
			//var sd,ed time.Time
			if disseminationfilter.DateDisseminationRange.From != nil {
				sd := time.Date(disseminationfilter.DateDisseminationRange.From.Year(), disseminationfilter.DateDisseminationRange.From.Month(), disseminationfilter.DateDisseminationRange.From.Day(), 0, 0, 0, 0, disseminationfilter.DateDisseminationRange.From.Location())
				ed := time.Date(disseminationfilter.DateDisseminationRange.From.Year(), disseminationfilter.DateDisseminationRange.From.Month(), disseminationfilter.DateDisseminationRange.From.Day(), 23, 59, 59, 0, disseminationfilter.DateDisseminationRange.From.Location())
				if disseminationfilter.DateDisseminationRange.To != nil {
					ed = time.Date(disseminationfilter.DateDisseminationRange.To.Year(), disseminationfilter.DateDisseminationRange.To.Month(), disseminationfilter.DateDisseminationRange.To.Day(), 23, 59, 59, 0, disseminationfilter.DateDisseminationRange.To.Location())
				}
				query = append(query, bson.M{"dateOfDissemination": bson.M{"$gte": sd, "$lte": ed}})

			}
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": bson.M{
			"state":    "$ref.content.indexingData.STATE",
			"district": "$ref.content.indexingData.DISTRICT",
			"type":     "$ref.content.type",
		},
			"farmer": bson.M{"$sum": "$numberOFfarmers"},

			"dessiminations": bson.M{"$sum": 1},
		}})
	mainPipeline = append(mainPipeline, bson.M{

		"$group": bson.M{"_id": bson.M{
			"state":    "$_id.state",
			"district": "$_id.district",
		},

			"farmer": bson.M{"$sum": "$farmer"},

			"dessiminations": bson.M{"$sum": "$dessiminations"},
			"sms":            bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.type", "S"}}, "then": "$dessiminations", "else": 0}}},
			"voice":          bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.type", "V"}}, "then": "$dessiminations", "else": 0}}},
			"poster":         bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.type", "P"}}, "then": "$dessiminations", "else": 0}}},
			"document":       bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.type", "D"}}, "then": "$dessiminations", "else": 0}}},
			"video":          bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []string{"$_id.type", "U"}}, "then": "$dessiminations", "else": 0}}},
			"count":          bson.M{"$sum": "$dessiminations"},
		}})

	mainPipeline = append(mainPipeline, bson.M{

		"$group": bson.M{
			"_id": bson.M{
				"district": "$_id.district",
			},
			"state":          bson.M{"$first": "$_id.state"},
			"farmer":         bson.M{"$sum": "$farmer"},
			"dessiminations": bson.M{"$sum": "$dessiminations"},
			"sms":            bson.M{"$sum": "$sms"},
			"voice":          bson.M{"$sum": "$voice"},
			"poster":         bson.M{"$sum": "$poster"},
			"document":       bson.M{"$sum": "$document"},
			"video":          bson.M{"$sum": "$video"},
			"count":          bson.M{"$sum": "$count"},
		}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "_id.district", "_id", "district", "district")...)

	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{
			"_id":       bson.M{"state": "$state"},
			"districts": bson.M{"$push": "$$ROOT"},
		}})
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$addFields": bson.M{"districts": bson.M{"$arrayElemAt": []interface{}{"$districts", 0}}}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "_id.state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Dissemination query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISSEMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var disseminations []models.RefDisseminationReport
	if err = cursor.All(context.TODO(), &disseminations); err != nil {
		return nil, err
	}
	return disseminations, nil
}
