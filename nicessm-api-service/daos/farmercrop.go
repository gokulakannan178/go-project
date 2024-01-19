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

//SaveFarmerCrop :""
func (d *Daos) SaveFarmerCrop(ctx *models.Context, farmerCrop *models.FarmerCrop) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROP).InsertOne(ctx.CTX, farmerCrop)
	if err != nil {
		return err
	}
	farmerCrop.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateFarmerCrop : ""
func (d *Daos) UpdateFarmerCrop(ctx *models.Context, farmerCrop *models.FarmerCrop) error {

	selector := bson.M{"_id": farmerCrop.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": farmerCrop}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROP).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableFarmerCrop :""
func (d *Daos) EnableFarmerCrop(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCROPSTATUSWIP}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmerCrop :""
func (d *Daos) DisableFarmerCrop(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCROPSTATUSDONE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmerCrop :""
func (d *Daos) DeleteFarmerCrop(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERCROPSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleFarmerCrop : ""
func (d *Daos) GetSingleFarmerCrop(ctx *models.Context, UniqueID string) (*models.RefFarmerCrop, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "crop", "_id", "ref.crop", "ref.crop")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "interCrop", "_id", "ref.interCrop", "ref.interCrop")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "season", "_id", "ref.season", "ref.season")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "variety", "_id", "ref.variety", "ref.variety")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmerCrops []models.RefFarmerCrop
	var farmerCrop *models.RefFarmerCrop
	if err = cursor.All(ctx.CTX, &farmerCrops); err != nil {
		return nil, err
	}
	if len(farmerCrops) > 0 {
		farmerCrop = &farmerCrops[0]
	}
	return farmerCrop, nil
}

//FilterFarmerCrop : ""
func (d *Daos) FilterFarmerCrop(ctx *models.Context, farmerCropfilter *models.FarmerCropFilter, pagination *models.Pagination) ([]models.RefFarmerCrop, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "crop", "_id", "ref.crop", "ref.crop")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "ref.crop.category", "_id", "ref.category", "ref.category")...)
	query := []bson.M{}
	if farmerCropfilter != nil {
		if len(farmerCropfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": farmerCropfilter.Status}})
		}
		if !farmerCropfilter.Farmer.IsZero() {
			query = append(query, bson.M{"farmer": farmerCropfilter.Farmer})
		}
		if len(farmerCropfilter.Category) > 0 {
			query = append(query, bson.M{"ref.crop.category": bson.M{"$in": farmerCropfilter.Category}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "season", "_id", "ref.season", "ref.season")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "variety", "_id", "ref.variety", "ref.variety")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "interCrop", "_id", "ref.interCrop", "ref.interCrop")...)

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("farmercrop Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROP).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("FarmerCrop query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var FarmerCrops []models.RefFarmerCrop
	if err = cursor.All(context.TODO(), &FarmerCrops); err != nil {
		return nil, err
	}
	return FarmerCrops, nil
}

//UpdateFarmerCropDone : ""
func (d *Daos) UpdateFarmerCropDone(ctx *models.Context, farmerCrop *models.FarmerCrop) error {

	selector := bson.M{"_id": farmerCrop.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"status": constants.FARMERCROPSTATUSDONE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROP).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetFarmerCropCount(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := []bson.M{}
	query = append(query, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

		{"$eq": []string{"$status", "WIP"}},
		{"$eq": []interface{}{"$farmer", id}},
	}}}})
	query = append(query, bson.M{"$group": bson.M{
		"_id":  nil,
		"crop": bson.M{"$sum": 1},
	}})
	d.Shared.BsonToJSONPrintTag("GetFarmerCropCount query =>", query)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERCROP).Aggregate(ctx.CTX, query, nil)
	if err != nil {
		return err
	}
	var crops []models.FarmerCropCount
	var crop *models.FarmerCropCount
	if err = cursor.All(context.TODO(), &crops); err != nil {
		return err
	}
	if len(crops) > 0 {
		crop = &crops[0]
	} else {
		crop = new(models.FarmerCropCount)
	}
	query2 := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"cropCount": crop.Crop}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query2, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return nil
}
