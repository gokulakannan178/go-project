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

//SaveNewsLike : ""
func (d *Daos) SaveNewsLike(ctx *models.Context, newslike *models.NewsLike) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).InsertOne(ctx.CTX, newslike)
	if err != nil {
		return err
	}
	newslike.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateNewsLike : ""
func (d *Daos) UpdateNewsLike(ctx *models.Context, news *models.NewsLike) error {
	selector := bson.M{"uniqueId": news.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": news}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleNewsLike : ""
func (d *Daos) GetSingleNewsLike(ctx *models.Context, uniqueID string) (*models.RefNewsLike, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var NewsLikes []models.RefNewsLike
	var NewsLike *models.RefNewsLike
	if err = cursor.All(ctx.CTX, &NewsLikes); err != nil {
		return nil, err
	}
	if len(NewsLikes) > 0 {
		NewsLike = &NewsLikes[0]
	}
	return NewsLike, err
}

// EnableNewsLike : ""
func (d *Daos) EnableNewsLike(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NEWSLIKESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableNewsLike : ""
func (d *Daos) DisableNewsLike(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.NEWSLIKESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteNewsLike :""
func (d *Daos) DeleteNewsLike(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NEWSLIKESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterNewsLike : ""
func (d *Daos) FilterNewsLike(ctx *models.Context, newslike *models.FilterNewsLike, pagination *models.Pagination) ([]models.RefNewsLike, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if newslike != nil {
		if len(newslike.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": newslike.Status}})
		}
		if len(newslike.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": newslike.OrganisationId}})
		}
		//Regex
		// if news.Regex.Title != "" {
		// 	query = append(query, bson.M{"title": primitive.Regex{Pattern: news.Regex.Title, Options: "xi"}})
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNEWSLIKE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var newslikeFilter []models.RefNewsLike
	if err = cursor.All(context.TODO(), &newslikeFilter); err != nil {
		return nil, err
	}
	return newslikeFilter, nil
}
