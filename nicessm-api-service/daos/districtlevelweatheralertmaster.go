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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveDisease :""
func (d *Daos) SaveDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlertMaster *models.DistrictWeatherAlertMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).InsertOne(ctx.CTX, DistrictWeatherAlertMaster)
	if err != nil {
		return err
	}
	DistrictWeatherAlertMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveDistrictWeatherAlertMasterWithUpsert(ctx *models.Context, district *models.DistrictWeatherAlertMaster) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"district": district.District, "parameterid": district.ParameterId, "month": district.Month, "severityType": district.SeverityType}
	updateData := bson.M{"$set": district}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) UpdateDistrictWeatherAlertMasterUpsertwithMin(ctx *models.Context, district *models.DistrictWeatherAlertMaster) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"district": district.District, "parameterid": district.ParameterId, "month": district.Month, "severityType": district.SeverityType}
	updateData := bson.M{"$set": bson.M{"min": district.Min}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) UpdateDistrictWeatherAlertMasterUpsertwithMax(ctx *models.Context, district *models.DistrictWeatherAlertMaster) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"district": district.District, "parameterid": district.ParameterId, "month": district.Month, "severityType": district.SeverityType}
	updateData := bson.M{"$set": bson.M{"max": district.Max}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleDisease : ""
func (d *Daos) GetSingleDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherAlertMaster, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlertMasters []models.RefDistrictWeatherAlertMaster
	var DistrictWeatherAlertMaster *models.RefDistrictWeatherAlertMaster
	if err = cursor.All(ctx.CTX, &DistrictWeatherAlertMasters); err != nil {
		return nil, err
	}
	if len(DistrictWeatherAlertMasters) > 0 {
		DistrictWeatherAlertMaster = &DistrictWeatherAlertMasters[0]
	}
	return DistrictWeatherAlertMaster, nil
}

//UpdateDistrictWeatherAlertMaster : ""
func (d *Daos) UpdateDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlertMaster *models.DistrictWeatherAlertMaster) error {

	selector := bson.M{"_id": DistrictWeatherAlertMaster.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": DistrictWeatherAlertMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDistrictWeatherAlertMaster : ""
func (d *Daos) FilterDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlertMasterfilter *models.DistrictWeatherAlertMasterFilter, pagination *models.Pagination) ([]models.RefDistrictWeatherAlertMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DistrictWeatherAlertMasterfilter != nil {

		if len(DistrictWeatherAlertMasterfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": DistrictWeatherAlertMasterfilter.ActiveStatus}})
		}
		if len(DistrictWeatherAlertMasterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DistrictWeatherAlertMasterfilter.Status}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if DistrictWeatherAlertMasterfilter != nil {
		if DistrictWeatherAlertMasterfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{DistrictWeatherAlertMasterfilter.SortBy: DistrictWeatherAlertMasterfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("DistrictWeatherAlertMaster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlertMasters []models.RefDistrictWeatherAlertMaster
	if err = cursor.All(context.TODO(), &DistrictWeatherAlertMasters); err != nil {
		return nil, err
	}
	return DistrictWeatherAlertMasters, nil
}

//EnableDistrictWeatherAlertMaster :""
func (d *Daos) EnableDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTMASTERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTMASTERSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrictWeatherAlertMaster :""
func (d *Daos) DeleteDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTMASTERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) DistrictWeatherAlertMasterQuery(ctx *models.Context, DistrictWeatherAlertMasterfilter *models.DistrictWeatherAlertMasterFilter) []bson.M {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DistrictWeatherAlertMasterfilter != nil {

		if len(DistrictWeatherAlertMasterfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": DistrictWeatherAlertMasterfilter.ActiveStatus}})
		}
		if len(DistrictWeatherAlertMasterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DistrictWeatherAlertMasterfilter.Status}})
		}
		if len(DistrictWeatherAlertMasterfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": DistrictWeatherAlertMasterfilter.District}})
		}
		if len(DistrictWeatherAlertMasterfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": DistrictWeatherAlertMasterfilter.District}})
		}
		if len(DistrictWeatherAlertMasterfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": DistrictWeatherAlertMasterfilter.Block}})
		}
		if len(DistrictWeatherAlertMasterfilter.ParameterId) > 0 {
			query = append(query, bson.M{"parameterid": bson.M{"$in": DistrictWeatherAlertMasterfilter.ParameterId}})
		}
		if len(DistrictWeatherAlertMasterfilter.Season) > 0 {
			query = append(query, bson.M{"season": bson.M{"$in": DistrictWeatherAlertMasterfilter.Season}})
		}
		if len(DistrictWeatherAlertMasterfilter.Month) > 0 {
			query = append(query, bson.M{"month": bson.M{"$in": DistrictWeatherAlertMasterfilter.Month}})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	return mainPipeline
}
func (d *Daos) GetDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlertMasterfilter *models.DistrictWeatherAlertMasterFilter) ([]models.GetWeatherAlertMaster, error) {
	var err error
	mainPipeline := []bson.M{}
	mainPipeline = d.DistrictWeatherAlertMasterQuery(ctx, DistrictWeatherAlertMasterfilter)
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{
			"_id": "$block",
			// "type":"$severityType"
			"values": bson.M{
				"$push": "$$ROOT"},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"_id": 1, "values.max": 1, "values.min": 1, "values.severityType": 1}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from":         constants.COLLECTIONBLOCK,
			"as":           "block",
			"localField":   "_id",
			"foreignField": "_id",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"block": bson.M{"$arrayElemAt": []interface{}{"$block.name", 0}}}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("DistrictWeatherAlertMaster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlertMasters []models.GetWeatherAlertMaster
	if err = cursor.All(context.TODO(), &DistrictWeatherAlertMasters); err != nil {
		return nil, err
	}
	return DistrictWeatherAlertMasters, nil
}
func (d *Daos) GetSingleDistrictWeatherAlertMasterWithSpecialIds(ctx *models.Context, month string, parameter string, district string) ([]models.DistrictWeatherAlertMaster, error) {
	monthid, err := primitive.ObjectIDFromHex(month)
	if err != nil {
		return nil, err
	}
	parameterid, err := primitive.ObjectIDFromHex(parameter)
	if err != nil {
		return nil, err
	}
	districtid, err := primitive.ObjectIDFromHex(district)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"month": monthid, "parameterid": parameterid, "district": districtid}})

	d.Shared.BsonToJSONPrintTag("GetSingleDistrictWeatherAlertMasterWithSpecialIds query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlertMasters []models.DistrictWeatherAlertMaster
	//var DistrictWeatherAlertMaster *models.RefDistrictWeatherAlertMaster
	if err = cursor.All(ctx.CTX, &DistrictWeatherAlertMasters); err != nil {
		return nil, err
	}

	return DistrictWeatherAlertMasters, nil
}
