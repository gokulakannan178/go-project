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

//SaveNewsComment : ""
func (d *Daos) SaveNewsComment(ctx *models.Context, newscomment *models.NewsComment) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).InsertOne(ctx.CTX, newscomment)
	if err != nil {
		return err
	}
	newscomment.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateNewsComment : ""
func (d *Daos) UpdateNewsComment(ctx *models.Context, news *models.NewsComment) error {
	selector := bson.M{"uniqueId": news.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": news}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleNewsComment : ""
func (d *Daos) GetSingleNewsComment(ctx *models.Context, uniqueID string) (*models.RefNewsComment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var NewsComments []models.RefNewsComment
	var NewsComment *models.RefNewsComment
	if err = cursor.All(ctx.CTX, &NewsComments); err != nil {
		return nil, err
	}
	if len(NewsComments) > 0 {
		NewsComment = &NewsComments[0]
	}
	return NewsComment, err
}

// EnableNewsComment : ""
func (d *Daos) EnableNewsComment(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NEWSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableNewsComment : ""
func (d *Daos) DisableNewsComment(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NEWSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteNewsComment :""
func (d *Daos) DeleteNewsComment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NEWSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterNewsComment : ""
func (d *Daos) FilterNewsComment(ctx *models.Context, newscomment *models.FilterNewsComment, pagination *models.Pagination) ([]models.RefNewsComment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if newscomment != nil {
		if len(newscomment.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": newscomment.Status}})
		}
		if len(newscomment.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": newscomment.OrganisationId}})
		}
		//Regex
		if newscomment.Regex.Comment != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: newscomment.Regex.Comment, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNEWSCOMMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var newsCommentFilter []models.RefNewsComment
	if err = cursor.All(context.TODO(), &newsCommentFilter); err != nil {
		return nil, err
	}
	return newsCommentFilter, nil
}
