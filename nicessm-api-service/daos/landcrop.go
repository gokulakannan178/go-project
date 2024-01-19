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

//SaveLandCrop :""
func (d *Daos) SaveLandCrop(ctx *models.Context, landCrop *models.LandCrop) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONLANDCROP).InsertOne(ctx.CTX, landCrop)
	if err != nil {
		return err
	}
	landCrop.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateLandCrop : ""
func (d *Daos) UpdateLandCrop(ctx *models.Context, landCrop *models.LandCrop) error {

	selector := bson.M{"_id": landCrop.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": landCrop}
	_, err := ctx.DB.Collection(constants.COLLECTIONLANDCROP).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableLandCrop :""
func (d *Daos) EnableLandCrop(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LANDCROPSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANDCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLandCrop :""
func (d *Daos) DisableLandCrop(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LANDCROPSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANDCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLandCrop :""
func (d *Daos) DeleteLandCrop(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LANDCROPSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANDCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleLandCrop : ""
func (d *Daos) GetSingleLandCrop(ctx *models.Context, UniqueID string) (*models.RefLandCrop, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "crop", "_id", "ref.crop", "ref.crop")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMERLAND, "farmerLand", "_id", "ref.farmerLand", "ref.farmerLand")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "season", "_id", "ref.season", "ref.season")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANDCROP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var landCrops []models.RefLandCrop
	var landCrop *models.RefLandCrop
	if err = cursor.All(ctx.CTX, &landCrops); err != nil {
		return nil, err
	}
	if len(landCrops) > 0 {
		landCrop = &landCrops[0]
	}
	return landCrop, nil
}

//FilterLandCrop : ""
func (d *Daos) FilterLandCrop(ctx *models.Context, landCropfilter *models.LandCropFilter, pagination *models.Pagination) ([]models.RefLandCrop, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if landCropfilter != nil {
		if len(landCropfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": landCropfilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "crop", "_id", "ref.crop", "ref.crop")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMERLAND, "farmerLand", "_id", "ref.farmerLand", "ref.farmerLand")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "season", "_id", "ref.season", "ref.season")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLANDCROP).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("LandCrop query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANDCROP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var landCrops []models.RefLandCrop
	if err = cursor.All(context.TODO(), &landCrops); err != nil {
		return nil, err
	}
	return landCrops, nil
}
