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
func (d *Daos) SaveStateWeatherAlertDissimination(ctx *models.Context, StateWeatherAlertDissimination *models.StateWeatherAlertDissimination) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).InsertOne(ctx.CTX, StateWeatherAlertDissimination)
	if err != nil {
		return err
	}
	StateWeatherAlertDissimination.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) (*models.RefStateWeatherAlertDissimination, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	d.Shared.BsonToJSONPrintTag("StateWeatherAlertDissimination query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlertDissiminations []models.RefStateWeatherAlertDissimination
	var StateWeatherAlertDissimination *models.RefStateWeatherAlertDissimination
	if err = cursor.All(ctx.CTX, &StateWeatherAlertDissiminations); err != nil {
		return nil, err
	}
	if len(StateWeatherAlertDissiminations) > 0 {
		StateWeatherAlertDissimination = &StateWeatherAlertDissiminations[0]
	}
	return StateWeatherAlertDissimination, nil
}

//UpdateStateWeatherAlertDissimination : ""
func (d *Daos) UpdateStateWeatherAlertDissimination(ctx *models.Context, StateWeatherAlertDissimination *models.StateWeatherAlertDissimination) error {

	selector := bson.M{"_id": StateWeatherAlertDissimination.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": StateWeatherAlertDissimination}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterStateWeatherAlertDissimination : ""
func (d *Daos) FilterStateWeatherAlertDissimination(ctx *models.Context, StateWeatherAlertDissiminationfilter *models.StateWeatherAlertDissiminationFilter, pagination *models.Pagination) ([]models.RefStateWeatherAlertDissimination, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline, err := d.StateWeatherAlertDissiminationQuery(ctx, StateWeatherAlertDissiminationfilter)
	if err != nil {
		return nil, err
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("StateWeatherAlertDissimination query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlertDissiminations []models.RefStateWeatherAlertDissimination
	if err = cursor.All(context.TODO(), &StateWeatherAlertDissiminations); err != nil {
		return nil, err
	}
	return StateWeatherAlertDissiminations, nil
}
func (d *Daos) StateWeatherAlertDissiminationQuery(ctx *models.Context, StateWeatherAlertDissiminationfilter *models.StateWeatherAlertDissiminationFilter) ([]bson.M, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if StateWeatherAlertDissiminationfilter != nil {

		if len(StateWeatherAlertDissiminationfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": StateWeatherAlertDissiminationfilter.ActiveStatus}})
		}
		if len(StateWeatherAlertDissiminationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": StateWeatherAlertDissiminationfilter.Status}})
		}
		//Regex
		if StateWeatherAlertDissiminationfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: StateWeatherAlertDissiminationfilter.SearchBox.Name, Options: "xi"}})
		}

	}
	//daterange
	if StateWeatherAlertDissiminationfilter.DateDisseminationRange != nil {
		//var sd,ed time.Time
		if StateWeatherAlertDissiminationfilter.DateDisseminationRange.From != nil {
			sd := time.Date(StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Year(), StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Month(), StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Day(), 0, 0, 0, 0, StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Location())
			ed := time.Date(StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Year(), StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Month(), StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Day(), 23, 59, 59, 0, StateWeatherAlertDissiminationfilter.DateDisseminationRange.From.Location())
			if StateWeatherAlertDissiminationfilter.DateDisseminationRange.To != nil {
				ed = time.Date(StateWeatherAlertDissiminationfilter.DateDisseminationRange.To.Year(), StateWeatherAlertDissiminationfilter.DateDisseminationRange.To.Month(), StateWeatherAlertDissiminationfilter.DateDisseminationRange.To.Day(), 23, 59, 59, 0, StateWeatherAlertDissiminationfilter.DateDisseminationRange.To.Location())
			}
			query = append(query, bson.M{"dateOfDissemination": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if StateWeatherAlertDissiminationfilter != nil {
		if StateWeatherAlertDissiminationfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{StateWeatherAlertDissiminationfilter.SortBy: StateWeatherAlertDissiminationfilter.SortOrder}})

		}

	}
	return mainPipeline, nil
}
func (d *Daos) FilterStateWeatherAlertDissiminationReport(ctx *models.Context, StateWeatherAlertDissiminationfilter *models.StateWeatherAlertDissiminationFilter, pagination *models.Pagination) ([]models.RefStateWeatherAlertDissimination, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline, err := d.StateWeatherAlertDissiminationQuery(ctx, StateWeatherAlertDissiminationfilter)
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{
			"farmers": 0,
			"users":   0,
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$sort": bson.M{"date": 1},
	})
	if err != nil {
		return nil, err
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("StateWeatherAlertDissimination query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlertDissiminations []models.RefStateWeatherAlertDissimination
	if err = cursor.All(context.TODO(), &StateWeatherAlertDissiminations); err != nil {
		return nil, err
	}
	return StateWeatherAlertDissiminations, nil
}

//EnableStateWeatherAlertDissimination :""
func (d *Daos) EnableStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTDISSIMINATIONSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTDISSIMINATIONSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteStateWeatherAlertDissimination :""
func (d *Daos) DeleteStateWeatherAlertDissimination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTDISSIMINATIONSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
