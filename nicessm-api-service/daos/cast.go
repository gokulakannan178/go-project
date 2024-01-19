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
func (d *Daos) SaveCast(ctx *models.Context, Cast *models.Cast) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCAST).InsertOne(ctx.CTX, Cast)
	if err != nil {
		return err
	}
	Cast.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleCast(ctx *models.Context, UniqueID string) (*models.RefCast, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCAST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Casts []models.RefCast
	var Cast *models.RefCast
	if err = cursor.All(ctx.CTX, &Casts); err != nil {
		return nil, err
	}
	if len(Casts) > 0 {
		Cast = &Casts[0]
	}
	return Cast, nil
}

//UpdateCast : ""
func (d *Daos) UpdateCast(ctx *models.Context, Cast *models.Cast) error {

	selector := bson.M{"_id": Cast.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Cast}
	_, err := ctx.DB.Collection(constants.COLLECTIONCAST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCast : ""
func (d *Daos) FilterCast(ctx *models.Context, Castfilter *models.CastFilter, pagination *models.Pagination) ([]models.RefCast, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Castfilter != nil {

		if len(Castfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Castfilter.ActiveStatus}})
		}
		if len(Castfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Castfilter.Status}})
		}
		//Regex
		if Castfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Castfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if Castfilter != nil {
		if Castfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{Castfilter.SortBy: Castfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCAST).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Cast query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCAST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Casts []models.RefCast
	if err = cursor.All(context.TODO(), &Casts); err != nil {
		return nil, err
	}
	return Casts, nil
}

//EnableCast :""
func (d *Daos) EnableCast(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CASTSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCAST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableCast(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CASTSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCAST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCast :""
func (d *Daos) DeleteCast(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CASTSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCAST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
