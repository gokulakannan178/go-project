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

//SaveHonoriffic :""
func (d *Daos) SaveHonoriffic(ctx *models.Context, honoriffic *models.Honoriffic) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).InsertOne(ctx.CTX, honoriffic)
	return err
}

//GetSingleHonoriffic : ""
func (d *Daos) GetSingleHonoriffic(ctx *models.Context, UniqueID string) (*models.RefHonoriffic, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var honoriffics []models.RefHonoriffic
	var honoriffic *models.RefHonoriffic
	if err = cursor.All(ctx.CTX, &honoriffics); err != nil {
		return nil, err
	}
	if len(honoriffics) > 0 {
		honoriffic = &honoriffics[0]
	}
	return honoriffic, nil
}

//UpdateHonoriffic : ""
func (d *Daos) UpdateHonoriffic(ctx *models.Context, honoriffic *models.Honoriffic) error {
	selector := bson.M{"uniqueId": honoriffic.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": honoriffic, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterHonoriffic : ""
func (d *Daos) FilterHonoriffic(ctx *models.Context, honorifficfilter *models.HonorifficFilter, pagination *models.Pagination) ([]models.RefHonoriffic, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if honorifficfilter != nil {

		if len(honorifficfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": honorifficfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("honoriffic query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var honoriffics []models.RefHonoriffic
	if err = cursor.All(context.TODO(), &honoriffics); err != nil {
		return nil, err
	}
	return honoriffics, nil
}

//EnableHonoriffic :""
func (d *Daos) EnableHonoriffic(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HONORIFFICTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableHonoriffic :""
func (d *Daos) DisableHonoriffic(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HONORIFFICTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteHonoriffic :""
func (d *Daos) DeleteHonoriffic(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HONORIFFICTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHONORIFFIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
