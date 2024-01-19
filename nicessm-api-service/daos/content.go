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

func (d *Daos) SaveContent(ctx *models.Context, content *models.Content) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).InsertOne(ctx.CTX, content)
	if err != nil {
		return err
	}
	content.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleContent(ctx *models.Context, UniqueID string) (*models.RefContent, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// Lookup
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
	})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var contents []models.RefContent
	var content *models.RefContent
	if err = cursor.All(ctx.CTX, &contents); err != nil {
		return nil, err
	}
	if len(contents) > 0 {
		content = &contents[0]
	}
	return content, nil
}

func (d *Daos) UpdateContent(ctx *models.Context, content *models.Content) error {

	selector := bson.M{"_id": content.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": content}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterContent(ctx *models.Context, contentfilter *models.ContentFilter, pagination *models.Pagination) ([]models.RefContent, error) {
	mainPipeline := []bson.M{}

	mainPipeline = d.ContentFilter(ctx, contentfilter)

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("Content Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	//Lookup
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
	})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("content query =>", mainPipeline)
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

func (d *Daos) EnableContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) ApprovedContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.APPROVED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) RejectedContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.REJECTED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) ContentManager(ctx *models.Context, contentfilter *models.ContentFilter, pagination *models.Pagination) ([]models.RefContent, error) {
	mainPipeline := []bson.M{}

	mainPipeline = d.ContentFilter(ctx, contentfilter)

	//Adding $match from filter

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("Content Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "project", "_id", "ref.project", "ref.project")...)
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Aidlocation query =>", mainPipeline)
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
func (d *Daos) ContentFilter(ctx *models.Context, contentfilter *models.ContentFilter) []bson.M {
	if contentfilter != nil {
		if len(contentfilter.TranslationStatus) <= 0 {
			contentfilter.TranslationStatus = make([]string, 0)
		}
	}
	mainPipeline := []bson.M{}
	if contentfilter != nil {
		if contentfilter != nil {
			contenttranslationpipeline := []bson.M{}
			contenttranslationpipeline = append(contenttranslationpipeline, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$$contentid", "$content"}},
				{"$in": []interface{}{"$status", contentfilter.TranslationStatus}},
			}}}})
			contenttranslationpipeline = append(contenttranslationpipeline, d.CommonLookup(constants.COLLECTIONUSER, "translator", "_id", "ref.translator", "ref.translator")...)
			contenttranslationpipeline = append(contenttranslationpipeline, d.CommonLookup(constants.COLLECTIONLANGAUAGE, "language", "_id", "ref.language", "ref.language")...)
			contenttranslationpipeline = append(contenttranslationpipeline, d.CommonLookup(constants.COLLECTIONUSER, "reviewedBy", "_id", "ref.reviewedBy", "ref.reviewedBy")...)

			mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
				"from":     constants.COLLECTIONCONTENTTRANSLATION,
				"as":       "ref.translatedContents",
				"let":      bson.M{"contentid": "$_id"},
				"pipeline": contenttranslationpipeline}})
		}
	}
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
			query = append(query, bson.M{"indexingData.STATE": bson.M{"$in": contentfilter.State}})
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
			query = append(query, bson.M{"indexingData.DISTRICT": bson.M{"$in": contentfilter.District}})
		}
		if len(contentfilter.Block) > 0 {
			query = append(query, bson.M{"indexingData.BLOCK": bson.M{"$in": contentfilter.Block}})
		}
		if len(contentfilter.Season) > 0 {
			query = append(query, bson.M{"indexingData.SEASON": bson.M{"$in": contentfilter.Season}})
		}
		if len(contentfilter.Stage) > 0 {
			query = append(query, bson.M{"indexingData.STAGE": bson.M{"$in": contentfilter.Stage}})
		}
		if len(contentfilter.Village) > 0 {
			query = append(query, bson.M{"indexingData.VILLAGE": bson.M{"$in": contentfilter.Village}})
		}
		if len(contentfilter.Gram_panchayat) > 0 {
			query = append(query, bson.M{"indexingData.GRAM_PANCHAYAT": bson.M{"$in": contentfilter.Gram_panchayat}})
		}
		if len(contentfilter.Function) > 0 {
			query = append(query, bson.M{"indexingData.FUNCTION": bson.M{"$in": contentfilter.Function}})
		}
		if len(contentfilter.Market) > 0 {
			query = append(query, bson.M{"indexingData.MARKET": bson.M{"$in": contentfilter.Market}})
		}
		if len(contentfilter.Causative) > 0 {
			query = append(query, bson.M{"indexingData.CAUSATIVE": bson.M{"$in": contentfilter.Causative}})
		}
		if len(contentfilter.Variety) > 0 {
			query = append(query, bson.M{"indexingData.VARIETY": bson.M{"$in": contentfilter.Variety}})
		}
		if len(contentfilter.Category) > 0 {
			query = append(query, bson.M{"indexingData.CATEGORY": bson.M{"$in": contentfilter.Category}})
		}
		if len(contentfilter.Commodity) > 0 {
			query = append(query, bson.M{"indexingData.COMMODITY": bson.M{"$in": contentfilter.Commodity}})
		}
		if len(contentfilter.Cause) > 0 {
			query = append(query, bson.M{"indexingData.CAUSE": bson.M{"$in": contentfilter.Cause}})
		}
		if len(contentfilter.Cause_type) > 0 {
			query = append(query, bson.M{"indexingData.CAUSE_TYPE": bson.M{"$in": contentfilter.Cause_type}})
		}
		if len(contentfilter.Classfication) > 0 {
			query = append(query, bson.M{"indexingData.CLASSIFICATION": bson.M{"$in": contentfilter.Classfication}})
		}
		if len(contentfilter.ContentTitle) > 0 {
			query = append(query, bson.M{"contentTitle": bson.M{"$in": contentfilter.ContentTitle}})
		}
		if len(contentfilter.SubVariety) > 0 {
			query = append(query, bson.M{"indexingData.SUB_VARIETY": bson.M{"$in": contentfilter.SubVariety}})
		}
		if len(contentfilter.SoilData) > 0 {
			query = append(query, bson.M{"indexingData.SOIL_DATA": bson.M{"$in": contentfilter.SoilData}})
		}
		if len(contentfilter.Irrigation) > 0 {
			query = append(query, bson.M{"indexingData.IRRIGATION": bson.M{"$in": contentfilter.SoilData}})
		}
		//Regex
		if contentfilter.SearchBox.Content != "" {
			query = append(query, bson.M{"content": primitive.Regex{Pattern: contentfilter.SearchBox.Content, Options: "i"}})
		}
		if contentfilter.SearchBox.ContentTitle != "" {
			query = append(query, bson.M{"contentTitle": primitive.Regex{Pattern: contentfilter.SearchBox.ContentTitle, Options: "xi"}})
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
		if len(contentfilter.TranslationStatus) > 0 {
			query = append(query, bson.M{"ref.translatedContents.status": bson.M{"$in": contentfilter.TranslationStatus}})

		}
		if contentfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{contentfilter.SortBy: contentfilter.SortOrder}})
		}
		if contentfilter.CreatedFrom.StartDate != nil {
			var sd, ed time.Time
			var sdcondition, edcondition string = "gte", "lte"
			sd = time.Date(contentfilter.CreatedFrom.StartDate.Year(), contentfilter.CreatedFrom.StartDate.Month(), contentfilter.CreatedFrom.StartDate.Day(), 0, 0, 0, 0, contentfilter.CreatedFrom.StartDate.Location())
			ed = time.Date(contentfilter.CreatedFrom.EndDate.Year(), contentfilter.CreatedFrom.EndDate.Month(), contentfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, contentfilter.CreatedFrom.EndDate.Location())

			if contentfilter.CreatedFrom.EndDate != nil {
				ed = time.Date(contentfilter.CreatedFrom.EndDate.Year(), contentfilter.CreatedFrom.EndDate.Month(), contentfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, contentfilter.CreatedFrom.EndDate.Location())
				//edcondition = contentfilter.CreatedTo.Condition
			}
			fmt.Println("sd==>", sd)
			fmt.Println("ed==>", ed)
			query = append(query, bson.M{"dateCreated": bson.M{"$" + sdcondition: sd, "$" + edcondition: ed}})
		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	return mainPipeline
}
func (d *Daos) EditApprovedContent(ctx *models.Context, content *models.ApprovedContent) error {

	selector := bson.M{"_id": content.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"status": constants.APPROVED, "reviewedBy": content.ReviewedBy, "dateCreated": t, "content": content.Content}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) EditRejectedContent(ctx *models.Context, content *models.RejectedContent) error {

	selector := bson.M{"_id": content.ID}
	t := time.Now()
	updateInterface := bson.M{"$set": bson.M{"status": constants.REJECTED, "reviewedBy": content.ReviewedBy, "dateCreated": t, "note": content.Note}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) ContentViewCountIncrement(ctx *models.Context, content *models.ContentViewCount) error {

	selector := bson.M{"_id": content.ContentId}
	//t := time.Now()
	updateInterface := bson.M{"$inc": bson.M{content.UserType: 1}}
	d.Shared.BsonToJSONPrintTag("ContentViewCountIncrement query =>", updateInterface)
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
