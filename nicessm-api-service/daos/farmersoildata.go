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

//SaveFarmerSoilData :""
func (d *Daos) SaveFarmerSoilData(ctx *models.Context, farmerSoilData *models.FarmerSoilData) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).InsertOne(ctx.CTX, farmerSoilData)
	if err != nil {
		return err
	}
	farmerSoilData.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateFarmerSoilData : ""
func (d *Daos) UpdateFarmerSoilData(ctx *models.Context, farmerSoilData *models.FarmerSoilData) error {

	selector := bson.M{"_id": farmerSoilData.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": farmerSoilData}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableFarmerSoilData :""
func (d *Daos) EnableFarmerSoilData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSOILDATASTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmerSoilData :""
func (d *Daos) DisableFarmerSoilData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSOILDATASTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmerSoilData :""
func (d *Daos) DeleteFarmerSoilData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSOILDATASTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleFarmerSoilData : ""
func (d *Daos) GetSingleFarmerSoilData(ctx *models.Context, UniqueID string) (*models.RefFarmerSoilData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMERLAND, "farmerLand", "_id", "ref.farmerLand", "ref.farmerLand")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONNUTRIENTS, "microNutrients", "_id", "ref.microNutrients", "ref.microNutrients")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONNUTRIENTS, "nutrients_macroNutrients", "_id", "ref.macroNutrients", "ref.macroNutrients")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAIDLOCATION, "labName", "_id", "ref.labName", "ref.labName")...)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmerSoilDatas []models.RefFarmerSoilData
	var farmerSoilData *models.RefFarmerSoilData
	if err = cursor.All(ctx.CTX, &farmerSoilDatas); err != nil {
		return nil, err
	}
	if len(farmerSoilDatas) > 0 {
		farmerSoilData = &farmerSoilDatas[0]
	}
	return farmerSoilData, nil
}

//FilterFarmerSoilData : ""
func (d *Daos) FilterFarmerSoilData(ctx *models.Context, farmerSoilDatafilter *models.FarmerSoilDataFilter, pagination *models.Pagination) ([]models.RefFarmerSoilData, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if farmerSoilDatafilter != nil {
		if len(farmerSoilDatafilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": farmerSoilDatafilter.Status}})
		}
		// if farmerSoilDatafilter.SearchBox.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: farmerSoilDatafilter.SearchBox.Name, Options: "xi"}})
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMERLAND, "farmerLand", "_id", "ref.farmerLand", "ref.farmerLand")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONNUTRIENTS, "nutrients_microNutrients", "_id", "ref.microNutrients", "ref.microNutrients")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONNUTRIENTS, "nutrients_macroNutrients", "_id", "ref.macroNutrients", "ref.macroNutrients")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAIDLOCATION, "labName", "_id", "ref.labName", "ref.labName")...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("FarmerSoilData query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERSOILDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var FarmerSoilDatas []models.RefFarmerSoilData
	if err = cursor.All(context.TODO(), &FarmerSoilDatas); err != nil {
		return nil, err
	}
	return FarmerSoilDatas, nil
}
