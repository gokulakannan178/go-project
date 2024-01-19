package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveAdvertisement :""
func (d *Daos) SaveAdvertisement(ctx *models.Context, advertisement *models.Advertisement) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).InsertOne(ctx.CTX, advertisement)
	return err
}

//GetSingleAdvertisement : ""
func (d *Daos) GetSingleAdvertisement(ctx *models.Context, UniqueID string) (*models.RefAdvertisement, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var advertisements []models.RefAdvertisement
	var advertisement *models.RefAdvertisement
	if err = cursor.All(ctx.CTX, &advertisements); err != nil {
		return nil, err
	}
	if len(advertisements) > 0 {
		advertisement = &advertisements[0]
	}
	return advertisement, nil
}

//UpdateAdvertisement : ""
func (d *Daos) UpdateAdvertisement(ctx *models.Context, advertisement *models.Advertisement) error {
	selector := bson.M{"uniqueId": advertisement.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": advertisement}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterAdvertisement : ""
func (d *Daos) FilterAdvertisement(ctx *models.Context, advertisementfilter *models.AdvertisementFilter, pagination *models.Pagination) ([]models.RefAdvertisement, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if advertisementfilter != nil {

		if len(advertisementfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": advertisementfilter.Status}})
		}
		//Regex
		if advertisementfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: advertisementfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("advertisement query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var advertisements []models.RefAdvertisement
	if err = cursor.All(context.TODO(), &advertisements); err != nil {
		return nil, err
	}
	return advertisements, nil
}

//EnableAdvertisement :""
func (d *Daos) EnableAdvertisement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ADVERTISEMENTOWNERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAdvertisement :""
func (d *Daos) DisableAdvertisement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ADVERTISEMENTOWNERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAdvertisement :""
func (d *Daos) DeleteAdvertisement(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ADVERTISEMENTOWNERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONADVERTISEMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
