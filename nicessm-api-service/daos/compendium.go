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

func (d *Daos) SaveCompendium(ctx *models.Context, content *models.Compendium) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).InsertOne(ctx.CTX, content)
	if err != nil {
		return err
	}
	content.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleCompendium(ctx *models.Context, UniqueID string) (*models.RefCompendium, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "user", "_id", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var contents []models.RefCompendium
	var content *models.RefCompendium
	if err = cursor.All(ctx.CTX, &contents); err != nil {
		return nil, err
	}
	if len(contents) > 0 {
		content = &contents[0]
	}
	return content, nil
}

func (d *Daos) UpdateCompendium(ctx *models.Context, content *models.Compendium) error {

	selector := bson.M{"_id": content.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": content}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterCompendium(ctx *models.Context, contentfilter *models.CompendiumFilter, pagination *models.Pagination) ([]models.RefCompendium, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if contentfilter != nil {

		if len(contentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": contentfilter.Status}})
		}
		if len(contentfilter.Content) > 0 {
			query = append(query, bson.M{"content": bson.M{"$in": contentfilter.Content}})
		}
		//Regex
		if contentfilter.SearchBox.Comment != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: contentfilter.SearchBox.Comment, Options: "xi"}})
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
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "user", "_id", "ref.user", "ref.user")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Aidlocation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Contents []models.RefCompendium
	if err = cursor.All(context.TODO(), &Contents); err != nil {
		return nil, err
	}
	return Contents, nil
}

func (d *Daos) EnableCompendium(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMPENDIUMSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableCompendium(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMPENDIUMSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteCompendium(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMPENDIUMSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMPENDIUM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
