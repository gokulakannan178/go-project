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

func (d *Daos) SaveContentTranslation(ctx *models.Context, contenttranslation *models.ContentTranslation) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).InsertOne(ctx.CTX, contenttranslation)
	if err != nil {
		return err
	}
	contenttranslation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleContentTranslation(ctx *models.Context, UniqueID string) (*models.RefContentTranslation, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "user", "_id", "ref.translator", "ref.translator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLANGAUAGE, "language", "_id", "ref.language", "ref.language")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "reviewedBy", "_id", "ref.reviewedBy", "ref.reviewedBy")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var contenttranslations []models.RefContentTranslation
	var contenttranslation *models.RefContentTranslation
	if err = cursor.All(ctx.CTX, &contenttranslations); err != nil {
		return nil, err
	}
	if len(contenttranslations) > 0 {
		contenttranslation = &contenttranslations[0]
	}
	return contenttranslation, nil
}

func (d *Daos) UpdateContentTranslation(ctx *models.Context, contenttranslation *models.ContentTranslation) error {

	selector := bson.M{"_id": contenttranslation.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": contenttranslation}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterContentTranslation(ctx *models.Context, contenttranslationfilter *models.ContentTranslationFilter, pagination *models.Pagination) ([]models.RefContentTranslation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCONTENT, "content", "_id", "ref.content", "ref.content")...)
	query := []bson.M{}
	if contenttranslationfilter != nil {

		if len(contenttranslationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": contenttranslationfilter.Status}})
		}
		if len(contenttranslationfilter.Content) > 0 {
			query = append(query, bson.M{"content": bson.M{"$in": contenttranslationfilter.Content}})
		}
		if len(contenttranslationfilter.ContentType) > 0 {
			query = append(query, bson.M{"ref.content.type": bson.M{"$in": contenttranslationfilter.ContentType}})
		}
		if len(contenttranslationfilter.ContentOrganisation) > 0 {
			query = append(query, bson.M{"ref.content.organisation": bson.M{"$in": contenttranslationfilter.ContentOrganisation}})
		}
		if len(contenttranslationfilter.ContentProject) > 0 {
			query = append(query, bson.M{"ref.content.project": bson.M{"$in": contenttranslationfilter.ContentProject}})
		}
		if len(contenttranslationfilter.ContentStatus) > 0 {
			query = append(query, bson.M{"ref.content.status": bson.M{"$in": contenttranslationfilter.ContentStatus}})
		}
		//Regex
		if contenttranslationfilter.SearchBox.TranslatedContent != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: contenttranslationfilter.SearchBox.TranslatedContent, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if contenttranslationfilter != nil {
		if contenttranslationfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{contenttranslationfilter.SortBy: contenttranslationfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("contenttranslation pagenation query =>", paginationPipeline)
		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "translator", "_id", "ref.translator", "ref.translator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLANGAUAGE, "language", "_id", "ref.language", "ref.language")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "reviewedBy", "_id", "ref.reviewedBy", "ref.reviewedBy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "ref.content.organisation", "_id", "ref.organisation", "ref.organisation")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Aidlocation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Contenttranslations []models.RefContentTranslation
	if err = cursor.All(context.TODO(), &Contenttranslations); err != nil {
		return nil, err
	}
	return Contenttranslations, nil
}

func (d *Daos) EnableContentTranslation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTTRANSLATIONSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableContentTranslation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTTRANSLATIONSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteContentTranslation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTTRANSLATIONSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) ApprovedContentTranslation(ctx *models.Context, Content *models.ApprovedContentTranslation) error {
	selector := bson.M{"_id": Content.ID}
	t := time.Now()
	updateInterface := bson.M{"$set": bson.M{"status": constants.APPROVED, "reviewedBy": Content.ReviewedBy, "dateCreated": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) RejectedContentTranslation(ctx *models.Context, Content *models.RejectedContentTranslation) error {
	selector := bson.M{"_id": Content.ID}
	t := time.Now()
	updateInterface := bson.M{"$set": bson.M{"status": constants.REJECTED, "reviewedBy": Content.ReviewedBy, "dateCreated": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENTTRANSLATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
