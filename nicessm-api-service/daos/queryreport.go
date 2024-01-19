package daos

import (
	"context"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//FilterQueryReport : ""
func (d *Daos) FilterQueryReport(ctx *models.Context, filter *models.QueryReportFilter, pagination *models.Pagination) ([]models.RefQuery, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if !filter.Village.ID.IsZero() {
			query = append(query, bson.M{"village": bson.M{"$" + filter.Village.Condition: filter.Village.ID}})
		}
		if !filter.State.ID.IsZero() {
			query = append(query, bson.M{"state": bson.M{"$" + filter.State.Condition: filter.State.ID}})
		}
		if !filter.District.ID.IsZero() {
			query = append(query, bson.M{"district": bson.M{"$" + filter.District.Condition: filter.District.ID}})
		}
		if !filter.GramPanchayat.ID.IsZero() {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$" + filter.GramPanchayat.Condition: filter.GramPanchayat.ID}})
		}
		if !filter.Block.ID.IsZero() {
			query = append(query, bson.M{"block": bson.M{"$" + filter.Block.Condition: filter.Block.ID}})
		}
		//Regex
		// if Queryfilter.Regex.Query != "" {
		// 	query = append(query, bson.M{"query": primitive.Regex{Pattern: Queryfilter.Regex.Query, Options: "xi"}})
		// }
		if filter.CreatedFrom.Date != nil {
			var sd, ed time.Time
			var sdcondition, edcondition string = "gte", "lte"
			sd = time.Date(filter.CreatedFrom.Date.Year(), filter.CreatedFrom.Date.Month(), filter.CreatedFrom.Date.Day(), 0, 0, 0, 0, filter.CreatedFrom.Date.Location())
			ed = time.Date(filter.CreatedFrom.Date.Year(), filter.CreatedFrom.Date.Month(), filter.CreatedFrom.Date.Day(), 23, 59, 59, 0, filter.CreatedFrom.Date.Location())
			sdcondition = filter.CreatedFrom.Condition

			if filter.CreatedTo.Date != nil {
				ed = time.Date(filter.CreatedTo.Date.Year(), filter.CreatedTo.Date.Month(), filter.CreatedTo.Date.Day(), 23, 59, 59, 0, filter.CreatedTo.Date.Location())
				edcondition = filter.CreatedTo.Condition
			}
			query = append(query, bson.M{"date": bson.M{"$" + sdcondition: sd, "$" + edcondition: ed}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$" + filter.Condition: query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("user pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONQUERY).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "assignedTo", "_id", "ref.assignedTo", "ref.assignedTo")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "resolvedBy", "_id", "ref.resolvedBy", "ref.resolvedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "createdBy", "_id", "ref.createdByUser", "ref.createdByUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "createdBy", "_id", "ref.createdByFarmer", "ref.createdByFarmer")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Query query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONQUERY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Querys []models.RefQuery
	if err = cursor.All(context.TODO(), &Querys); err != nil {
		return nil, err
	}
	return Querys, nil
}
