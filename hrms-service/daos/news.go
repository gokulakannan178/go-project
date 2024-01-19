package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveNews : ""
func (d *Daos) SaveNews(ctx *models.Context, news *models.News) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONNEWS).InsertOne(ctx.CTX, news)
	if err != nil {
		return err
	}
	news.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateNews : ""
func (d *Daos) UpdateNews(ctx *models.Context, news *models.News) error {
	selector := bson.M{"uniqueId": news.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": news}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleNews : ""
func (d *Daos) GetSingleNews(ctx *models.Context, uniqueID string) (*models.RefNews, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONNEWSLIKE, "uniqueId", "newsId", "ref.newslike", "ref.newslike")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONNEWSCOMMENT, "uniqueId", "newsId", "ref.newscomment", "ref.newscomment")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "createdBy", "uniqueId", "ref.createdBy", "ref.createdBy")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNEWS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Newss []models.RefNews
	var News *models.RefNews
	if err = cursor.All(ctx.CTX, &Newss); err != nil {
		return nil, err
	}
	if len(Newss) > 0 {
		News = &Newss[0]
	}
	return News, err
}

// GetSingleNewsWithoutRef : ""
func (d *Daos) GetSingleNewsWithoutRef(ctx *models.Context, uniqueID string) (*models.News, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNEWS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Newss []models.News
	var News *models.News
	if err = cursor.All(ctx.CTX, &Newss); err != nil {
		return nil, err
	}
	if len(Newss) > 0 {
		News = &Newss[0]
	}
	return News, err
}

// EnableNews : ""
func (d *Daos) EnableNews(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NEWSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWS).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableNews : ""
func (d *Daos) DisableNews(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NEWSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWS).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteNews :""
func (d *Daos) DeleteNews(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NEWSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterNews : ""
func (d *Daos) FilterNews(ctx *models.Context, news *models.FilterNews, pagination *models.Pagination) ([]models.RefNews, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if news != nil {
		if len(news.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": news.Status}})
		}
		if len(news.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": news.OrganisationId}})
		}
		if len(news.Employee) > 0 {
			query = append(query, bson.M{"sendTo.employee": bson.M{"$in": news.Employee}})
		}
		if len(news.DepartmentId) > 0 {
			query = append(query, bson.M{"sendTo.departmentId": bson.M{"$in": news.DepartmentId}})
		}
		//Regex
		if news.Regex.Title != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: news.Regex.Title, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if news != nil {
		if news.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{news.SortBy: news.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNEWS).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONNEWSLIKE, "uniqueId", "newsId", "ref.newslike", "ref.newslike")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONNEWSCOMMENT, "uniqueId", "newsId", "ref.newscomment", "ref.newscomment")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "createdBy", "uniqueId", "ref.createdBy", "ref.createdBy")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNEWS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var newsFilter []models.RefNews
	if err = cursor.All(context.TODO(), &newsFilter); err != nil {
		return nil, err
	}
	return newsFilter, nil
}
func (d *Daos) PublishedNews(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NEWSSTATUSPUBLISHED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
