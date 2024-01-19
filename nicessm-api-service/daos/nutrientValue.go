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

//SaveNutrientValue :""
func (d *Daos) SaveNutrientValue(ctx *models.Context, NutrientValue *models.NutrientValue) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).InsertOne(ctx.CTX, NutrientValue)
	if err != nil {
		return err
	}
	NutrientValue.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateNutrientValue : ""
func (d *Daos) UpdateNutrientValue(ctx *models.Context, NutrientValue *models.NutrientValue) error {

	selector := bson.M{"_id": NutrientValue.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": NutrientValue}
	_, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableNutrientValue :""
func (d *Daos) EnableNutrientValue(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NUTRIENTVALUESTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableNutrientValue :""
func (d *Daos) DisableNutrientValue(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NUTRIENTVALUESTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteNutrientValue :""
func (d *Daos) DeleteNutrientValue(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NUTRIENTVALUESTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleNutrientValue : ""
func (d *Daos) GetSingleNutrientValue(ctx *models.Context, UniqueID string) (*models.RefNutrientValue, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMERSOILDATA, "farmerSoilData", "_id", "ref.farmerSoilData", "ref.farmerSoilData")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONNUTRIENTS, "nutrient", "_id", "ref.nutrient", "ref.nutrient")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var NutrientValues []models.RefNutrientValue
	var NutrientValue *models.RefNutrientValue
	if err = cursor.All(ctx.CTX, &NutrientValues); err != nil {
		return nil, err
	}
	if len(NutrientValues) > 0 {
		NutrientValue = &NutrientValues[0]
	}
	return NutrientValue, nil
}

//FilterNutrientValue : ""
func (d *Daos) FilterNutrientValue(ctx *models.Context, NutrientValuefilter *models.NutrientValueFilter, pagination *models.Pagination) ([]models.RefNutrientValue, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if NutrientValuefilter != nil {
		if len(NutrientValuefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": NutrientValuefilter.Status}})
		}
		// if NutrientValuefilter.SearchBox.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: NutrientValuefilter.SearchBox.Name, Options: "xi"}})
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMERSOILDATA, "farmerSoilData", "_id", "ref.farmerSoilData", "ref.farmerSoilData")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONNUTRIENTS, "nutrient", "_id", "ref.nutrient", "ref.nutrient")...)

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("NutrientValue query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTVALUE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var NutrientValues []models.RefNutrientValue
	if err = cursor.All(context.TODO(), &NutrientValues); err != nil {
		return nil, err
	}
	return NutrientValues, nil
}
