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

//SavePenalty :""
func (d *Daos) SavePenalty(ctx *models.Context, penalty *models.Penalty) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).InsertOne(ctx.CTX, penalty)
	return err
}

//GetSinglePenalty : ""
func (d *Daos) GetSinglePenalty(ctx *models.Context, UniqueID string) (*models.RefPenalty, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var penaltys []models.RefPenalty
	var penalty *models.RefPenalty
	if err = cursor.All(ctx.CTX, &penaltys); err != nil {
		return nil, err
	}
	if len(penaltys) > 0 {
		penalty = &penaltys[0]
	}
	return penalty, nil
}

//UpdatePenalty : ""
func (d *Daos) UpdatePenalty(ctx *models.Context, penalty *models.Penalty) error {
	selector := bson.M{"uniqueId": penalty.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": penalty, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPenalty : ""
func (d *Daos) FilterPenalty(ctx *models.Context, penaltyfilter *models.PenaltyFilter, pagination *models.Pagination) ([]models.RefPenalty, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if penaltyfilter != nil {

		if len(penaltyfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": penaltyfilter.Status}})
		}
		if len(penaltyfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": penaltyfilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("penalty query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var penaltys []models.RefPenalty
	if err = cursor.All(context.TODO(), &penaltys); err != nil {
		return nil, err
	}
	return penaltys, nil
}

//EnablePenalty :""
func (d *Daos) EnablePenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PENALTYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePenalty :""
func (d *Daos) DisablePenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PENALTYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePenalty :""
func (d *Daos) DeletePenalty(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PENALTYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPENALTY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
