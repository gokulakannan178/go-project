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
func (d *Daos) SaveSeason(ctx *models.Context, Season *models.Season) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSEASON).InsertOne(ctx.CTX, Season)
	if err != nil {
		return err
	}
	Season.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleSeason(ctx *models.Context, UniqueID string) (*models.RefSeason, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSEASON).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Seasons []models.RefSeason
	var Season *models.RefSeason
	if err = cursor.All(ctx.CTX, &Seasons); err != nil {
		return nil, err
	}
	if len(Seasons) > 0 {
		Season = &Seasons[0]
	}
	return Season, nil
}

//UpdateSeason : ""
func (d *Daos) UpdateSeason(ctx *models.Context, Season *models.Season) error {

	selector := bson.M{"_id": Season.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Season}
	_, err := ctx.DB.Collection(constants.COLLECTIONSEASON).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterSeason : ""
func (d *Daos) FilterSeason(ctx *models.Context, Seasonfilter *models.SeasonFilter, pagination *models.Pagination) ([]models.RefSeason, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Seasonfilter != nil {

		if len(Seasonfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Seasonfilter.ActiveStatus}})
		}
		if len(Seasonfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Seasonfilter.Status}})
		}
		//Regex
		if Seasonfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Seasonfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if Seasonfilter != nil {
		if Seasonfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{Seasonfilter.SortBy: Seasonfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSEASON).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Season query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSEASON).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Seasons []models.RefSeason
	if err = cursor.All(context.TODO(), &Seasons); err != nil {
		return nil, err
	}
	return Seasons, nil
}

//EnableSeason :""
func (d *Daos) EnableSeason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SEASONSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableSeason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SEASONSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteSeason :""
func (d *Daos) DeleteSeason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SEASONSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
