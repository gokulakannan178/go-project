package daos

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveDisease :""
func (d *Daos) SaveMonthSeason(ctx *models.Context, MonthSeason *models.MonthSeason) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).InsertOne(ctx.CTX, MonthSeason)
	if err != nil {
		return err
	}
	MonthSeason.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleMonthSeason(ctx *models.Context, UniqueID string) (*models.RefMonthSeason, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var MonthSeasons []models.RefMonthSeason
	var MonthSeason *models.RefMonthSeason
	if err = cursor.All(ctx.CTX, &MonthSeasons); err != nil {
		return nil, err
	}
	if len(MonthSeasons) > 0 {
		MonthSeason = &MonthSeasons[0]
	}
	return MonthSeason, nil
}

//UpdateMonthSeason : ""
func (d *Daos) UpdateMonthSeason(ctx *models.Context, MonthSeason *models.MonthSeason) error {

	selector := bson.M{"_id": MonthSeason.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": MonthSeason}
	_, err := ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterMonthSeason : ""
func (d *Daos) FilterMonthSeason(ctx *models.Context, MonthSeasonfilter *models.MonthSeasonFilter, pagination *models.Pagination) ([]models.RefMonthSeason, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if MonthSeasonfilter != nil {

		if len(MonthSeasonfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": MonthSeasonfilter.ActiveStatus}})
		}
		if len(MonthSeasonfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": MonthSeasonfilter.Status}})
		}
		if len(MonthSeasonfilter.Season) > 0 {
			query = append(query, bson.M{"season": bson.M{"$in": MonthSeasonfilter.Season}})
		}
		//Regex
		if MonthSeasonfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: MonthSeasonfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if MonthSeasonfilter != nil {
		if MonthSeasonfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{MonthSeasonfilter.SortBy: MonthSeasonfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("MonthSeason query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var MonthSeasons []models.RefMonthSeason
	if err = cursor.All(context.TODO(), &MonthSeasons); err != nil {
		return nil, err
	}
	return MonthSeasons, nil
}

//EnableMonthSeason :""
func (d *Daos) EnableMonthSeason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.MONTHSEASONSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableMonthSeason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.MONTHSEASONSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteMonthSeason :""
func (d *Daos) DeleteMonthSeason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.MONTHSEASONSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetCurrentMonthSeason(ctx *models.Context) (*models.MonthSeason, error) {
	t := time.Now()
	UniqueID := fmt.Sprintf("%v", int64(t.Month()))
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	d.Shared.BsonToJSONPrintTag("GetCurrentMonthSeason =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMONTHSEASON).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var MonthSeasons []models.MonthSeason
	var MonthSeason *models.MonthSeason
	if err = cursor.All(ctx.CTX, &MonthSeasons); err != nil {
		return nil, err
	}
	if len(MonthSeasons) > 0 {
		MonthSeason = &MonthSeasons[0]
	}
	return MonthSeason, nil

}
