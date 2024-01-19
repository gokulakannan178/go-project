package daos

import (
	"context"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FilterContentReport : ""
func (d *Daos) FilterContentReport(ctx *models.Context, filter *models.ContentReportFilter, pagination *models.Pagination) ([]models.RefContent, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if !filter.Village.ID.IsZero() {
			query = append(query, bson.M{"indexingData.VILLAGE": bson.M{"$" + filter.Village.Condition: filter.Village.ID}})
		}
		if !filter.State.ID.IsZero() {
			query = append(query, bson.M{"indexingData.STATE": bson.M{"$" + filter.State.Condition: filter.State.ID}})
		}
		if !filter.District.ID.IsZero() {
			query = append(query, bson.M{"indexingData.DISTRICT": bson.M{"$" + filter.District.Condition: filter.District.ID}})
		}
		if !filter.GramPanchayat.ID.IsZero() {
			query = append(query, bson.M{"indexingData.GRAM_PANCHAYAT": bson.M{"$" + filter.GramPanchayat.Condition: filter.GramPanchayat.ID}})
		}
		if !filter.Block.ID.IsZero() {
			query = append(query, bson.M{"indexingData.BLOCK": bson.M{"$" + filter.Block.Condition: filter.Block.ID}})
		}
		if !filter.KnowledgeDomain.ID.IsZero() {
			query = append(query, bson.M{"knowledgeDomain": bson.M{"$" + filter.KnowledgeDomain.Condition: filter.KnowledgeDomain.ID}})
		}
		if !filter.SubDomain.ID.IsZero() {
			query = append(query, bson.M{"subDomain": bson.M{"$" + filter.SubDomain.Condition: filter.SubDomain.ID}})
		}
		if !filter.SubTopic.ID.IsZero() {
			query = append(query, bson.M{"subTopic": bson.M{"$" + filter.SubTopic.Condition: filter.SubTopic.ID}})
		}
		if !filter.Topic.ID.IsZero() {
			query = append(query, bson.M{"topic": bson.M{"$" + filter.Topic.Condition: filter.Topic.ID}})
		}
		if filter.Classfication.ID != "" {
			query = append(query, bson.M{"indexingData.CLASSIFICATION": bson.M{"$" + filter.Classfication.Condition: filter.Classfication.ID}})
		}
		if !filter.Season.ID.IsZero() {
			query = append(query, bson.M{"indexingData.SEASON": bson.M{"$" + filter.Season.Condition: filter.Season.ID}})
		}
		if !filter.Market.ID.IsZero() {
			query = append(query, bson.M{"indexingData.MARKET": bson.M{"$" + filter.Market.Condition: filter.Market.ID}})
		}
		if !filter.Organisation.ID.IsZero() {
			query = append(query, bson.M{"organisation": bson.M{"$" + filter.Organisation.Condition: filter.Organisation.ID}})
		}
		if !filter.Project.ID.IsZero() {
			query = append(query, bson.M{"project": bson.M{"$" + filter.Project.Condition: filter.Project.ID}})
		}
		if filter.Status.ID != "" {
			query = append(query, bson.M{"status": bson.M{"$" + filter.Status.Condition: filter.Status.ID}})
		}
		if !filter.Soil_type.ID.IsZero() {
			query = append(query, bson.M{"indexingData.SOIL_TYPE": bson.M{"$" + filter.Soil_type.Condition: filter.Soil_type.ID}})
		}
		if !filter.Commodity.ID.IsZero() {
			query = append(query, bson.M{"indexingData.COMMODITY": bson.M{"$" + filter.Commodity.Condition: filter.Commodity.ID}})
		}
		if filter.Type.ID != "" {
			query = append(query, bson.M{"type": bson.M{"$" + filter.Type.Condition: filter.Type.ID}})
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
			query = append(query, bson.M{"dateCreated": bson.M{"$" + sdcondition: sd, "$" + edcondition: ed}})
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
		d.Shared.BsonToJSONPrintTag("Content pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "project", "_id", "ref.project", "ref.project")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTOPIC, "topic", "_id", "ref.topic", "ref.topic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBTOPIC, "subTopic", "_id", "ref.subTopic", "ref.subTopic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "indexingData.STATE", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "indexingData.DISTRICT", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "indexingData.BLOCK", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "indexingData.SEASON", "_id", "ref.season", "ref.season")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYSTAGE, "indexingData.STAGE", "_id", "ref.stage", "ref.stage")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "indexingData.VILLAGE", "_id", "ref.village", "ref.village")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "indexingData.GRAM_PANCHAYAT", "_id", "ref.gram_panchayat", "ref.gram_panchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYFUNCTION, "indexingData.FUNCTION", "_id", "ref.function", "ref.function")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSOILTYPE, "indexingData.SOIL_TYPE", "_id", "ref.soil_type", "ref.soil_type")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMARKET, "indexingData.MARKET", "_id", "ref.market", "ref.market")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "indexingData.VARIETY", "_id", "ref.variety", "ref.variety")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "author", "_id", "ref.author", "ref.author")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "reviewedBy", "_id", "ref.reviewedBy", "ref.reviewedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "indexingData.CATEGORY", "_id", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "indexingData.COMMODITY", "_id", "ref.commodity", "ref.commodity")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONINSECT, "indexingData.CAUSATIVE", "_id", "ref.causativeInsect", "ref.causativeInsect")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISEASE, "indexingData.CAUSATIVE", "_id", "ref.causativeDisease", "ref.causativeDisease")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYSUBVARIETY, "indexingData.SUB_VARIETY", "_id", "ref.subVariety", "ref.subVariety")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMERSOILDATA, "indexingData.SOIL_DATA", "_id", "ref.soilData", "ref.soilData")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONPROJECT,
			"as":   "ref.project",
			"let":  bson.M{"projectId": "$project"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$_id", "$$projectId"}},
				}}}},
				{"$project": bson.M{"farmers": 0}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"ref.project": bson.M{"$arrayElemAt": []interface{}{"$ref.project", 0}}},
	}) //Aggregation
	d.Shared.BsonToJSONPrintTag("Content query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Contents []models.RefContent
	if err = cursor.All(context.TODO(), &Contents); err != nil {
		return nil, err
	}
	return Contents, nil
}

//FilterDuplicateContentReport
func (d *Daos) FilterDuplicateContentReport(ctx *models.Context, contentfilter *models.ContentFilter, pagination *models.Pagination) ([]models.DuplicateContentReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if contentfilter != nil {

		if len(contentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": contentfilter.Status}})
		}
		if len(contentfilter.OmitStatus) > 0 {
			query = append(query, bson.M{"status": bson.M{"$nin": contentfilter.OmitStatus}})
		}
		if len(contentfilter.KnowledgeDomain) > 0 {
			query = append(query, bson.M{"knowledgeDomain": bson.M{"$in": contentfilter.KnowledgeDomain}})
		}
		if len(contentfilter.State) > 0 {
			query = append(query, bson.M{"STATE": bson.M{"$in": contentfilter.State}})
		}
		if len(contentfilter.Soil_type) > 0 {
			query = append(query, bson.M{"SOIL_TYPE": bson.M{"$in": contentfilter.Soil_type}})
		}
		if len(contentfilter.Organisation) > 0 {
			query = append(query, bson.M{"organisation": bson.M{"$in": contentfilter.Organisation}})
		}
		if len(contentfilter.Source) > 0 {
			query = append(query, bson.M{"source": bson.M{"$in": contentfilter.Source}})
		}
		if len(contentfilter.Author) > 0 {
			query = append(query, bson.M{"author": bson.M{"$in": contentfilter.Author}})
		}
		if len(contentfilter.SubDomain) > 0 {
			query = append(query, bson.M{"subDomain": bson.M{"$in": contentfilter.SubDomain}})
		}
		if len(contentfilter.SubTopic) > 0 {
			query = append(query, bson.M{"subTopic": bson.M{"$in": contentfilter.SubTopic}})
		}
		if len(contentfilter.Topic) > 0 {
			query = append(query, bson.M{"Topic": bson.M{"$in": contentfilter.Topic}})
		}
		if len(contentfilter.ReviewedBy) > 0 {
			query = append(query, bson.M{"reviewedBy": bson.M{"$in": contentfilter.ReviewedBy}})
		}
		if len(contentfilter.Project) > 0 {
			query = append(query, bson.M{"project": bson.M{"$in": contentfilter.Project}})
		}
		if len(contentfilter.SmsType) > 0 {
			query = append(query, bson.M{"smsType": bson.M{"$in": contentfilter.SmsType}})
		}
		if len(contentfilter.SmsContentType) > 0 {
			query = append(query, bson.M{"smsContentType": bson.M{"$in": contentfilter.SmsContentType}})
		}
		if len(contentfilter.RecordId) > 0 {
			query = append(query, bson.M{"recordId": bson.M{"$in": contentfilter.RecordId}})
		}
		if len(contentfilter.District) > 0 {
			query = append(query, bson.M{"DISTRICT": bson.M{"$in": contentfilter.District}})
		}
		if len(contentfilter.Block) > 0 {
			query = append(query, bson.M{"BLOCK": bson.M{"$in": contentfilter.Block}})
		}
		if len(contentfilter.Season) > 0 {
			query = append(query, bson.M{"SEASON": bson.M{"$in": contentfilter.Season}})
		}
		if len(contentfilter.Stage) > 0 {
			query = append(query, bson.M{"STAGE": bson.M{"$in": contentfilter.Stage}})
		}
		if len(contentfilter.Village) > 0 {
			query = append(query, bson.M{"VILLAGE": bson.M{"$in": contentfilter.Village}})
		}
		if len(contentfilter.Gram_panchayat) > 0 {
			query = append(query, bson.M{"GRAM_PANCHAYAT": bson.M{"$in": contentfilter.Gram_panchayat}})
		}
		if len(contentfilter.Function) > 0 {
			query = append(query, bson.M{"FUNCTION": bson.M{"$in": contentfilter.Function}})
		}
		if len(contentfilter.Market) > 0 {
			query = append(query, bson.M{"MARKET": bson.M{"$in": contentfilter.Market}})
		}
		if len(contentfilter.Causative) > 0 {
			query = append(query, bson.M{"CAUSATIVE": bson.M{"$in": contentfilter.Causative}})
		}
		if len(contentfilter.Variety) > 0 {
			query = append(query, bson.M{"VARIETY": bson.M{"$in": contentfilter.Variety}})
		}
		if len(contentfilter.Category) > 0 {
			query = append(query, bson.M{"CATEGORY": bson.M{"$in": contentfilter.Category}})
		}
		if len(contentfilter.Commodity) > 0 {
			query = append(query, bson.M{"COMMODITY": bson.M{"$in": contentfilter.Commodity}})
		}
		if len(contentfilter.Cause) > 0 {
			query = append(query, bson.M{"CAUSE": bson.M{"$in": contentfilter.Cause}})
		}
		if len(contentfilter.Cause_type) > 0 {
			query = append(query, bson.M{"CAUSE_TYPE": bson.M{"$in": contentfilter.Cause_type}})
		}
		if len(contentfilter.Classfication) > 0 {
			query = append(query, bson.M{"CLASSIFICATION": bson.M{"$in": contentfilter.Classfication}})
		}
		if len(contentfilter.ContentTitle) > 0 {
			query = append(query, bson.M{"contentTitle": bson.M{"$in": contentfilter.ContentTitle}})
		}
		if len(contentfilter.SubVariety) > 0 {
			query = append(query, bson.M{"SUB_VARIETY": bson.M{"$in": contentfilter.SubVariety}})
		}
		if len(contentfilter.SoilData) > 0 {
			query = append(query, bson.M{"SOIL_DATA": bson.M{"$in": contentfilter.SoilData}})
		}
		if len(contentfilter.Irrigation) > 0 {
			query = append(query, bson.M{"IRRIGATION": bson.M{"$in": contentfilter.SoilData}})
		}
		//Regex
		if contentfilter.SearchBox.Content != "" {
			query = append(query, bson.M{"content": primitive.Regex{Pattern: contentfilter.SearchBox.Content, Options: "xi"}})
		}
		if contentfilter.SearchBox.Comment != "" {
			query = append(query, bson.M{"comment": primitive.Regex{Pattern: contentfilter.SearchBox.Comment, Options: "xi"}})
		}
		if contentfilter.SearchBox.RecordId != "" {
			query = append(query, bson.M{"recordId": primitive.Regex{Pattern: contentfilter.SearchBox.RecordId, Options: "xi"}})
		}
		if len(contentfilter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": contentfilter.Type}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if contentfilter != nil {
		if contentfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{contentfilter.SortBy: contentfilter.SortOrder}})
		}
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{"_id": 1, "knowledgeDomain": 1, "subDomain": 1, "topic": 1, "subTopic": 1, "indexingData.SEASON": 1, "recordId": 1},
	})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
		"_id": bson.M{
			"kd":       "$knowledgeDomain",
			"sd":       "$subDomain",
			"topic":    "$topic",
			"subtopic": "$subTopic",
			"season":   "$indexingData.SEASON",
		},
		"contents": bson.M{"$push": bson.M{"recordId": "$recordId"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"count": bson.M{"$size": "$contents"}}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"count": bson.M{"$gt": 1}}})
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}

		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("dupicatecontent pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "_id.kd", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTOPIC, "_id.topic", "_id", "ref.topic", "ref.topic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "_id.sd", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBTOPIC, "_id.subtopic", "_id", "ref.subTopic", "ref.subTopic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "_id.season", "_id", "ref.season", "ref.season")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("duplicateContent query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Contents []models.DuplicateContentReport
	if err = cursor.All(context.TODO(), &Contents); err != nil {
		return nil, err
	}
	return Contents, nil
}
