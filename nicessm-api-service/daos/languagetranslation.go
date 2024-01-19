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

//SaveLanguageTranslation :""
func (d *Daos) SaveLanguageTranslation(ctx *models.Context, languagetranslation *models.LanguageTranslations) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).InsertOne(ctx.CTX, languagetranslation)
	if err != nil {
		return err
	}
	languagetranslation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleLanguageTranslation : ""
func (d *Daos) GetSingleLanguageTranslation(ctx *models.Context, UniqueID string) (*models.RefLanguageTranslation, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var languagetranslations []models.RefLanguageTranslation
	var languagetranslation *models.RefLanguageTranslation
	if err = cursor.All(ctx.CTX, &languagetranslations); err != nil {
		return nil, err
	}
	if len(languagetranslations) > 0 {
		languagetranslation = &languagetranslations[0]
	}
	return languagetranslation, nil
}

//UpdateLanguageTranslation : ""
func (d *Daos) UpdateLanguageTranslation(ctx *models.Context, languagetranslation *models.LanguageTranslations) error {

	selector := bson.M{"_id": languagetranslation.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": languagetranslation}
	_, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterLanguageTranslation : ""
func (d *Daos) FilterLanguageTranslation(ctx *models.Context, languagetranslationfilter *models.LanguageTranslationFilter, pagination *models.Pagination) ([]models.RefLanguageTranslation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if languagetranslationfilter != nil {

		if len(languagetranslationfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": languagetranslationfilter.ActiveStatus}})
		}
		if len(languagetranslationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": languagetranslationfilter.Status}})
		}
		//Regex
		if languagetranslationfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: languagetranslationfilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if languagetranslationfilter != nil {
		if languagetranslationfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{languagetranslationfilter.SortBy: languagetranslationfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("languagetranslation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var languagetranslations []models.RefLanguageTranslation
	if err = cursor.All(context.TODO(), &languagetranslations); err != nil {
		return nil, err
	}
	return languagetranslations, nil
}

//EnableLanguageTranslation :""
func (d *Daos) EnableLanguageTranslation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.LANGUAGETRANSLATIONSTATUSTRUE, "status": constants.LANGUAGETRANSLATIONSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLanguageTranslation :""
func (d *Daos) DisableLanguageTranslation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.LANGUAGETRANSLATIONSTATUSFALSE, "status": constants.LANGUAGETRANSLATIONSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLanguageTranslation :""
func (d *Daos) DeleteLanguageTranslation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.LANGUAGETRANSLATIONSTATUSFALSE, "status": constants.LANGUAGETRANSLATIONSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleLanguageTranslationWithType(ctx *models.Context, UniqueID string) (*models.RefLanguageTranslation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"languageType": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGETRANSLATIONS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var languagetranslations []models.RefLanguageTranslation
	var languagetranslation *models.RefLanguageTranslation
	if err = cursor.All(ctx.CTX, &languagetranslations); err != nil {
		return nil, err
	}
	if len(languagetranslations) > 0 {
		languagetranslation = &languagetranslations[0]
	}
	return languagetranslation, nil
}
