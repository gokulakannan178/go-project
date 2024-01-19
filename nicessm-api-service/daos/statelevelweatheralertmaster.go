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
func (d *Daos) SaveStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlertMaster *models.StateWeatherAlertMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).InsertOne(ctx.CTX, StateWeatherAlertMaster)
	if err != nil {
		return err
	}
	StateWeatherAlertMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveStateWeatherAlertMasterWithUpsert(ctx *models.Context, state *models.StateWeatherAlertMaster) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"state": state.State, "parameterid": state.ParameterId, "month": state.Month, "severityType": state.SeverityType}
	updateData := bson.M{"$set": state}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) UpdateStateWeatherAlertMasterUpsertwithMin(ctx *models.Context, state *models.StateWeatherAlertMaster) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"state": state.State, "parameterid": state.ParameterId, "month": state.Month, "severityType": state.SeverityType}
	updateData := bson.M{"$set": bson.M{"min": state.Min}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) UpdateStateWeatherAlertMasterUpsertwithMax(ctx *models.Context, state *models.StateWeatherAlertMaster) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"state": state.State, "parameterid": state.ParameterId, "month": state.Month, "severityType": state.SeverityType}
	updateData := bson.M{"$set": bson.M{"max": state.Max}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleDisease : ""
func (d *Daos) GetSingleStateWeatherAlertMaster(ctx *models.Context, UniqueID string) (*models.RefStateWeatherAlertMaster, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlertMasters []models.RefStateWeatherAlertMaster
	var StateWeatherAlertMaster *models.RefStateWeatherAlertMaster
	if err = cursor.All(ctx.CTX, &StateWeatherAlertMasters); err != nil {
		return nil, err
	}
	if len(StateWeatherAlertMasters) > 0 {
		StateWeatherAlertMaster = &StateWeatherAlertMasters[0]
	}
	return StateWeatherAlertMaster, nil
}

//UpdateStateWeatherAlertMaster : ""
func (d *Daos) UpdateStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlertMaster *models.StateWeatherAlertMaster) error {

	selector := bson.M{"_id": StateWeatherAlertMaster.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": StateWeatherAlertMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterStateWeatherAlertMaster : ""
func (d *Daos) FilterStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlertMasterfilter *models.StateWeatherAlertMasterFilter, pagination *models.Pagination) ([]models.RefStateWeatherAlertMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if StateWeatherAlertMasterfilter != nil {

		if len(StateWeatherAlertMasterfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": StateWeatherAlertMasterfilter.ActiveStatus}})
		}
		if len(StateWeatherAlertMasterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": StateWeatherAlertMasterfilter.Status}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if StateWeatherAlertMasterfilter != nil {
		if StateWeatherAlertMasterfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{StateWeatherAlertMasterfilter.SortBy: StateWeatherAlertMasterfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("StateWeatherAlertMaster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlertMasters []models.RefStateWeatherAlertMaster
	if err = cursor.All(context.TODO(), &StateWeatherAlertMasters); err != nil {
		return nil, err
	}
	return StateWeatherAlertMasters, nil
}

//EnableStateWeatherAlertMaster :""
func (d *Daos) EnableStateWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTMASTERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableStateWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTMASTERSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteStateWeatherAlertMaster :""
func (d *Daos) DeleteStateWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTMASTERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) StateWeatherAlertMasterQuery(ctx *models.Context, StateWeatherAlertMasterfilter *models.StateWeatherAlertMasterFilter) []bson.M {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if StateWeatherAlertMasterfilter != nil {

		if len(StateWeatherAlertMasterfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": StateWeatherAlertMasterfilter.ActiveStatus}})
		}
		if len(StateWeatherAlertMasterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": StateWeatherAlertMasterfilter.Status}})
		}
		if len(StateWeatherAlertMasterfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": StateWeatherAlertMasterfilter.State}})
		}
		if len(StateWeatherAlertMasterfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": StateWeatherAlertMasterfilter.District}})
		}
		if len(StateWeatherAlertMasterfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": StateWeatherAlertMasterfilter.Block}})
		}
		if len(StateWeatherAlertMasterfilter.ParameterId) > 0 {
			query = append(query, bson.M{"parameterid": bson.M{"$in": StateWeatherAlertMasterfilter.ParameterId}})
		}
		if len(StateWeatherAlertMasterfilter.Season) > 0 {
			query = append(query, bson.M{"season": bson.M{"$in": StateWeatherAlertMasterfilter.Season}})
		}
		if len(StateWeatherAlertMasterfilter.Month) > 0 {
			query = append(query, bson.M{"month": bson.M{"$in": StateWeatherAlertMasterfilter.Month}})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	return mainPipeline
}
func (d *Daos) GetStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlertMasterfilter *models.StateWeatherAlertMasterFilter) ([]models.GetWeatherAlertMaster, error) {
	var err error
	mainPipeline := []bson.M{}
	mainPipeline = d.StateWeatherAlertMasterQuery(ctx, StateWeatherAlertMasterfilter)
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
	d.Shared.BsonToJSONPrintTag("StateWeatherAlertMaster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlertMasters []models.GetWeatherAlertMaster
	if err = cursor.All(context.TODO(), &StateWeatherAlertMasters); err != nil {
		return nil, err
	}
	return StateWeatherAlertMasters, nil
}
func (d *Daos) GetSingleStateWeatherAlertMasterWithSpecialIds(ctx *models.Context, month string, parameter string, state string) ([]models.StateWeatherAlertMaster, error) {
	monthid, err := primitive.ObjectIDFromHex(month)
	if err != nil {
		return nil, err
	}
	parameterid, err := primitive.ObjectIDFromHex(parameter)
	if err != nil {
		return nil, err
	}
	stateid, err := primitive.ObjectIDFromHex(state)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"month": monthid, "parameterid": parameterid, "state": stateid}})

	d.Shared.BsonToJSONPrintTag("GetSingleStateWeatherAlertMasterWithSpecialIds query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlertMasters []models.StateWeatherAlertMaster
	//var StateWeatherAlertMaster *models.RefStateWeatherAlertMaster
	if err = cursor.All(ctx.CTX, &StateWeatherAlertMasters); err != nil {
		return nil, err
	}

	return StateWeatherAlertMasters, nil
}
