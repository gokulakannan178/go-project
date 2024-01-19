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
func (d *Daos) SaveDistrictweatheralertdissimination(ctx *models.Context, Districtweatheralertdissimination *models.DistrictWeatherAlertDissimination) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).InsertOne(ctx.CTX, Districtweatheralertdissimination)
	if err != nil {
		return err
	}
	Districtweatheralertdissimination.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleDistrictweatheralertdissimination(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherAlertDissimination, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	d.Shared.BsonToJSONPrintTag("Districtweatheralertdissimination query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Districtweatheralertdissiminations []models.RefDistrictWeatherAlertDissimination
	var Districtweatheralertdissimination *models.RefDistrictWeatherAlertDissimination
	if err = cursor.All(ctx.CTX, &Districtweatheralertdissiminations); err != nil {
		return nil, err
	}
	if len(Districtweatheralertdissiminations) > 0 {
		Districtweatheralertdissimination = &Districtweatheralertdissiminations[0]
	}
	return Districtweatheralertdissimination, nil
}

//UpdateDistrictweatheralertdissimination : ""
func (d *Daos) UpdateDistrictweatheralertdissimination(ctx *models.Context, Districtweatheralertdissimination *models.DistrictWeatherAlertDissimination) error {

	selector := bson.M{"_id": Districtweatheralertdissimination.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Districtweatheralertdissimination}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDistrictweatheralertdissimination : ""
func (d *Daos) FilterDistrictweatheralertdissimination(ctx *models.Context, Districtweatheralertdissiminationfilter *models.DistrictWeatherAlertDissiminationFilter, pagination *models.Pagination) ([]models.RefDistrictWeatherAlertDissimination, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline, err := d.DistrictweatheralertdissiminationQuery(ctx, Districtweatheralertdissiminationfilter)
	if err != nil {
		return nil, err
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Districtweatheralertdissimination query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Districtweatheralertdissiminations []models.RefDistrictWeatherAlertDissimination
	if err = cursor.All(context.TODO(), &Districtweatheralertdissiminations); err != nil {
		return nil, err
	}
	return Districtweatheralertdissiminations, nil
}
func (d *Daos) DistrictweatheralertdissiminationQuery(ctx *models.Context, Districtweatheralertdissiminationfilter *models.DistrictWeatherAlertDissiminationFilter) ([]bson.M, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Districtweatheralertdissiminationfilter != nil {

		if len(Districtweatheralertdissiminationfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Districtweatheralertdissiminationfilter.ActiveStatus}})
		}
		if len(Districtweatheralertdissiminationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Districtweatheralertdissiminationfilter.Status}})
		}
		//Regex
		if Districtweatheralertdissiminationfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Districtweatheralertdissiminationfilter.SearchBox.Name, Options: "xi"}})
		}

	}
	//daterange
	if Districtweatheralertdissiminationfilter.DateDisseminationRange != nil {
		//var sd,ed time.Time
		if Districtweatheralertdissiminationfilter.DateDisseminationRange.From != nil {
			sd := time.Date(Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Year(), Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Month(), Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Day(), 0, 0, 0, 0, Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Location())
			ed := time.Date(Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Year(), Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Month(), Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Day(), 23, 59, 59, 0, Districtweatheralertdissiminationfilter.DateDisseminationRange.From.Location())
			if Districtweatheralertdissiminationfilter.DateDisseminationRange.To != nil {
				ed = time.Date(Districtweatheralertdissiminationfilter.DateDisseminationRange.To.Year(), Districtweatheralertdissiminationfilter.DateDisseminationRange.To.Month(), Districtweatheralertdissiminationfilter.DateDisseminationRange.To.Day(), 23, 59, 59, 0, Districtweatheralertdissiminationfilter.DateDisseminationRange.To.Location())
			}
			query = append(query, bson.M{"dateOfDissemination": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if Districtweatheralertdissiminationfilter != nil {
		if Districtweatheralertdissiminationfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{Districtweatheralertdissiminationfilter.SortBy: Districtweatheralertdissiminationfilter.SortOrder}})

		}

	}
	return mainPipeline, nil
}
func (d *Daos) FilterDistrictweatheralertdissiminationReport(ctx *models.Context, Districtweatheralertdissiminationfilter *models.DistrictWeatherAlertDissiminationFilter, pagination *models.Pagination) ([]models.RefDistrictWeatherAlertDissimination, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline, err := d.DistrictweatheralertdissiminationQuery(ctx, Districtweatheralertdissiminationfilter)
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Districtweatheralertdissimination query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Districtweatheralertdissiminations []models.RefDistrictWeatherAlertDissimination
	if err = cursor.All(context.TODO(), &Districtweatheralertdissiminations); err != nil {
		return nil, err
	}
	return Districtweatheralertdissiminations, nil
}

//EnableDistrictweatheralertdissimination :""
func (d *Daos) EnableDistrictweatheralertdissimination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTDISSIMINATIONSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableDistrictweatheralertdissimination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTDISSIMINATIONSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrictweatheralertdissimination :""
func (d *Daos) DeleteDistrictweatheralertdissimination(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTDISSIMINATIONSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTDISSIMINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
