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

//SaveCommonLanguageTranslations :""
func (d *Daos) SaveCommonLanguageTranslations(ctx *models.Context, commonlanguagetranslations *models.CommonLanguageTranslationss) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).InsertOne(ctx.CTX, commonlanguagetranslations)
	if err != nil {
		return err
	}
	commonlanguagetranslations.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleCommonLanguageTranslations : ""
func (d *Daos) GetSingleCommonLanguageTranslations(ctx *models.Context, UniqueID string) (*models.RefCommonLanguageTranslations, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var commonlanguagetranslationss []models.RefCommonLanguageTranslations
	var commonlanguagetranslations *models.RefCommonLanguageTranslations
	if err = cursor.All(ctx.CTX, &commonlanguagetranslationss); err != nil {
		return nil, err
	}
	if len(commonlanguagetranslationss) > 0 {
		commonlanguagetranslations = &commonlanguagetranslationss[0]
	}
	return commonlanguagetranslations, nil
}

//UpdateCommonLanguageTranslations : ""
func (d *Daos) UpdateCommonLanguageTranslations(ctx *models.Context, commonlanguagetranslations *models.CommonLanguageTranslationss) error {

	selector := bson.M{"_id": commonlanguagetranslations.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": commonlanguagetranslations}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCommonLanguageTranslations : ""
func (d *Daos) FilterCommonLanguageTranslations(ctx *models.Context, commonlanguagetranslationsfilter *models.CommonLanguageTranslationsFilter, pagination *models.Pagination) ([]models.RefCommonLanguageTranslations, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if commonlanguagetranslationsfilter != nil {

		if len(commonlanguagetranslationsfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": commonlanguagetranslationsfilter.ActiveStatus}})
		}
		if len(commonlanguagetranslationsfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": commonlanguagetranslationsfilter.Status}})
		}
		//Regex
		if commonlanguagetranslationsfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: commonlanguagetranslationsfilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if commonlanguagetranslationsfilter != nil {
		if commonlanguagetranslationsfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{commonlanguagetranslationsfilter.SortBy: commonlanguagetranslationsfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("commonlanguagetranslations query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var commonlanguagetranslationss []models.RefCommonLanguageTranslations
	if err = cursor.All(context.TODO(), &commonlanguagetranslationss); err != nil {
		return nil, err
	}
	return commonlanguagetranslationss, nil
}

//EnableCommonLanguageTranslations :""
func (d *Daos) EnableCommonLanguageTranslations(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.COMMONLANGUAGETRANSLATIONSSTATUSTRUE, "status": constants.COMMONLANGUAGETRANSLATIONSSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommonLanguageTranslations :""
func (d *Daos) DisableCommonLanguageTranslations(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.COMMONLANGUAGETRANSLATIONSSTATUSFALSE, "status": constants.COMMONLANGUAGETRANSLATIONSSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommonLanguageTranslations :""
func (d *Daos) DeleteCommonLanguageTranslations(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.COMMONLANGUAGETRANSLATIONSSTATUSFALSE, "status": constants.COMMONLANGUAGETRANSLATIONSSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleCommonLanguageTranslationsWithType(ctx *models.Context, UniqueID string) (*models.RefCommonLanguageTranslations, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"languageType": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLANGUAGETRANSLATIONS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var commonlanguagetranslationss []models.RefCommonLanguageTranslations
	var commonlanguagetranslations *models.RefCommonLanguageTranslations
	if err = cursor.All(ctx.CTX, &commonlanguagetranslationss); err != nil {
		return nil, err
	}
	if len(commonlanguagetranslationss) > 0 {
		commonlanguagetranslations = &commonlanguagetranslationss[0]
	}
	return commonlanguagetranslations, nil
}
