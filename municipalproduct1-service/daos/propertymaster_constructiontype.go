package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveConstructionType :""
func (d *Daos) SaveConstructionType(ctx *models.Context, constructionType *models.ConstructionType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).InsertOne(ctx.CTX, constructionType)
	return err
}

//GetSingleConstructionType : ""
func (d *Daos) GetSingleConstructionType(ctx *models.Context, UniqueID string) (*models.RefConstructionType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var constructionTypes []models.RefConstructionType
	var constructionType *models.RefConstructionType
	if err = cursor.All(ctx.CTX, &constructionTypes); err != nil {
		return nil, err
	}
	if len(constructionTypes) > 0 {
		constructionType = &constructionTypes[0]
	}
	return constructionType, nil
}

//UpdateConstructionType : ""
func (d *Daos) UpdateConstructionType(ctx *models.Context, constructionType *models.ConstructionType) error {
	selector := bson.M{"uniqueId": constructionType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": constructionType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterConstructionType : ""
func (d *Daos) FilterConstructionType(ctx *models.Context, constructionTypefilter *models.ConstructionTypeFilter, pagination *models.Pagination) ([]models.RefConstructionType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if constructionTypefilter != nil {

		if len(constructionTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": constructionTypefilter.Status}})
		}
		if len(constructionTypefilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": constructionTypefilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("constructionType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var constructionTypes []models.RefConstructionType
	if err = cursor.All(context.TODO(), &constructionTypes); err != nil {
		return nil, err
	}
	return constructionTypes, nil
}

//EnableConstructionType :""
func (d *Daos) EnableConstructionType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CONSTRUCTIONTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableConstructionType :""
func (d *Daos) DisableConstructionType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CONSTRUCTIONTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteConstructionType :""
func (d *Daos) DeleteConstructionType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CONSTRUCTIONTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONSTRUCTIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
