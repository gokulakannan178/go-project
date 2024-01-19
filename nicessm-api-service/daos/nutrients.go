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

//SaveDisease :""
func (d *Daos) SaveNutrients(ctx *models.Context, Nutrients *models.Nutrients) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).InsertOne(ctx.CTX, Nutrients)
	if err != nil {
		return err
	}
	Nutrients.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleNutrients(ctx *models.Context, UniqueID string) (*models.RefNutrients, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Nutrientss []models.RefNutrients
	var Nutrients *models.RefNutrients
	if err = cursor.All(ctx.CTX, &Nutrientss); err != nil {
		return nil, err
	}
	if len(Nutrientss) > 0 {
		Nutrients = &Nutrientss[0]
	}
	return Nutrients, nil
}

//UpdateNutrients : ""
func (d *Daos) UpdateNutrients(ctx *models.Context, Nutrients *models.Nutrients) error {

	selector := bson.M{"_id": Nutrients.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Nutrients}
	_, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterNutrients : ""
func (d *Daos) FilterNutrients(ctx *models.Context, Nutrientsfilter *models.NutrientsFilter, pagination *models.Pagination) ([]models.RefNutrients, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Nutrientsfilter != nil {

		if len(Nutrientsfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Nutrientsfilter.ActiveStatus}})
		}
		if len(Nutrientsfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Nutrientsfilter.Status}})
		}
		if Nutrientsfilter.Type != "" {
			query = append(query, bson.M{"type": Nutrientsfilter.Type})
		}
		//Regex
		if Nutrientsfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Nutrientsfilter.SearchBox.Name, Options: "xi"}})
		}
		if Nutrientsfilter.SearchBox.Code != "" {
			query = append(query, bson.M{"code": primitive.Regex{Pattern: Nutrientsfilter.SearchBox.Code, Options: "xi"}})
		}
		if Nutrientsfilter.SearchBox.Type != "" {
			query = append(query, bson.M{"type": primitive.Regex{Pattern: Nutrientsfilter.SearchBox.Type, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if Nutrientsfilter != nil {
		if Nutrientsfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{Nutrientsfilter.SortBy: Nutrientsfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Nutrients query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Nutrientss []models.RefNutrients
	if err = cursor.All(context.TODO(), &Nutrientss); err != nil {
		return nil, err
	}
	return Nutrientss, nil
}

//EnableNutrients :""
func (d *Daos) EnableNutrients(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NUTRIENTSSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableNutrients(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NUTRIENTSSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteNutrients :""
func (d *Daos) DeleteNutrients(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NUTRIENTSSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNUTRIENTS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
